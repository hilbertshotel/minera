"use strict";
const getById = (id) => {
    return document.getElementById(id);
};
const getInputWithId = (id) => {
    return document.getElementById(id);
};
const newElement = (type) => {
    return document.createElement(type);
};
const newInputElement = () => {
    return document.createElement("input");
};
const newScriptElement = () => {
    return document.createElement("script");
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
