const addNewCategory = async (output: HTMLElement) => {
    const newCategoryName = getInputWithId("new").value
    if (!newCategoryName) { output.innerHTML = "ВЪВЕДЕТЕ ИМЕ"; return }
    
    const url = `${IP}/NewCategory`
    const data = newPackage("POST", newCategoryName)
    const request = await fetch(url, data)
    if (request.ok) { loadEditor() }
}


const insertCategory = (category: Category, content: HTMLElement) => {
    const button = newElement("button")
    button.id = "categoryButton"
    button.innerHTML = category.Name
    button.onclick = () => { fetchItems(category.Id, content) }
    const br = newElement("br")

    content.appendChild(button)
    content.appendChild(br)
}


const loadCategories = (categories: Category[], content: HTMLElement) => {
    for (const category of categories) {
        insertCategory(category, content)
    }

    const input = newInputElement()
    input.placeholder = "Нова категория"
    input.id = "new"

    const button = newElement("button")
    button.innerHTML = "ДОБАВИ"
    
    const output = newElement("p")
    output.id = "output"
    button.onclick = () => { addNewCategory(output) }

    content.appendChild(input)
    content.appendChild(button)
    content.appendChild(output)

    input.focus()
}


const fetchCategories = async (content: HTMLElement) => {
    const url = `${IP}/LoadCategories`
    const request = await fetch(url)
    if (request.ok) {
        const categories = await request.json()
        loadCategories(categories, content)
    }
}


const loadEditor = () => {
    const content = getById("content")
    clear(content)
    fetchCategories(content)
}


loadEditor()