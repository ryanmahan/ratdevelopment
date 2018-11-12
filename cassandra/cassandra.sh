#!/bin/bash

echo -e Running apt-get update/upgrade...
sudo apt-get update -y
sudo apt-get upgrade -y

echo -e Adding Apache Cassandra repository to package sources...
echo "deb http://www.apache.org/dist/cassandra/debian 310x main" | sudo tee -a /etc/apt/sources.list.d/cassandra.sources.list

echo -e Adding Apache Cassandra repository keys...
curl https://www.apache.org/dist/cassandra/KEYS | sudo apt-key add -

echo -e Running apt-get update again...
sudo apt-get update -y

echo -e Installing Cassandra...
sudo apt-get install cassandra -y

echo -e Configuring Cassandra...
sudo service cassandra stop

sudo rm -rf /var/lib/cassandra/data/system/*
sudo cp -f /etc/cassandra/cassandra.yaml /etc/cassandra/cassandra.yaml.backup

sudo sed -r -i "s/(cluster_name):(.*)/\1: 'Cluster $1'/g" /etc/cassandra/cassandra.yaml
sudo sed -r -i "s/(listen_address):(.*)/\1: \"10.10.10.$((30+$1))\"/g" /etc/cassandra/cassandra.yaml
sudo sed -r -i "s/(- seeds):(.*)/\1: \"10.10.10.31\"/g" /etc/cassandra/cassandra.yaml
sudo sed -r -i "s/(rpc_address):(.*)/\1: 10.10.10.$((30+$1))/g" /etc/cassandra/cassandra.yaml
echo "auto_bootstrap: false" >> /etc/cassandra/cassandra.yaml

echo -e Starting Cassandra service...
sudo service cassandra start

#sudo chmod u+x /vagrant/cassandra/*.sh

echo -e Cassandra service started!
bash /vagrant/scripts/migrate.sh "10.10.10.$((30+$1))"
