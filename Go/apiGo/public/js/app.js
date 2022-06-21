//variables for player
let formPlayer = document.getElementById("form-player");
let btnEditPlayer = document.getElementById('btn-edit-player')
let btnSavePlayer = document.getElementById('btn-create-player')
let btnCancelPlayer = document.getElementById('btn-cancel-player')
let btnDeletePlayer = document.getElementById('btn-delete-player-modal')
let showModalDelete = document.getElementsByClassName('delete-show-modal-player')


//variables for teams
let formTeam = document.getElementById("form-team");
let btnEditTeam = document.getElementById('btn-edit-team')
let btnSaveTeam = document.getElementById('btn-create-team')
let btnCancelTeam = document.getElementById('btn-cancel-team')
let btnDeleteTeam = document.getElementById('btn-delete-modal-team')
let showModalDeleteTeam = document.getElementsByClassName('delete-show-modal-team')



// //Show player and teams when stay on a page
// let windowsPlayer = document.getElementById('player')
// let windowsTeams = document.getElementById('teams')
window.onload = loadInformationTabs

function loadInformationTabs() {
    let arrayTeams = []
    const ruta2 = 'http://localhost:3000/api/teams'
    fetch(ruta2, {
        method: 'GET',
        cache: 'default'
    })
        .then(res => {
            return res.json()
        })
        .then((data) => {
            arrayTeams = data.Data
            getAllPlayers(arrayTeams)
            getAllTeams(arrayTeams)

        })
}


//***Functions for players */

// List all players 

function getAllPlayers(data) {

    const ruta = 'http://localhost:3000/api/players'
    fetch(ruta, {
        method: 'GET',
        cache: 'default'
    })
        .then(res => {
            return res.json()
        })
        .then((data) => {
            let html = ''

            //list of players
            for (let index = 0; index < data.Data.length; index++) {
                const player = data.Data[index];
                let team = ''
                let club = ''
                for (let index = 0; index < player.Teams.length; index++) {
                    const element = player.Teams[index];
                    const name = element.Name.charAt(0).toUpperCase() + element.Name.slice(1);
                    if (element.Type == "club") {

                        club = `${name}`
                    } else {
                        team = `${name}`
                    }

                }
                html += `
            <tr>
                <th scope="row">${player.ID}</th>
                <td>${player.FirstName}</td>
                <td>${player.LastName}</td>
                <td>${player.Level}</td>
                <td>${player.Age}</td>
                <td>${player.Position}</td>
                <td>${player.PhysicalCondition}</td>
                <td>${club}</td>
                <td>${team}</td>
                <td class="d-flex py-5 py-xl-4 py-md-5 justify-content-center align-items-center">
                
                    <button  type="button" class="btn btn-success ms-2 edit-player"><i class="fas fa-edit"></i></button>
                    <button type="button" class="btn btn-danger ms-2  delete-show-modal-player" data-bs-toggle="modal" data-bs-target="#exampleModal"><i class="far fa-trash-alt "></i></button>
                   
                </td>
                
            </tr>
            `
            }
            document.getElementById('body-players').innerHTML = html

            //Show One Player
            //let form = document.getElementById('form-player')
            let idPlayer = ""
            let showInformationPlayer = document.getElementsByClassName('edit-player')
            for (let index = 0; index < showInformationPlayer.length; index++) {
                const element = showInformationPlayer[index];
                element.addEventListener('click', () => {
                    btnEditPlayer.classList.remove('d-none')
                    btnSavePlayer.classList.add('d-none')
                    document.getElementById('first-name').value = element.parentElement.parentElement.children[1].innerHTML
                    document.getElementById('last-name').value = element.parentElement.parentElement.children[2].innerHTML
                    document.getElementById('level').value = element.parentElement.parentElement.children[3].innerHTML
                    document.getElementById('age').value = element.parentElement.parentElement.children[4].innerHTML

                    for (let item = 0; item < document.getElementById('position').length; item++) {
                        const i = document.getElementById('position').options[item]
                        if (i.value.toUpperCase() == element.parentElement.parentElement.children[5].innerHTML.replace(" ", "").toUpperCase()) {
                            i.selected = true
                        }
                    }

                    for (let item = 0; item < document.getElementById('condition').length; item++) {
                        const i = document.getElementById('condition').options[item]
                        if (i.value.toUpperCase() == element.parentElement.parentElement.children[6].innerHTML.toUpperCase()) {
                            i.selected = true
                        }
                    }


                    for (let item = 0; item < document.getElementById('club').length; item++) {
                        const i = document.getElementById('club').options[item]
                        const club = element.parentElement.parentElement.children[7].innerHTML
                        if (club != "") {
                            if (i.text.toUpperCase() == club.toUpperCase()) {
                                i.selected = true
                            }
                        } else {
                            document.getElementById('club').options[0].selected = true
                        }
                    }

                    for (let item = 0; item < document.getElementById('national').length; item++) {
                        const i = document.getElementById('national').options[item]
                        const national = element.parentElement.parentElement.children[8].innerHTML
                        if (national != "") {
                            if (i.text.toUpperCase() == national.toUpperCase()) {
                                i.selected = true
                            }
                        } else {
                            document.getElementById('national').options[0].selected = true
                        }
                    }

                    window.location.href = '#form-player'

                })
            }
            btnEditPlayer.addEventListener('click', (e) => {
                editPlayer(e, idPlayer)
            })


        })


    let clubSelect = document.getElementById("club");
    let nationalSelect = document.getElementById("national");
    let html = ''
    html = `
            <option selected disabled value="">Choose...</option>
            `
    clubSelect.innerHTML += html
    nationalSelect.innerHTML += html
    for (let index = 0; index < data.length; index++) {

        const team = data[index];
        const name = team.Name.charAt(0).toUpperCase() + team.Name.slice(1);
        if (team.Type == "club") {
            html = `
            <option value="${team.Name}">${name}</option>
            `
            clubSelect.innerHTML += html
        } else {
            html = `
            <option value="${team.Name}">${name}</option>
            `
            nationalSelect.innerHTML += html


        }

        //

    }

    //List Position
    let positionSelect = document.getElementById("position");
    let html3 = ''
    html3 = `
            <option selected disabled value="">Choose...</option>
            `
    positionSelect.innerHTML += html3
    const ruta2 = 'http://localhost:3000/api/positions'
    fetch(ruta2, {
        method: 'GET',
    }).then(res => res.json()
    ).then((data) => {

        for (let index = 0; index < data.length; index++) {
            let element = data[index];

            element = element.toUpperCase()
            html3 += `
            <option value="${data[index]}">${element}</option>
            `

        }
        positionSelect.innerHTML = html3
    })

    //List Physical Condition

    let physicalConditionSelect = document.getElementById("condition");
    let html4 = ''
    html4 = `
            <option selected disabled value="">Choose...</option>
            `
    physicalConditionSelect.innerHTML += html4

    const ruta3 = 'http://localhost:3000/api/conditions'
    fetch(ruta3, {
        method: 'GET',
    }).then(res => res.json()
    ).then((data) => {

        for (let index = 0; index < data.length; index++) {
            let element = data[index];
            element = element.toUpperCase()
            html4 += `
            <option value="${data[index]}">${element}</option>
            `

        }
        physicalConditionSelect.innerHTML = html4
    })


}


