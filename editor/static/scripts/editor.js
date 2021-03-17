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
const addNewItem = (category) => __awaiter(void 0, void 0, void 0, function* () {
    const name = getInputValue("name");
    const description = getInputValue("description");
    const images = getById("images");
    if (!name) { } // ADD ERROR MSG
    if (!description) { } // ADD ERROR MSG
    let filenameArray = [];
    const imgData = new FormData();
    const files = images.files;
    if (!files) {
        return;
    } // ADD ERROR MSG
    if (files.length === 0) {
        return;
    } // ADD ERROR MSG
    if (files.length > 3) {
        return;
    } // ADD ERROR MSG
    for (const file of files) {
        imgData.append("files", file);
        filenameArray.push(file.name);
    }
    const itemData = {
        id: category.Id,
        name: name,
        description: description,
        images: filenameArray
    };
    // send data to backend
    const data = { method: "POST", body: imgData };
    const request1 = yield fetch(`${IP}/NewItemImages`, data);
    if (request1.ok) {
        const data = newPackage("POST", itemData);
        const request2 = yield fetch(`${IP}/NewItem`, data);
        if (request2.ok) {
            fetchItems(category);
        }
    }
});
// LOAD ITEMS ON PAGE
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
    addButton.onclick = () => { addNewItem(category); };
    appendChildren([header, name, br(), description, br(), images, br(), addButton], content);
};
const insertItem = (id, item, content) => {
    const input = newInput();
    input.value = item.Name;
    input.maxLength = 50;
    input.placeholder = item.Name;
    const button1 = newButton("РЕДАКЦИЯ ИМЕ");
    button1.onclick = () => { }; // edit item name
    const button2 = newButton("ИЗТРИЙ");
    button2.id = "delete";
    button2.onclick = () => { }; // delete item with a warning
    const textarea = newTextArea();
    textarea.value = item.Description;
    textarea.maxLength = 300;
    textarea.placeholder = item.Description;
    const button3 = newButton("РЕДАКЦИЯ ОПИСАНИЕ");
    button3.onclick = () => { }; // edit item description
    const children = [input, button1, button2, br(), textarea, button3, br()];
    appendChildren(children, content);
    for (let i = 0; i < 3; i++) {
        const input = newInput();
        input.type = "file";
        const button = newButton(`ДОБАВИ СНИМКА ${i + 1}`);
        button.onclick = () => { }; // change item image (using index)
        appendChildren([input, button, br()], content);
    }
    // ADD EDIT ITEM OUTPUT FIELD
    content.appendChild(hr());
};
const loadItems = (category, items) => {
    const content = getById("content");
    clear(content);
    const input = newInput();
    input.id = "editCategoryName";
    input.maxLength = 50;
    input.value = category.Name;
    input.placeholder = category.Name;
    const button1 = newButton("РЕДАКЦИЯ ИМЕ");
    button1.onclick = () => { editCategoryName(category); };
    const button2 = newButton("ИЗТРИЙ");
    button2.id = "delete";
    button2.onclick = () => { deleteCategory(category.Id); };
    // EDIT CATEGORY OUTPUT FIELD
    appendChildren([input, button1, button2, hr()], content);
    if (items !== null) {
        for (const item of items) {
            insertItem(category.Id, item, content);
        }
    }
    input.focus();
    insertNewItemSection(category, content);
    // NEW ITEM OUTPUT FIELD
    // ADD BACK BUTTON
};
const fetchItems = (category) => __awaiter(void 0, void 0, void 0, function* () {
    const data = newPackage("POST", category.Id);
    const request = yield fetch(`${IP}/LoadItems`, data);
    if (request.ok) {
        const items = yield request.json();
        loadItems(category, items);
    }
});
// CATEGORY METHODS
const deleteCategory = (id) => __awaiter(void 0, void 0, void 0, function* () {
    const data = newPackage("DELETE", id);
    const request = yield fetch(`${IP}/DeleteCategory`, data);
    if (request.ok) {
        fetchCategories();
    }
});
const editCategoryName = (category) => __awaiter(void 0, void 0, void 0, function* () {
    const newName = getInputValue("editCategoryName");
    if (newName === category.Name) {
        return;
    }
    if (!newName) {
        return;
    } // ERROR NO NAME ENTERED
    const info = { id: category.Id, newName: newName };
    const data = newPackage("PUT", info);
    const request = yield fetch(`${IP}/EditCategoryName`, data);
    if (request.ok) {
        const editedCategory = yield request.json();
        fetchItems(editedCategory);
    }
});
const addNewCategory = (output) => __awaiter(void 0, void 0, void 0, function* () {
    const newCategory = getInputValue("new");
    if (!newCategory) {
        output.innerHTML = "ВЪВЕДЕТЕ ИМЕ";
        return;
    }
    const url = `${IP}/NewCategory`;
    const data = newPackage("POST", newCategory);
    const request = yield fetch(url, data);
    if (request.ok) {
        fetchCategories();
    }
});
// LOAD CATEGORIES ON PAGE
const insertCategory = (category, content) => {
    const button = newButton(category.Name);
    button.id = "category";
    button.onclick = () => { fetchItems(category); };
    content.appendChild(button);
};
const loadCategories = (categories, content) => {
    for (const category of categories) {
        insertCategory(category, content);
    }
    const input = newInput();
    input.id = "new";
    input.maxLength = 50;
    input.placeholder = "Нова категория";
    const button = newButton("ДОБАВИ");
    const output = outputField();
    button.onclick = () => { addNewCategory(output); };
    appendChildren([br(), input, button, output], content);
    input.focus();
};
const fetchCategories = () => __awaiter(void 0, void 0, void 0, function* () {
    const content = getById("content");
    clear(content);
    const request = yield fetch(`${IP}/LoadCategories`);
    if (request.ok) {
        const categories = yield request.json();
        loadCategories(categories, content);
    }
});
fetchCategories();
