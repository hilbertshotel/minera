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
    const titleTag = newElement("h1");
    titleTag.innerHTML = item.Name;
    const textTag = newElement("pre");
    textTag.innerHTML = item.Description;
    const imagesDiv = newElement("div");
    imagesDiv.className = "images";
    appendChildren([titleTag, textTag, imagesDiv], content);
    for (const src of item.Images) {
        const imgTag = newImgElement(src);
        imagesDiv.appendChild(imgTag);
    }
};
const loadItems = (items) => {
    const content = getById("content");
    clear(content);
    for (const item of items) {
        insertItem(item, content);
    }
    const backButton = newButton("ОБРАТНО");
    backButton.id = "button";
    backButton.onclick = () => { fetchCategories(); };
    content.appendChild(backButton);
};
const fetchItems = (id) => __awaiter(void 0, void 0, void 0, function* () {
    const request = yield fetch(`${IP}/Catalog/Items/${id}`);
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
    const element = newElement("div");
    element.className = "category";
    element.innerHTML = category.Name;
    element.onclick = () => { fetchItems(category.Id); };
    content.appendChild(element);
};
const loadCategories = (categories) => {
    const content = getById("content");
    clear(content);
    if (categories !== null) {
        for (const category of categories) {
            insertCategory(category, content);
        }
    }
};
const fetchCategories = () => __awaiter(void 0, void 0, void 0, function* () {
    const request = yield fetch(`${IP}/Catalog/Categories`);
    if (request.ok) {
        const categories = yield request.json();
        loadCategories(categories);
    }
});
// MAIN
const IP = "http://127.0.0.1:5252"; // development
// const IP = "https://catalog.minera.bg" // production
fetchCategories();
