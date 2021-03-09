import psycopg2
import errors
import bcrypt

CONNECTION_STRING = "dbname=minera_catalog user=postgres"

# VERIFY PASSWORD
def verify(password):
    try:
        connection = psycopg2.connect(CONNECTION_STRING)
    except Exception as err:
        errors.log("editor.py", "verify", "9", err)
        return

    cursor = connection.cursor()

    try:
        cursor.execute("SELECT password FROM login;")
    except Exception as err:
        errors.log("editor.py", "verify", "17", err)
        return

    # response = cursor.fetchall()
    print(cursor.fetchone())

    # if bcrypt.checkpw(password, hashed):
    #     print("It Matches!")
    # else:
    #     print("It Does not Match :(")

    if password == "asd":
        return {"msg": "ok"}
    return {"msg": "error"}
    # verify password against hash in database
    # if correct return session id
    # else return error "incorrect password"

    # MAX 5 PASSWORD ATTEMPTS FOR 30 MINUTES
    # TRY AND CREATE A SESSION ID with a time limit