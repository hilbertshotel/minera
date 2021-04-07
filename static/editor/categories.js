// GET
const getSubCategories = async (id) => {
    goto(`${IP}/editor/${id}`)
}


// POST
const addCategory = async () => {
    const newCategory = document.getElementById("newCategoryInput").value
    if (!newCategory) {
        output.innerHTML = "ВЪВЕДЕТЕ ИМЕ"
        return
    }

    const data = newPackage("POST", newCategory)
    const response = await fetch(`${IP}/editor/`, data)
    if (response.ok) { goto(`${IP}/editor`) }
}


// PUT
const renameCategory = async (id, oldName) => {
    const inputField = document.getElementById(id)
    const newName = inputField.value
    if (newName === oldName || !newName) { return }

    const data = newPackage("PUT", { id: id, name: newName })
    const response = await fetch(`${IP}/editor`, data)
    if (response.ok) { goto(`${IP}/editor`) }
}


// DELETE
const deleteCategory = async (id, button) => {
    if (button.id === "deleteButton") {
        const data = newPackage("DELETE", id)
        const response = await fetch(`${IP}/editor`, data)
        if (response.ok) { goto(`${IP}/editor`) }
    }

    button.id = "deleteButton"
    await new Promise((resolve) => setTimeout(resolve, 3000))
    button.id = ""
}
