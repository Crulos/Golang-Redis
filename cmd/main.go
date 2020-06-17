package main

import (
	redis "god/src"
)

func main() {

	redis.SetRedis()
	redis.GetRedis()
}
