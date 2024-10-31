include dev.env

up:
	@echo "starting container..."
	docker-compose up --build -d --remove-orphans

down:
	@echo "stoping container..."
	docker-compose down

build:
	go build -o ${BINARY} .

start:
	./${BINARY}

restart: build start