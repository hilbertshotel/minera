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
const loadItems = (items) => {
    const content = getById("content");
    clear(content);
    // add category section
    // edit category name
    // delete category with WARNING
    // add items section
    // list all items
    // edit item title
    // edit item description
    // edit item images
    // add new item
    console.log(items);
};
const fetchItems = (id) => __awaiter(void 0, void 0, void 0, function* () {
    const url = `${IP}/LoadItems`;
    const data = newPackage("POST", id);
    const request = yield fetch(url, data);
    if (request.ok) {
        const items = yield request.json();
        loadItems(items);
    }
});
// CATEGORIES
const addNewCategory = () => __awaiter(void 0, void 0, void 0, function* () {
    const output = getById("output");
    const newCategory = getInputWithId("new").value;
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
    const [id, name] = [category.Id, category.Name];
    const element = `<button id="category" onclick="fetchItems(${id})">${name}</button><br>`;
    content.innerHTML += element;
};
const loadCategories = (categories, content) => {
    for (const category of categories) {
        insertCategory(category, content);
    }
    const newCategorySection = `
    <input id="new" placeholder="Нова категория">
    <button onclick="addNewCategory()">ДОБАВИ</button>
    <p id="output"></p>`;
    content.innerHTML += newCategorySection;
    getById("new").focus();
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
