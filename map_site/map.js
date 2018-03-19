;(() => {
    const load = (filepath) => new Promise((resolve, reject) => {
        let xhr = new XMLHttpRequest();

        xhr.open('GET', `/${filepath}`);

        xhr.onload = () => {
            if (xhr.status == 200) {
                resolve(JSON.parse(xhr.responseText));
            }
        };

        xhr.send();
    });

    mapboxgl.accessToken = 'MAPBOXTOKEN';

    const map = new mapboxgl.Map({
        container: 'map',
        style: 'mapbox://styles/mapbox/light-v9',
        zoom: 0,
        keyboard: false
    });

    const layersToRemove = [
        'admin-3-4-boundaries',
        'admin-3-4-boundaries-bg',
        'place-city-sm',
        'place-city-md-n',
        'place-city-md-s',
        'state-label-sm',
        'state-label-md',
        'state-label-lg'
    ];

    const countriesP = load('data/countries.json');
    const mapP = new Promise((resolve, reject) => {
        map.on('load', () => {
            for (let layer of layersToRemove) {
                map.removeLayer(layer);
            }

            resolve();
        });
    });

    currCCA3 = ''; // global for puppeteer
    let currZoom = 0;
    let currLayerID = '';
    const displayCountry = () => {
        Promise.all([countriesP, mapP]).then(([countries]) => {
            const [nextCCA3, zoom=3] = location.hash.substr(1).split(',');

            if (nextCCA3 == '') {
                return;
            }

            load(`data/geojson/${nextCCA3.toLowerCase()}.geo.json`).then((geojson) => {
                if (zoom != currZoom) {
                    currZoom = zoom;

                    map.setZoom(zoom);
                }

                const country = countries.find(({cca3}) => cca3 == nextCCA3);
                if (!country) {
                    return;
                }

                const [lat, lng] = country.latlng;

                map.setCenter([lng, lat]);

                if (nextCCA3 == currCCA3) {
                    return;
                }

                if (currLayerID != '') {
                    map.removeLayer(currLayerID);
                }

                currLayerID = '' + Date.now();

                map.addLayer({
                    'id': currLayerID,
                    'type': 'fill',
                    'source': {
                        'type': 'geojson',
                        'data': geojson,
                    },
                    'layout': {},
                    'paint': {
                        'fill-outline-color': '#000',
                        'fill-color': '#ddd',
                        'fill-opacity': 0.6
                    }
                });

                currCCA3 = nextCCA3;
            });
        });
    };

    window.onhashchange = displayCountry;

    window.onkeyup = ({ key }) => {
        countriesP.then((countries) => {
            switch (key) {
            case 'ArrowUp':
                location.hash = `#${currCCA3},${Math.min(24, parseInt(currZoom)+1)}`;

                return;
            case 'ArrowDown':
                location.hash = `#${currCCA3},${Math.max(0, parseInt(currZoom)-1)}`;

                return;
            }

            const i = countries.findIndex(({cca3}) => cca3 == currCCA3);
            if (i == -1) {
                return;
            }

            switch (key) {
            case 'ArrowLeft':
                location.hash = `#${countries[(i-1+countries.length)%countries.length].cca3},${currZoom}`;

                return;
            case 'ArrowRight':
                location.hash = `#${countries[(i+1+countries.length)%countries.length].cca3},${currZoom}`;

                return;
            }
        });
    };

    displayCountry();
})();
