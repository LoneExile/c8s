
.PHONY: install
install:
	go mod download
	go install github.com/cosmtrek/air@latest
	go install github.com/a-h/templ/cmd/templ@latest
	curl https://unpkg.com/htmx.org@1.9.12/dist/htmx.min.js --create-dirs -o ./static/script/htmx.min.js
	curl https://unpkg.com/htmx.org@1.9.12/dist/ext/response-targets.js --create-dirs -o ./static/script/response-targets.js
	pnpm install
	make tailwind
# npx tailwindcss -i tailwind.css -o output.css

.PHONY: tailwind
tailwind:
	./node_modules/.bin/tailwindcss -i ./static/css/input.css -o ./static/css/style.css -c ./tailwind.config.js

.PHONY: tailwind-watch
tailwind-watch:
	./node_modules/.bin/tailwindcss -i ./static/css/input.css -o ./static/css/style.css --watch -c ./tailwind.config.js

.PHONY: tailwind-build
tailwind-build:
	./node_modules/.bin/tailwindcss -i ./static/css/input.css -o ./static/css/style.min.css --minify -c ./tailwind.config.js

.PHONY: templ-generate
templ-generate:
	templ generate

.PHONY: templ-watch
templ-watch:
	templ generate --watch
	
.PHONY: dev
dev:
	make tailwind
	go build -o ./tmp/c8s ./cmd/c8s/main.go && air

.PHONY: build
build:
	make tailwind-build
	make templ-generate
	go build -ldflags "-X main.Environment=production" -o ./bin/c8s ./cmd/c8s/main.go

.PHONY: vet
vet:
	go vet ./...

.PHONY: lint
lint:
	golangci-lint run --verbose
