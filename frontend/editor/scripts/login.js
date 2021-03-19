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
// LOGIN
const verifyUser = (url, userData, output) => __awaiter(void 0, void 0, void 0, function* () {
    const data = newPackage("POST", userData);
    const request = yield fetch(url, data);
    if (request.ok) {
        const status = yield request.json();
        if (status !== "ok") {
            output.innerHTML = status;
            return;
        }
        addScript("scripts/editor.js");
    }
});
const validateInput = (output) => {
    output.innerHTML = "";
    const username = getInputValue("username");
    if (!username) {
        output.innerHTML = "МОЛЯ ВЪВЕДЕТЕ ПОТРЕБИТЕЛ";
        return;
    }
    const password = getInputValue("password");
    if (!password) {
        output.innerHTML = "МОЛЯ ВЪВЕДЕТЕ ПАРОЛА";
        return;
    }
    const data = { "Username": username, "Password": password };
    verifyUser(`${IP}/Authentication`, data, output);
};
const openLogin = () => {
    const content = getById("content");
    const username = newInput();
    username.id = "username";
    username.maxLength = 20;
    username.placeholder = "Потребител";
    const password = newInput();
    password.type = "password";
    password.id = "password";
    password.maxLength = 20;
    password.placeholder = "Парола";
    const enter = newButton("ВХОД");
    const output = outputField();
    enter.onclick = () => { validateInput(output); };
    appendChildren([username, password, enter, output], content);
    username.focus();
};
// MAIN
const IP = "http://127.0.0.1:5252";
openLogin();
