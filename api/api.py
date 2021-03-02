from fastapi import FastAPI, Request
import uvicorn

app = FastAPI()

# add static files here

@app.get("/categories")
async def categories():
    return ["Категория 1", "Категория 2", "Категория 3", "Категория 4"]


@app.post("/items")
# @app.post("/{category}")
async def items(request: Request):
    data = await request.body()
    print(data.decode())
    return [{"Title": "Item 1",
        "Text": "some description",
        "Images": ["", "", ""]}]


# @app.get("/{id}")
# def get_id(id):
#     return {"id": id}


if __name__ == '__main__':
    uvicorn.run("api:app", host="127.0.0.1", port=5000, reload=True)
