
function setToday() {
    let today = new Date().toJSON().slice(0, 10);
    return today;
}

function createView(checks, groups) {
    let date = setToday();

    let len = checks.length;
    let dateChecks = [];
    for (let i = 0 ; i < len; i++) {
        if (checks[i].Date == date) {
            console.log(checks[i].Group, checks[i].Name)
            dateChecks.push(checks[i]);
        }
    }
    console.log(dateChecks);

    len = groups.length;
    let len1 = dateChecks.length;
    let grChecks = [];
    for (i = 0 ; i < len; i++) {
        for (let j = 0 ; j < len1; j++) {
            if (groups[i] == dateChecks[j].Group) {
                console.log("GR:"+groups[i]+"NAME:"+dateChecks[j].Name);
                grChecks.push(dateChecks[j]);
            }
        }
        genHTML(groups[i], grChecks);
        grChecks = [];
    }
}

function genHTML(gr, checks) {

    let html = '<h5>'+gr+'</h5>';
    let len = checks.length;
    for (let i = 0 ; i < len; i++) {
        html = html+'<div class="col-md-auto"><a href="/add/'+checks[i].ID+'">';
        if (checks[i].Count) {
            html = html+'<button class="btn btn-lg my-btn-lg" style="background-color: '+checks[i].Color+';">'+checks[i].Name+'</button><button class="btn btn-lg" style="background-color: '+checks[i].Color+';">'+checks[i].Count+'</button></div></a>';
        } else {
            html = html+'<button class="btn btn-lg btn-outline-primary my-btn-lg">'+checks[i].Name+'</button><button class="btn btn-lg btn-outline-primary">'+checks[i].Count+'</button></div></a>';
        }
    }

    document.getElementById('checkList').insertAdjacentHTML('beforeend', html);
}