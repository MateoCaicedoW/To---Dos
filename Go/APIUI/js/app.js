


//Show player and teams when stay on page
let arrayTeams = []
const ruta2 = 'http://localhost:3000/teams'
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



// List all players 
function getAllPlayers(data) {
    
    const ruta = 'http://localhost:3000/players'
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
                
                    <button type="button" class="btn btn-primary "><i class="far fa-eye"></i></button>
                    <button  type="button" class="btn btn-success ms-2 edit-player"><i class="fas fa-edit"></i></button>
                    <button type="button" class="btn btn-danger ms-2" data-bs-toggle="modal" data-bs-target="#exampleModal"><i class="far fa-trash-alt "></i></button>
                   
                </td>
                
            </tr>
            `
            }
            document.getElementById('body-players').innerHTML = html


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
            <option value="${team.ID}">${name}</option>
            `
            clubSelect.innerHTML += html
        }else{
            html = `
            <option value="${team.ID}">${name}</option>
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
    const ruta2 = 'http://localhost:3000/positions'
    fetch(ruta2, {
        method: 'GET',
    }).then(res =>  res.json()
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
   
    const ruta3 = 'http://localhost:3000/conditions'
    fetch(ruta3, {
        method: 'GET',
    }).then(res =>  res.json()
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

                <button type="button" class="btn btn-primary "><i class="far fa-eye"></i></button>
                <button type="button" class="btn btn-success ms-2"><i class="fas fa-edit"></i></button>
                <button type="button" class="btn btn-danger ms-2" data-bs-toggle="modal" data-bs-target="#modal-Team"><i class="far fa-trash-alt "></i></button>

            </td>

        </tr>
        `


    }
    document.getElementById('body-teams').innerHTML = html2
           
        
}



//Show One Player
//let form = document.getElementById('form-player')
let idPlayer =""
setTimeout(() => {
    let btnEditPlayer = document.getElementsByClassName('edit-player')
    if (btnEditPlayer) {
        for (let index = 0; index < btnEditPlayer.length; index++) {
            const element = btnEditPlayer[index];
            element.addEventListener('click', () => {
                idPlayer = element.parentElement.parentElement.children[0].innerHTML
            
                document.getElementById('first-name').value = element.parentElement.parentElement.children[1].innerHTML
                document.getElementById('last-name').value = element.parentElement.parentElement.children[2].innerHTML
                document.getElementById('level').value = element.parentElement.parentElement.children[3].innerHTML
                document.getElementById('age').value = element.parentElement.parentElement.children[4].innerHTML

                for(let item= 0;item< document.getElementById('position').length;item++){
                    const i = document.getElementById('position').options[item]
                        if(i.value.toUpperCase() == element.parentElement.parentElement.children[5].innerHTML.replace(" ","").toUpperCase()){
                            i.selected = true
                        }
                }

                for(let item= 0;item< document.getElementById('condition').length;item++){
                    const i = document.getElementById('condition').options[item]
                        if(i.value.toUpperCase() == element.parentElement.parentElement.children[6].innerHTML.toUpperCase()){
                            i.selected = true
                        }
                }


                for(let item= 0;item< document.getElementById('club').length;item++){
                    const i = document.getElementById('club').options[item]
                    const club = element.parentElement.parentElement.children[7].innerHTML
                    if (club != "") {
                        if (i.text.toUpperCase() == club.toUpperCase())  {
                            i.selected = true
                        }
                    }else{
                        document.getElementById('club').options[0].selected = true
                    }
                }

                for(let item= 0;item< document.getElementById('national').length;item++){
                    const i = document.getElementById('national').options[item]
                    const national = element.parentElement.parentElement.children[8].innerHTML
                    if (national != "") {
                        if (i.text.toUpperCase() == national.toUpperCase())  {
                            i.selected = true
                        }
                    }else{
                        document.getElementById('national').options[0].selected = true
                    }
                }

                window.location.href = '#form-player'
            })
           
            
                //console.log(form)
            
        }
    }
    
},1000)


//validate form and edir
let btnCreatePlayer = document.getElementById("btn-create-player")
btnCreatePlayer.addEventListener('click', (e) => {
    e.preventDefault();
    let form = document.getElementById("form-player");
    if (!form.checkValidity()) {
        e.preventDefault()
        e.stopPropagation()
    }

    form.classList.add('was-validated')

    let ruta = 'http://localhost:3000/players/' + idPlayer
    console.log(ruta);
    fetch(ruta, {
        method: 'PUT',
        mode: 'cors',
        body: JSON.stringify({
            FirstName: document.getElementById('first-name').value,
            LastName: document.getElementById('last-name').value,
            Level: document.getElementById('level').value,
            Age: document.getElementById('age').value,
            Position: document.getElementById('position').value,
            PhysicalCondition: document.getElementById('condition').value,
        }), 

    }).then(res => res.json())
    .then((data) => {
        console.log(data)
        window.location.href = '#list-players'
    })


})
