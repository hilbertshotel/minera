// INTERFACES
interface Item {
    Id: number,
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
    const titleTag = newElement("h1")
    titleTag.innerHTML = item.Name

    const textTag = newElement("pre")
    textTag.innerHTML = item.Description

    const imagesDiv = newElement("div")
    imagesDiv.className = "images"

    appendChildren([titleTag, textTag, imagesDiv], content)

    for (const src of item.Images) {
        const imgTag = newImgElement(src)
        imagesDiv.appendChild(imgTag)
    }
}


const loadItems = (items: Item[]) => {
    const content = getById("content")!
    clear(content)

    for (const item of items) {
        insertItem(item, content)
    }

    const backButton = newButton("ОБРАТНО")
    backButton.id = "button"
    backButton.onclick = () => { fetchCategories() }
    content.appendChild(backButton)
}


const fetchItems = async (id: number) => {
    const request = await fetch(`${IP}/Catalog/Items/${id}`)
    if (request.ok) {
        const items = await request.json()
        if (items === null) { loadItems([]); return }
        loadItems(items)
    }
}


// CATEGORIES
const insertCategory = (category: Category, content: HTMLElement) => {
    const element = newElement("div")
    element.className = "category"
    element.innerHTML = category.Name
    element.onclick = () => { fetchItems(category.Id) }
    content.appendChild(element)
}
  

const loadCategories = (categories: Category[]) => {
    const content = getById("content")
    clear(content)
    if (categories !== null) {
        for (const category of categories) {
            insertCategory(category, content)
        }
    }
}
  

const fetchCategories = async () => {
    const request = await fetch(`${IP}/Catalog/Categories`)
    if (request.ok) {
        const categories = await request.json()
        loadCategories(categories)
    }
}
  

// MAIN
const IP = "http://127.0.0.1:5252"
fetchCategories()
