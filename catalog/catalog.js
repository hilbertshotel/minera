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
//////////////////// ITEMS ////////////////////
const createItem = (item, content) => {
    const titleTag = document.createElement("h1");
    titleTag.innerHTML = item.name;
    content.appendChild(titleTag);
    const textTag = document.createElement("pre");
    textTag.innerHTML = item.description;
    content.appendChild(textTag);
    const imagesTag = document.createElement("div");
    imagesTag.className = "images";
    content.appendChild(imagesTag);
    for (const img of item.images) {
        const imgTag = document.createElement("img");
        imgTag.src = img;
        imagesTag.appendChild(imgTag);
    }
};
const loadItems = (items) => {
    const content = document.getElementById("content");
    // clear content div
    while (content.firstChild) {
        content.removeChild(content.firstChild);
    }
    // list all items
    for (const item of items) {
        createItem(item, content);
    }
    // create back button
    const backButton = document.createElement("button");
    backButton.innerHTML = "ОБРАТНО";
    backButton.id = "button";
    backButton.onclick = fetchCategories;
    document.getElementById("mainWindow").appendChild(backButton);
};
const fetchItems = (category_id) => __awaiter(void 0, void 0, void 0, function* () {
    const body = {
        method: "POST",
        header: { "content-type": "application/json" },
        body: JSON.stringify(category_id)
    };
    const request = yield fetch(`${IP}/items`, body);
    if (request.ok) {
        const items = yield request.json();
        if (items === null) {
            loadItems([]);
        }
        else {
            loadItems(items);
        }
    }
});
//////////////////// CATEGORIES ////////////////////
const createCategory = (category, content) => {
    const [id, name] = category;
    const div = document.createElement("div");
    div.className = "category";
    // div.id = `${id}`
    div.innerHTML = name;
    div.onclick = () => { fetchItems(id); };
    content.appendChild(div);
};
const loadCategories = (categories) => {
    var _a;
    const content = document.getElementById("content");
    // remove back button
    (_a = document.getElementById("button")) === null || _a === void 0 ? void 0 : _a.remove();
    // clear content div
    while (content.firstChild) {
        content.removeChild(content.firstChild);
    }
    // list all categories
    for (const category of categories) {
        createCategory(category, content);
    }
};
const fetchCategories = () => __awaiter(void 0, void 0, void 0, function* () {
    const request = yield fetch(`${IP}/categories`);
    if (request.ok) {
        const data = yield request.json();
        loadCategories(data);
    }
});
//////////////////// MAIN ////////////////////
const IP = "http://127.0.0.1";
fetchCategories();
let isLogged = false;
isLogged = true;
