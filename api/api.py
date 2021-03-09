from fastapi import FastAPI, Request
import uvicorn
import catalog
import editor

app = FastAPI()


# CATALOG
@app.get("/categories")
async def categories():
    response = catalog.get_categories()
    return response


@app.post("/items")
async def items(request: Request):
    # client_IP = request.client.host
    category_id = await request.json()
    response = catalog.get_items(category_id)
    return response


# EDITOR
@app.post("/login")
async def login(request: Request):
    password = await request.json()
    response = editor.verify(password)
    return response


if __name__ == '__main__':
    uvicorn.run("api:app", host="127.0.0.1", port=5000, reload=True)
