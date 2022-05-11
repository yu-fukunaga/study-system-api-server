.PHONY: setup
setup:
	go install github.com/cosmtrek/air@latest

.PHONY: env
env:
	cp .envrc-develop .envrc
	direnv allow .

.PHONY: mysql 
mysql: 
	docker-compose exec mysql mysql -u root --password=passwd sample_db
