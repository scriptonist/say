build:
	GO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .
	docker build -t scriptonist/say .
	rm -f app
