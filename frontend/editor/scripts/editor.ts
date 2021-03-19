// ITEM METHODS
const deleteItem = async (id: number, category: Category) => {
    const data = newPackage("DELETE", "")
    const request = await fetch(`${IP}/Editor/Items/${id}`, data)
    if (request.ok) { getItems(category) }
}


const postItem = async (category: Category) => {
    const name = getInputValue("name")
    const description = getInputValue("description")
    const imagesTag = <HTMLInputElement>getById("images")!
    const output = getById("output")

    if (!name) { output.innerHTML = "ВЪВЕДЕТЕ ИМЕ"; return }
    if (!description) {
        output.innerHTML = "ВЪВЕДЕТЕ ОПИСАНИЕ"
        output.focus()
        return
    }

    let filenameArray: string[] = []
    const images = new FormData()
    const files = imagesTag.files

    if (!files) { return }
    if (files.length === 0) { 
        output.innerHTML = "ДОБАВЕТЕ ПОНЕ ЕДНА СНИМКА"
        output.focus()
        return
    }
    if (files.length > 3) {
        output.innerHTML = "ИЗБРАЛИ СТЕ ПОВЕЧЕ ОТ ТРИ ИЗОБРАЖЕНИЯ"
        output.focus()
        return
    }

    for (const file of files) {
        images.append("files", file)
        filenameArray.push(file.name)
    }

    const newItem = {
        id: category.Id,
        name: name,
        description: description,
        images: filenameArray
    }

    const data = { method: "POST", body: images }
    const request1 = await fetch(`${IP}/FileTransfer`, data)
    if (request1.ok) {
        const data = newPackage("POST", newItem)
        const request2 = await fetch(`${IP}/Editor/Items/0`, data)
        if (request2.ok) { getItems(category) }
    }
}


const getItems = async (category: Category) => {
    const request = await fetch(`${IP}/Editor/Items/${category.Id}`)
    if (request.ok) {
        const items = await request.json()
        loadItems(category, items)
    }
}


// ITEM PAGE CONSTRUCTION
const deleteItemButton = async (id: number, button: HTMLElement, category: Category) => {
    if (button.id === "deleteButton") { deleteItem(id, category) }
    button.id = "deleteButton"
    await new Promise((resolve) => setTimeout(resolve, 3000))
    button.id = ""
}

const insertNewItemSection = (category: Category, content: HTMLElement) => {
    const header = newElement("h1")
    header.innerHTML = "Добавяне на нов артикул"

    const name = newInput()
    name.id = "name"
    name.maxLength = 50
    name.placeholder = "Име на артикула"

    const description = newTextArea()
    description.id = "description"
    description.maxLength = 300
    description.placeholder = "Описание на артикула"

    const images = newInput()
    images.type = "file"
    images.id = "images"
    images.setAttribute("multiple", "")
    images.accept = "image/*"
    
    const addButton = newButton("ДОБАВИ")
    addButton.onclick = () => { postItem(category) }

    appendChildren([header, name, br(), description, br(), images, br(), addButton], content)
    name.focus()
}


const insertItem = (item: Item, content: HTMLElement, category: Category) => {
    const nameTag = newButton(item.Name)
    nameTag.className = "itemName"

    const deleteButton = newButton("Изтрий")
    deleteButton.onclick = () => { deleteItemButton(item.Id, deleteButton, category) } 

    appendChildren([nameTag, deleteButton, br()], content)
}


const loadItems = (category: Category, items: Item[]) => {
    const content = getById("content")
    clear(content)
    getById("h").innerHTML = category.Name

    if (items !== null) {
        for (const item of items) {
            insertItem(item, content, category)
        }
    }

    insertNewItemSection(category, content)

    const output = outputField()
    const backButton = newButton("ВЪРНИ СЕ")
    backButton.id = "backButton"
    backButton.onclick = getCategories

    appendChildren([output, backButton], content)
}


// CATEGORY METHODS
const deleteCategory = async (id: number) => {
    const data = newPackage("DELETE", id)
    const request = await fetch(`${IP}/Editor/Categories`, data)
    if (request.ok) { getCategories() }
}


const putCategory = async (category: Category, nameTag: HTMLInputElement) => {
    const newName = nameTag.value
    if (newName === category.Name || !newName) { return }

    const info = { id: category.Id, newName: newName }
    const data = newPackage("PUT", info)
    const request = await fetch(`${IP}/Editor/Categories`, data)
    if (request.ok) { getCategories() }
}


const postCategory = async () => {
    const newCategory = getInputValue("newCategory")
    if (!newCategory) {
        const output = getById("output")
        output.innerHTML = "ВЪВЕДЕТЕ ИМЕ"
        output.focus()
        return
    }
    
    const data = newPackage("POST", newCategory)
    const request = await fetch(`${IP}/Editor/Categories`, data)
    if (request.ok) { getCategories() }
}

const getCategories = async () => {
    getById("h").innerHTML = "РЕДАКТОР"
    const content = getById("content")
    clear(content)
    const request = await fetch(`${IP}/Editor/Categories`)
    if (request.ok) {
        const categories = await request.json()
        loadCategories(categories, content)
    }
}


// CATEGORY PAGE CONSTRUCTION
const deleteCategoryButton = async (id: number, button: HTMLElement) => {
    if (button.id === "deleteButton") { deleteCategory(id) }
    button.id = "deleteButton"
    await new Promise((resolve) => setTimeout(resolve, 3000))
    button.id = ""
}


const insertCategory = (category: Category, content: HTMLElement) => {
    const nameTag = newInput()
    nameTag.className = "categoryName"
    nameTag.value = category.Name
    nameTag.placeholder = category.Name

    const itemsButton = newButton("Преглед")
    itemsButton.onclick = () => { getItems(category) }
    
    const editButton = newButton("Ново име")
    editButton.onclick = () => { putCategory(category, nameTag) } 
    
    const deleteButton = newButton("Изтрий")
    deleteButton.onclick = () => { deleteCategoryButton(category.Id, deleteButton) } 

    appendChildren([nameTag, itemsButton, editButton, deleteButton, br()], content)
}


const loadCategories = (categories: Category[], content: HTMLElement) => {
    if (categories !== null) {
        for (const category of categories) {
            insertCategory(category, content)
        }
    }

    const input = newInput()
    input.id = "newCategory"
    input.maxLength = 50
    input.placeholder = "Нова категория"

    const button = newButton("ДОБАВИ")
    const output = outputField()
    button.onclick = () => { postCategory() }

    appendChildren([br(), input, button, output], content)
    input.focus()
}


getCategories()
