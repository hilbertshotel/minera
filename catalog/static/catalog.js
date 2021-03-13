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
    content.appendChild(titleTag);
    const textTag = newElement("pre");
    textTag.innerHTML = item.Description;
    content.appendChild(textTag);
    const imagesTag = newElement("div");
    imagesTag.className = "images";
    content.appendChild(imagesTag);
    for (const image of item.Images) {
        const imgTag = newImgElement();
        imgTag.src = image;
        imagesTag.appendChild(imgTag);
    }
};
const loadItems = (items) => {
    const content = getById("content");
    clear(content);
    for (const item of items) {
        insertItem(item, content);
    }
    const backButton = newElement("button");
    backButton.innerHTML = "ОБРАТНО";
    backButton.id = "button";
    backButton.onclick = fetchCategories;
    getById("mainWindow").appendChild(backButton);
};
const fetchItems = (id) => __awaiter(void 0, void 0, void 0, function* () {
    const url = `${IP}/LoadItems`;
    const data = newPackage("POST", id);
    const request = yield fetch(url, data);
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
    const id = category.Id;
    const name = category.Name;
    const div = newElement("div");
    div.className = "category";
    div.innerHTML = name;
    div.onclick = () => { fetchItems(id); };
    content.appendChild(div);
};
const loadCategories = (categories) => {
    var _a;
    const content = getById("content");
    (_a = getById("button")) === null || _a === void 0 ? void 0 : _a.remove();
    clear(content);
    for (const category of categories) {
        insertCategory(category, content);
    }
};
const fetchCategories = () => __awaiter(void 0, void 0, void 0, function* () {
    const url = `${IP}/LoadCategories`;
    const request = yield fetch(url);
    if (request.ok) {
        const categories = yield request.json();
        loadCategories(categories);
    }
});
const IP = "http://127.0.0.1";
fetchCategories();
