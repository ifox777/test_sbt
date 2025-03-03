BINARY_NAME = testapp

build:
	docker build -t $(BINARY_NAME) .

run:
	docker run $(BINARY_NAME)
	 
up:
	docker-compose up --build

down:
	docker-compose down