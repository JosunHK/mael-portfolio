.PHONY: tailwind-watch 
tailwind-watch:
	./tailwindcss -i ./web/static/input.css -o ./web/static/css/style.css --watch

.PHONY: tailwind-build
tailwind-build:
	./tailwindcss -i ./web/static/input.css -o ./web/static/css/style.min.css --minify

.PHONY: templ-generate
templ-generate:
	templ generate

.PHONY: templ-watch
templ-watch:
	templ generate --watch

.PHONY: sqlc-generate
sqlc-watch:
	sqlc generate 

.PHONY: dev
dev:
	go build -o ./tmp/$(APP_NAME) ./cmd/$(APP_NAME)/main.go && air

.PHONY: build
build:
	make tailwind-build
	templ generate
	sqlc generate
	go build -ldflags "-X main.Environment=production" -o ./bin/ ./cmd/main.go

.PHONY: init-build
build:
	curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/download/v3.4.19/tailwindcss-linux-x64
	chmod +x tailwindcss-linux-x64
	mv tailwindcss-linux-x64 tailwindcss
	make tailwind-build
	go install github.com/a-h/templ/cmd/templ@latest
	templ generate
	go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
	sqlc generate
	go build -ldflags "-X main.Environment=production" -o ./bin/ ./cmd/main.go

.PHONY: start
start:
	./bin/main

.PHONY: vet
vet:
	go vet ./...

.PHONY: staticcheck
staticcheck:
	staticcheck ./...

.PHONY: test
test:
	  go test -race -v -timeout 30s ./...
