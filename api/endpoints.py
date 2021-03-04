from fastapi import FastAPI, Request
import uvicorn
import queries

app = FastAPI()


@app.get("/categories")
async def categories():
    response = queries.get_categories()
    return response


@app.post("/items")
async def items(request: Request):
    # client_IP = request.client.host
    category_id = await request.json()
    response = queries.get_items(category_id)
    return response


if __name__ == '__main__':
    uvicorn.run("endpoints:app", host="127.0.0.1", port=5000, reload=True)
