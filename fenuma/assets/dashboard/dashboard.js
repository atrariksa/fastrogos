const mainBody = document.getElementById("id-main-body");
const manageUser = document.getElementById("id-manage-user");

manageUser.addEventListener("click", (e) => {
    e.preventDefault();
    sibling = manageUser.nextSibling;
    if (sibling === null) {
        let createUser = document.createElement("li")
        createUser.setAttribute("id", "id-manage-user-create")
        createUser.setAttribute("class", "material-icons")
        createUser.innerText = "add"
        let editUser = document.createElement("li")
        editUser.setAttribute("id", "id-manage-user-edit")
        editUser.setAttribute("class", "material-icons")
        editUser.innerText = "edit"
        let deleteUser = document.createElement("li")
        deleteUser.setAttribute("id", "id-manage-user-delete")
        deleteUser.setAttribute("class", "material-icons")
        deleteUser.innerText = "remove"
    
        let expandedList = document.createElement("ol")
        expandedList.appendChild(createUser)
        expandedList.appendChild(editUser)
        expandedList.appendChild(deleteUser)
    
        manageUser.parentNode.insertBefore(expandedList, manageUser.nextSibling)
        regCreateUserEventListener()
    } else {
        manageUser.parentNode.removeChild(sibling)
    }

    
})

function regCreateUserEventListener() {
    const createUserEventListener = document.getElementById("id-manage-user-create");
    createUserEventListener.addEventListener("click", (e) => {
        e.preventDefault();
        while (mainBody.firstChild !== null) {
            mainBody.removeChild(mainBody.lastChild)
        }
        let createUserComponent = document.createElement("div")
        createUserComponent.setAttribute("id", "id-main-manage-user-create")
        createUserComponent.setAttribute("class", "material-icons")
        createUserComponent.innerHTML = 
        `<form id="id-form-create-user">
            <label for="id-input-username">Username : </label>
            <input type="text" name="input-username" id="id-input-username">

            <label for="id-input-email">Email : </label>
            <input type="email" name="input-email" id="id-input-email">

            <label for="id-input-password">Password : </label>
            <input type="password" name="input-password" id="id-input-password">

            <p>Role : </p>
            <label for="id-input-role-admin">ADMIN</label>
            <input type="radio" name="input-role" id="id-input-role-admin" value="ADMIN">

            <label for="id-input-role-editor">EDITOR</label>
            <input type="radio" name="input-role" id="id-input-role-editor" value="EDITOR">

            <label for="id-input-role-author">AUTHOR</label>
            <input type="radio" name="input-role" id="id-input-role-author" value="AUTHOR">

            <a href="#" class="form-button">Clear</a>
            <input class="form-button" type="submit" value="Submit" id="id-form-create-user-submit" />
        </form>
        `
        mainBody.append(createUserComponent)
    })
}



