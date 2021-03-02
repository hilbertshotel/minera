from fastapi import FastAPI, Request
import uvicorn

app = FastAPI()


@app.get("/categories")
def categories():
    return ["Категория 1", "Категория 2", "Категория 3", "Категория 4"]


@app.post("/items")
def items(request: Request):
    data = request.body()
    print(data.decode())
    return [{"Title": "Item 1",
        "Text": "some description",
        "Images": ["", "", ""]}]


# @app.get("/{id}")
# def get_id(id):
#     return {"id": id}


if __name__ == '__main__':
    uvicorn.run("main:app", host="127.0.0.1", port=5000, reload=True)
