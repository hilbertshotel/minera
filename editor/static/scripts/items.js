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
const loadItems = (items, content) => {
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
const fetchItems = (id, content) => __awaiter(void 0, void 0, void 0, function* () {
    const url = `${IP}/LoadItems`;
    const data = newPackage("POST", id);
    const request = yield fetch(url, data);
    if (request.ok) {
        const items = yield request.json();
        loadItems(items, content);
    }
});
