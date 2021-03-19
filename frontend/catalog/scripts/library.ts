// INTERFACES
interface Item {
    Id: number,
    Name: string,
    Description: string,
    Images: string[]
}

interface Category {
    Id: number,
    Name: string
}


// FUNCTIONS
const getById = (id: string): HTMLElement => {
    return document.getElementById(id)!
}

const newElement = (type: string): HTMLElement => {
    return document.createElement(type)
}

const newImgElement = (src: string): HTMLImageElement => {
    const img = <HTMLImageElement>document.createElement("img")
    img.src = src
    return img
}

const newButton = (name: string): HTMLButtonElement => {
    const button = document.createElement("button")
    button.innerHTML = name
    return button
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

const appendChildren = (children: HTMLElement[], parent: HTMLElement) => {
    for (const child of children) {
        parent.appendChild(child)
    }
}
