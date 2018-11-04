docker build -t ratdevelopment-backend /go/src/ratdevelopment-backend
docker run --name ratdev-application --link cdaemon:cassandra -d ratdevelopment-backend
