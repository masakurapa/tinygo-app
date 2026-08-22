build-docker:
	docker compose exec -w /app/apps/timer tinygo make build

build:
	tinygo build --size short --target wioterminal -o timer.uf2 .
