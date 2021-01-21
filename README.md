# Simple Echo Framework App Using Facebook ent Framework with Casbin

Simple App with echo v4 Framework, Facebook ent and Casbin

## Installation:

Go to Project Folder then run the following command :

```
go generate ./ent
```

## Running Project Command:
After generating ent schema resources, run the following command

```
go run .
```

## EndPoints Instructions:

```
First, you have to send a request to login endpoint to generate a new
JWT Token which will be used later for authenticating yourself to other
endpoints.

```

## User Roles:

```
UserA -> has access to EndPoint /app/hello/{name} only
UserB -> has access to both EndPoints
```

## Simple Request:

```
curl -d '{"username":"mostafa", "password":"testtest"}' -H "Content-Type: application/json" -X POST http://localhost:8080/app/login
```

## Credits:

I've used [casbinrest adapter](https://github.com/prongbang/casbinrest) (with some updates to work with echo v4) for integrating echo with casbin