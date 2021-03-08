let loggedIn = (): boolean => {
    return false
}

if (loggedIn()) {
    loadEditor()
} else {
    openLogin()
}