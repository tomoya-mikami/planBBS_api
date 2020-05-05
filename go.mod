module github.com/tomoya-paseri/planBBS_api

go 1.12

replace local.packages/plan => ./src/Domain/Plan

replace local.packages/handler => ./src/Handler

replace local.packages/src => ./src

require (
	cloud.google.com/go/firestore v1.2.0
	github.com/gofiber/fiber v1.9.3
	github.com/klauspost/compress v1.10.5 // indirect
	github.com/yuin/goldmark v1.1.30 // indirect
	golang.org/x/crypto v0.0.0-20200429183012-4b2356b1ed79 // indirect
	golang.org/x/net v0.0.0-20200501053045-e0ff5e5a1de5 // indirect
	golang.org/x/sys v0.0.0-20200501145240-bc7a7d42d5c3 // indirect
	golang.org/x/tools v0.0.0-20200504022951-6b6965ac5dd1 // indirect
	google.golang.org/api v0.22.0
	local.packages/handler v0.0.0-00010101000000-000000000000
	local.packages/plan v0.0.0-00010101000000-000000000000
	local.packages/src v0.0.0-00010101000000-000000000000
)
