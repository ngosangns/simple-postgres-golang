include wire/.env
export

cp-env:
	cp wire/.env.example wire/.env

# Generators
gen-sql:
	go run github.com/sqlc-dev/sqlc/cmd/sqlc -f sqlc.yaml generate
gen-mig:
	atlas schema diff \
		--from "postgres://${POSTGRES_USER}:${POSTGRES_PASS}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_NAME}?sslmode=disable" \
		--to "file://postgres/schema" \
		--dev-url "docker://postgres/latest/test" \
		--format "{{ sql . }}" \
		> ./postgres/migration_gen/$(shell date +%Y%m%d%H%M%S).sql

# Migration
# https://atlasgo.io/versioned/apply
mig-up:
	atlas schema apply \
		--url "postgres://${POSTGRES_USER}:${POSTGRES_PASS}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_NAME}?sslmode=disable" \
		--to "file://postgres/migration_gen" \
		--dev-url "docker://postgres/latest/test"
