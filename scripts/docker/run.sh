cd ../..
echo run th build
docker build . -t go-er-service -f scripts/docker/Dockerfile
echo run docker
docker run -p 8080:8080 go-er-service