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
    
    const files = document.getElementById("images").files
    if (files.length === 0) { 
        output.innerHTML = "ДОБАВЕТЕ ПОНЕ ЕДНО ИЗОБРАЖЕНИЕ"
        return
    }
    if (files.length > 3) {
        output.innerHTML = "ИЗБРАЛИ СТЕ ПОВЕЧЕ ОТ ТРИ ИЗОБРАЖЕНИЯ"
        return
    }

    const newProduct = {
        id: subCategoryId,
        name: name,
        description: description,
        images: []
    }

    const images = new FormData()
    for (const file of files) {
        images.append("files", file)
        newProduct.images.push(file.name)
    }

    const imagesPackage = { method: "POST", body: images }
    const response1 = await fetch(`${IP}/files`, imagesPackage)
    if (response1.ok) {
        sendProductData("POST", newProduct, categoryId, subCategoryId)
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

    const files = document.getElementById(`images${id}`).files
    if (newName === oldName && newDescription === oldDescription && files.length === 0) { return }

    if (files.length > 3) {
        out.innerHTML = "ИЗБРАЛИ СТЕ ПОВЕЧЕ ОТ ТРИ ИЗОБРАЖЕНИЯ"
        return
    }
    
    const productData = {
        id: id,
        name: newName,
        description: newDescription,
        images: []
    }

    if (files.length > 0) {
        const images = new FormData()

        for (const file of files) {
            images.append("files", file)
            productData.images.push(file.name)
        }

        const imagesPackage = { method: "POST", body: images }
        const response = await fetch(`${IP}/files`, imagesPackage)
        if (response.ok) {
            sendProductData("PUT", productData, categoryId, subCategoryId)
        }
        return
    }
    
    sendProductData("PUT", productData, categoryId, subCategoryId)
}


// DELETE
const deleteProduct = async (categoryId, subCategoryId, id, button) => {
    if (button.id === "deleteButton") {
        sendProductData("DELETE", id, categoryId, subCategoryId)
    }

    button.id = "deleteButton"
    await new Promise((resolve) => setTimeout(resolve, 3000))
    button.id = ""
}
