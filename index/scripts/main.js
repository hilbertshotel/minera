"use strict";
const IP = "http://127.0.0.1";
// on load
fetchCategories();
// on event
const handleKey = (event) => {
    const k = event.key;
    if (event.ctrlKey && k == "F1") {
        console.log("asd");
        // login -> pre configured (hashed in DB)
        // if at index -> manage category buttons (new, edit, delete)
        // if at category -> manage item buttons (new, edit, delete)
        // -> add item name, add description, add 3 image paths (input file)
    }
};
window.addEventListener("keydown", handleKey);
