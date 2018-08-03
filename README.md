# go\_restful\_playground

### api server
```bash
go get -u golang.org/x/lint/golint
go get -u github.com/gin-gonic/gin
go get -u github.com/jinzhu/gorm
go get -u github.com/mattn/go-sqlite3
go get -u github.com/lib/pq
go get -u github.com/satori/go.uuid
go get -u github.com/tsenart/vegeta
go run main.go
```

#### dockerization
```bash
./create-binary.sh
docker-compose build
docker-compose up -d

cd benchmarks
./post.sh

docker container stop $(docker ps -a -q)
docker rm $(docker ps -a -q)
docker rmi $(docker images -q)
```

### scanner
```bash
go get -u golang.org/x/lint/golint
go get -u github.com/hashicorp/consul
go get -u github.com/denisenkom/go-mssqldb
go get -u github.com/go-sql-driver/mysql
go get -u github.com/gin-gonic/gin
go get -u github.com/glaslos/ssdeep
go get -u github.com/mitchellh/mapstructure
```

### A very good example of go api server
https://github.com/hugomd/go-todo
