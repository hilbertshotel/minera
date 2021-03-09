import bcrypt
import psycopg2

def main():
    password = b"R0m1K0l4"
    hash = bcrypt.hashpw(password, bcrypt.gensalt())
    hashed_password = hash.decode("UTF-8")
    
    CONNECTION_STRING = "dbname=minera_catalog user=postgres"

    try:
        connection = psycopg2.connect(CONNECTION_STRING)
    except:
        print("connection failure")
        return

    cursor = connection.cursor()

    query_string = "INSERT INTO login (password) VALUES (%s);"
    query_params = (password.decode("ascii"),)
    # COMMIT?

    try:
        cursor.execute(query_string, query_params)
    except:
        print("execution error")
        return

    print("ok")

main()

