### UserShip

Building a demo about REST API in Go with Gorilla Mux and PostgreSQL

#### Prerequisites
- go 1.13 or later
- postgreSQL 10 or later


#### Installation
1.download code
```shell script
git clone https://github.com/sdw2330976/usership.git
```
2.import sql into postgresSQL
```shell script
psql -d {database_name} -f {path of sql} {username}
```
note: use yourself name,then should modify `config/config.json`

3.build and run
```shell script
go build
./usership
```

