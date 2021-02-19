build:
	go build -o ./bin/rootly ./cmd/rootly

docker-build:
	docker build -f docker/Dockerfile -t rootly/cli .

clean:
	rm -r ./bin
