import psycopg2

CONNECTION_STRING = "dbname=minera_catalog user=postgres"

# get_items :: String -> IO [{String}]
def get_items(category):
    connection = psycopg2.connect(CONNECTION_STRING) # try clause for connection
    cursor = connection.cursor()
    
    # get category id from database
    query_string = "SELECT id FROM categories WHERE name = %s;"
    query_params = (category,)
    cursor.execute(query_string, query_params) # try clause for execute
    category_id = cursor.fetchone()[0]
    
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


# get_categories :: [String]
def get_categories():
    connection = psycopg2.connect(CONNECTION_STRING) # try clause for connection
    cursor = connection.cursor()

    cursor.execute("SELECT name FROM categories;") # try clause for connection
    categories = [x[0] for x in cursor.fetchall()]
    connection.close()
    return categories
