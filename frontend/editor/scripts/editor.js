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
const deleteItem = (id, category) => __awaiter(void 0, void 0, void 0, function* () {
    const data = newPackage("DELETE", "");
    const request = yield fetch(`${IP}/Editor/Items/${id}`, data);
    if (request.ok) {
        getItems(category);
    }
});
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
const deleteItemButton = (id, button, category) => __awaiter(void 0, void 0, void 0, function* () {
    if (button.id === "deleteButton") {
        deleteItem(id, category);
    }
    button.id = "deleteButton";
    yield new Promise((resolve) => setTimeout(resolve, 3000));
    button.id = "";
});
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
const insertItem = (item, content, category) => {
    const nameTag = newButton(item.Name);
    nameTag.className = "itemName";
    const deleteButton = newButton("Изтрий");
    deleteButton.onclick = () => { deleteItemButton(item.Id, deleteButton, category); };
    appendChildren([nameTag, deleteButton, br()], content);
};
const loadItems = (category, items) => {
    const content = getById("content");
    clear(content);
    getById("h").innerHTML = category.Name;
    if (items !== null) {
        for (const item of items) {
            insertItem(item, content, category);
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
const putCategory = (category, nameTag) => __awaiter(void 0, void 0, void 0, function* () {
    const newName = nameTag.value;
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
        const output = getById("output");
        output.innerHTML = "ВЪВЕДЕТЕ ИМЕ";
        output.focus();
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
const deleteCategoryButton = (id, button) => __awaiter(void 0, void 0, void 0, function* () {
    if (button.id === "deleteButton") {
        deleteCategory(id);
    }
    button.id = "deleteButton";
    yield new Promise((resolve) => setTimeout(resolve, 3000));
    button.id = "";
});
const insertCategory = (category, content) => {
    const nameTag = newInput();
    nameTag.className = "categoryName";
    nameTag.value = category.Name;
    nameTag.placeholder = category.Name;
    const itemsButton = newButton("Преглед");
    itemsButton.onclick = () => { getItems(category); };
    const editButton = newButton("Ново име");
    editButton.onclick = () => { putCategory(category, nameTag); };
    const deleteButton = newButton("Изтрий");
    deleteButton.onclick = () => { deleteCategoryButton(category.Id, deleteButton); };
    appendChildren([nameTag, itemsButton, editButton, deleteButton, br()], content);
};
const loadCategories = (categories, content) => {
    if (categories !== null) {
        for (const category of categories) {
            insertCategory(category, content);
        }
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
