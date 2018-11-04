docker stop cdaemon
docker rm cdaemon
docker run --name cdaemon -d cassandra:latest
