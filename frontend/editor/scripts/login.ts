// LOGIN
const verifyUser = async (url: string, userData: Object, output: HTMLElement) => {
    const data = newPackage("POST", userData)
    const request = await fetch(url, data)
    
    if (request.ok) {
        const status = await request.json()
        if (status !== "ok") {
            output.innerHTML = status
            return
        }
        addScript("scripts/editor.js")
    }
}


const validateInput = (output: HTMLElement) => {
    output.innerHTML = ""

    const username = getInputValue("username")
    if (!username) { output.innerHTML = "МОЛЯ ВЪВЕДЕТЕ ПОТРЕБИТЕЛ"; return }
    const password = getInputValue("password")
    if (!password) { output.innerHTML = "МОЛЯ ВЪВЕДЕТЕ ПАРОЛА"; return }

    const data = { "Username": username, "Password": password }
    verifyUser(`${IP}/Authentication`, data, output)
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

    appendChildren([username, password, enter, output], content)
    username.focus()
}


// MAIN
const IP = "http://127.0.0.1:5252"
openLogin()
