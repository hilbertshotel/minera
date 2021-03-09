import threading
from datetime import datetime

lock = threading.Lock()

def log(file, function, line, err):
    global lock
    lock.acquire()

    now = datetime.now()

    with open("logs/error.log", "a") as f:
        f.write(
            "________________________________________"
            f"\n{now.strftime('%d %B %y - %H:%M:%S')}\n"
            f"file name: {file}\n"
            f"function name: {function}\n"
            f"line number: {line}\n"
            f"{err}"
        )
    lock.release()
