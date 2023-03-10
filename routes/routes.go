package routes

import (
	"GoJWT/controllers"

	"os"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
)


func UserRoutes(app *fiber.App) {

    //without protection

    app.Post("/login", controllers.Login) //si hay token redir /, si no hay procede a el registro

	app.Post("/signup",controllers.Signup) //si hay token redir /, " " " "

    app.Get("/free",func(c *fiber.Ctx) error { //este llevara la proteccion y solo si hay token puede acceder
        return c.SendString("free works!")
    }) 

    //bearer protection
    app.Use(jwtware.New(jwtware.Config{ //recibe el token
		SigningKey: []byte(os.Getenv("SECRET")),
	}))

    app.Get("/",func(c *fiber.Ctx) error { //este llevara la proteccion y solo si hay token puede acceder
        user:=c.Locals("user").(*jwt.Token)
        claims := user.Claims.(jwt.MapClaims)
        nombre := claims["user"].(string)
        return c.SendString(nombre)
    })

    

	

}