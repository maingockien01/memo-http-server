# Overview
This project is solely for academic purpose. 
Im still owrking on my project due to many bugs still on my applications. 

# Architecture
I implemented my webserver based on serveral layers:
- `tcp` : Transportation layer which is in charge of handling sockets.
- `http` : Middleware layer which is in charge of handling HTTP protocol.
- `application` : Application layer which is in charge of handling data and get data.
- `filesystem` : Database layer which is in charge of handling data access to database.

# Web server
- My webserver currently is able to serve static file based on the link request.
- My webserver can server some API path:
    - GET /api/memo
    - POST /api/memo
    - DELETE /api/memeo/{memoId}
    - UPDATE /api/memo

# Command
- build server: `make build`
- run server: `make runServer`

# Testing
- my computer is running: Manjaro OS and I test my website with Brave Browser (which is almost identical to Chrome)

