build:
	GOOS=linux GOARCH=amd64 go build -o product-cli
	docker build -t product-cli .

run:
	docker run product-cli