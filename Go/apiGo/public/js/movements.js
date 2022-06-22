//variables for movements
let formMovement = document.getElementById("form-sign")
let btnSign = document.getElementById("btn-submit-sign")
let btnUnsign = document.getElementById("btn-submit-unsign")
let btnTransfer = document.getElementById("btn-submit-transfer")
let btnCancelMovement = document.getElementById("btn-cancel-sign")
const divSuccess = document.getElementById('success-div-sign')
const divWarning = document.getElementById('warning-div-sign')
const closeButtonError = document.getElementById('close-button-sign-error')
const closeButtonSuccess = document.getElementById('close-button-sign-success')
//btnCancel
function resetFormMovement() {

    btnCancelMovement.addEventListener('click', () => {
        formMovement.reset()
    })
}

//btnSignPlayer
function signPlayer() {
    btnSign.addEventListener('click', (e) => {
        e.preventDefault()

        let sign = {
            TeamID: formMovement["team-sign"].value,
            PlayerID: formMovement["player-sign"].value,
        }
        console.log(sign);
        const rute = 'http://localhost:3000/api/movements/sign-player'

        fetch(rute, {
            method: 'POST',
            body: JSON.stringify(sign),
        }).then((response) => response.json())
            .then((data) => {
                validateMovements(data)
            })

    })
}


//btnUnsignPlayer
function unsignPlayer() {
    btnUnsign.addEventListener('click', (e) => {
        e.preventDefault()

        let unsign = {
            TeamID: formMovement["team-sign"].value,
            PlayerID: formMovement["player-sign"].value,
        }

        const rute = 'http://localhost:3000/api/movements/unsign-player'

        fetch(rute, {
            method: 'DELETE',
            body: JSON.stringify(unsign),
        }).then((response) => response.json())
            .then((data) => {
                validateMovements(data)
            })
    })
}


function transferPlayer() {
    btnTransfer.addEventListener('click', (e) => {
        e.preventDefault()
        let transfer = {
            TeamID: formMovement["team-sign"].value,
            PlayerID: formMovement["player-sign"].value,
        }

        const rute = 'http://localhost:3000/api/movements/transfer-player'

        fetch(rute, {
            method: 'PUT',
            body: JSON.stringify(transfer),
        }).then((response) => response.json())
            .then((data) => {
                validateMovements(data)
            })

    });
}


//validate movements
function validateMovements(data) {

    let messageWarning = document.getElementById('warning-message-sign')
    let messagesuccess = document.getElementById('success-message-sign')

    if (data.Status != 200) {
        divWarning.classList.remove('d-none')
        divSuccess.classList.add('d-none')   //show warning div
        messageWarning.innerHTML = data.Message

    } else {
        divSuccess.classList.remove('d-none')
        divWarning.classList.add('d-none')  //show warning div
        messagesuccess.innerHTML = data.Message
    }
}
function close(){
    closeButtonError.addEventListener('click', (e) => {
        e.preventDefault();
        divWarning.classList.add('d-none')
    })

    closeButtonSuccess.addEventListener('click', (e) => {
        e.preventDefault();
        divSuccess.classList.add('d-none')
    })

}


export default {resetFormMovement,signPlayer,unsignPlayer,transferPlayer,close}
