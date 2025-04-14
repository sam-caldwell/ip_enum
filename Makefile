.PHONY=build

all: clean build
	@echo 'ok'

build:
	@mkdir build || true &> /dev/null
	@go build -o build/ip_enum main.go

clean:
	@rm -rf build || true

