import psycopg2
import errors

CONNECTION_STRING = "dbname=minera_catalog user=postgres"


# GET ITEMS
def get_items(category_id):
    try:
        connection = psycopg2.connect(CONNECTION_STRING)
    except Exception as err:
        errors.log("catalog.py", "get_items", "10", err)
        return
    
    cursor = connection.cursor()
    query_string = "SELECT name, description, img1, img2, img3 FROM items WHERE category_id = %s;"
    query_params = (category_id,)
    
    try:
        cursor.execute(query_string, query_params)
    except Exception as err:
        errors.log("catalog.py", "get_items", "20", err)
        return

    items = cursor.fetchall()
    response = []
    
    for item in items:
        name, description, *images = item
        item_object = {
            "name": name,
            "description": description,
            "images": [img for img in images if img != None]
        }
        response.append(item_object)

    connection.close()
    return response # [{String:String}]


# GET CATEGORIES
def get_categories():
    try:
        connection = psycopg2.connect(CONNECTION_STRING)
    except Exception as err:
        errors.log("catalog.py", "get_categories", "44", err)
        return

    cursor = connection.cursor()

    try:
        cursor.execute("SELECT id, name FROM categories;")
    except Exception as err:
        errors.log("catalog.py", "get_categories", "52", err)
        return

    response = cursor.fetchall()
    connection.close()
    return response # [(Int, String)]
