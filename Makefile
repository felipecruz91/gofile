graph:
	go run cmd/gofile/main.go -graph | buildctl debug dump-llb | jq .

graph-dot:
	go run cmd/gofile/main.go -graph | buildctl debug dump-llb --dot | dot -T png -o out.png

push-gateway-img:
	docker buildx build -t felipecruz/gofile . --platform linux/amd64,linux/arm64 --push

build-demo:
	docker buildx build -f Gofile.yaml -t felipecruz/gofile-demo . --load

push-demo:
	docker buildx build -f Gofile.yaml -t felipecruz/gofile-demo . --platform linux/amd64,linux/arm64 --push
