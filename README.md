# ratdevelopment

For the CS 320 Class at UMass

# golang-vagrant

Before you can do anything, you need to set the password for the vm (if you havent already done so)
```
sudo passwd ubuntu
```

### Cassandra Basics

Cassandra should typically start when your machine starts
But if it isnt running, just use the following to start it
```
service cassandra start
```

Alternatively, to stop Cassandra use
```
service cassandra stop
```

When cassandra is running, use the following to get into the cluster to start messing around with things
Do note that you cant use GoCQL to mess with the schema at all, that has to be done through cqlsh
```
cqlsh
```

First your going to have to create is a keyspace
I'm going to recomend the following:
```
CREATE KEYSPACE IF NOT EXISTS keyspace_name WITH replication = {'class': 'SimpleStrategy', 'replication_factor': 1};
```
replication_factor refers to the number of replicas of data kept on multiple nodes. Make sure this never exceeds your node count, or all operations will fail. (required for SimpleStrategy)
Class refers to the replication strategy used, two main types:
SimpleStrategy - all nodeclusters will use the same replication factor
NetworkTopologyStrategy - allows you to set the replication factor on a per datacenter basis (also has some other quirks involving physical location of the nodes)

NetworkTopologyStrategy is typically recommended (easier when exapnding in the future), but for the purposes of cassandra on our local machines it wont really matter so its easier to just use SimpleStrategy

Next, need to create a table. It uses the following format
```
CREATE TABLE IF NOT EXISTS keyspace_name.table_name (
    field_name_one datatype,
    field_name_two datatype,
    field_name_three datatype,
    ...,
    PRIMARY KEY ((field_name_one), field_name_three)
    ) WITH CLUSTERING ORDER BY (field_name_three DESC);
```

Now you should be ready to enter data (using GoCQL too)
Basic CQL Command formats:
```
INSERT INTO table_name (field_name_one, field_name_two) VALUES (data_one, data_two);
SELECT * FROM keyspace_name.table_name WHERE primary_key = 'some_value' AND (other key/value relations);
UPDATE table_name SET field_name_one = 'new_data' WHERE primary_key = 'some_value';
DELETE FROM table_name WHERE (some relations);
```
Refer to http://cassandra.apache.org/doc/latest/cql/dml.html to go more in depth into CQL


For my quick demo to work, you need to run:
```
CREATE KEYSPACE IF NOT EXISTS test_space 
    WITH replication = {'class': 'SimpleStrategy', 'replication_factor': 1};
CREATE TABLE IF NOT EXISTS test_space.comments (
    comment varchar,
    name varchar,
    datetime timestamp,
    PRIMARY KEY ((name), datetime)
    ) WITH CLUSTERING ORDER BY (datetime DESC);
```

### Potnetially useful tidbits

Use the following from within cqlsh to see all of your clusters...
```
SELECT cluster_name, listen_address FROM system.local;
```

TODO: could be a good idea to maintain the schema using scripts of some kind... (i think i remember seeing a gocql addon project that maintained the schema, need to look into it)