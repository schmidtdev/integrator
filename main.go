package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/gofiber/template/handlebars/v2"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	engine := handlebars.New("./views", ".hbs")

	app := fiber.New(fiber.Config{
		Prefork:       false,
		CaseSensitive: false,
		AppName:       "Integrator",
		Views:         engine,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title": "Hello, World!",
		})
	})

	app.Post("/login", func(c *fiber.Ctx) error {
		if c.FormValue("username") == "admin" && c.FormValue("password") == "password" {
			return c.Redirect("/dashboard")
		}
		return c.Render("error", fiber.Map{
			"errorMessage": "Invalid credentials",
		})
	})

	app.Get("/dashboard", func(c *fiber.Ctx) error {
		return c.Render("home", fiber.Map{
			"Title": "Welcome, admin!",
			"transactionLogs": []fiber.Map{
				{
					"date":   "01/10/2023",
					"id":     "550e8400-e29b-41d4-a716-446655440000",
					"amount": 100.50,
					"status": "bi-check-circle", // Completed
				},
				{
					"date":   "02/10/2023",
					"id":     "550e8400-e29b-41d4-a716-446655440001",
					"amount": 200.75,
					"status": "bi-hourglass-split", // Pending
				},
				{
					"date":   "03/10/2023",
					"id":     "550e8400-e29b-41d4-a716-446655440002",
					"amount": 150.00,
					"status": "bi-x-circle", // Failed
				},
				{
					"date":   "04/10/2023",
					"id":     "550e8400-e29b-41d4-a716-446655440003",
					"amount": 300.20,
					"status": "bi-check-circle", // Completed
				},
			},
		}, "layouts/main")
	})

	app.Get("/integrations", func(c *fiber.Ctx) error {
		return c.Render("integrations", fiber.Map{
			"Title": "Integrations",
			"integrations": []fiber.Map{
				{
					"name":      "WebGR",
					"logo":      "https://webgr.com.br/wp-content/uploads/2024/01/favicon-aplicativo-webgr.png",
					"connected": true,
				},
				{
					"name":      "WebMais",
					"logo":      "https://scontent-gru1-2.xx.fbcdn.net/v/t39.30808-6/361196835_572669078372693_3865494450512108304_n.png?_nc_cat=103&ccb=1-7&_nc_sid=6ee11a&_nc_eui2=AeEwk2CFdxcghG7Da8ox2CwbbfX92rt8w9Zt9f3au3zD1kcXGy5V_qrqKhJytaY4ur1ouOUCqKOPo-PKC0OJ2OcL&_nc_ohc=V17D3OqSksgQ7kNvgGVqHcE&_nc_ht=scontent-gru1-2.xx&_nc_gid=A1MehCL8qQnoWWAnuxfiDQ0&oh=00_AYCQXm_R9DszqeP1i_QHpjpjicDjIe5aXo4mFZNo1dDv6Q&oe=67061FEE",
					"connected": false,
				},
				{
					"name":      "Neogrid",
					"logo":      "https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Fraichu-uploads.s3.amazonaws.com%2Flogo_neogrid_abBtZa.png&f=1&nofb=1&ipt=130b4444a79b02501f759e7483f30b00b4107bc4cd368059117dfd811be83285&ipo=images",
					"connected": false,
				},
				{
					"name":      "PowerBI",
					"logo":      "https://img.icons8.com/?size=100&id=3sGOUDo9nJ4k&format=png&color=000000",
					"connected": false,
				},
			},
		}, "layouts/main")
	})

	app.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			"admin": "password",
		},
	}))

	app.Get("/api", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Hello, World!"})
	})

	app.Listen(":3000")
}
