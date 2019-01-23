.PHONY: up down rebuild

up:
	docker build -t builtin/backend -f Dockerfile.backend .
	docker build -t builtin/frontend -f Dockerfile.frontend ./frontend
	docker-compose up -d

down:
	docker-compose down

rebuild:
	docker-compose restart
