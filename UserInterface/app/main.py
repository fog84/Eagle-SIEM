from fastapi import FastAPI, Response, Request, HTTPException
from fastapi.responses import FileResponse
from fastapi.staticfiles import StaticFiles
from pathlib import Path
import os

import classes
import auth
import bdd

bdd.init()
app = FastAPI()

static_dir = os.path.join(os.path.dirname(__file__), "static")
app.mount("/static", StaticFiles(directory=static_dir), name="static")

@app.get("/login_ui")
def login_ui():
    return FileResponse(os.path.join(static_dir, "login.html"))

@app.get("/siem_ui")
def siem_ui(request: Request):
    token = request.cookies.get('token')
    username = auth.get_username_from_jwt(token)

    if username is None:
        raise HTTPException(status_code=401, detail="Unauthorized")

    return FileResponse(os.path.join(static_dir, "siem.html"))

@app.get("/save_indexer_api_key_readlogs_ui")
def siem_ui(request: Request):
    token = request.cookies.get('token')
    username = auth.get_username_from_jwt(token)

    if username is None:
        raise HTTPException(status_code=401, detail="Unauthorized")

    return FileResponse(os.path.join(static_dir, "add_api_key_readlogs_in_cookie.html"))

@app.get("/whoami")
def whoami(request: Request):
    token = request.cookies.get('token')
    return auth.get_username_from_jwt(token)

@app.get("/save_indexer_api_key_readlogs_in_cookie")
def save_indexer_api_key_readlogs_in_cookie(request: Request, api_key_readlogs: str, response: Response):
    token = request.cookies.get('token')
    username = auth.get_username_from_jwt(token)

    if username is None:
        raise HTTPException(status_code=401, detail="Unauthorized")
    
    response.set_cookie(key="indexer_api_key_readlogs", value=api_key_readlogs, httponly=False) # httponly would be better
    return api_key_readlogs

@app.post("/login")
async def login(login_form: classes.LoginForm, response: Response):
    username = login_form.username
    password = auth.hash_password(login_form.password)
    bdd.cursor.execute("SELECT username FROM users WHERE username=? AND password=?", (username, password))
    row = bdd.cursor.fetchone()
    if row:
        response.set_cookie(key="token", value=auth.generate_jwt(username), httponly=True)
        return "Connected"
    else:
        return "Bad password or username"