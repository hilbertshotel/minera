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
const newCategory = () => {
};
const deleteCategory = (categoryId) => {
};
const editCategory = (categoryId, oldName) => __awaiter(void 0, void 0, void 0, function* () {
    const input = document.getElementById(categoryId);
    const newName = input.value;
    if (newName !== oldName) {
        const body = {
            method: "PUT",
            header: { "content-type": "application/json" },
            body: JSON.stringify([newName, oldName])
        };
        const request = yield fetch(`${IP}/edit_category`, body);
        if (request.ok) {
            input.value = yield request.json();
        }
    }
});
const loadCategory = (category, content) => {
    const id = category[0];
    const name = category[1];
    const input = document.createElement("input");
    input.id = id;
    input.value = name;
    const button1 = document.createElement("button");
    button1.innerHTML = "ЗАПАЗИ";
    button1.onclick = () => { editCategory(id, name); };
    const button2 = document.createElement("button");
    button2.innerHTML = "ИЗТРИЙ";
    button2.onclick = () => { deleteCategory(id); };
    const br = document.createElement("br");
    content.appendChild(input);
    content.appendChild(button1);
    content.appendChild(button2);
    content.appendChild(br);
};
const loadCategories = (categories, content) => {
    for (const category of categories) {
        loadCategory(category, content);
    }
    const input = document.createElement("input");
    input.placeholder = "Нова категория";
    input.id = "new";
    const button = document.createElement("button");
    button.innerHTML = "ДОБАВИ";
    button.onclick = newCategory;
    const output = document.createElement("p");
    output.id = "output";
    content.appendChild(input);
    content.appendChild(button);
    content.appendChild(output);
};
const getCategories = (content) => __awaiter(void 0, void 0, void 0, function* () {
    const request = yield fetch(`${IP}/categories`);
    if (request.ok) {
        const categories = yield request.json();
        loadCategories(categories, content);
    }
});
const loadEditor = () => {
    const content = document.getElementById("content");
    while (content.firstChild) {
        content.removeChild(content.firstChild);
    }
    getCategories(content);
};
