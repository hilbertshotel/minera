// GET
const getProducts = async (categoryId, id) => {
    goto(`${IP}/editor/${categoryId}/${id}`)
}


// POST
const addSubCategory = async (categoryId) => {
    const newSubCategory = document.getElementById("newSubCategoryInput").value
    if (!newSubCategory) {
        output.innerHTML = "ВЪВЕДЕТЕ ИМЕ"
        return
    }

    const data = newPackage("POST", newSubCategory)
    const response = await fetch(`${IP}/editor/${categoryId}`, data)
    if (response.ok) { goto(`${IP}/editor/${categoryId}`) }
}


// PUT
const renameSubCategory = async (categoryId, id, oldName) => {
    const inputField = document.getElementById(id)
    const newName = inputField.value
    if (newName === oldName || !newName) { return }

    const data = newPackage("PUT", { id: id, name: newName })
    const response = await fetch(`${IP}/editor/${categoryId}`, data)
    if (response.ok) { goto(`${IP}/editor/${categoryId}`) }
}


// DELETE
const deleteSubCategory = async (categoryId, id, button) => {
    if (button.id === "deleteButton") {
        const data = newPackage("DELETE", id)
        const response = await fetch(`${IP}/editor/${categoryId}`, data)
        if (response.ok) { goto(`${IP}/editor/${categoryId}`) }
    }

    button.id = "deleteButton"
    await new Promise((resolve) => setTimeout(resolve, 3000))
    button.id = ""
}
