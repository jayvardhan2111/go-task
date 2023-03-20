package main

import (
	// "fmt"
    "github.com/gofiber/fiber/v2"
    "os/exec"
	
)

func main() {
    app := fiber.New()

    app.Post("/api/cmd", postHandler)

	app.Use(otherHandler)

    app.Listen(":5000")
}


func postHandler(c *fiber.Ctx) error {
	
	cmd := c.FormValue("cmd")

	if cmd == ""{
		return c.Status(400).SendString("Please enter the value of CMD")
	}

	out, err := exec.Command("sh","-c",cmd).Output()
	
	if err != nil {
		return c.Status(400).SendString(" Error : "+cmd+" is invalid command ")
	}

	return c.Status(200).SendString(string(out))
}

func otherHandler(c *fiber.Ctx) error {

	return c.Status(404).SendString("The Requested API or Method is nod Found !! ")
}

