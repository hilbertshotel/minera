"use strict";
const getById = (id) => {
    return document.getElementById(id);
};
const newElement = (type) => {
    return document.createElement(type);
};
const newImgElement = () => {
    return document.createElement("img");
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
