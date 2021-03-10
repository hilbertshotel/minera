import bcrypt
import psycopg2

CONNECTION_STRING = "dbname=minera_catalog user=postgres"

def main():
    password = b"asd"
    hashed_password = bcrypt.hashpw(password, bcrypt.gensalt()).decode("UTF-8")

    try:
        connection = psycopg2.connect(CONNECTION_STRING)
        cursor = connection.cursor()
    except:
        print("connection failure")
        return

    query_string = "INSERT INTO login (password) VALUES (%s);"
    query_params = (hashed_password,)

    try:
        cursor.execute(query_string, query_params)
        connection.commit()
    except:
        print("execution error")
        return

    print("ok")

main()
