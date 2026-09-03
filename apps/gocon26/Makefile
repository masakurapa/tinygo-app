.PHONY: gen

build-docker:
	docker compose exec -w /app/apps/gocon26 tinygo make build

build:
	tinygo build --size short --opt s --target wioterminal -o gocon26.uf2 .

display:
	docker compose exec -w /app/apps/gocon26 tinygo go run .

gen:
	docker compose exec -w /app/apps/gocon26/gen tinygo sh -c "GOWORK=off /usr/local/go/bin/go run main.go gopher    > ../gopher_data.go"
	docker compose exec -w /app/apps/gocon26/gen tinygo sh -c "GOWORK=off /usr/local/go/bin/go run main.go qr        > ../qr_data.go"
	docker compose exec -w /app/apps/gocon26/gen tinygo sh -c "GOWORK=off /usr/local/go/bin/go run main.go mygopher  > ../mygopher_data.go"
