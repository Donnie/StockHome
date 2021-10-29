dev:
	docker-compose --env-file ./.env.local up

sql:
	sqlite3 sql.db

clean:
	@echo "Cleaning Docker environment..."
	docker-compose stop
	docker-compose down -v

# CI
build:
	@echo "Building for prod"
	docker build -t donnieashok/stockhome:prod .

deploy: build
	- echo "$(DOCKER_PASSWORD)" | docker login -u "$(DOCKER_USERNAME)" --password-stdin
	docker push donnieashok/stockhome:prod
	@echo "Deployed!"

# Prod
live:
	ssh root@stockhome docker pull donnieashok/stockhome:prod
	- ssh root@stockhome docker stop stockhome
	- ssh root@stockhome docker rm stockhome
	scp -r ./.env root@stockhome:/root/
	ssh root@stockhome docker run -d --restart on-failure -v /home/stockhome/:/db/ --env-file /root/.env -p 1339:8080 --name stockhome donnieashok/stockhome:prod
	ssh root@stockhome docker cp .env stockhome:/.env
	ssh root@stockhome rm /root/.env
	@echo "Is live"
