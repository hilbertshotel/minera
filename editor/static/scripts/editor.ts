// ITEMS
const insertItem = (item: Item, content: HTMLElement) => {
    const input = newInput()
    input.value = item.Name
    input.placeholder = item.Name
    const button1 = newButton("РЕДАКЦИЯ")

    const textarea = newTextArea()
    textarea.value = item.Description
    textarea.placeholder = item.Description
    const button2 = newButton("РЕДАКЦИЯ")

    const children = [input, button1, br(), textarea, button2, br()]
    appendChildren(children, content)

    for (const src of item.Images) {
        const p = document.createElement("p")
        p.innerHTML = src
        content.appendChild(p)
    }

    content.appendChild(hr())
}


const loadItems = (category: Category, items: Item[]) => {
    const content = getById("content")
    clear(content)

    const input = newInput()
    input.value = category.Name
    input.placeholder = category.Name

    const button = newButton("РЕДАКЦИЯ")

    appendChildren([input, button, hr()], content)

    for (const item of items) {
        insertItem(item, content)
    }

    input.focus()
}


const fetchItems = async (category: Category) => {
    const url = `${IP}/LoadItems`
    const data = newPackage("POST", category.Id)
    const request = await fetch(url, data)
    if (request.ok) {
        const items = await request.json()
        loadItems(category, items)
    }
}


// CATEGORIES
const addNewCategory = async (output: HTMLElement) => {
    const newCategory = getInputValue("new")
    if (!newCategory) { output.innerHTML = "ВЪВЕДЕТЕ ИМЕ"; return }
    
    const url = `${IP}/NewCategory`
    const data = newPackage("POST", newCategory)
    const request = await fetch(url, data)
    if (request.ok) { fetchCategories() }
}


const insertCategory = (category: Category, content: HTMLElement) => {
    const button = newButton(category.Name)
    button.id = "category"
    button.onclick = () => { fetchItems(category) }
    content.appendChild(button)
}


const loadCategories = (categories: Category[], content: HTMLElement) => {
    for (const category of categories) {
        insertCategory(category, content)
    }

    const input = newInput()
    input.id = "new"
    input.placeholder = "Нова категория"
    const button = newButton("ДОБАВИ")
    const output = outputField()
    button.onclick = () => { addNewCategory(output) }

    appendChildren([br(), input, button, output], content)
    input.focus()
}


const fetchCategories = async () => {
    const content = getById("content")
    clear(content)
    const request = await fetch(`${IP}/LoadCategories`)
    if (request.ok) {
        const categories = await request.json()
        loadCategories(categories, content)
    }
}


fetchCategories()
