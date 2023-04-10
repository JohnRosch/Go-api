package main

import "github.com/gofiber/fiber/v2"

//Dummy data
type Car struct {
	Id    int
	Name  string
	Model string
}

//Function
func getCar(c *fiber.Ctx) error {
	Apicar := []Car{
		{
			Id:    1,
			Name:  "hondaEsi",
			Model: "2010",
		},
		{
			Id:    2,
			Name:  "hondaHatchback",
			Model: "2001",
		},
		{
			Id:    3,
			Name:  "hondaCivic",
			Model: "1998",
		},
		{
			Id:    4,
			Name:  "honda",
			Model: "1999",
		},
	}
	return c.Status(fiber.StatusOK).JSON(Apicar)
}

func main() {
	app := fiber.New()
	//root route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Home page")
	})

	app.Get("/car", getCar)

	app.Listen(":3000")

}
