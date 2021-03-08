const verifyPassword = async (password: string) => {
    
}


const validateInput = (output: HTMLElement) => {
    const password = (<HTMLInputElement>document.getElementById("password")).value
    if (!password) {
        output.innerHTML = "НЕ СТЕ ВЪВЕЛИ ПАРОЛА"
    } else {
        verifyPassword(password)
    }
}


const addLoginField = (mainWindow: HTMLElement) => {
    const div = document.createElement("div")
    div.className = "login"
    const input = document.createElement("input")
    input.type = "password"
    input.id = "password"
    input.placeholder = "Парола"
    const button = document.createElement("button")
    button.innerHTML = "ВХОД"
    const p = document.createElement("p")
    p.id = "output"
    button.onclick = () => { validateInput(p) }
    div.appendChild(input)
    div.appendChild(button)
    div.appendChild(p)
    mainWindow.appendChild(div)
    input.focus()
}


const openLogin = () => {
    const mainWindow = document.getElementById("main")!
    addLoginField(mainWindow)
}
