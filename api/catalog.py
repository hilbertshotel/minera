import psycopg2
import errors

CONNECTION_STRING = "dbname=minera_catalog user=postgres"


# GET ITEMS
def get_items(category_id):
    try:
        connection = psycopg2.connect(CONNECTION_STRING)
        cursor = connection.cursor()
    except Exception as err:
        errors.log("catalog.py", "get_items", "10", err)
        return
    
    query_string = "SELECT name, description, img1, img2, img3 FROM items WHERE category_id = %s;"
    query_params = (category_id,)
    
    try:
        cursor.execute(query_string, query_params)
        items = cursor.fetchall()
    except Exception as err:
        errors.log("catalog.py", "get_items", "20", err)
        return

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
        cursor = connection.cursor()
    except Exception as err:
        errors.log("catalog.py", "get_categories", "43", err)
        return

    try:
        cursor.execute("SELECT id, name FROM categories;")
        response = cursor.fetchall()
    except Exception as err:
        errors.log("catalog.py", "get_categories", "50", err)
        return

    connection.close()
    return response # [(Int, String)]
