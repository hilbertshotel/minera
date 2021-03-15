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
const verifyUser = (userData, output) => __awaiter(void 0, void 0, void 0, function* () {
    const data = newPackage("POST", userData);
    const request = yield fetch(`${IP}/VerifyUser`, data);
    if (request.ok) {
        const status = yield request.json();
        if (status !== "ok") {
            output.innerHTML = status;
            getInputWithId("pass").value = "";
            return;
        }
        // THIS IS A SECURITY RISK WHICH NEEDS TO BE FIXED
        addScript("scripts/editor.js");
    }
});
const validateInput = () => {
    const output = getById("output");
    output.innerHTML = "";
    const username = getInputWithId("user").value;
    const password = getInputWithId("pass").value;
    if (!username) {
        output.innerHTML = "МОЛЯ ВЪВЕДЕТЕ ПОТРЕБИТЕЛ";
        return;
    }
    if (!password) {
        output.innerHTML = "МОЛЯ ВЪВЕДЕТЕ ПАРОЛА";
        return;
    }
    const data = { "Username": username, "Password": password };
    verifyUser(data, output);
};
const openLogin = () => {
    const content = getById("content");
    const userInput = `<input id="user" maxlength="20" placeholder="Потребител">`;
    const passInput = `<input type="password" id="pass" maxlength="20" placeholder="Парола">`;
    const button = `<button onclick="validateInput()">ВХОД</button>`;
    const output = `<p id="output"></p>`;
    content.innerHTML += userInput;
    content.innerHTML += passInput;
    content.innerHTML += button;
    content.innerHTML += output;
    getById("user").focus();
};
// MAIN
const IP = "http://127.0.0.1:5252";
openLogin();
