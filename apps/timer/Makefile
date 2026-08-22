build-docker:
	docker compose exec -w /app/apps/timecounter tinygo make build

build:
	tinygo build --size short --target wioterminal -o timecounter.uf2 .
