.PHONY: build
docker-build:
	docker build -t plan-bbs-api .

.PHONY: run
run:
	docker run -p 8080:8080 -p 8812:8812 -v $(PWD):/workdir --name api --rm plan-bbs-api

.PHONY: stop
stop:
	docker stop api

.PHONY: firestore-start
firestore-start:
	gcloud beta emulators firestore start --host-port=localhost:8812

.PHONY: deploy
deploy:
	gcloud app deploy
