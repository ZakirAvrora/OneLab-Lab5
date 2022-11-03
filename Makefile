build:
	docker build --rm -t crud-web .
	docker image prune --filter label=stage=builder -f
run:
	docker run --rm --name crud-web -p 8080:8080 crud-web