const getById = (id: string): HTMLElement => {
    return document.getElementById(id)!
}

const newElement = (type: string): HTMLElement => {
    return document.createElement(type)
}

const newPackage = <T>(method: string, data: T): Object => {
    return {
        method: method,
        header: {"content-type": "application/json"},
        body: JSON.stringify(data)
    }
}

const clear = (element: HTMLElement) => {
    while(element.firstChild) {
        element.removeChild(element.firstChild)
    }
}