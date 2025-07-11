import base64
import hashlib
import jwt

SALT = b'7qauRuO3QGyJmiBDd9aUaw'
SECRET = "ijz25f8z14s2fs5f95W6F58F5G8Q55df6q2fz448z"
algorithm = 'HS256'

def hash_password(password):
    t_sha = hashlib.sha512()
    t_sha.update(password.encode('utf-8') + SALT)
    hashed_password = base64.urlsafe_b64encode(t_sha.digest())
    return str(hashed_password.decode('utf-8'))

def generate_jwt(username):
    payload = {"username": username}
    token = jwt.encode(payload, SECRET, algorithm)
    return token

def get_username_from_jwt(token):
    try:
        data = jwt.decode(token, key=SECRET, algorithms=algorithm)
        return data
    except:
        return None