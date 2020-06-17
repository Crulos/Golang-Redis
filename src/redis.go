package redis1

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

var rdb = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "", // no password set
	DB:       0,  // use default DB
})

type Person struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Age       int    `json:"ageName"`
}

func SetRedis() {
	value := []Person{
		{
			FirstName: "10",
			LastName:  "20",
			Age:       30,
		},
		{
			FirstName: "20",
			LastName:  "30",
			Age:       40,
		},
		{
			FirstName: "30",
			LastName:  "40",
			Age:       50,
		},
	}

	// encode struct to []byte  && set []byte to redis
	val, errVal := json.Marshal(value)
	if errVal != nil {
		panic(errVal)
	}

	err := rdb.Set(ctx, "key", val, 0).Err()
	if err != nil {
		panic(err)
	}

}

func GetRedis() {

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}

	fmt.Println("String Struct From Redis : ", string(val))

	// unmarshal and convert to []byte
	var res []Person
	json.Unmarshal([]byte(val), &res)
	fmt.Println("String Struct From []Byte : ", string([]byte(val)))
	fmt.Println("after decode and ready to use this :", res)
}
