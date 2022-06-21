let btn_player =  document.getElementById('player')
let btn_team = document.getElementById('teams')
let contentTeam = document.getElementById('content-team')
let contentPlayer = document.getElementById('content-player')

btn_player.addEventListener('click', () => {
    contentTeam.classList.add('d-none')
    contentPlayer.classList.remove('d-none')
})
btn_team.addEventListener('click', () => {
    contentPlayer.classList.add('d-none')
    contentTeam.classList.remove('d-none')
})




//active
var btnContainer = document.getElementById("navbarNavDropdown");

// Get all buttons with class="btn" inside the container
var btns = btnContainer.getElementsByClassName("nav-link");


// Loop through the buttons and add the active class to the current/clicked button
for (var i = 0; i < btns.length; i++) {
  btns[i].addEventListener("click", function() {
    var current = document.getElementsByClassName("active");
    
    current[0].className = current[0].className.replace(" active", "");
    this.className += " active";
  });
}