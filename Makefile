test:
	go test -v ./...

# Local storage
build:
	docker build -t urlsapp . 

run: build
	docker run -dp 80:80 --name urlsapp urlsapp -db local

destroy:
	docker stop urlsapp
	docker rm urlsapp
	docker rmi urlsapp

restart: destroy build run


# PostgreSQL
start:
	docker-compose -f docker-compose.yml up -d

stop:
	docker-compose -f docker-compose.yml stop

stoprm: stop
	docker rm urlsapp
	docker rmi url-shortener_urlsapp

down: 
	docker-compose -f docker-compose.yml down 
	docker rmi url-shortener_urlsapp

