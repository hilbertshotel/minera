const IP = "http://127.0.0.1:5252"
const output = document.getElementById("output")

const newPackage = (method, data) => {
    return {
        header: {"content-type": "application/json"},
        method: method,
        credentials: "include",
        body: JSON.stringify(data)
    }
}

const goto = (url) => {
    window.location.href = url
}
