
import show from "./showCont.js";
import players from "./players.js";
import teams from "./teams.js";
import movements from "./movements.js";

//Show Content-Functions
show.activeFunction()
show.indexPlayerClick()
show.indexSignClick()
show.indexTeamClick()
show.indexUnsignClick()
show.indexTransferClick()

// //Show player and teams when stay on a page
window.onload = loadInformationTabs

function loadInformationTabs() {
    let arrayTeams = []
    const rute = 'http://localhost:3000/api/teams'
    fetch(rute, {
        method: 'GET',
    })
        .then(res => {
            return res.json()
        })
        .then((data) => {
            arrayTeams = data.Data
            players.getAllPlayers(arrayTeams)
            teams.getAllTeams(arrayTeams)
            //List Physical Condition

            let Teams = document.getElementById("team-sign");
            let html = ''
            html = `
                    <option selected disabled value="">Choose...</option>
                    `
                    Teams.innerHTML += html

            for (let index = 0; index < arrayTeams.length; index++) {
                let element = arrayTeams[index];
                element = element.Name.charAt(0).toUpperCase() + element.Name.slice(1);
                html += `
                <option value="${arrayTeams[index].ID}">${element}</option>
                `
            }
            Teams.innerHTML = html
        })

    
}


//***Functions for players */
// List all players 
players.savePlayer()
players.deletePlayer()
players.resetFormPlayer()
players.closeWarning()


//**Function for teams */
//List all teams
teams.saveTeam()
teams.deleteTeam()
teams.resetFormTeam()
teams.editTeam()
teams.closeWarning()
//***End Function for teams */


//Function for movements

movements.resetFormMovement()
movements.signPlayer()
movements.transferPlayer()
movements.unsignPlayer()
movements.close()





