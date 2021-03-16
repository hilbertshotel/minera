// LOGIN
const verifyUser = async (userData: Object, output: HTMLElement) => {
    const data = newPackage("POST", userData)
    const request = await fetch(`${IP}/VerifyUser`, data)
    
    if (request.ok) {
        const status = await request.json()
        if (status !== "ok") {
            output.innerHTML = status
            return
        }
        // THIS IS A SECURITY RISK WHICH NEEDS TO BE FIXED
        addScript("scripts/editor.js")
    }
}


const validateInput = (output: HTMLElement) => {
    output.innerHTML = ""
    const username = getInputValue("username")
    const password = getInputValue("password")
    if (!username) { output.innerHTML = "МОЛЯ ВЪВЕДЕТЕ ПОТРЕБИТЕЛ"; return }
    if (!password) { output.innerHTML = "МОЛЯ ВЪВЕДЕТЕ ПАРОЛА"; return }
    const data = { "Username": username, "Password": password }
    verifyUser(data, output)
}


const openLogin = () => {
    const content = getById("content")

    const username = newInput()
    username.id = "username"
    username.maxLength = 20
    username.placeholder = "Потребител"

    const password = newInput()
    password.type = "password"
    password.id = "password"
    password.maxLength = 20
    password.placeholder = "Парола"

    const enter = newButton("ВХОД")
    const output = outputField()
    enter.onclick = () => { validateInput(output) }

    const children = [username, password, enter, output]
    appendChildren(children, content)
    username.focus()
}


// MAIN
const IP = "http://127.0.0.1:5252"
openLogin()
