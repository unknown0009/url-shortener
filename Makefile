# Local storage
build:
	docker build -t go-docker . 

run: build
	docker run -dp 80:80 --name go-docker go-docker -db local

destroy:
	docker stop go-docker
	docker rm go-docker
	docker rmi go-docker

restart: destroy build run


# PostgreSQL
start:
	docker-compose -f docker-compose.yml up -d

stop:
	docker-compose -f docker-compose.yml stop

stoprm: stop
	docker rm urlsapp
	docker rmi docker_urlsapp

down: 
	docker-compose -f docker-compose.yml down 
	docker rm urlsapp
	docker rm urls_container
	docker rmi docker_urlsapp

