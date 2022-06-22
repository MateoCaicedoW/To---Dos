//variables for teams
import validate from "./validate.js";

let formTeam = document.getElementById("form-team");
let btnEditTeam = document.getElementById('btn-edit-team')
let btnSaveTeam = document.getElementById('btn-create-team')
let btnCancelTeam = document.getElementById('btn-cancel-team')
let btnDeleteTeam = document.getElementById('btn-delete-modal-team')
let btnClose =document.getElementById('close-button-teams');
const divWarning = document.getElementById('warning-div-team')
function getAllTeams(data) {
    let html2 = ''
    // list of teams
    for (let index = 0; index < data.length; index++) {
        const team = data[index];
        const name = team.Name.charAt(0).toUpperCase() + team.Name.slice(1);
        const type = team.Type.charAt(0).toUpperCase() + team.Type.slice(1);
        const country = team.Country.charAt(0).toUpperCase() + team.Country.slice(1);
        html2 += `
        <tr>
            <th scope="row">${team.ID}</th>
            <td>${name}</td>
            <td>${type}</td>
            <td>${country}</td>

            <td class="d-flex py-5 py-xl-4 py-md-5 justify-content-center align-items-center">

                <button type="button"  class="btn btn-success ms-2 edit-team"><i class="fas fa-edit icon-edit-team"></i></button>
                <button type="button" class="btn btn-danger ms-2 delete-show-modal-team" data-bs-toggle="modal" data-bs-target="#modal-Team"><i class="far fa-trash-alt icon-delete-team "></i></button>

            </td>

        </tr>
        `
    }
    document.getElementById('body-teams').innerHTML = html2


    //Show information Team team
    let select = document.getElementById("types");
   

    let html4 = ''
    html4 = `
            <option selected disabled value="">Choose...</option>
            `
    select.innerHTML += html4
    const route = 'http://localhost:3000/api/types'
    fetch(route, {
        method: 'GET',
    }).then(res => res.json()
    ).then((data) => {
        for (let index = 0; index < data.length; index++) {
            let element = data[index];
            element = element.charAt(0).toUpperCase() + element.slice(1);
            html4 += `
            <option value="${data[index]}">${element}</option>
            `
        }
        select.innerHTML = html4
    })

    select.addEventListener('change', () => {
        if (select.value == "national") {
            document.getElementById('country-team').disabled = true
        } else {
            document.getElementById('country-team').disabled = false
        }
    })



}

function resetFormTeam(){
    btnCancelTeam.addEventListener('click', () => {
        formTeam.reset()
        btnSaveTeam.classList.remove('d-none')
        btnEditTeam.classList.add('d-none')
        document.getElementById('country-team').disabled = false
    
    })
}



let idTeam = ""
document.getElementById('body-teams').addEventListener('click', (e)=>{
    let ev = e.target
    let select = document.getElementById("types");
    let type
    if (ev.classList.contains('edit-team')){
        document.getElementById('name-team').value = ev.parentElement.parentElement.children[1].innerHTML
        document.getElementById('country-team').value = ev.parentElement.parentElement.children[3].innerHTML
        idTeam = ev.parentElement.parentElement.children[0].innerHTML;
        type = ev.parentElement.parentElement.children[2].innerHTML.toLowerCase()
        SetTypes(select,type)

    }else if(ev.classList.contains('delete-show-modal-team')){
        idTeam = ev.parentElement.parentElement.children[0].innerHTML;

    }else if(ev.classList.contains('icon-delete-team')){
        idTeam = ev.parentElement.parentElement.parentElement.children[0].innerHTML

    }else if (ev.classList.contains('icon-edit-team')){
        idTeam = ev.parentElement.parentElement.parentElement.children[0].innerHTML
        document.getElementById('name-team').value = ev.parentElement.parentElement.parentElement.children[1].innerHTML
        document.getElementById('country-team').value = ev.parentElement.parentElement.parentElement.children[3].innerHTML
        idTeam = ev.parentElement.parentElement.parentElement.children[0].innerHTML;
        type = ev.parentElement.parentElement.parentElement.children[2].innerHTML.toLowerCase()
        SetTypes(select,type)
    }
})


function SetTypes(select,type){
    btnSaveTeam.classList.add('d-none')
    btnEditTeam.classList.remove('d-none')
    for (let index = 0; index < select.options.length; index++) {
        const element = select.options[index];
        if (element.value == type) {
            select.selectedIndex = index
            if (type == "national") {
                document.getElementById('country-team').disabled = true
            } else {
                document.getElementById('country-team').disabled = false
            }
        }
    }
}


function deleteTeam() {
    btnDeleteTeam.addEventListener('click', (e) => {
        e.preventDefault();
        let route = 'http://localhost:3000/api/teams/' + idTeam
        console.log(route);
        fetch(route, {
            method: 'DELETE',
            headers: {
                'Content-Type': 'application/json',
            },
            cors: 'no-cors'
            
        }).then(res => {
            res.json()
            
        }).then(() => {
            window.location.href = '/'
        }).catch(err => console.log(err))
    })
}
//Delete team


//Edit team

function editTeam(){
    btnEditTeam.addEventListener('click', (e) => {
        e.preventDefault();
    
        let route = 'http://localhost:3000/api/teams/' + idTeam
        let team = {
            Name: formTeam["name-team"].value,
            Type: formTeam["types"].value,
            Country: formTeam["country-team"].value,
        }
        console.log(team);
        fetch(route, {
            method : "PUT",
            body : JSON.stringify(team),
        }).then(res => res.json())
        .then((data) => {
            let message = document.getElementById('warning-message-team')
            validate(divWarning, message, data, formTeam, btnSaveTeam, btnEditTeam)
        }).catch(err => console.log(err))
    })
}

function saveTeam(){
    btnSaveTeam.addEventListener('click', (e) => {
        e.preventDefault();
        let route = 'http://localhost:3000/api/teams'
        let team = {
            Name: formTeam["name-team"].value,
            Type: formTeam["types"].value,
            Country: formTeam["country-team"].value,
        }
        console.log(team);
        fetch(route, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'Alow-Control-Allow-Origin': '*',
            },
            
            body: JSON.stringify(team)
        }).then(res => res.json()
        ).then((data) => {
            let message = document.getElementById('warning-message-team')
            validate(divWarning, message, data, formTeam, btnSaveTeam, btnEditTeam)
        }).catch(err => console.log(err))
        // window.location.href = '/'
    
    })
}


function closeWarning(){
    btnClose.addEventListener('click', (e) => {
        e.preventDefault();
        divWarning.classList.add('d-none')
    })

}
export default {getAllTeams, resetFormTeam, saveTeam, editTeam, deleteTeam, closeWarning}