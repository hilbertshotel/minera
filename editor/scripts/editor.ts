const newCategory = () => {

}


const deleteCategory = (categoryId: string) => {

}


const editCategory = async (categoryId: string, oldName: string) => {
    const input = <HTMLInputElement>document.getElementById(categoryId)
    const newName = input.value

    if (newName !== oldName) {
        const body = {
            method: "PUT",
            header: { "content-type": "application/json" },
            body: JSON.stringify([newName, oldName])
        }
        const request = await fetch(`${IP}/edit_category`, body)
        if (request.ok) {
            input.value = await request.json()
        }
    }
}


const loadCategory = (category: string[], content: HTMLElement) => {
    const id = category[0]
    const name = category[1]

    const input = document.createElement("input")
    input.id = id
    input.value = name

    const button1 = document.createElement("button")
    button1.innerHTML = "ЗАПАЗИ"
    button1.onclick = () => { editCategory(id, name) }

    const button2 = document.createElement("button")
    button2.innerHTML = "ИЗТРИЙ"
    button2.onclick = () => { deleteCategory(id) }

    const br = document.createElement("br")

    content.appendChild(input)
    content.appendChild(button1)
    content.appendChild(button2)
    content.appendChild(br)
}


const loadCategories = (categories: string[][], content: HTMLElement) => {
    for (const category of categories) {
        loadCategory(category, content)
    }

    const input = document.createElement("input")
    input.placeholder = "Нова категория"
    input.id = "new"

    const button = document.createElement("button")
    button.innerHTML = "ДОБАВИ"
    button.onclick = newCategory
    
    const output = document.createElement("p")
    output.id = "output"

    content.appendChild(input)
    content.appendChild(button)
    content.appendChild(output)
}


const getCategories = async (content: HTMLElement) => {
    const request = await fetch(`${IP}/categories`)
    if (request.ok) {
        const categories = await request.json()
        loadCategories(categories, content)
    }
}


const loadEditor = () => {
    const content = document.getElementById("content")!

    while(content.firstChild) {
        content.removeChild(content.firstChild)
    }

    getCategories(content)
}