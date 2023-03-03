var loginForm = document.getElementById("login-form")
var sumbitButton = document.getElementById("login-submit-button")

sumbitButton.addEventListener("click", function(event){
    event.preventDefault()
    
    var json = formToJson(loginForm)
    $.ajax({
        type: "POST",
        url: "/api/registration",
        data: json,
        success: function(result){

        },
        dataType: "json",
        contentType: "application/json; charset=utf-8"
    })

})