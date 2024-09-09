front(id) -> routes.go(id) -> handler.go(id) -> service/task.go(id) -> db(Находит task по id) if success

db(task) -> service/task.go(task) -> handler.go(task) -> front(task)


Success status codes:
200 - OK 
201 - Created

Fail status codes:
- client errors
400 - BadRequest
401 - Unauthorized
404 - NotFound
- server error
500 - InternalServerError


// Atoi = ASCII to Int
// Itoa = Int to ASCII