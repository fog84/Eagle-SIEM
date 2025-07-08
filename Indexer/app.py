from fastapi import FastAPI, Depends
from sqlalchemy.orm import Session
from sqlalchemy import text
from database import SessionLocal

app = FastAPI()

def get_db():
    db = SessionLocal()
    try:
        yield db
    finally:
        db.close()

@app.get("/events")
def get_events(db: Session = Depends(get_db)):
    result = db.execute(text("SELECT * FROM events"))
    rows = result.fetchall()
    return [{"id": row[0], "log": row[1]} for row in rows]
