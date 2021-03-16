// INTERFACES
interface Item {
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

const newInput = (): HTMLInputElement => {
    return <HTMLInputElement>document.createElement("input")
}

const outputField = (): HTMLParagraphElement => {
    const output = document.createElement("p")
    output.id = "output"
    return output
}

const br = (): HTMLBRElement => {
    return document.createElement("br")
}

const hr = (): HTMLHRElement => {
    return document.createElement("hr")
}

const newTextArea = (): HTMLTextAreaElement => {
    return document.createElement("textarea")
}

const newButton = (name: string): HTMLButtonElement => {
    const button = document.createElement("button")
    button.innerHTML = name
    return button
}

const getInputValue = (id: string): string => {
    const input = <HTMLInputElement>document.getElementById(id)!
    return input.value
}

const addScript = (source: string) => {
    const script = document.createElement("script")
    script.src = source
    document.body.appendChild(script)
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