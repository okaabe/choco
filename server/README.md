# Choco API
That's the api of the choco, where the authentication and content management are done, if you are interested on it, check more information about it below

## Routes
- [x] `/api/auth/signin` <strong>POST</strong>
- [x] `/api/auth/signup` <strong>POST</strong>
- [x] `/api/auth/rewoke` <strong>GET</strong>

- [ ] `/api/content/user/communities` <strong>GET</strong>


- [ ] `/api/content/community` <strong>POST</strong>
- [ ] `/api/content/community/:id` <strong>GET</strong>
- [ ] `/api/content/community/:id/posts` <strong>POST</strong>
- [ ] `/api/content/community/:id/posts` <strong>GET</strong>
- [ ] `/api/content/community/:id/posts/:postId` <strong>GET</strong>

## Stack
On the server-side of the choco, i have used:

- Golang

- Docker
- Docker Compose

- Gin Gonic (Http Framework)

- Postgresql(production db) and SQlite(db to run the tests)
- Gorm (as sql orm)

- TDD (Test driven development)
- DDD (Domain driven design)
- DOD (Data oriented design)


## Links
- http://jamesmcm.github.io/blog/2020/07/25/intro-dod/
- https://airbrake.io/blog/software-design/domain-driven-design
- https://en.wikipedia.org/wiki/Test-driven_development
- https://gorm.io
- https://github.com/gin-gonic/gin
- https://www.docker.com/
- https://docs.docker.com/compose/
- https://golang.org/
