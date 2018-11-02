/*
    The CQL statements in this file will be executed when running migrate.sh.
    All tables and keyspaces should be set up in this file.
*/

CREATE KEYSPACE IF NOT EXISTS defaultks
    WITH replication = {'class': 'SimpleStrategy', 'replication_factor': 1};

/*
    Latest snapshots by tentant
*/
CREATE TABLE IF NOT EXISTS defaultks.latest_snapshots_by_tenant (
    tenant text,
    serial_number int,
    snapshot blob,
    PRIMARY KEY ((tenant), serial_number)
) WITH CLUSTERING ORDER BY (serial_number DESC);

/*
    Snapshots by serial number
*/
CREATE TABLE IF NOT EXISTS defaultks.snapshots_by_serial_number (
    tenant text,
    serial_number int,
    time timestamp,
    snapshot blob,
    PRIMARY KEY ((tenant, serial_number), time)
) WITH CLUSTERING ORDER BY (time DESC);
