"use strict";
// FUNCTIONS
const getById = (id) => {
    return document.getElementById(id);
};
const newElement = (type) => {
    return document.createElement(type);
};
const newImgElement = (src) => {
    const img = document.createElement("img");
    img.src = src;
    return img;
};
const newButton = (name) => {
    const button = document.createElement("button");
    button.innerHTML = name;
    return button;
};
const newPackage = (method, data) => {
    return {
        method: method,
        header: { "content-type": "application/json" },
        body: JSON.stringify(data)
    };
};
const clear = (element) => {
    while (element.firstChild) {
        element.removeChild(element.firstChild);
    }
};
const appendChildren = (children, parent) => {
    for (const child of children) {
        parent.appendChild(child);
    }
};
