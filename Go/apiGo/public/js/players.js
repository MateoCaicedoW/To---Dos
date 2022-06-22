import validate from "./validate.js";

//variables for player
let formPlayer = document.getElementById("form-player");
let btnEditPlayer = document.getElementById('btn-edit-player')
let btnSavePlayer = document.getElementById('btn-create-player')
let btnCancelPlayer = document.getElementById('btn-cancel-player')
let btnDeletePlayer = document.getElementById('btn-delete-player-modal')
let showModalDelete = document.getElementsByClassName('delete-show-modal-player')
let btnClose =document.getElementById('close-button-player');
const divWarning = document.getElementById('warning-div')
//editPlayer Function

function getAllPlayers(data) {

    const ruta = 'http://localhost:3000/api/players'
    fetch(ruta, {
        method: 'GET'
    })
        .then(res => {
            return res.json()
        })
        .then((data) => {
            let html = ''
            let html2 = ''
            let dropPlayer = document.getElementById("player-sign");
            //list of players
            html2 = `
                        <option selected disabled value="">Choose...</option>
                        `
                        dropPlayer.innerHTML += html2
            for (let index = 0; index < data.Data.length; index++) {
                const player = data.Data[index];
                

                    let element = player.FirstName + ' ' + player.LastName;
                    html2 += `
                    <option value="${data.Data[index].ID}">${element}</option>
                    `
                
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
            dropPlayer.innerHTML = html2

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
                    idPlayer = element.parentElement.parentElement.children[0].innerHTML
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
        let team = {
            Name: formPlayer["club"].value,
        }
        player.Teams.push(team)
    }
    if (formPlayer["national"].value != "") {
        let team = {
            Name: formPlayer["national"].value,
        }
        player.Teams.push(team)
    }

    console.log(player);

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
            console.log(data);
            let message = document.getElementById('warning-message')
            validate(divWarning, message, data, formPlayer, btnSavePlayer, btnEditPlayer)
        }).catch(err => console.log(err))

}

function resetFormPlayer(){
    btnCancelPlayer.addEventListener('click', () => {
        formPlayer.reset()
        btnSavePlayer.classList.remove('d-none')
        btnEditPlayer.classList.add('d-none')
    })
}

function savePlayer(){
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
            body: JSON.stringify(player),
        }).then(res => res.json())
            .then((data) => {
    
                
                let message = document.getElementById('warning-message')
                validate(divWarning, message, data, formPlayer, btnSavePlayer, btnEditPlayer)
    
            }).catch(err => console.log('aa', err))
        //window.location.href = '/'
    })
}


//Delete player fuction
let idDeletePlayer = ""
setTimeout(() => {
    for (let index = 0; index < showModalDelete.length; index++) {
        const element = showModalDelete[index];
        element.addEventListener('click', () => {
            idDeletePlayer = element.parentElement.parentElement.children[0].innerHTML

        })

    }
}, 10 * 10)

function deletePlayer(){
    btnDeletePlayer.addEventListener('click', (e) => {
        e.preventDefault();
        let route = 'http://localhost:3000/api/players/' + idDeletePlayer
        fetch(route, {
            method: 'DELETE',
            headers: {
                'Content-Type': 'application/json',
            },
        }).then(res => {
            console.log(res)
        }).then((data) => {
            console.log(data)
        }).catch(err => console.log(err))
        window.location.href = '/'
    
    })
}

function closeWarning(){
    btnClose.addEventListener('click', (e) => {
        e.preventDefault();
        divWarning.classList.add('d-none')
    })

}

//end Delete player fuction

//**End Function for players */

export default {getAllPlayers, resetFormPlayer, deletePlayer, savePlayer, closeWarning}