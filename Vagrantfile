# -*- mode: ruby -*-
# vi: set ft=ruby :

Vagrant.configure("2") do |config|
  config.vm.box = "acloudfan/hlfdev2.0-0"

  # Define the "active" machine.
  config.vm.define "active" do |active|
    active.vm.hostname = "active"
    active.vm.network "private_network", ip: "192.168.56.10"

    # Forwarding commonly used Fabric ports (adjust if needed)
   active.vm.network "forwarded_port", guest: 7054, host: 7054 # Orderer
    active.vm.network "forwarded_port", guest: 7051, host: 7051  # Peer
    active.vm.network "forwarded_port", guest: 7050, host: 7050  # Peer

    #active.vm.network "forwarded_port", guest: 7052, host: 7052  # Chaincode (dev mode)
    #active.vm.network "forwarded_port", guest: 7053, host: 7053  # Event service

    active.vm.provider "virtualbox" do |vb|
      vb.memory = "1024"
    end
  end

  # Define the "passive" machine.
  config.vm.define "passive" do |passive|
    passive.vm.hostname = "passive"
    passive.vm.network "private_network", ip: "192.168.56.11"

    # Forwarding ports for second peer/orderer if needed
    passive.vm.network "forwarded_port", guest: 8050, host: 8050
    passive.vm.network "forwarded_port", guest: 8051, host: 8051

    passive.vm.provider "virtualbox" do |vb|
      vb.memory = "1024"
    end
  end
end
