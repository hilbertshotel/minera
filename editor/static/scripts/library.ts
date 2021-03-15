const getById = (id: string): HTMLElement => {
    return document.getElementById(id)!
}

const getInputWithId = (id: string): HTMLInputElement => {
    return <HTMLInputElement>document.getElementById(id)!
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