//editPlayer Function
function editPlayer(e, idPlayer) {
    e.preventDefault();

    let player = {
        FirstName: formPlayer["first-name"].value,
        LastName: formPlayer["last-name"].value,
        Level: parseInt(formPlayer["level"].value),
        Age: parseInt(formPlayer["age"].value),
        Position: formPlayer["position"].value,
        PhysicalCondition: formPlayer["condition"].value,
        Teams: []
    }
    if (formPlayer["club"].value != "") {
        team = {
            Name: formPlayer["club"].value,
        }
        player.Teams.push(team)
    }
    if (formPlayer["national"].value != "") {
        team = {
            Name: formPlayer["national"].value,
        }
        player.Teams.push(team)
    }

    let route = 'http://localhost:3000/api/players/' + idPlayer
    fetch(route, {
        method: 'PUT',
        headers: {
            'Content-Type': 'application/json',
            'Alow-Control-Allow-Origin': '*',
        },
        body: JSON.stringify(player),

    }).then(res => res.json())
        .then((data) => {
            const divWarning = document.getElementById('warning-div')
            let message = document.getElementById('warning-message')
            validate(divWarning, message, data, formPlayer, btnSavePlayer, btnEditPlayer)
        }).catch(err => console.log(err))

}

btnCancelPlayer.addEventListener('click', () => {
    formPlayer.reset()
    btnSavePlayer.classList.remove('d-none')
    btnEditPlayer.classList.add('d-none')
})

