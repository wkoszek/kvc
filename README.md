# RCLI

Easy to use Redis CLI that can be used from the command line mode.

## Build:
```
git clone https://github.com/ramin0x53/rcli.git
cd ./rcli
make build
```
## Docker:
```
docker build -t rcli-img .
docker run -d --name rcli-test rcli-img 
docker exec -it rcli-test bash
rcli configcl localhost:6379
```

## Supported commands
```
      configcl         rcli configcl [ip:port] [password] [db(integer)]
      set              rcli set key value
      get              rcli get key
      del              rcli del key
      lpop             rcli lpop key
      lpush            rcli lpush key value
      lrange           rcli lrange key start stop
      hmset            rcli hmset key field value
      hgetall          rcli hgetall key
      hget             rcli hget key field
      hdel             rcli hdel key field
```