const verifyPassword = async (password: string, output: HTMLElement) => {
    const body = {
        method: "POST",
        header: {"content-type": "application/json"},
        body: JSON.stringify(password)
    }

    const request = await fetch(`${IP}/login`, body)
    
    if (request.ok) {
        const status = await request.json()
        if (status.msg === "error") {
            output.innerHTML = "ГРЕШНА ПАРОЛА"
        } else {
            loadEditor()
        }
    }

}


const validateInput = (output: HTMLElement) => {
    const password = (<HTMLInputElement>document.getElementById("password")).value
    if (!password) {
        output.innerHTML = "МОЛЯ ВЪВЕДЕТЕ ПАРОЛА"
    } else {
        verifyPassword(password, output)
    }
}


const openLogin = () => {
    const mainWindow = document.getElementById("main")!

    const div = document.createElement("div")
    div.className = "login"
    
    const input = document.createElement("input")
    input.type = "password"
    input.id = "password"
    input.maxLength = 20
    input.placeholder = "Парола"
    
    const button = document.createElement("button")
    button.innerHTML = "ВХОД"
    const p = document.createElement("p")
    button.onclick = () => { validateInput(p) }
    
    div.appendChild(input)
    div.appendChild(button)
    div.appendChild(p)
    mainWindow.appendChild(div)

    input.focus()
}
