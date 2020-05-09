default: build run

PROOT=$(shell pwd)

build:
	@echo "build application with docker"
	docker build -t ns-stored:latest .

run:
	@echo "run application with docker"
	docker run --restart=always -dit --name ns-serv ns-stored:latest