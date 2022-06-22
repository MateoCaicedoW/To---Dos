
let indexPlayer = document.getElementById('player')
let indexTeam = document.getElementById('teams')
let indexSign = document.getElementById('sign')
let indexUnsign = document.getElementById('un-sign')
let indexTransfer = document.getElementById('transfer')

let contentTeam = document.getElementById('content-team')
let contentPlayer = document.getElementById('content-player')
let contentSign = document.getElementById('sign-player')

const signSubmit = document.getElementById('btn-submit-sign')
const unSignSubmit = document.getElementById('btn-submit-unsign')
const transferSubmit = document.getElementById('btn-submit-transfer')

const titleIndex = document.getElementById('title-index')


function indexPlayerClick(){
  indexPlayer.addEventListener('click', () => {
    contentTeam.classList.add('d-none')
    contentPlayer.classList.remove('d-none')
    contentSign.classList.add('d-none')
  
  })
}

function indexTeamClick() {
  indexTeam.addEventListener('click', () => {
    contentPlayer.classList.add('d-none')
    contentTeam.classList.remove('d-none')
    contentSign.classList.add('d-none')
  })
}

function indexSignClick() {
  indexSign.addEventListener('click', () => {
    contentPlayer.classList.add('d-none')
    contentTeam.classList.add('d-none')
    contentSign.classList.remove('d-none')
    signSubmit.classList.remove('d-none')
    unSignSubmit.classList.add('d-none')
    transferSubmit.classList.add('d-none')
    titleIndex.innerText = 'Sign Player'

  });
}

function indexUnsignClick() {
  indexUnsign.addEventListener('click', () => {
    contentPlayer.classList.add('d-none')
    contentTeam.classList.add('d-none')
    contentSign.classList.remove('d-none')
    signSubmit.classList.add('d-none')
    unSignSubmit.classList.remove('d-none')
    transferSubmit.classList.add('d-none')
    titleIndex.innerText = 'Unsign Player'
  })
}

function indexTransferClick() {
  indexTransfer.addEventListener('click', () => {
    contentPlayer.classList.add('d-none')
    contentTeam.classList.add('d-none')
    contentSign.classList.remove('d-none')
    signSubmit.classList.add('d-none')
    unSignSubmit.classList.add('d-none')
    transferSubmit.classList.remove('d-none')
    titleIndex.innerText = 'Transfer Player'
  })
}

function activeFunction() {
  var btnContainer = document.getElementById("navbarNavDropdown");
  // Get all buttons with class="btn" inside the container
  var btns = btnContainer.getElementsByClassName("nav-link");
  for (var i = 0; i < btns.length; i++) {
    btns[i].addEventListener("click", function () {
      var current = document.getElementsByClassName("active");

      current[0].className = current[0].className.replace(" active", "");
      this.className += " active";
    });
  }
}

export default {activeFunction, indexPlayerClick, indexTeamClick,indexUnsignClick, indexSignClick, indexTransferClick}


// Loop through the buttons and add the active class to the current/clicked button
