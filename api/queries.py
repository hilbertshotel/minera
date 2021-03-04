import psycopg2

CONNECTION_STRING = "dbname=minera_catalog user=postgres"

# get_items :: Int -> IO [{String}]
def get_items(category_id):
    connection = psycopg2.connect(CONNECTION_STRING) # try clause for connection
    cursor = connection.cursor()
    
    # get items from database
    query_string = "SELECT name, description, img1, img2, img3 FROM items WHERE category_id = %s;"
    query_params = (category_id,)
    cursor.execute(query_string, query_params) # try clause for execute
    items = cursor.fetchall()

    # load items into response array
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
    return response


# get_categories :: IO [(Int, String)]
def get_categories():
    connection = psycopg2.connect(CONNECTION_STRING) # try clause for connection
    cursor = connection.cursor()

    cursor.execute("SELECT id, name FROM categories;") # try clause for connection
    response = cursor.fetchall()

    connection.close()
    return response
