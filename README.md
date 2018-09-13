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
go get -u github.com/golang/protobuf/protoc-gen-go
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

```bash
http://192.168.7.140:8100/api/v1/contract?role=provider
http://192.168.7.140:8100/api/v0/ds/servers/
cd benchmarks
curl -i -H "Content-Type: application/json" --data @login.json http://192.168.7.140:8100/api/v1/login/
curl -i -H "Xauth: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJYcm9sZSI6ImFkbWluIiwiWHVzZXIiOiJoZWxpb3MifQ.lNA0CQiMmdF40rmwEpKFBmzTUYfhtaIwQiNuPNdIKc0" -H "Content-Type: application/json" http://192.168.7.140:8100/api/v0/ds/services
curl -i -H "Xauth: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJYcm9sZSI6InVzZXIiLCJYdXNlciI6InRpbSJ9.9JuBN2dHXgk_0krlqq-qc9s1fVXmJlzdUvZdkfUbme4" -H "Content-Type: application/json" http://192.168.7.140:8100/api/v0/ds/services
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
go get -u github.com/gocql/gocql
```

### svcmgr
```
go get -u github.com/satori/go.uuid
```

### pam
```bash
yum install pam-devel
```

### test
```bash
cd test
sudo GOPATH=$GOPATH $(which go) test -v
```
