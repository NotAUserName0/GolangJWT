package controllers

import (
	"GoJWT/databases"
	"GoJWT/helpers"
	"GoJWT/models"

	"golang.org/x/crypto/bcrypt"

	"github.com/gofiber/fiber/v2"
)

func Signup(c *fiber.Ctx) error {

	newUser := new(models.User)

	err := c.BodyParser(newUser)

	if err != nil {
		return c.Status(400).JSON(&fiber.Map{
			"success":"false",
			"message": err.Error(),
		})
	}

	/*
      Password check min lenght 8, min caracter mayus and minus, min 1 special character
	*/

	hash, err := bcrypt.GenerateFromPassword([]byte(newUser.Password),10)

	if err != nil {
		return c.Status(501).JSON(&fiber.Map{
			"success":"imposible to hash password",
		})
	}

	newUser.Password = string(hash);

	result := databases.Connection().Select("nombre", "apellido", "edad","genero","email","password","rol").Create(&newUser)

	if result.Error != nil {
		return c.Status(501).JSON(&fiber.Map{
			"success":"imposible to create user",
			"message": result.Error.Error(),
		})
	}

	return c.Status(200).JSON(&fiber.Map{
		"success":"true",
	})
}

func Login (c *fiber.Ctx) error {
	
	user := new(models.UserLogin)

	err := c.BodyParser(user)

	if err != nil {
		return c.Status(400).JSON(&fiber.Map{
			"success":"false",
			"message": err.Error(),
		})
	}

	var usr models.User //aqui guarda la consulta a users
	r := databases.Connection().First(&usr,"email = ?", user.Email)

	if r.Error != nil {
		return c.Status(404).JSON(&fiber.Map{
			"success":"User Incorrect",
		})
	}

	cmp := bcrypt.CompareHashAndPassword([]byte(usr.Password),[]byte(user.Password))

	if cmp != nil {
		return c.Status(404).JSON(&fiber.Map{
			"success":"Password incorrect",
		})
	}

	//enviamos usr el gen token
	token, err := helpers.GenerateToken(usr)

	if err != nil {
		return c.Status(406).JSON(&fiber.Map{
			"success":"Token didn't generate",
			"message":err.Error(),
		})
	}

	return c.Status(200).JSON(&fiber.Map{
		"token":token,
	})

}