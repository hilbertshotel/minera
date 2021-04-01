const output = document.getElementById("output")
document.getElementById("newCategoryInput").focus()


// GET
const getSubCategories = (id) => {
    console.log(id)
}


// POST
const addCategory = async () => {
    const newCategory = document.getElementById("newCategoryInput").value
    if (!newCategory) {
        output.innerHTML = "ВЪВЕДЕТЕ ИМЕ"
        return
    }

    const data = newPackage("POST", newCategory)
    const request = await fetch(`${IP}/editor/`, data)
    if (request.ok) { location.reload() }
}


// PUT
const renameCategory = async (id, oldName) => {
    const inputField = document.getElementById(id)
    const newName = inputField.value
    if (newName === oldName || !newName) { return }

    const data = newPackage("PUT", { id: id, name: newName })
    const request = await fetch(`${IP}/editor`, data)
    if (request.ok) { location.reload() }
}


// DELETE
const deleteCategory = async (id, button) => {
    if (button.id === "deleteButton") {
        const data = newPackage("DELETE", id)
        const request = await fetch(`${IP}/editor`, data)
        if (request.ok) { location.reload() }
    }

    button.id = "deleteButton"
    await new Promise((resolve) => setTimeout(resolve, 3000))
    button.id = ""
}
