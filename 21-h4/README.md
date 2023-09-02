We start app:

```shell
$ go run ./cmd/web 
```
To test app:

```shell
$ curl http://localhost:8080/list
shoes: $50.00
socks: $5.00

$ curl localhost:8080/create?item=ties\&price=13
added ties with price $13.00

$ curl http://localhost:8080/read?item=socks
item socks has price $5.00

$ curl http://localhost:8080/update?item=socks\&price=6
new price $6.00 for socks

$ curl http://localhost:8080/delete?item=shose
dropped shoes

$ kill %1
[1]+  Terminated: 15          go run ./part1
```

Note that in bash we must escape the ampersand `\&` so the shell doesn't think we're asking to run a background process.
