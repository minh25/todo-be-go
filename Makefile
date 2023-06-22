.PHONY: start

start:
	go run main.go start

migrate-up:
	go run main.go migrate up

migrate-down-1:
	go run main.go migrate down 1

migrate:
	echo \# make migrate name="$(name)"
	go run main.go migrate create $(name)