btnSavePlayer.addEventListener('click', (e) => {
    e.preventDefault();
    let route = 'http://localhost:3000/api/players'
    let player = {
        FirstName: formPlayer["first-name"].value,
        LastName: formPlayer["last-name"].value,
        Level: parseInt(formPlayer["level"].value),
        Age: parseInt(formPlayer["age"].value),
        Position: formPlayer["position"].value,
        PhysicalCondition: formPlayer["condition"].value,
        Teams: []
    }
    if (formPlayer["club"].value != "") {
        team = {
            Name: formPlayer["club"].value,
        }
        player.Teams.push(team)
    }
    if (formPlayer["national"].value != "") {
        team = {
            Name: formPlayer["national"].value,
        }
        player.Teams.push(team)
    }
    console.log(player);
    fetch(route, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'Alow-Control-Allow-Origin': '*',
        },
        cache: 'default',
        body: JSON.stringify(player),
    }).then(res => res.json())
        .then((data) => {

            const divWarning = document.getElementById('warning-div')
            let message = document.getElementById('warning-message')
            validate(divWarning, message, data, formPlayer, btnSavePlayer, btnEditPlayer)

        }).catch(err => console.log('aa', err))
    //window.location.href = '/'
})

//Delete player fuction
let idDeletePlayer = ""
setTimeout(() => {
    for (let index = 0; index < showModalDelete.length; index++) {
        const element = showModalDelete[index];
        element.addEventListener('click', (e) => {
            ev = e.target
            idDeletePlayer = element.parentElement.parentElement.children[0].innerHTML

        })

    }
}, 10 * 10)

btnDeletePlayer.addEventListener('click', (e) => {
    e.preventDefault();
    let route = 'http://localhost:3000/api/players/' + idDeletePlayer
    fetch(route, {
        method: 'DELETE',
        headers: {
            'Content-Type': 'application/json',
        },
        cache: 'default',
    }).then(res => {
        console.log(res)
    }).then((data) => {
        console.log(data)
    }).catch(err => console.log(err))
    window.location.href = '/'

})

//end Delete player fuction

//**End Function for players */


//**Function for teams */
//List all teams
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


btnCancelTeam.addEventListener('click', () => {
    formTeam.reset()
    btnSaveTeam.classList.remove('d-none')
    btnEditTeam.classList.add('d-none')
    document.getElementById('country-team').disabled = false

})


let idTeam = ""
let tableBodyTeam = document.getElementById('body-teams').addEventListener('click', (e)=>{
    let ev = e.target
    let select = document.getElementById("types");
    let type
    if (ev.classList.contains('edit-team') || ev.classList.contains('delete-team')){
        document.getElementById('name-team').value = ev.parentElement.parentElement.children[1].innerHTML
        document.getElementById('country-team').value = ev.parentElement.parentElement.children[3].innerHTML
        idTeam = ev.parentElement.parentElement.children[0].innerHTML;
        type = ev.parentElement.parentElement.children[2].innerHTML.toLowerCase()
        SetTypes(select,type)
    }else if (ev.classList.contains('icon-edit-team') || ev.classList.contains('icon-delete-team')){
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

//Delete team
btnDeleteTeam.addEventListener('click', (e) => {
    e.preventDefault();
    let route = 'http://localhost:3000/api/teams/' + idTeam
    console.log(route);
    fetch(route, {
        method: 'DELETE',
        headers: {
            'Content-Type': 'application/json',
        },
        cors: 'no-cors',
        cache: 'default',
    }).then(res => {
        console.log(res)
    }).then((data) => {
        console.log(data)
    }).catch(err => console.log(err))
    window.location.href = '/'
})

//Edit team
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
        console.log(data.Message);
        const divWarning = document.getElementById('warning-div-team')
        let message = document.getElementById('warning-message-team')
        validate(divWarning, message, data, formTeam, btnSaveTeam, btnEditTeam)
    }).catch(err => console.log(err))
})

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
        cache: 'default',
        body: JSON.stringify(team),
    }).then(res => res.json()
    ).then((data) => {

        const divWarning = document.getElementById('warning-div-team')
        let message = document.getElementById('warning-message-team')
        validate(divWarning, message, data, formTeam, btnSaveTeam, btnEditTeam)
    }).catch(err => console.log(err))
    // window.location.href = '/'

})




//***End Function for teams */


//validate function

function validate(divWarning, message, data, form, btnSave, btnEdit) {

    if (data.Message != "") {
        divWarning.classList.remove('d-none')  //show warning div
        message.innerHTML = data.Message
    } else {
        divWarning.classList.add('d-none')  //show warning div
        form.reset()
        btnSave.classList.add('d-none')
        btnEdit.classList.remove('d-none')

        console.log(data.Message);
        window.location.href = '/'
    }


}
