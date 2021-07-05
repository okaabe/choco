# Choco server-side
That's the choco's server-side code which was writen in golang

## Info
Here it is all the information about what you need to know to use this api

## Addresses
The addresses of the rest-api

### /api/auth/signin
Route to authenticate the user, you need to write on the body of the http request the json below

```json
{
    email: "the user's email"
    password: "the user's password"
}
```

The server will return a json that has the property err that will inform if everything happened as planned or if ocurred some error

```json
{
    err: "Something rlly wrong happened"
}
```

But if everything happend as planned, the json will have the property "data", that will be the jwt:

```json
{
    err: null,
    data: "jwt here"
}
```

### /auth/auth/signup
Route to create the user's account, necessary the following body:

```json
{
    username: "user's username",
    email: "user's email",
    password: "user's password",
}
```

the response body of a successful operation will be:

```json
{
    err: null,
    data: "jwt"
}
```

and about a failed operation:

```json
{
    err: "Something rlly happened",
}
```
