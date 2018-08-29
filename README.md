# go\_restful\_playground

### api server with jwt

#### go packages
```bash
go get -u golang.org/x/lint/golint
go get -u github.com/gin-gonic/gin
go get -u github.com/jinzhu/gorm
go get -u github.com/mattn/go-sqlite3
go get -u github.com/lib/pq
go get -u github.com/satori/go.uuid
go get -u github.com/dgrijalva/jwt-go
go get -u github.com/tsenart/vegeta
go run main.go
```
#### install docker on ubuntu
https://www.digitalocean.com/community/tutorials/how-to-install-and-use-docker-on-ubuntu-16-04
https://www.digitalocean.com/community/tutorials/how-to-install-docker-compose-on-ubuntu-16-04

#### dockerization
```bash
./create-binary.sh
docker-compose -f ./authserver-compose.yml up -d

cd benchmarks
./post.sh

docker container stop $(docker ps -a -q)
docker rm $(docker ps -a -q)
docker rmi $(docker images -q)
```

```bash
docker container cp auth-server-binary gorestfulplayground_auth_server_1:/
docker container restart gorestfulplayground_auth_server_1
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
