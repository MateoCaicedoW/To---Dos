window.addEventListener('click',getAllPlayers)

let array = []
function getAllPlayers() {
    const ruta = 'http://localhost:3000/players'
    fetch(ruta,{
        method: 'GET',
         
        cache: 'default'
    })
    .then(res=>{
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
                
                if(element.Type == "club"){
                    club = `${element.Name}`
                }else{
                    team = `${element.Name}`
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
                    <button type="button" class="btn btn-success ms-2"><i class="fas fa-edit"></i></button>
                    <button type="button" class="btn btn-danger ms-2" data-bs-toggle="modal" data-bs-target="#exampleModal"><i class="far fa-trash-alt "></i></button>
                   
                </td>
                
            </tr>
            `
        }
        document.getElementById('body-players').innerHTML = html
       
        
      })


      array = getAllTeams()
      console.log(array);
      for (let index = 0; index < array.length; index++) {
           const team = data.Data[index];
            let select = document.getElementById("team");
            let option = document.createElement("option");
            option.text = team.Name;
            option.value = team.ID;
            select.add(option);

          
      }

        
}


function getAllTeams() {
    let arrayTeams = []

    const ruta2 = 'http://localhost:3000/teams'
    fetch(ruta2,{
        method: 'GET',
         
        cache: 'default'
    })
    .then(res=>{
        return res.json()
    })
    .then((data) => {
        console.log(data)
        // let html2 = ''
     
        //list of teams
        for (let index = 0; index < data.Data.length; index++) {
            
            let teams = { ID: data.Data[index].ID, Name : data.Data[index].Name, Type:data.Data[index].Type, Country:data.Data[index].Country} ;
            arrayTeams.push(teams)
            // html2 += `
            // <tr>
            //     <th scope="row">${team.ID}</th>
            //     <td>${team.Name}</td>
            //     <td>${team.Type}</td>
            //     <td>${team.Country}</td>
                
            //     <td class="d-flex py-5 py-xl-4 py-md-5 justify-content-center align-items-center">
                
            //         <button type="button" class="btn btn-primary "><i class="far fa-eye"></i></button>
            //         <button type="button" class="btn btn-success ms-2"><i class="fas fa-edit"></i></button>
            //         <button type="button" class="btn btn-danger ms-2" data-bs-toggle="modal" data-bs-target="#exampleModal"><i class="far fa-trash-alt "></i></button>
                   
            //     </td>
                
            // </tr>
            // `
           
        }
    
        // document.getElementById('body-teams').innerHTML = html2
       
        
        
      })

      return arrayTeams
}





(() => {
    'use strict'
  
    // Fetch all the forms we want to apply custom Bootstrap validation styles to
    const forms = document.querySelectorAll('.needs-validation')
  
    // Loop over them and prevent submission
    Array.from(forms).forEach(form => {
      form.addEventListener('submit', event => {
        if (!form.checkValidity()) {
          event.preventDefault()
          event.stopPropagation()
        }
  
        form.classList.add('was-validated')
      }, false)
    })
  })()