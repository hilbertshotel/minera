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
    const titleTag = `<h1>${item.Name}</h1>`;
    content.innerHTML += titleTag;
    const textTag = `<pre>${item.Description}</pre>`;
    content.innerHTML += textTag;
    const imagesTag = newElement("div");
    imagesTag.className = "images";
    content.appendChild(imagesTag);
    for (const image of item.Images) {
        const imgTag = `<img src="${image}">`;
        imagesTag.innerHTML += imgTag;
    }
};
const loadItems = (items) => {
    const content = getById("content");
    clear(content);
    for (const item of items) {
        insertItem(item, content);
    }
    const backButton = `<button id="button" onclick="fetchCategories()">ОБРАТНО</button>`;
    content.innerHTML += backButton;
};
const fetchItems = (id) => __awaiter(void 0, void 0, void 0, function* () {
    const data = newPackage("POST", id);
    const request = yield fetch(`${IP}/LoadItems`, data);
    if (request.ok) {
        const items = yield request.json();
        if (items === null) {
            loadItems([]);
            return;
        }
        loadItems(items);
    }
});
// CATEGORIES
const insertCategory = (category, content) => {
    const [id, name] = [category.Id, category.Name];
    const element = `<div class="category" onclick="fetchItems(${id})">${name}</div>`;
    content.innerHTML += element;
};
const loadCategories = (categories) => {
    const content = getById("content");
    clear(content);
    for (const category of categories) {
        insertCategory(category, content);
    }
};
const fetchCategories = () => __awaiter(void 0, void 0, void 0, function* () {
    const request = yield fetch(`${IP}/LoadCategories`);
    if (request.ok) {
        const categories = yield request.json();
        loadCategories(categories);
    }
});
// MAIN
const IP = "http://127.0.0.1";
fetchCategories();
