const createCategory = (file: string, content: HTMLElement) => {
  const category = document.createElement("div")
  category.className = "category"
  category.innerHTML = file
  category.onclick = () => { fetchItems(file) }
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
  for (const file of categories) {
    createCategory(file, content)
  }
}


const fetchCategories = async () => {
  console.log("1")
  const request = await fetch(`${IP}/categories`)
  if (request.ok) {
    const data = await request.json()
    loadCategories(data)
  }
}
