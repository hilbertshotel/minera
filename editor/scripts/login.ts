const verifyPassword = async (password: string, output: HTMLElement) => {
    const body = {
        method: "POST",
        header: {"content-type": "application/json"},
        body: JSON.stringify(password)
    }

    const request = await fetch(`${IP}/login`, body)
    
    if (request.ok) {
        const status = await request.json()
        if (status.msg !== "ok") {
            output.innerHTML = status.msg
            const input = <HTMLInputElement>document.getElementById("password")!
            input.value = ""
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
    const content = document.getElementById("content")!

    const input = document.createElement("input")
    input.type = "password"
    input.id = "password"
    input.maxLength = 20
    input.placeholder = "Парола"
    
    const button = document.createElement("button")
    button.innerHTML = "ВХОД"
    const output = document.createElement("p")
    output.id = "output"
    button.onclick = () => { validateInput(output) }
    
    content.appendChild(input)
    content.appendChild(button)
    content.appendChild(output)

    input.focus()
}
