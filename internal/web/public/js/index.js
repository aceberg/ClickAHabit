var curDate = '';

setToday();

function setToday() {
    let today = new Date().toJSON().slice(0, 10);
    createView(today);
}

function getGroupMap(checks) {
    const groupMap = new Map();
    let tmpChecks = [];

    let len = checks.length;

    for (let i = 0 ; i < len; i++) {
        if (groupMap.has(checks[i].Group)) {
            tmpChecks = groupMap.get(checks[i].Group);
        } else {
            tmpChecks = [];
        }
        tmpChecks.push(checks[i]);
        groupMap.set(checks[i].Group, tmpChecks);
    }
    return groupMap
}

async function createView(date) {
    curDate = date;

    if (document.getElementById('checkList')) {
        document.getElementById('checkList').innerHTML = '';
    }

    let groupMap = new Map();
    let checks = [];
    let url = '/date/'+date;

    checks = await (await fetch(url)).json();
    groupMap = getGroupMap(checks);

    groupMap.forEach (function(value, key) {
        genHTML(key, value);
    })
}

function genHTML(gr, checks) {
    let btn = '';
    let html = `<p>
                <h5>${gr}</h5>`;
    let len = checks.length;
    for (let i = 0 ; i < len; i++) {
        btn = `btn btn-lg`;
        if (checks[i].Count) {
            btn = btn + ` btn-primary"`;
        } else {
            btn = btn + ` btn-outline-primary"`;
        }
        html = html + `
        <div class="col-md-auto">
            <a href="#" onclick="addOne(${checks[i].ID})">
                <button id="btn${checks[i].ID}" class="my-btn-lg ${btn}>
                    <i class="bi bi-circle-fill" style="color: ${checks[i].Color};"></i>&nbsp;${checks[i].Name}
                </button>
                <button id="count${checks[i].ID}" class="${btn}>
                    ${checks[i].Count}
                </button>
            </a>
        </div>`;
    }
    html = html + `</p>`;

    document.getElementById('checkList').insertAdjacentHTML('beforeend', html);
}

async function addOne(id) {
    let resp = '';
    let url = '/add/'+id;
    resp = await (await fetch(url)).json();

    document.getElementById('count'+id).innerHTML = resp;

    if (resp == 1) {
        document.getElementById('btn'+id).classList.remove('btn-outline-primary');
        document.getElementById('count'+id).classList.remove('btn-outline-primary');
        document.getElementById('btn'+id).classList.add('btn-primary');
        document.getElementById('count'+id).classList.add('btn-primary');
    }
}