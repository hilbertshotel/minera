"use strict";
var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
// ITEM METHODS
const postItem = (category) => __awaiter(void 0, void 0, void 0, function* () {
    const name = getInputValue("name");
    const description = getInputValue("description");
    const imagesTag = getById("images");
    const output = getById("output");
    if (!name) {
        output.innerHTML = "ВЪВЕДЕТЕ ИМЕ";
        return;
    }
    if (!description) {
        output.innerHTML = "ВЪВЕДЕТЕ ОПИСАНИЕ";
        output.focus();
        return;
    }
    let filenameArray = [];
    const images = new FormData();
    const files = imagesTag.files;
    if (!files) {
        return;
    }
    if (files.length === 0) {
        output.innerHTML = "ДОБАВЕТЕ ПОНЕ ЕДНА СНИМКА";
        output.focus();
        return;
    }
    if (files.length > 3) {
        output.innerHTML = "ИЗБРАЛИ СТЕ ПОВЕЧЕ ОТ ТРИ ИЗОБРАЖЕНИЯ";
        output.focus();
        return;
    }
    for (const file of files) {
        images.append("files", file);
        filenameArray.push(file.name);
    }
    const newItem = {
        id: category.Id,
        name: name,
        description: description,
        images: filenameArray
    };
    const data = { method: "POST", body: images };
    const request1 = yield fetch(`${IP}/FileTransfer`, data);
    if (request1.ok) {
        const data = newPackage("POST", newItem);
        const request2 = yield fetch(`${IP}/Editor/Items/0`, data);
        if (request2.ok) {
            getItems(category);
        }
    }
});
const getItems = (category) => __awaiter(void 0, void 0, void 0, function* () {
    const request = yield fetch(`${IP}/Editor/Items/${category.Id}`);
    if (request.ok) {
        const items = yield request.json();
        loadItems(category, items);
    }
});
// ITEM PAGE CONSTRUCTION
const insertNewItemSection = (category, content) => {
    const header = newElement("h1");
    header.innerHTML = "Добавяне на нов артикул";
    const name = newInput();
    name.id = "name";
    name.maxLength = 50;
    name.placeholder = "Име на артикула";
    const description = newTextArea();
    description.id = "description";
    description.maxLength = 300;
    description.placeholder = "Описание на артикула";
    const images = newInput();
    images.type = "file";
    images.id = "images";
    images.setAttribute("multiple", "");
    images.accept = "image/*";
    const addButton = newButton("ДОБАВИ");
    addButton.onclick = () => { postItem(category); };
    appendChildren([header, name, br(), description, br(), images, br(), addButton], content);
    name.focus();
};
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
const insertItem = (item, content) => {
    const itemTag = newButton(item.Name);
    content.appendChild(itemTag);
};
const loadItems = (category, items) => {
    const content = getById("content");
    clear(content);
    getById("h").innerHTML = category.Name;
    if (items !== null) {
        for (const item of items) {
            insertItem(item, content);
        }
    }
    insertNewItemSection(category, content);
    const output = outputField();
    const backButton = newButton("ВЪРНИ СЕ");
    backButton.id = "backButton";
    backButton.onclick = getCategories;
    appendChildren([output, backButton], content);
};
// CATEGORY METHODS
const deleteCategory = (id) => __awaiter(void 0, void 0, void 0, function* () {
    const data = newPackage("DELETE", id);
    const request = yield fetch(`${IP}/Editor/Categories`, data);
    if (request.ok) {
        getCategories();
    }
});
const putCategory = (category) => __awaiter(void 0, void 0, void 0, function* () {
    const newName = getInputValue("categoryName");
    if (newName === category.Name || !newName) {
        return;
    }
    const info = { id: category.Id, newName: newName };
    const data = newPackage("PUT", info);
    const request = yield fetch(`${IP}/Editor/Categories`, data);
    if (request.ok) {
        getCategories();
    }
});
const postCategory = () => __awaiter(void 0, void 0, void 0, function* () {
    const newCategory = getInputValue("newCategory");
    if (!newCategory) {
        getById("output").innerHTML = "ВЪВЕДЕТЕ ИМЕ";
        return;
    }
    const data = newPackage("POST", newCategory);
    const request = yield fetch(`${IP}/Editor/Categories`, data);
    if (request.ok) {
        getCategories();
    }
});
const getCategories = () => __awaiter(void 0, void 0, void 0, function* () {
    getById("h").innerHTML = "РЕДАКТОР";
    const content = getById("content");
    clear(content);
    const request = yield fetch(`${IP}/Editor/Categories`);
    if (request.ok) {
        const categories = yield request.json();
        loadCategories(categories, content);
    }
});
// CATEGORY PAGE CONSTRUCTION
const insertCategory = (category, content) => {
    const nameTag = newInput();
    nameTag.id = "categoryName";
    nameTag.value = category.Name;
    nameTag.placeholder = category.Name;
    const itemsButton = newButton("Преглед");
    itemsButton.onclick = () => { getItems(category); };
    const editButton = newButton("Ново име");
    editButton.onclick = () => { putCategory(category); };
    const deleteButton = newButton("Изтрий");
    deleteButton.id = "deleteButton";
    deleteButton.onclick = () => { deleteCategory(category.Id); };
    appendChildren([nameTag, itemsButton, editButton, deleteButton, br()], content);
};
const loadCategories = (categories, content) => {
    for (const category of categories) {
        insertCategory(category, content);
    }
    const input = newInput();
    input.id = "newCategory";
    input.maxLength = 50;
    input.placeholder = "Нова категория";
    const button = newButton("ДОБАВИ");
    const output = outputField();
    button.onclick = () => { postCategory(); };
    appendChildren([br(), input, button, output], content);
    input.focus();
};
getCategories();
