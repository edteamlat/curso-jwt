import jwt

token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyIjoiQWxleHlzIExvemFkYSIsImVtYWlsIjoiYWxleHlzLmxvemFkYUBnbWFpbC5jb20iLCJpYXQiOjE2OTY2OTMxMTAsImV4cCI6MTY5NjY5MzExMX0.MfbvwKVf2INWr-pReB6JD0ixntBB4Jgvl8etVQ2WqTM"

secret = "this-is-a-very-long-and-secure-secret"

try:
    payload = jwt.decode(token, secret, algorithms=["HS256"])
except jwt.exceptions.ExpiredSignatureError:
    print("el token ha expirado")
except jwt.exceptions.InvalidSignatureError:
    print("la firma no es válida")
except jwt.exceptions.DecodeError:
    print("el token no es válido")
else:
    print("Bienvenido!")

