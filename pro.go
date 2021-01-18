package main

import (
	"fmt"
	"strconv"
)

func main() {
	//app := fiber.New()
	//app.Get("/hello", func(ctx *fiber.Ctx) error {
	//	return ctx.SendString("world")
	//})
	//test := app.Group("/test")
	//{
	//	test.Get("/hello", func(ctx *fiber.Ctx) error {
	//		return ctx.JSON("test")
	//	})
	//}
	//
	//app.Listen(":3000")
	//fmt.Print(strconv.Itoa(0))
	//
	//fmt.Print(time.Now().AddDate(0, 0, -37).Format("20060102"))
	fmt.Println(len(""))
	publishIdInt, err := strconv.Atoi("")
	fmt.Println("err value:", err)
	fmt.Println("publishIdInt value:", publishIdInt)

	//ticker := time.NewTicker(time.Second * 3)
	////for range ticker.C {
	////	fmt.Println(22222)
	////}
	//for {
	//	select {
	//	case <-ticker.C:
	//		fmt.Print(11111)
	//	}
	//}

}
