# Project outline - provide easy to user redis features

## Redis data types

### String
#### Set
`set key value`: `set user:12:name skhammari`
#### Get
`get key`: `get user:12:name`

#### Increment
`incr key`: `incr product:1:quantity`
#### Increment by
`incrby key amount`: `incrby product:1:quantity -1`

Redis is atomic, so we can use these commands to prevent race condition

#### getset
`getset key value`: `getset product:1:quantity 10`

Sample usecase: for global locking in distributed systems

#### mget & mset
`mget post:12:title post:12:content post:13:title post:13:content`

`mset post:12:title "Hello World" post:12:content "Hello World content" post:13:title "Hello World" post:13:content "Hello World content"`

using these commands will reduce network latency as we can get the data in one request

### Hash

#### Hset
`hset key field value`: `hset page:20 title "Hello World" content "Hello World content"`
#### Hget
`hget key field`: `hget page:20 title`
#### hgetall
`hgetall key`: `hgetall page:20`
#### hmget
`hmget key field1 field2`: `hmget page:20 title content`
#### hmset
`hmset key field1 value1 field2 value2`: `hmset page:20 title "Hello World" content "Hello World content"`

### Set

#### Sadd
`sadd key value`: `sadd meat kebab pizza burger`
`sadd key value`: `sadd veg salad corn`
`sadd key value`: `sadd iranian kebab fesenjoon`
`sadd key value`: `sadd fastfood pizza burger`
`sadd key value`: `sadd healthy salad`
#### Sinter
`sinter key1 key2`: `sinter meat iranian`
`sinter key1 key2`: `sinter meat fastfood`
#### Sunion
`sunion key1 key2`: `sunion meat iranian`

#### Sinterstore
`sinterstore name key1 key2 `: `sinterstore iranian-meat meat iranian`

#### Sunionstore
`sunionstore name key1 key2 `: `sunionstore iranian-or-meat meat iranian`

#### Smembers
`smembers key`: `smembers iranian-or-meat`

#### Sismember
`sismember key value`: `sismember meat kebab`

### List

#### Lpush
`lpush key value`: `lpush myList 1 2 3`

#### Rpush
`rpush key value`: `rpush myList 4`

#### Lpop
`lpop key`: `lpop myList`

#### Rpop
`rpop key`: `rpop myList`

#### Lrange
`lrange key 0 -1`: `lrange myList 0 -1`

Sample:

Fan out

global redis: `lpush job:createImage '{}'`
```redis
BRPOP job:createImage 0
```

### Sorted Set

#### Zadd
`zadd key score value`

```redis
ZADD x 1 dave x 2 john x 2.2 sally x 3.2
```

#### Zrange
`zrange key 0 -1`: `zrange x 0 -1 WITHSCORES`

#### Zrevrange
`zrevrange key 0 -1`: `zrevrange x 0 -1 WITHSCORES`

Sample in go code:
https://gist.github.com/mhrlife/535c82e71deda8285fd3c9d574f65a42

### Stream

#### Xadd - produce
`xadd key value`: `xadd stream:1 * name dave age 25`
`xadd key value`: `xadd stream:1 * name john age 30`

#### Xrange - consume
`xrange key 0 -1`: `xrange stream:1 - + count 2`
`xrange key 0 -1`: `xrange stream:1 (1705174856520-0 + count 2`

#### xgroup create - create consumer group
`xgroup create key groupname id`: `xgroup create stream:1 group:saveDB $`

#### xreadgroup - read from consumer group
`xreadgroup group groupname id count 1 consumername key`: `XREADGROUP GROUP group:saveDB consumer:1 COUNT 5 STREAMS stream:1 >`

#### xreadgroup - read unaknowledged messages from consumer group
`xreadgroup group groupname id count 1 consumername key`: `XREADGROUP GROUP group:saveDB consumer:1 COUNT 5 STREAMS stream:1 0`

#### xreadgroup - read from consumer group with blocking
`xreadgroup group groupname id count 1 consumername key`: `XREADGROUP GROUP group:saveDB consumer:1 COUNT 5 BLOCK 100000 STREAMS stream:1 >`

#### xack - acknowledge
`xack key groupname id`: `xack stream:1 group:saveDB 1705174856520-0`

### Pipeline

