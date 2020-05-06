package src

import (
	"log"

	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"

	Handler "local.packages/handler"
)

type Server struct {
	planHandler *Handler.PlanHandler
}

func NewServer(planHandler *Handler.PlanHandler) *Server {
	server := new(Server)
	server.planHandler = planHandler

	return server
}

func (s Server) Start() error {
	app := fiber.New()

	app.Use(cors.New())

	app.Get("/", func(c *fiber.Ctx) {
		c.Send("Hello, World!")
	})

	app.Get("/plans", s.planHandler.PlanList)
	app.Get("/plan", s.planHandler.PlanShow)
	app.Post("/plan", s.planHandler.PlanAdd)

	err := app.Listen(8080)
	if err != nil {
		log.Fatalln(err)
	}

	return err
}
