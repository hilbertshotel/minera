"use strict";
let loggedIn = () => {
    return false;
};
if (loggedIn()) {
    loadEditor();
}
else {
    openLogin();
}
