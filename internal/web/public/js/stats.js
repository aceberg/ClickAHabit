
async function getMore(key) {
    let checks = [];
    let url = '/smore/'+key;

    heat = await (await fetch(url)).json();

    document.getElementById('statsHeader').innerHTML = key;
    document.getElementById('statsTable').innerHTML = '';

    html = `<div class="horiz-scroll">
                <div style="max-height: 150px; width: 1200px;">
                    <canvas id="matrix-chart" style="height: 100%; width: 100%;"></canvas>
                </div>
            </div>`;

    document.getElementById('statsTable').insertAdjacentHTML('beforeend', html);

    let color = getComputedStyle(document.body).getPropertyValue('--bs-primary');

    makeChart(heat, color);
}