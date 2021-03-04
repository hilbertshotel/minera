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
const createCategory = (name, content) => {
    const category = document.createElement("div");
    category.className = "category";
    category.innerHTML = name;
    category.onclick = () => { fetchItems(name); };
    content.appendChild(category);
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
    for (const name of categories) {
        createCategory(name, content);
    }
};
const fetchCategories = () => __awaiter(void 0, void 0, void 0, function* () {
    const request = yield fetch(`${IP}/categories`);
    if (request.ok) {
        const data = yield request.json();
        loadCategories(data);
    }
});
