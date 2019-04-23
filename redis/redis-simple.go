package main

import (
	"fmt"
	"github.com/go-redis/redis"
)


func main(){
	client := redis.NewClient(&redis.Options{
		Addr: 		"localhost:6379",
		Password: 	"", //no password set
		DB:			0, //use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	//set key value
	err = client.Set("key", "value", 0).Err()
	if err != nil{
		panic(err)
	}
	//get key
	val, err := client.Get("key").Result()
	if err != nil{
		panic(err)
	}
	fmt.Println("get key,",val)

	//get key which is not exist
	val2, err := client.Get("key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
}

//output
/*

root@wxtest047:/home/cxy/redis/code/demos# ./redis-demo
PONG <nil>
get key, value
key2 does not exist


 */