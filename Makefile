include .env #выгружаем все данные с .env и они будут в каждой теме участвовать и запускаться вместе с ней
export
service:
	go run main.go 
	migrate -path migrations -database ${CONN_STRING} up