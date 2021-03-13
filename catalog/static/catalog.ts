// ITEMS
const insertItem = (item: Item, content: HTMLElement) => {
    const titleTag = newElement("h1")
    titleTag.innerHTML = item.Name
    content.appendChild(titleTag)

    const textTag = newElement("pre")
    textTag.innerHTML = item.Description
    content.appendChild(textTag)

    const imagesTag = newElement("div")
    imagesTag.className = "images"
    content.appendChild(imagesTag)

    for (const image of item.Images) {
        const imgTag = newImgElement()
        imgTag.src = image
        imagesTag.appendChild(imgTag)
    }
}


const loadItems = (items: Item[]) => {
    const content = getById("content")!
    clear(content)

    for (const item of items) {
        insertItem(item, content)
    }

    const backButton = newElement("button")
    backButton.innerHTML = "ОБРАТНО"
    backButton.id = "button"
    backButton.onclick = fetchCategories
    getById("mainWindow")!.appendChild(backButton)
}


const fetchItems = async (id: number) => {
    const url = `${IP}/LoadItems`
    const data = newPackage("POST", id)

    const request = await fetch(url, data)
    if (request.ok) {
        const items = await request.json()
        if (items === null) { loadItems([]); return }
        loadItems(items)
    }
}


// CATEGORIES
const insertCategory = (category: Category, content: HTMLElement) => {
    const id = category.Id;
    const name = category.Name;
    const div = newElement("div")
    div.className = "category"
    div.innerHTML = name
    div.onclick = () => { fetchItems(id) }
    content.appendChild(div)
}
  
  
const loadCategories = (categories: Category[]) => {
    const content = getById("content")
    getById("button")?.remove()
    clear(content)
    
    for (const category of categories) {
        insertCategory(category, content)
    }
}
  
  
const fetchCategories = async () => {
    const url = `${IP}/LoadCategories`
    const request = await fetch(url)
    if (request.ok) {
        const categories = await request.json()
        loadCategories(categories)
    }
}
  

// MAIN
interface Item {
    Name: string,
    Description: string,
    Images: string[]
}

interface Category {
    Id: number,
    Name: string
}

const IP = "http://127.0.0.1"
fetchCategories()
