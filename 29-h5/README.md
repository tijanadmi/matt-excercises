We start app:

```shell
$ cd api 
$ go run -race . 
```
To test app we can do:

```shell
$ cd test
$ go run main.go

got item=shoes&price=46 = 200 (no err)
got item=shoes&price=46 = 400 (no err)
got item=shoes = 200 (no err)
got item=socks&price=6 = 200 (no err)
got item=socks&price=6 = 400 (no err)
got item=socks = 200 (no err)
got item=sandals = 404 (no err)
got item=sandals&price=27 = 404 (no err)
got item=sandals&price=27 = 200 (no err)
got item=clogs&price=36 = 404 (no err)
got item=clogs = 404 (no err)
got item=clogs&price=36 = 200 (no err)
got item=pants&price=30 = 200 (no err)
got item=pants = 200 (no err)
got item=pants&price=30 = 404 (no err)
got item=shorts = 404 (no err)
got item=shorts&price=20 = 200 (no err)
got item=shorts&price=20 = 200 (no err)
got item=shoes = 404 (no err)
got item=shoes&price=46 = 404 (no err)
got item=shoes&price=46 = 200 (no err)
got item=socks = 404 (no err)
got item=socks&price=6 = 404 (no err)
got item=socks&price=6 = 200 (no err)
got item=sandals&price=27 = 200 (no err)
got item=sandals = 200 (no err)
got item=sandals&price=27 = 200 (no err)
got item=clogs = 200 (no err)
got item=clogs&price=36 = 404 (no err)
got item=clogs&price=36 = 200 (no err)
got item=pants&price=30 = 200 (no err)
got item=pants&price=30 = 200 (no err)
got item=pants = 200 (no err)
got item=shorts = 200 (no err)
got item=shorts&price=20 = 404 (no err)
got item=shorts&price=20 = 200 (no err)
got item=shoes&price=46 = 200 (no err)
got item=shoes = 200 (no err)
got item=shoes&price=46 = 200 (no err)
got item=socks&price=6 = 200 (no err)
got item=socks = 200 (no err)
got item=socks&price=6 = 200 (no err)
got item=sandals = 200 (no err)
got item=sandals&price=27 = 200 (no err)
got item=sandals&price=27 = 200 (no err)
got item=clogs = 200 (no err)
got item=clogs&price=36 = 200 (no err)
got item=clogs&price=36 = 200 (no err)
```


