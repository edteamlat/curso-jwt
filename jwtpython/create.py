import jwt
import datetime

now = datetime.datetime.now()
iat = int(now.timestamp())

tomorrow = now + datetime.timedelta(days=1)
expireAt = int(tomorrow.timestamp())

user = "Alexys Lozada"
email = "alexys.lozada@gmail.com"

secret = "this-is-a-very-long-and-secure-secret"

userdata = {
        "user": user,
        "email": email,
        "iat": iat,
        "exp": expireAt,
}

token = jwt.encode(userdata, secret, algorithm="HS256")

print(token)
