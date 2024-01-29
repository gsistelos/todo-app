COMPOSE = docker-compose

VOLUMES = mysql_data

up: volumes
	$(COMPOSE) up -d

start:
	$(COMPOSE) start

down:
	$(COMPOSE) down

stop:
	$(COMPOSE) stop

restart:
	$(COMPOSE) restart

volumes:
	mkdir -p $(VOLUMES)

clean:
	-docker rm -f $$(docker ps -aq)
	-docker image rm -f $$(docker images -aq)
	-docker volume rm -f $$(docker volume ls -q)
	-docker network rm -f $$(docker network ls -q)

.PHONY: up start down stop restart volumes clean
