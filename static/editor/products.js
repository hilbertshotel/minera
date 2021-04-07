// POST
const addProduct = async (categoryId, subCategoryId) => {
    const name = document.getElementById("name").value
    if (!name) {
        output.innerHTML = "ВЪВЕДЕТЕ ИМЕ"
        return
    }
    
    const description = document.getElementById("description").value
    if (!description) {
        output.innerHTML = "ВЪВЕДЕТЕ ОПИСАНИЕ"
        return
    }
    
    const imagesTag = document.getElementById("images")
    const files = imagesTag.files
    if (files.length === 0) { 
        output.innerHTML = "ДОБАВЕТЕ ПОНЕ ЕДНО ИЗОБРАЖЕНИЕ"
        return
    }
    if (files.length > 3) {
        output.innerHTML = "ИЗБРАЛИ СТЕ ПОВЕЧЕ ОТ ТРИ ИЗОБРАЖЕНИЯ"
        return
    }

    let filenameArray = []
    const images = new FormData()
    for (const file of files) {
        images.append("files", file)
        filenameArray.push(file.name)
    }

    const newProduct = {
        id: subCategoryId,
        name: name,
        description: description,
        images: filenameArray
    }

    const data = { method: "POST", body: images }
    const response1 = await fetch(`${IP}/files`, data)
    if (response1.ok) {
        const data = newPackage("POST", newProduct)
        const response2 = await fetch(`${IP}/editor/${categoryId}/${subCategoryId}`, data)
        if (response2.ok) { goto(`${IP}/editor/${categoryId}/${subCategoryId}`) }
    }
}


// PUT
const editProduct = async (categoryId, subCategoryId, id, oldName, oldDescription) => {
    const out = document.getElementById(`out${id}`)

    const newName = document.getElementById(id).value
    if (!newName) {
        out.innerHTML = "ВЪВЕДЕТЕ ИМЕ"
        return
    }
    
    const newDescription = document.getElementById(`textarea${id}`).value
    if (!newDescription) {
        out.innerHTML = "ВЪВЕДЕТЕ ОПИСАНИЕ"
        return
    }

    if (newName === oldName && newDescription === oldDescription) { return }

    const productData = {
        id: id,
        name: newName,
        description: newDescription,
        images: []
    }

    const data = newPackage("PUT", productData)
    const response = await fetch(`${IP}/editor/${categoryId}/${subCategoryId}`, data)
    if (response.ok) { goto(`${IP}/editor/${categoryId}/${subCategoryId}`) }
}


// DELETE
const deleteProduct = async (categoryId, subCategoryId, id, button) => {
    if (button.id === "deleteButton") {
        const data = newPackage("DELETE", id)
        const response = await fetch(`${IP}/editor/${categoryId}/${subCategoryId}`, data)
        if (response.ok) { goto(`${IP}/editor/${categoryId}/${subCategoryId}`) }
    }

    button.id = "deleteButton"
    await new Promise((resolve) => setTimeout(resolve, 3000))
    button.id = ""
}
