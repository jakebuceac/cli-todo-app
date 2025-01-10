.PHONY: build

up:
	@echo "Starting Docker images..."
	docker-compose up -d
	@echo "Docker images started!"

down:
	@echo "Stopping docker compose..."
	docker-compose down
	@echo "Done!"

build:
	@echo "Building tasks binary..."
	cd build && go build -o tasks ..
	@echo "Done!"