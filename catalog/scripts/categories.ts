const createCategory = (category: [number, string], content: HTMLElement) => {
  const [id, name] = category;
  const div = document.createElement("div")
  div.className = "category"
  // div.id = `${id}`
  div.innerHTML = name
  div.onclick = () => { fetchItems(id) }
  content.appendChild(div)
}


const loadCategories = (categories: [[number, string]]) => {
  const content = document.getElementById("content")!

  // remove back button
  document.getElementById("button")?.remove()

  // clear content div
  while(content.firstChild){
    content.removeChild(content.firstChild)
  }

  // list all categories
  for (const category of categories) {
    createCategory(category, content)
  }
}


const fetchCategories = async () => {
  const request = await fetch(`${IP}/categories`)
  if (request.ok) {
    const data = await request.json()
    loadCategories(data)
  }
}
