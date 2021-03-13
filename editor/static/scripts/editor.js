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
const addNewCategory = (output) => __awaiter(void 0, void 0, void 0, function* () {
    const newCategoryName = getInputWithId("new").value;
    if (!newCategoryName) {
        output.innerHTML = "ВЪВЕДЕТЕ ИМЕ";
        return;
    }
    const url = `${IP}/NewCategory`;
    const data = newPackage("POST", newCategoryName);
    const request = yield fetch(url, data);
    if (request.ok) {
        loadEditor();
    }
});
const insertCategory = (category, content) => {
    const button = newElement("button");
    button.id = "categoryButton";
    button.innerHTML = category.Name;
    button.onclick = () => { fetchItems(category.Id, content); };
    const br = newElement("br");
    content.appendChild(button);
    content.appendChild(br);
};
const loadCategories = (categories, content) => {
    for (const category of categories) {
        insertCategory(category, content);
    }
    const input = newInputElement();
    input.placeholder = "Нова категория";
    input.id = "new";
    const button = newElement("button");
    button.innerHTML = "ДОБАВИ";
    const output = newElement("p");
    output.id = "output";
    button.onclick = () => { addNewCategory(output); };
    content.appendChild(input);
    content.appendChild(button);
    content.appendChild(output);
    input.focus();
};
const fetchCategories = (content) => __awaiter(void 0, void 0, void 0, function* () {
    const url = `${IP}/LoadCategories`;
    const request = yield fetch(url);
    if (request.ok) {
        const categories = yield request.json();
        loadCategories(categories, content);
    }
});
const loadEditor = () => {
    const content = getById("content");
    clear(content);
    fetchCategories(content);
};
loadEditor();
