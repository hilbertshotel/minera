// DATA
interface Item {
    Name: string,
    Description: string,
    Images: string[]
}

interface Category {
    Id: number,
    Name: string
}


// LOGIN
const verifyUser = async (userData: Object, output: HTMLElement) => {
    const data = newPackage("POST", userData)
    const request = await fetch(`${IP}/VerifyUser`, data)
    
    if (request.ok) {
        const status = await request.json()
        if (status !== "ok") {
            output.innerHTML = status
            getInputWithId("pass").value = ""
            return
        }
        // THIS IS A SECURITY RISK WHICH NEEDS TO BE FIXED
        addScript("scripts/editor.js")
    }
}


const validateInput = () => {
    const output = getById("output")
    output.innerHTML = ""
    const username = getInputWithId("user").value
    const password = getInputWithId("pass").value
    if (!username) { output.innerHTML = "МОЛЯ ВЪВЕДЕТЕ ПОТРЕБИТЕЛ"; return }
    if (!password) { output.innerHTML = "МОЛЯ ВЪВЕДЕТЕ ПАРОЛА"; return }
    const data = { "Username": username, "Password": password }
    verifyUser(data, output)
}


const openLogin = () => {
    const content = getById("content")

    const userInput = `<input id="user" maxlength="20" placeholder="Потребител">`
    const passInput = `<input type="password" id="pass" maxlength="20" placeholder="Парола">`
    const button = `<button onclick="validateInput()">ВХОД</button>`
    const output = `<p id="output"></p>`

    content.innerHTML += userInput
    content.innerHTML += passInput
    content.innerHTML += button
    content.innerHTML += output

    getById("user").focus()
}


// MAIN
const IP = "http://127.0.0.1:5252"
openLogin()
