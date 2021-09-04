build:
	docker build . -f Dockerfile -t go-devdou:0.1
appup: 
	docker-compose -f docker-compose.yml --env-file ./.env.development up -d
appdown:
	docker-compose down --volume --rmi local
install_swagger:
	sudo wget https://github.com/go-swagger/go-swagger/releases/download/v0.27.0/swagger_linux_amd64 -O /usr/local/bin/swagger 
	sudo chmod +x /usr/local/bin/swagger
gen:
	swagger generate spec -o ./swagger.yaml --scan-models
docui: gen
	swagger serve -F=swagger swagger.yaml -p=8001