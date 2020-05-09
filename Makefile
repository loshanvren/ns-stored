default: build run clean

PROOT=$(shell pwd)

build:
	@echo "build application with docker"
	docker build -t ns-stored:latest .

run:
	@echo "run application with docker"
	docker run --restart=always -dit --name ns-serv ns-stored:latest

stop:
	@echo "stop application with docker"
	docker stop ns-serv

clean: stop
	docker container rm ns-serv
