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
const verifyPassword = (password) => __awaiter(void 0, void 0, void 0, function* () {
});
const validateInput = (output) => {
    const password = document.getElementById("password").value;
    if (!password) {
        output.innerHTML = "НЕ СТЕ ВЪВЕЛИ ПАРОЛА";
    }
    else {
        verifyPassword(password);
    }
};
const addLoginField = (mainWindow) => {
    const div = document.createElement("div");
    div.className = "login";
    const input = document.createElement("input");
    input.type = "password";
    input.id = "password";
    input.placeholder = "Парола";
    const button = document.createElement("button");
    button.innerHTML = "ВХОД";
    const p = document.createElement("p");
    p.id = "output";
    button.onclick = () => { validateInput(p); };
    div.appendChild(input);
    div.appendChild(button);
    div.appendChild(p);
    mainWindow.appendChild(div);
    input.focus();
};
const openLogin = () => {
    const mainWindow = document.getElementById("main");
    addLoginField(mainWindow);
};
