module github.com/tomoya-paseri/planBBS_api

go 1.12

replace local.packages/plan => ./src/Domain/Plan

replace local.packages/handler => ./src/Handler

replace local.packages/src => ./src

require (
	cloud.google.com/go/firestore v1.2.0
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/gofiber/cors v0.0.3 // indirect
	github.com/google/wire v0.4.0
	github.com/klauspost/compress v1.10.5 // indirect
	github.com/yuin/goldmark v1.1.30 // indirect
	golang.org/x/crypto v0.0.0-20200429183012-4b2356b1ed79 // indirect
	golang.org/x/net v0.0.0-20200505041828-1ed23360d12c // indirect
	golang.org/x/sys v0.0.0-20200501145240-bc7a7d42d5c3 // indirect
	golang.org/x/tools v0.0.0-20200505023115-26f46d2f7ef8 // indirect
	local.packages/handler v0.0.0-00010101000000-000000000000
	local.packages/plan v0.0.0-00010101000000-000000000000
	local.packages/src v0.0.0-00010101000000-000000000000
)
