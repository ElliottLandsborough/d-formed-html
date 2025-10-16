all: tidy run copy-assets

.PHONY: tidy run copy-assets

tidy:
	go mod tidy

run:
	go run main.go

copy-assets:
	mkdir -p public/css public/js public/img
	cp -R assets/* public/