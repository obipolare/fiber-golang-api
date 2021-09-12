package main

import (
	"fmt"
	"reflect"

	"github.com/gofiber/fiber/v2"
)

type Movie struct {
	Title string `json:"title"`
	Id    int    `json:"id"`
}

func main() {

	app := fiber.New()

	// fiber.Map is a shorcut for a map in golang

	movies := []*Movie{
		{
			Title: "movie1",
			Id:    4,
		},
		{
			Title: "movie2",
			Id:    5,
		},
		{
			Title: "movie3",
			Id:    6,
		},
	}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"results": movies,
		})
	})

	app.Get("/:id", func(c *fiber.Ctx) error {

		id, err := c.ParamsInt("id")

		if err != nil {
			return c.Status(404).JSON(fiber.Map{
				"results": "Id not found",
			})
		}

		user := movies[id-1]

		fmt.Println(reflect.TypeOf(id))

		return c.JSON(fiber.Map{
			"results": user,
		})
	})

	type Request struct {
		Name string `json:"name"`
	}

	app.Post("/", func(c *fiber.Ctx) error {

		var body Request

		err := c.BodyParser(&body)

		if err != nil {
			return c.Status(404).JSON(fiber.Map{
				"Error": "There's a error, there's no movie",
			})
		}

		newUser := Movie{
			Title: body.Name,
			Id:    len(movies) + 1,
		}

		movies = append(movies, &newUser)

		return c.JSON(fiber.Map{
			"results": movies,
		})
	})

	app.Put("/:id", func(c *fiber.Ctx) error {

		id, err := c.ParamsInt("id")

		if err != nil {
			return c.Status(404).JSON(fiber.Map{
				"results": "error",
			})
		}

		var body *Request

		c.BodyParser(&body)

		var movie *Movie

		for _, m := range movies {
			if (*m).Id == id {
				fmt.Println("value of found movie", m)
				movie = m
				break
			}
		}

		fmt.Println("value of movie pointer", movie)
		(*movie).Title = (*body).Name

		return c.JSON(fiber.Map{
			"results": movies,
		})
	})

	app.Delete("/:id", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")

		if err != nil {
			return c.Status(404).JSON(fiber.Map{
				"results": "error",
			})
		}

		for i, m := range movies {
			if (*m).Id == id {
				movies = append(movies[:i], movies[i+1:]...)
				break
			}
		}

		return c.JSON(fiber.Map{
			"results": movies,
		})
	})

	err := app.Listen(":3000")

	if err != nil {
		panic("The port is not working")
	}
}
