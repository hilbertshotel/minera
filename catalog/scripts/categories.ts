const createCategory = (name: string, content: HTMLElement) => {
  const category = document.createElement("div")
  category.className = "category"
  category.innerHTML = name
  category.onclick = () => { fetchItems(name) }
  content.appendChild(category)
}


const loadCategories = (categories: string[]) => {
  const content = document.getElementById("content")!

  // remove back button
  document.getElementById("button")?.remove()

  // clear content div
  while(content.firstChild){
    content.removeChild(content.firstChild)
  }

  // list all categories
  for (const name of categories) {
    createCategory(name, content)
  }
}


const fetchCategories = async () => {
  const request = await fetch(`${IP}/categories`)
  if (request.ok) {
    const data = await request.json()
    loadCategories(data)
  }
}
