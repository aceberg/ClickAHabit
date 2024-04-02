var layout;
var tabname;

function setToday() {
    let today = new Date().toLocaleDateString('en-CA');
    layout = localStorage.getItem("layout");

    tabname = localStorage.getItem("tabname");
    if (tabname != 'weeks') {
        tabname = "checks"
    }
    
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

    let mapAsc = new Map([...groupMap.entries()].sort());

    return mapAsc
}

async function createView(date) {
    document.getElementById('realDate').value = date;
    document.getElementById('checkList').innerHTML = '';

    let groupMap = new Map();
    let checks = [];
    let url = '/date/'+tabname+'/'+date;

    checks = await (await fetch(url)).json();
    if (checks) {
        groupMap = getGroupMap(checks);

        groupMap.forEach (function(value, key) {
            genHTML(key, value);
        })
    }
}

function genHTML(gr, checks) {
    let btn = ''; vcol = ''; vdiv = ''; hcol = ''; hdiv = ''; icon = '';
    if (layout == 'vert') {
        vcol = '<div class="col-md-auto">';
        vdiv = '</div>';
    } else {
        hcol = '<div class="col-md-auto">';
        hdiv = '</div>';
    }
    let html = vcol + `<a href="/plan/?gr=${gr}"><h5>${gr}</h5></a>`;
    let len = checks.length;

    for (let i = 0 ; i < len; i++) {
        btn = `btn btn-lg`;
        if (checks[i].Count) {
            btn = btn + ` btn-primary"`;
        } else {
            btn = btn + ` btn-outline-primary"`;
        }
        if (checks[i].Icon) {
            icon = `<img src="${checks[i].Icon}" style="height:1.3em;"/>&nbsp;`;
        } else {
            icon = '';
        }
        html = html + hcol + `
            </p>
                <div class="btn-group btn-group-lg" oncontextmenu="showMenu(event,${checks[i].ID}, '${checks[i].Link}');">
                    <button id="btn${checks[i].ID}" onclick="addOne(${checks[i].ID});" class="my-btn-lg ${btn} style="border-left-width: thick; border-left-color: ${checks[i].Color};">
                        ${icon}
                        ${checks[i].Name}
                    </button>
                    <button id="count${checks[i].ID}" onclick="showMenu(event,${checks[i].ID}, '${checks[i].Link}');" class="${btn}>
                        ${checks[i].Count}
                    </button>
                </div>
            </p>
        ` + hdiv;
    }
    html = html + vdiv;

    document.getElementById('checkList').insertAdjacentHTML('beforeend', html);
}

async function addOne(id) {
    let resp = '';
    let url = '/add/'+tabname+'/'+id;
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
    let url = '/update/'+tabname+'/'+date;
    resp = await (await fetch(url)).json();

    createView(date);
}

function toggleLayout() {

    if (layout == 'vert') {
        layout = '';
    } else {
        layout = 'vert';
    }
    localStorage.setItem("layout", layout);

    setToday();
}

function toggleTabname() {

    if (tabname == 'weeks') {
        tabname = 'checks';
    } else {
        tabname = 'weeks';
    }
    localStorage.setItem("tabname", tabname);

    setFormDate(0);
}

function showMenu(e, id, link) {
    let linkBtn = '';

    // console.log("LINK ="+ link);

    e.preventDefault()
    let menu = document.createElement("div");
    menu.id = "ctxmenu"
    menu.className ="btn-group-vertical";
    menu.style.position= "fixed";
    menu.style.top = e.pageY + 'px';
    menu.style.left = e.pageX + 'px';
    menu.onmouseleave = () => ctxmenu.outerHTML = '';
    menu.onpointerleave = () => ctxmenu.outerHTML = '';

    document.getElementById('indexBody').ontouchmove = () => ctxmenu.outerHTML = '';
    
    if (link !== '') {
        linkBtn = `<button class="btn" onclick="window.open('${link}', '_blank');">Open Link in a New Tab</button>`;
    }
    menu.innerHTML = linkBtn +`
        <button class="btn" onclick="window.open('/stats/${tabname}/${id}', '_self');">Statistics</button>
        <button class="btn" onclick="histDel(${id});">Reset Todays Counter</button>`
    
    document.getElementById('checkList').appendChild(menu);
}

async function histDel(id) {

    let url = '/del/'+tabname+'/'+id;
    resp = await (await fetch(url));

    setFormDate(0);
}