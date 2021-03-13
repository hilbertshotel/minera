const loadItems = (items: Item[], content: HTMLElement) => {
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


const fetchItems = async (id: number, content: HTMLElement) => {
    const url = `${IP}/LoadItems`
    const data = newPackage("POST", id)
    const request = await fetch(url, data)
    if (request.ok) {
        const items = await request.json()
        loadItems(items, content)
    }
}