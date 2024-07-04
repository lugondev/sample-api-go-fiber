.PHONY: build
include .env.prod

PATH_MAIN := "cmd/main.go"
PATH_CURRENT := $(shell pwd)
GIT_COMMIT := $(shell git log --oneline -1 HEAD)

ENV_DEPLOY = \
			--allow-unauthenticated \
			--set-env-vars DB_HOST="${DB_HOST}" \
			--set-env-vars DB_PORT="${DB_PORT}" \
			--set-env-vars DB_NAME="${DB_NAME}" \
			--set-env-vars DB_USER="${DB_USER}" \
			--set-env-vars DB_PASSWORD="${DB_PASSWORD}" \
			--set-env-vars GIT_COMMIT="${GIT_COMMIT}";

dev:
	go run ${PATH_MAIN} api

watch:
	reflex -s -r '\.go$$' make dev

swag:
	swag init --parseDependency --parseInternal --generalInfo ./cmd/app/api.go

pre-build:
	echo "current commit: ${GIT_COMMIT}"
	go mod tidy

build: pre-build clean-built
	env GOOS=linux GOARCH=amd64 go build -v -o ./build/bin -ldflags "-X 'main.GitCommitLog=${GIT_COMMIT}'" ${PATH_MAIN}

clean-built:
	rm -fr "$(PATH_CURRENT)/build/bin"; \
	echo "Clean built."

copy-docker-file:
	cp Dockerfile_build ./build/Dockerfile

deploy: build copy-docker-file
	gcloud run deploy ${SERVICE_NAME}-app --source ./build --region asia-southeast1 --project ${GCP_PROJECT} \
			--min-instances=1 --max-instances=10 --memory=1Gi --cpu=1 \
			${ENV_DEPLOY} \
