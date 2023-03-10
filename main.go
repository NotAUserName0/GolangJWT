package main

import (
	"GoJWT/databases"
	"GoJWT/helpers"
	"GoJWT/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func init(){ //starts before main
	helpers.LoadVariables()
	databases.Connection()
	helpers.MigrateDB()
}

func main(){
	app := fiber.New()

	app.Use(cors.New(cors.Config{
        AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin",
        AllowOrigins:     "*",
        AllowCredentials: true,
        AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
    }))

	routes.UserRoutes(app)

   app.Listen("localhost:3000")


}