.PHONY: list test benchmark build

list:
	@$(MAKE) -pRrq -f $(lastword $(MAKEFILE_LIST)) : 2>/dev/null | awk -v RS= -F: '/^# File/,/^# Finished Make data base/ {if ($$1 !~ "^[#.]") {print $$1}}' | sort | egrep -v -e '^[^[:alnum:]]' -e '^$@$$'

test:
	go test ./tests/... -v

benchmark:
	go test -bench=. -benchmem ./tests/...

build:
	go build -v ./...

