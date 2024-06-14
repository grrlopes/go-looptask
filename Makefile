dev:
	@cd src && air server --port 8080

test:
	@go test -v ./tests/*

clean:
	@rm -rf ./src/tmp && rm -rf ./tmp
