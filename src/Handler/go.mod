module github.com/tomoya-paseri/planBBS_api

go 1.12

replace local.packages/plan => ../Domain/Plan

require (
	github.com/gofiber/fiber v1.9.3
	local.packages/plan v0.0.0-00010101000000-000000000000
)
