CURRENT_DIR=$(shell pwd)

create-mig:
	migrate create -ext sql -dir migration -seq auth-service

mig-insert:
	migrate create -ext sql -dir migration -seq insert_table


proto-gen:
	./scripts/gen-proto.sh ${CURRENT_DIR}

mig-up:
	migrate -database 'postgres://postgres:1234@localhost:5432/yusolve?sslmode=disable' -path migration up

mig-down:
	migrate -database 'postgres://postgres:1234@localhost:5432/yusolve?sslmode=disable' -path migration down

mig-force:
	migrate -database 'postgres://postgres:1234@localhost:5432/yusolve?sslmode=disable' -path migration force 1

# Define paths to swag and swagger documentation output
SWAGGER := ~/go/bin/swag
SWAGGER_DOCS := api/docs
SWAGGER_INIT := $(SWAGGER) init -g ./api/router.go -o $(SWAGGER_DOCS)

# Target to generate swagger documentation
swag-gen:
	$(SWAGGER_INIT)
run:
	go run cmd/main.go
