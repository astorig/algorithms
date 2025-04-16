POSTGRES_USER = algorithms
POSTGRES_PASSWORD = algorithms
POSTGRES_DB = algorithms
VOLUME_NAME = algorithms-data
CONTAINER_NAME = algorithms-db
IMAGE_NAME = postgres:latest
PORT = 15555

run:
	docker run --name $(CONTAINER_NAME) \
	  -e POSTGRES_USER=$(POSTGRES_USER) \
	  -e POSTGRES_PASSWORD=$(POSTGRES_PASSWORD) \
	  -e POSTGRES_DB=$(POSTGRES_DB) \
	  -v $(VOLUME_NAME):/var/lib/postgresql/data \
	  -p $(PORT):5432 \
	  -d $(IMAGE_NAME)

stop:
	docker stop $(CONTAINER_NAME)

rm:
	docker rm $(CONTAINER_NAME)

rm-volume:
	docker volume rm $(VOLUME_NAME)

restart: stop rm run

ps:
	docker ps

volumes:
	docker volume ls