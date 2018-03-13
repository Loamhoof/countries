const load = (filepath) => new Promise((resolve, reject) => {
    let xhr = new XMLHttpRequest();

    xhr.open('GET', `http://localhost:8000/${filepath}`);

    xhr.onload = () => {
        if (xhr.status == 200) {
            resolve(JSON.parse(xhr.responseText));
        }
    };

    xhr.send();
});

mapboxgl.accessToken = '';

const map = new mapboxgl.Map({
    container: 'map',
    style: 'mapbox://styles/mapbox/light-v9',
    zoom: 3
});

const countriesP = load('countries.json');
const mapP = new Promise((resolve, reject) => {
    map.on('load', resolve);
});

let currLayerID = '';
const displayCountry = () => {
    Promise.all([countriesP, mapP]).then(([countries]) => {
        const cca3 = location.hash.substr(1);

        if (cca3 == '') {
            return;
        }

        load(`data/${cca3}.geo.json`).then((geojson) => {
            let country;
            for (country of countries) {
                if (country.cca3 == cca3.toUpperCase()) {
                    break;
                }
            }

            const [lat, lng] = country.latlng;

            if (currLayerID != '') {
                map.removeLayer(currLayerID);
            }

            currLayerID = cca3 + Date.now();

            map.setCenter([lng, lat]);
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
        });
    });
};

window.onhashchange = displayCountry;

displayCountry();
