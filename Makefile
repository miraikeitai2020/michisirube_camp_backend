DOCKER=docker-compose
RUNCMD=$(DOCKER) up
DOWNCMD=$(DOCKER) down

all:
	$(UPCMD) -d --build
clean:
	$(DOWNCMD) --rmi all