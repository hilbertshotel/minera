const createItem = (item: any, content: HTMLElement) => {
    const titleTag = document.createElement("h1")
    titleTag.innerHTML = item.name
    content.appendChild(titleTag)

    const textTag = document.createElement("pre")
    textTag.innerHTML = item.description
    content.appendChild(textTag)

    const imagesTag = document.createElement("div")
    imagesTag.className = "images"
    content.appendChild(imagesTag)

    for (const img of item.images) {
        const imgTag = document.createElement("img")
        imgTag.src = img
        imagesTag.appendChild(imgTag)
    }
    
}


const loadItems = (items: string[]) => {
    const content = document.getElementById("content")!

    // clear content div
    while(content.firstChild){
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
    const request = await fetch(`${IP}/items`, body)
    if (request.ok) {
        const items = await request.json()
        if (items === null) { loadItems([]) }
        else { loadItems(items) }
    }
}
