const puppeteer = require('puppeteer');
const countries = require('./countries.json');

(async () => {
    const browser = await puppeteer.launch();
    const page = await browser.newPage();

    page.setViewport({ width: 1920, height: 1080 });

    for (let country of countries) {
        for (let zoom=0; zoom<8; zoom++) {
            let cca3 = country.cca3.toLowerCase();

            console.log(cca3, zoom);

            try {
                await page.goto(`http://localhost:8000/maps/map.html#${cca3},${zoom}`, { timeout: 1 });
            } catch (_) {} // :D
            await page.waitForFunction(`window.currCCA3 == '${cca3}'`);
            await page.waitFor(2000);
            await page.screenshot({ path: `screenshots/${cca3}-${zoom}.png` });
        }
    }

    await browser.close();
})();
