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
// ITEMS
const insertItem = (item, content) => {
    const input = newInput();
    input.value = item.Name;
    input.placeholder = item.Name;
    const button1 = newButton("РЕДАКЦИЯ");
    const textarea = newTextArea();
    textarea.value = item.Description;
    textarea.placeholder = item.Description;
    const button2 = newButton("РЕДАКЦИЯ");
    const children = [input, button1, br(), textarea, button2, br()];
    appendChildren(children, content);
    for (const src of item.Images) {
        const p = document.createElement("p");
        p.innerHTML = src;
        content.appendChild(p);
    }
    content.appendChild(hr());
};
const loadItems = (category, items) => {
    const content = getById("content");
    clear(content);
    const input = newInput();
    input.value = category.Name;
    input.placeholder = category.Name;
    const button = newButton("РЕДАКЦИЯ");
    appendChildren([input, button, hr()], content);
    for (const item of items) {
        insertItem(item, content);
    }
    input.focus();
};
const fetchItems = (category) => __awaiter(void 0, void 0, void 0, function* () {
    const url = `${IP}/LoadItems`;
    const data = newPackage("POST", category.Id);
    const request = yield fetch(url, data);
    if (request.ok) {
        const items = yield request.json();
        loadItems(category, items);
    }
});
// CATEGORIES
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
