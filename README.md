## Commands

#### sqlc

docker-compose run --rm sqlc generate

#### golang-migrate

docker-compose run --rm golang-migrate

#### exec goapp bash

docker-compose exec goapp bash

---

comando de migração do banco de dados do Golang

migrate -source file:///go/app/sql/migrations -database 'mysql://root:root@tcp(mysql:3306)/cartola' up
