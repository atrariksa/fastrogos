const loginForm = document.getElementById("login-form");
const loginButton = document.getElementById("login-form-submit");
const loginErrorMsgHolder = document.getElementById("login-error-msg-holder")

loginButton.addEventListener("click", (e) => {
    e.preventDefault();
    const username = loginForm.username.value;
    const password = loginForm.password.value;

    fetch("http://localhost:7389/api/login", {
        method: 'POST',
        headers: {
            "Content-Type":"application/json"
        },
        redirect: 'follow',
        body: JSON.stringify({
            "username":username,
            "password":password
        })
    })
    .then(response => {
        if (response.redirected) {
            window.location.href = response.url;
        } else {
            response.json().then(body =>{
                let errorMessageP = document.createElement("p")
                errorMessageP.setAttribute("id", "login-error-msg")
                let errorMessageString = document.createTextNode(body.message)
                errorMessageP.append(errorMessageString)
                childCount = loginErrorMsgHolder.childElementCount
                if (childCount===0) {
                    loginErrorMsgHolder.append(errorMessageP)
                } else {
                    toDelete = loginErrorMsgHolder.children[0]
                    loginErrorMsgHolder.removeChild(toDelete)
                    loginErrorMsgHolder.append(errorMessageP)
                }
            })
        }
    })
    .catch(function(err) {
        console.info("Login Failed : " + err);
        let errorMessageP = document.createElement("p")
        errorMessageP.setAttribute("id", "login-error-msg")
        let errorMessageString = document.createTextNode("Invalid username and/or password")
        errorMessageP.append(errorMessageString)
        childCount = loginErrorMsgHolder.childElementCount
        if (childCount===0) {
            loginErrorMsgHolder.append(errorMessageP)
        }
    });
})