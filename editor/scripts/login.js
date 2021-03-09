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
const verifyPassword = (password, output) => __awaiter(void 0, void 0, void 0, function* () {
    const body = {
        method: "POST",
        header: { "content-type": "application/json" },
        body: JSON.stringify(password)
    };
    const request = yield fetch(`${IP}/login`, body);
    if (request.ok) {
        const status = yield request.json();
        if (status.msg === "error") {
            output.innerHTML = "ГРЕШНА ПАРОЛА";
        }
        else {
            loadEditor();
        }
    }
});
const validateInput = (output) => {
    const password = document.getElementById("password").value;
    if (!password) {
        output.innerHTML = "МОЛЯ ВЪВЕДЕТЕ ПАРОЛА";
    }
    else {
        verifyPassword(password, output);
    }
};
const openLogin = () => {
    const mainWindow = document.getElementById("main");
    const div = document.createElement("div");
    div.className = "login";
    const input = document.createElement("input");
    input.type = "password";
    input.id = "password";
    input.maxLength = 20;
    input.placeholder = "Парола";
    const button = document.createElement("button");
    button.innerHTML = "ВХОД";
    const p = document.createElement("p");
    button.onclick = () => { validateInput(p); };
    div.appendChild(input);
    div.appendChild(button);
    div.appendChild(p);
    mainWindow.appendChild(div);
    input.focus();
};
