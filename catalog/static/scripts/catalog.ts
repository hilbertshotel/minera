// DATA
interface Item {
    Name: string,
    Description: string,
    Images: string[]
}

interface Category {
    Id: number,
    Name: string
}


// ITEMS
const insertItem = (item: Item, content: HTMLElement) => {
    const titleTag = `<h1>${item.Name}</h1>`
    content.innerHTML += titleTag

    const textTag = `<pre>${item.Description}</pre>`
    content.innerHTML += textTag

    const imagesTag = newElement("div")
    imagesTag.className = "images"
    content.appendChild(imagesTag)

    for (const image of item.Images) {
        const imgTag = `<img src="${image}">`
        imagesTag.innerHTML += imgTag
    }
}


const loadItems = (items: Item[]) => {
    const content = getById("content")!
    clear(content)

    for (const item of items) {
        insertItem(item, content)
    }

    const backButton = `<button id="button" onclick="fetchCategories()">ОБРАТНО</button>`
    content.innerHTML += backButton
}


const fetchItems = async (id: number) => {
    const data = newPackage("POST", id)
    const request = await fetch(`${IP}/LoadItems`, data)
    if (request.ok) {
        const items = await request.json()
        if (items === null) { loadItems([]); return }
        loadItems(items)
    }
}


// CATEGORIES
const insertCategory = (category: Category, content: HTMLElement) => {
    const [id, name] = [category.Id, category.Name]
    const element = `<div class="category" onclick="fetchItems(${id})">${name}</div>`
    content.innerHTML += element
}
  

const loadCategories = (categories: Category[]) => {
    const content = getById("content")
    clear(content)
    for (const category of categories) {
        insertCategory(category, content)
    }
}
  

const fetchCategories = async () => {
    const request = await fetch(`${IP}/LoadCategories`)
    if (request.ok) {
        const categories = await request.json()
        loadCategories(categories)
    }
}
  

// MAIN
const IP = "http://127.0.0.1"
fetchCategories()
