"use strict";
const loadEditor = () => {
    const mainWindow = document.getElementById("main");
    while (mainWindow.firstChild) {
        mainWindow.removeChild(mainWindow.firstChild);
    }
    console.log("Editor loaded");
};
