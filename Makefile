DOCKER=docker-compose
RUNCMD=$(DOCKER) up
DOWNCMD=$(DOCKER) down

all:
	$(RUNCMD) -d --build
init-db:
	sh ./db/init-container.sh
clean:
	$(DOWNCMD) --rmi all