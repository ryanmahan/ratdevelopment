echo "CREATE KEYSPACE IF NOT EXISTS test_space WITH replication = {'class': 'SimpleStrategy', 'replication_factor': 1};" | docker exec -i cdaemon cqlsh
echo "CREATE TABLE IF NOT EXISTS test_space.comments (comment varchar, name varchar, datetime timestamp, PRIMARY KEY ((name), datetime)) WITH CLUSTERING ORDER BY (datetime DESC);" | docker exec -i cdaemon cqlsh
