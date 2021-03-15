"use strict";
const getById = (id) => {
    return document.getElementById(id);
};
const getInputWithId = (id) => {
    return document.getElementById(id);
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
