// // ITEM METHODS
// const addNewItem = async (category: Category) => {
//     const name = getInputValue("name")
//     const description = getInputValue("description")
//     const images = <HTMLInputElement>getById("images")!

//     if (!name) {} // ADD ERROR MSG
//     if (!description) {} // ADD ERROR MSG

//     let filenameArray: string[] = []
//     const imgData = new FormData()
//     const files = images.files

//     if (!files) { return } // ADD ERROR MSG
//     if (files.length === 0) { return } // ADD ERROR MSG
//     if (files.length > 3) { return } // ADD ERROR MSG

//     for (const file of files) {
//         imgData.append("files", file)
//         filenameArray.push(file.name)
//     }

//     const itemData = {
//         id: category.Id,
//         name: name,
//         description: description,
//         images: filenameArray
//     }

//     // send data to backend
//     const data = { method: "POST", body: imgData }
//     const request1 = await fetch(`${IP}/NewItemImages`, data)
//     if (request1.ok) {
//         const data = newPackage("POST", itemData)
//         const request2 = await fetch(`${IP}/NewItem`, data)
//         if (request2.ok) { fetchItems(category) }
//     }
// }


// // LOAD ITEMS ON PAGE
// const insertNewItemSection = (category: Category, content: HTMLElement) => {
//     const header = newElement("h1")
//     header.innerHTML = "Добавяне на нов артикул"

//     const name = newInput()
//     name.id = "name"
//     name.maxLength = 50
//     name.placeholder = "Име на артикула"

//     const description = newTextArea()
//     description.id = "description"
//     description.maxLength = 300
//     description.placeholder = "Описание на артикула"

//     const images = newInput()
//     images.type = "file"
//     images.id = "images"
//     images.setAttribute("multiple", "")
//     images.accept = "image/*"
    
//     const addButton = newButton("ДОБАВИ")
//     addButton.onclick = () => { addNewItem(category) }

//     appendChildren([header, name, br(), description, br(), images, br(), addButton], content)
// }


// const insertItem = (id: number, item: Item, content: HTMLElement) => {
//     const input = newInput()
//     input.value = item.Name
//     input.maxLength = 50
//     input.placeholder = item.Name

//     const button1 = newButton("РЕДАКЦИЯ ИМЕ")
//     button1.onclick = () => {  } // edit item name
//     const button2 = newButton("ИЗТРИЙ")
//     button2.id = "delete"
//     button2.onclick = () => {  } // delete item with a warning

//     const textarea = newTextArea()
//     textarea.value = item.Description
//     textarea.maxLength = 300
//     textarea.placeholder = item.Description

//     const button3 = newButton("РЕДАКЦИЯ ОПИСАНИЕ")
//     button3.onclick = () => {  } // edit item description

//     const children = [input, button1, button2, br(), textarea, button3, br()]
//     appendChildren(children, content)

//     for (let i=0; i<3;i++) {
//         const input = newInput()
//         input.type = "file"
//         const button = newButton(`ДОБАВИ СНИМКА ${i+1}`)
//         button.onclick = () => {  } // change item image (using index)
//         appendChildren([input, button, br()], content)
//     }

//     // ADD EDIT ITEM OUTPUT FIELD

//     content.appendChild(hr())
// }


// const loadItems = (category: Category, items: Item[]) => {
//     const content = getById("content")
//     clear(content)

//     const input = newInput()
//     input.id = "editCategoryName"
//     input.maxLength = 50
//     input.value = category.Name
//     input.placeholder = category.Name

//     const button1 = newButton("РЕДАКЦИЯ ИМЕ")
//     button1.onclick = () => { editCategoryName(category) }
//     const button2 = newButton("ИЗТРИЙ")
//     button2.id = "delete"
//     button2.onclick = () => { deleteCategory(category.Id) }

//     // EDIT CATEGORY OUTPUT FIELD

//     appendChildren([input, button1, button2, hr()], content)

//     if (items !== null) {
//         for (const item of items) {
//             insertItem(category.Id, item, content)
//         }
//     }

//     input.focus()
//     insertNewItemSection(category, content)
//     // NEW ITEM OUTPUT FIELD
//     // ADD BACK BUTTON
// }


const fetchItems = async (category: Category) => {
    // const request = await fetch(`${IP}/Editor/Items/${category.Id}`)
    // if (request.ok) {
    //     const items = await request.json()
    //     loadItems(category, items)
    // }
}


// CATEGORY METHODS
const deleteCategory = async (id: number) => {
    const data = newPackage("DELETE", id)
    const request = await fetch(`${IP}/Editor/Categories`, data)
    if (request.ok) { getCategories() }
}


const putCategory = async (category: Category) => {
    const newName = getInputValue("categoryName")
    if (newName === category.Name) { return }
    if (!newName) {
        getById("output").innerHTML = category.Name
        return
    }

    const info = { id: category.Id, newName: newName }
    const data = newPackage("PUT", info)
    const request = await fetch(`${IP}/Editor/Categories`, data)
    if (request.ok) { getCategories() }
}


const postCategory = async () => {
    const newCategory = getInputValue("newCategory")
    if (!newCategory) {
        getById("output").innerHTML = "ВЪВЕДЕТЕ ИМЕ"
        return
    }
    
    const data = newPackage("POST", newCategory)
    const request = await fetch(`${IP}/Editor/Categories`, data)
    if (request.ok) { getCategories() }
}


// CATEGORY PAGE CONSTRUCTION
const insertCategory = (category: Category, content: HTMLElement) => {
    const nameTag = newInput()
    nameTag.id = "categoryName"
    nameTag.value = category.Name
    nameTag.placeholder = category.Name

    const itemsButton = newButton("Преглед")
    itemsButton.onclick = () => { fetchItems(category) }
    
    const editButton = newButton("Ново име")
    editButton.onclick = () => { putCategory(category) } 
    
    const deleteButton = newButton("Изтрий")
    deleteButton.id = "deleteButton"
    editButton.onclick = () => { deleteCategory(category.Id) } 

    appendChildren([nameTag, itemsButton, editButton, deleteButton, br()], content)
}


const loadCategories = (categories: Category[], content: HTMLElement) => {
    for (const category of categories) {
        insertCategory(category, content)
    }

    const input = newInput()
    input.id = "newCategory"
    input.maxLength = 50
    input.placeholder = "Нова категория"

    const button = newButton("ДОБАВИ")
    const output = outputField()
    button.onclick = () => { postCategory() }

    appendChildren([br(), input, button, output], content)
    input.focus()
}


const getCategories = async () => {
    const content = getById("content")
    clear(content)
    const request = await fetch(`${IP}/Editor/Categories`)
    if (request.ok) {
        const categories = await request.json()
        loadCategories(categories, content)
    }
}


getCategories()
