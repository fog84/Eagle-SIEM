import sqlite3
import os

import auth

connexion_bdd = sqlite3.connect('users.db', check_same_thread=False)
cursor = connexion_bdd.cursor()

def init():
    if os.path.getsize('users.db') == 0:
        cursor.execute("""
        CREATE TABLE IF NOT EXISTS users(
            id INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE,
            username TEXT,
            password TEXT
        )
        """)
        cursor.execute("""
            INSERT INTO users(username, password)VALUES(?, ?)
            """, ("admin", auth.hash_password("dapasswordtochange")))
        connexion_bdd.commit()