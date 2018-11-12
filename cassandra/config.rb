# Size of the cluster
$instances = 1

# Base name of the VM
# Example: "cassandra-01", "cassandra-02", "cassandra-${instances}"
$instance_name_prefix = "cassandra"

# VM
$vm_memory = 1024
$vm_cpus = 1
$vm_cpuexecutioncap = 50

# Port forwarding from guest(s) to host machine
# Format: { 80 => 8080, ... }
#$forwarded_ports = {}
