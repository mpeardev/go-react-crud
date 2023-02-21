// REST API
package main // requerido

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
) // formateador de strings, mostrar msjs en consola

func main(){
port :=	os.Getenv("PORT")

if port == ""{
	port = "3000"
}

	app := fiber.New() // := ahorra declarar el tipo de variable como var

	app.Use(cors.New())

	app.Static("/", "./client/dist")

	app.Get("/users", func (c *fiber.Ctx) error  {
		return c.JSON(&fiber.Map{
			"data": "usuarios desde el backend",
		})
	})

	app.Listen(":" + port)  // creando servidor
	fmt.Println("Server on port 3000")
}