.PHONY: start
start:
	docker-compose up -d
.PHONY: stop
stop:
	docker-compose down

.PHONY: reset-db
reset-db:
	docker-compose down -v --remove-orphans
	docker-compose up -d db


