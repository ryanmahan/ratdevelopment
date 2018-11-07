CQL="CREATE KEYSPACE IF NOT EXISTS test_space WITH replication = {'class': 'SimpleStrategy', 'replication_factor': 1};
CREATE TABLE IF NOT EXISTS test_space.comments (comment varchar, name varchar, datetime timestamp, PRIMARY KEY ((name), datetime)) WITH CLUSTERING ORDER BY (datetime DESC);"

until echo $CQL | cqlsh; do
    echo "cqlsh: Cassandra is unavailable to initialize - will retry later"
    sleep 2
done &

echo "Running docker-entrypoint..."
exec /docker-entrypoint.sh "$@"
