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
const VerifyUser = (userData, output) => __awaiter(void 0, void 0, void 0, function* () {
    const url = `${IP}/VerifyUser`;
    const data = newPackage("POST", userData);
    const request = yield fetch(url, data);
    if (request.ok) {
        const status = yield request.json();
        if (status !== "ok") {
            output.innerHTML = status;
            getInputWithId("password").value = "";
            return;
        }
        const script = newScriptElement();
        script.src = "scripts/editor.js";
        document.body.appendChild(script);
    }
});
const validateInput = (output) => {
    output.innerHTML = "";
    const username = getInputWithId("username").value;
    const password = getInputWithId("password").value;
    if (!username) {
        output.innerHTML = "МОЛЯ ВЪВЕДЕТЕ ПОТРЕБИТЕЛ";
        return;
    }
    if (!password) {
        output.innerHTML = "МОЛЯ ВЪВЕДЕТЕ ПАРОЛА";
        return;
    }
    const data = { "Username": username, "Password": password };
    VerifyUser(data, output);
};
const openLogin = () => {
    const content = getById("content");
    const usernameInput = newInputElement();
    usernameInput.type = "username";
    usernameInput.id = "username";
    usernameInput.maxLength = 20;
    usernameInput.placeholder = "Потребител";
    const passwordInput = newInputElement();
    passwordInput.type = "password";
    passwordInput.id = "password";
    passwordInput.maxLength = 20;
    passwordInput.placeholder = "Парола";
    const button = newElement("button");
    button.innerHTML = "ВХОД";
    const output = newElement("p");
    output.id = "output";
    button.onclick = () => { validateInput(output); };
    content.appendChild(usernameInput);
    content.appendChild(passwordInput);
    content.appendChild(button);
    content.appendChild(output);
    usernameInput.focus();
};
