package Handler

import (
	"encoding/json"
	"github.com/gofiber/fiber"

	"local.packages/plan"
)

type PlanHandler struct {
	planService Plan.ServiceInterface
}

func NewPlanHadler(planService Plan.ServiceInterface) *PlanHandler {
	handler := new (PlanHandler)
	handler.planService = planService
	return handler
}

func (h PlanHandler) PlanList(c *fiber.Ctx) {
	var plans []Plan.Plan
	plans, _ = h.planService.FindAll()
	res, _ := json.Marshal(plans)

	c.Send(res)
}

func (h PlanHandler) PlanAdd(c *fiber.Ctx) {
	h.planService.Add()
	c.Send("{}")
}
