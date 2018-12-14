/*
    The CQL statements in this file will be executed when running migrate.sh.
    All tables and keyspaces should be set up in this file.
*/

DROP KEYSPACE IF EXISTS defaultks;

CREATE KEYSPACE IF NOT EXISTS defaultks
    WITH replication = {'class': 'SimpleStrategy', 'replication_factor': 1};

/*
    Latest snapshots by tentant
*/
CREATE TABLE IF NOT EXISTS defaultks.latest_snapshots_by_tenant (
    tenant text,
    serial_number int,
    tenant_search text,
    serial_string text,
    company_name text,
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


CREATE CUSTOM INDEX ser_string_index
      ON defaultks.latest_snapshots_by_tenant (serial_string)
      USING 'org.apache.cassandra.index.sasi.SASIIndex'
      WITH OPTIONS = {'mode': 'CONTAINS'};

CREATE CUSTOM INDEX tenant_index
      ON defaultks.latest_snapshots_by_tenant (tenant_search)
      USING 'org.apache.cassandra.index.sasi.SASIIndex'
      WITH OPTIONS = {'mode': 'CONTAINS'};

CREATE CUSTOM INDEX company_name_index
      ON defaultks.latest_snapshots_by_tenant (company_name)
      USING 'org.apache.cassandra.index.sasi.SASIIndex'
      WITH OPTIONS = {'mode': 'CONTAINS'};