build-docker:
	docker compose exec -w /app/apps/gocon26 tinygo make build

build:
	tinygo build --size short --target wioterminal -o gocon26.uf2 .

display:
	docker compose exec -w /app/apps/gocon26 tinygo go run .

gen:
	cd gen && GOWORK=off go run main.go gopher > ../gopher_data.go
	cd gen && GOWORK=off go run main.go qr    > ../qr_data.go
