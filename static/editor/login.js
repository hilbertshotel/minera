const output = document.getElementById("output")
document.getElementById("username").focus()


const validate_credentials = async () => {

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

    const url = `${IP}/authentication`
    const user_data = { username: username, password: password }
    const data = {
        method: "POST",
        header: { "content-type": "application/json" },
        credentials: "include",
        body: JSON.stringify(user_data)
    }

    const request = await fetch(url, data)
    if (request.ok) {
        const status = await request.json()
        if (status !== "ok") {
            output.innerHTML = status
            return
        }
    }

    window.location.replace(`${IP}/editor`)
}
