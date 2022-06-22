//validate functionA
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


export default validate