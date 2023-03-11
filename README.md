## Github view counter
Golang view counter for your github page

## Deployment
```
go mod download
```

Create configs/server.toml and redis.conf if you user docker. Fill them out following the examples

```
make run
```

## Using Docker
```
make compose-start
```

Then just go to http://127.0.0.1:8080/counter/?name={YourName}

## Run tests
```
make test
```