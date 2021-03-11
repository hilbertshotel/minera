// ITEMS
const createItem = (item: Item, content: HTMLElement) => {
    const titleTag = document.createElement("h1")
    titleTag.innerHTML = item.Name
    content.appendChild(titleTag)

    const textTag = document.createElement("pre")
    textTag.innerHTML = item.Description
    content.appendChild(textTag)

    const imagesTag = document.createElement("div")
    imagesTag.className = "images"
    content.appendChild(imagesTag)

    for (const image of item.Images) {
        const imgTag = document.createElement("img")
        imgTag.src = image
        imagesTag.appendChild(imgTag)
    }
}


const loadItems = (items: Item[]) => {
    const content = document.getElementById("content")!

    // clear content div
    while(content.firstChild) {
        content.removeChild(content.firstChild)
    }

    // list all items
    for (const item of items) {
        createItem(item, content)
    }

    // create back button
    const backButton = document.createElement("button")
    backButton.innerHTML = "ОБРАТНО"
    backButton.id = "button"
    backButton.onclick = fetchCategories
    document.getElementById("mainWindow")!.appendChild(backButton)
}


const fetchItems = async (category_id: number) => {
    const body = {
        method: "POST",
        header: {"content-type": "application/json"},
        body: JSON.stringify(category_id)
    }
    const request = await fetch(`${IP}/LoadItems`, body)
    if (request.ok) {
        const items = await request.json()
        if (items === null) { loadItems([]) }
        else { loadItems(items) }
    }
}


// CATEGORIES
const createCategory = (category: Category, content: HTMLElement) => {
    const id = category.Id;
    const name = category.Name;
    const div = document.createElement("div")
    div.className = "category"
    div.innerHTML = name
    div.onclick = () => { fetchItems(id) }
    content.appendChild(div)
}
  
  
const loadCategories = (categories: Category[]) => {
    const content = document.getElementById("content")!
  
    // remove back button
    document.getElementById("button")?.remove()
  
    // clear content div
    while(content.firstChild){
        content.removeChild(content.firstChild)
    }
  
    // list all categories
    for (const category of categories) {
        createCategory(category, content)
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
