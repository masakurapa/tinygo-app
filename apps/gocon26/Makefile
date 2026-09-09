.PHONY: gen

build-docker:
	docker compose exec -w /app/apps/gocon26 tinygo make build

build:
	tinygo build --size short --opt s --target wioterminal -o gocon26.uf2 .

display:
	docker compose exec -w /app/apps/gocon26 tinygo go run .

gen:
	docker compose exec -w /app/apps/gocon26 tinygo sh -c "mkdir -p img"
	docker compose exec -w /app/apps/gocon26/gen tinygo sh -c "GOWORK=off /usr/local/go/bin/go run main.go qr        > ../img/qr_data.go"
	docker compose exec -w /app/apps/gocon26/gen tinygo sh -c "GOWORK=off /usr/local/go/bin/go run main.go mygopher    > ../img/mygopher_data.go"
	docker compose exec -w /app/apps/gocon26/gen tinygo sh -c "GOWORK=off /usr/local/go/bin/go run main.go gopher_fall > ../img/gopher_fall_data.go"
