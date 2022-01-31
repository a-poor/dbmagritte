.PHONY: default
default:
	@echo ""

.PHONY: build
build:
	go build -o build

.PHONY: clean
clean:
	rm -rf build/* migrations/ dbmconf.yaml
