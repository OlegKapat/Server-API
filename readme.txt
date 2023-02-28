go get github.com/BurntSushi/toml
go get github.com/sirupsen/logrus
go get -u github.com/gorilla/mux
go get github.com/stretchr/testify
go get github.com/jmoiron/sqlx
go get github.com/lib/pq
go get github.com/golang-migrate/migrate
migrate create -ext sql -dir migrations create-users     
migrate -path migrations/ -database "postgres://postgres:1q2q3q4q@localhost:5432/restapi_dev?sslmode=disable" up


