# Load env vars from .env.deploy
include deploy/.env.deploy
export

API=api
GOLANGCI_LINT=$(shell go env GOPATH)/bin/golangci-lint

IMAGE_NAME?=aioz-roadmap-service

# ── Local ─────────────────────────────────────────────────────────────────────
build: lint
	@go build -o bin/${API} ./app/*.go

run: build
	@bin/${API}

lint:
	@${GOLANGCI_LINT} run

lint-fix:
	@${GOLANGCI_LINT} run --fix

docker-build:
	@echo "Building docker image..."
	@docker build -t ${IMAGE_NAME}:latest .

# ── Deploy ────────────────────────────────────────────────────────────────────
deploy: docker-build
	@echo "Saving image to tar file..."
	@docker save ${IMAGE_NAME}:latest | gzip > /tmp/${IMAGE_NAME}.tar.gz

	@echo "Copying image to server..."
	@scp /tmp/${IMAGE_NAME}.tar.gz ${SERVER_USER}@${SERVER_HOST}:${SERVER_DIR}/${IMAGE_NAME}.tar.gz

	@echo "Copying docker-compose.yml to server..."
	@scp docker-compose.yml ${SERVER_USER}@${SERVER_HOST}:${SERVER_DIR}/docker-compose.yml

	@echo "Loading and restarting on server..."
	@ssh ${SERVER_USER}@${SERVER_HOST} "\
		cd ${SERVER_DIR} && \
		docker load < ${IMAGE_NAME}.tar.gz && \
		docker compose up -d && \
		rm ${IMAGE_NAME}.tar.gz \
	"

	@echo "Cleaning up local tar file..."
	@rm /tmp/${IMAGE_NAME}.tar.gz

	@echo "Done! App is running on ${SERVER_HOST}"


# ── Helpers ───────────────────────────────────────────────────────────────────
restart:
	@ssh ${SERVER_USER}@${SERVER_HOST} "cd ${SERVER_DIR} && docker compose restart"

ssh:
	@ssh ${SERVER_USER}@${SERVER_HOST}

logs:
	@ssh ${SERVER_USER}@${SERVER_HOST} "cd ${SERVER_DIR} && docker compose logs -f"

status:
	@ssh ${SERVER_USER}@${SERVER_HOST} "cd ${SERVER_DIR} && docker compose ps"

stop:
	@ssh ${SERVER_USER}@${SERVER_HOST} "cd ${SERVER_DIR} && docker compose down"
