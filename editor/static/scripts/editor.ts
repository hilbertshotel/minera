// ITEMS
const loadItems = (items: Item[]) => {
    const content = getById("content")
    clear(content)
    // add category section
        // edit category name
        // delete category with WARNING
    // add items section
        // list all items
        // edit item title
        // edit item description
        // edit item images
        // add new item
    console.log(items)
}


const fetchItems = async (id: number) => {
    const url = `${IP}/LoadItems`
    const data = newPackage("POST", id)
    const request = await fetch(url, data)
    if (request.ok) {
        const items = await request.json()
        loadItems(items)
    }
}


// CATEGORIES
const addNewCategory = async () => {
    const output = getById("output")
    const newCategory = getInputWithId("new").value
    if (!newCategory) { output.innerHTML = "ВЪВЕДЕТЕ ИМЕ"; return }
    
    const url = `${IP}/NewCategory`
    const data = newPackage("POST", newCategory)
    const request = await fetch(url, data)
    if (request.ok) { fetchCategories() }
}


const insertCategory = (category: Category, content: HTMLElement) => {
    const [id, name] = [category.Id, category.Name]
    const element = `<button id="category" onclick="fetchItems(${id})">${name}</button><br>`
    content.innerHTML += element
}


const loadCategories = (categories: Category[], content: HTMLElement) => {
    for (const category of categories) {
        insertCategory(category, content)
    }

    const newCategorySection = `
    <input id="new" placeholder="Нова категория">
    <button onclick="addNewCategory()">ДОБАВИ</button>
    <p id="output"></p>`
    content.innerHTML += newCategorySection

    getById("new").focus()
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
