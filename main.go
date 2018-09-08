package main

import (
	"io/ioutil"
	"log"

	"github.com/go-redis/redis"
	"github.com/pkg/errors"
)

func run() error {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	script, err := ioutil.ReadFile("script.lua")
	if err != nil {
		return errors.WithStack(err)
	}

	sha, err := client.ScriptLoad(string(script)).Result()
	if err != nil {
		return err
	}

	for i := 0; i < 10; i++ {
		result, err := client.EvalSha(sha, []string{"test"}).Result()
		if err != nil {
			return err
		}
		log.Println(result)
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
