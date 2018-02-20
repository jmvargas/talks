# To run inside minimal-docker-images dir
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ../hello-world-http
docker build -t hello-scratch -f Dockerfile .
docker run -p "8080:8080" --rm -t hello-scratch

# To see image size
docker images | grep hello-scratch