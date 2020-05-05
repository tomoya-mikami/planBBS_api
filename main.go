// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// [START gae_go111_app]

// Sample helloworld is an App Engine app.
package main

// [START import]
import (
	"context"
	"log"

	"github.com/gofiber/fiber"
	"cloud.google.com/go/firestore"

	"local.packages/handler"
	"local.packages/plan"
)

func main() {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, "planbbs")
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()

	// 本当はここをDIコンテナでやりたい
	planRepository := Plan.NewRepository(client, ctx)
	planService := Plan.NewService(planRepository)
	planHandler := Handler.NewPlanHandler(planService)

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) {
		c.Send("Hello, World!")
	})

	app.Get("/plan", planHandler.PlanList)
	app.Post("/plan", planHandler.PlanAdd)

	err = app.Listen(8080)
	if err != nil {
		log.Fatalln(err)
	}
}
