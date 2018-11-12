# -*- mode: ruby -*-
# vi: set ft=ruby :

require 'fileutils'

CONFIG = File.join(File.dirname(__FILE__), File.join("cassandra","config.rb"))

# config.rb defaults
$instances = 3
$instance_name_prefix = "cassandra"
$vm_memory = 1024
$vm_cpus = 1
$vm_cpuexecutioncap = 50
$forwarded_ports = {}

if File.exist?(CONFIG)
  require CONFIG
end

Vagrant.configure("2") do |config|

  config.vm.box = "ubuntu/xenial64"

  #config.ssh.forward_agent = true
  #config.ssh.insert_key = false

  (1..$instances).each do |i|

    config.vm.define vm_name = "%s-%02d" % [$instance_name_prefix, i] do |config|

      config.vm.hostname = vm_name
      config.vm.network "private_network", ip: "10.10.10.#{i+30}"

      $forwarded_ports.each do |guest, host|
        config.vm.network "forwarded_port", guest: guest, host: host, auto_correct: true
      end

      config.vm.provision "shell", path: "cassandra/cassandra.sh", args: "#{i}"

      config.vm.provider "virtualbox" do |vm|
        vm.memory = $vm_memory
        vm.cpus = $vm_cpus
        vm.customize ["modifyvm", :id, "--cpuexecutioncap", "#{$vm_cpuexecutioncap}"]
      end

    end

  end

end
