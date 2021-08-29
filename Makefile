build:
	docker build . -f Dockerfile -t go-devdou:0.1
appup: 
	docker-compose -f docker-compose.yml --env-file ./.env.development up -d
appdown:
	docker-compose down --volume --rmi local

