const validateCredentials = async () => {

    const username = document.getElementById("username").value
    if (!username) {
        output.innerHTML = "МОЛЯ ВЪВЕДЕТЕ ПОТРЕБИТЕЛ"
        return  
    }

    const password = document.getElementById("password").value
    if (!password) {
        output.innerHTML = "МОЛЯ ВЪВЕДЕТЕ ПАРОЛА"
        return
    }

    const userData = { username: username, password: password }
    const data = newPackage("POST", userData)
    const response = await fetch(`${IP}/authentication`, data)
    if (response.ok) {
        const status = await response.json()
        if (status !== "ok") {
            output.innerHTML = status
            return
        }
    }

    document.location.replace(`${IP}/editor/`)
}