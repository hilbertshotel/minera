const ADDRESS = "http://127.0.0.1:5252"


const getCategories = async () => {
    const url = `${ADDRESS}/Catalog/Category`
    const response = await fetch(url)
    if (response.ok) {
        // const data = await response.json()
        // console.log(data)
        console.log("ok")
    }
}


const getItems = async (id) => {
    const url = `${ADDRESS}/Catalog/Item/${id}`
    const response = await fetch(url)
    if (response.ok) {
        // const data = await response.json()
        // console.log(data)
        console.log("ok")
    }
}
