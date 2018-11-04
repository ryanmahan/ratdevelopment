docker stop ratdev-application
docker rm ratdev-application
docker stop cdaemon
docker rm cdaemon
docker run --name cdaemon -d cassandra:latest
