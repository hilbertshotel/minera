const IP = "http://127.0.0.1:8000"

const newPackage = (method, data) => {
    return {
        header: {"content-type": "application/json"},
        method: method,
        credentials: "include",
        body: JSON.stringify(data)
    }
}