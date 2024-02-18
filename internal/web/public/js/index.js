
function setToday() {
    let today = new Date().toLocaleDateString('en-CA');
    
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
    document.getElementById('realDate').value = date;
    document.getElementById('checkList').innerHTML = '';

    let groupMap = new Map();
    let checks = [];
    let url = '/date/'+date;

    checks = await (await fetch(url)).json();
    if (checks) {
        groupMap = getGroupMap(checks);

        groupMap.forEach (function(value, key) {
            genHTML(key, value);
        })
    }
}

function genHTML(gr, checks) {
    let btn = '';
    let html = `<h5>${gr}</h5>`;
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
            <a href="#" onclick="addOne(${checks[i].ID})"><p>
                <div class="btn-group btn-group-lg">
                    <button id="btn${checks[i].ID}" class="my-btn-lg ${btn} style="border-left-width: thick; border-left-color: ${checks[i].Color};">
                        <img src="${checks[i].Icon}" style="height:1.3em;"/>&nbsp;
                        ${checks[i].Name}
                    </button>
                    <button id="count${checks[i].ID}" class="${btn}>
                        ${checks[i].Count}
                    </button>
                </div>
            </p></a>
        </div>`;
    }

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

function setFormDate(where) {
    dateStr = document.getElementById('realDate').value;

    if (where) {
        let year  = dateStr.substring(0,4);
        let month = dateStr.substring(5,7);
        let day   = dateStr.substring(8,10);
        var date  = new Date(year, month-1, day);

        date.setDate(date.getDate() + parseInt(where));
        let left = date.toLocaleDateString('en-CA');

        createView(left);
    } else {
        createView(dateStr);
    }
}

async function updatePlan() {
    date = document.getElementById('realDate').value;

    let resp = '';
    let url = '/update/'+date;
    resp = await (await fetch(url)).json();

    console.log("UPDATE:", resp);

    createView(date);
}