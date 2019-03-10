.PHONY: build dev predeploy deploy unit-test integration-test test

run: 
	go run ./main.go

dev: 
	./realize start

integration-test:
	go test ./... -tags=integration -count=1

unit-test:
	go test ./... -tags=unit -count=1

test: unit-test integration-test