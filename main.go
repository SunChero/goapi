package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gomodule/redigo/redis"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func getKey(c echo.Context) error {
	conn, err := redis.Dial("tcp", ":6379")
	if err != nil {
		log.Fatalf("Could not connect: %v\n", err)
	}
	defer conn.Close()
	name := c.Param("key")
	value, _ := redis.StringMap(conn.Do("HGETALL", name))
	js, _ := json.Marshal(value)
	fmt.Println(value)
	return c.String(200, "Got it from db "+name+"and Value is "+string(js))
}

func main() {

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/db/:key", getKey)
	e.Logger.Fatal(e.Start(":3000"))
}
