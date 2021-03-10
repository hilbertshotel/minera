import psycopg2
import errors
import bcrypt

CONNECTION_STRING = "dbname=minera_catalog user=postgres"

login_attempts = 0

# VERIFY PASSWORD
def verify(password):
    global login_attempts
    login_attempts += 1

    if login_attempts > 10:
        return {"msg": "ПРЕВИШИЛИ СТЕ ОПИТИТЕ ЗА ДОСТЪП"}

    try:
        connection = psycopg2.connect(CONNECTION_STRING)
        cursor = connection.cursor()
    except Exception as err:
        errors.log("editor.py", "verify", "18", err)
        return

    try:
        cursor.execute("SELECT password FROM login;")
        hashed = cursor.fetchone()[0]
    except Exception as err:
        errors.log("editor.py", "verify", "25", err)
        return

    if bcrypt.checkpw(password.encode("UTF-8"), hashed.encode("UTF-8")):
        login_attempts = 0
        return {"msg": "ok"}
    return {"msg": "ГРЕШНА ПАРОЛА"}


# EDIT CATEGORY
def edit_category(new_name, old_name):
    print(new_name, old_name)
    # UPDATE categories SET name = new_name WHERE name = old_name;
    return new_name
