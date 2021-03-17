"use strict";
// FUNCTIONS
const getById = (id) => {
    return document.getElementById(id);
};
const newInput = () => {
    return document.createElement("input");
};
const newElement = (type) => {
    return document.createElement(type);
};
const outputField = () => {
    const output = document.createElement("p");
    output.id = "output";
    return output;
};
const br = () => {
    return document.createElement("br");
};
const hr = () => {
    return document.createElement("hr");
};
const newTextArea = () => {
    return document.createElement("textarea");
};
const newButton = (name) => {
    const button = document.createElement("button");
    button.innerHTML = name;
    return button;
};
const getInputValue = (id) => {
    const input = document.getElementById(id);
    return input.value;
};
const addScript = (source) => {
    const script = document.createElement("script");
    script.src = source;
    document.body.appendChild(script);
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
