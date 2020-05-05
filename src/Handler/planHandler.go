package Handler

import (
	"encoding/json"

	"github.com/gofiber/fiber"

	Plan "local.packages/plan"
)

type PlanHandler struct {
	planService Plan.ServiceInterface
}

func NewPlanHandler(planService Plan.ServiceInterface) *PlanHandler {
	handler := new(PlanHandler)
	handler.planService = planService
	return handler
}

func (h PlanHandler) PlanList(c *fiber.Ctx) {
	responseTable := make(map[string][]Plan.Plan)
	plans, _ := h.planService.FindAll()
	responseTable["Plan"] = plans
	res, _ := json.Marshal(responseTable)

	c.Send(res)
}

func (h PlanHandler) PlanAdd(c *fiber.Ctx) {
	var plan Plan.Plan
	json.Unmarshal([]byte(c.Body()), &plan)

	err := h.planService.Add(plan)
	if err != nil {
		c.Status(500).Send(err)
		return
	}

	c.Send("{}")
}
