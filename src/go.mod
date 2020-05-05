module github.com/tomoya-paseri/planBBS_api

go 1.12

replace local.packages/plan => ./Domain/Plan

replace local.packages/handler => ./Handler

require (
	cloud.google.com/go/firestore v1.2.0
	github.com/gofiber/fiber v1.9.3
	github.com/google/wire v0.4.0
	local.packages/handler v0.0.0-00010101000000-000000000000
	local.packages/plan v0.0.0-00010101000000-000000000000
)
