.PHONY: build
docker-build:
	docker build -t plan-bbs-api .

.PHONY: run
run:
	docker run -d -p 8080:8080 -v $(PWD):/workdir --name api --rm plan-bbs-api

.PHONY: stop
stop:
	docker stop api

.PHONY: firestore-start
firestore-start:
	gcloud beta emulators firestore start --host-port=localhost:8812