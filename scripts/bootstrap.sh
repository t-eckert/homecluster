#!/bin/bash

echo "Starting Bootstrap, this script will install Docker daemon and set up a " \
"K3s cluster using the k3d tool."

echo "Removing any existing Docker daemons"
sudo apt-get remove docker docker-engine docker.io containerd runc

echo "Installing Docker"
sudo apt-get update
sudo apt-get install \
    ca-certificates \
    curl \
    gnupg \
    lsb-release
curl -fsSL https://download.docker.com/linux/debian/gpg | sudo gpg --dearmor -o /usr/share/keyrings/docker-archive-keyring.gpg
echo \
  "deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/docker-archive-keyring.gpg] https://download.docker.com/linux/debian \
  $(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
sudo apt-get update
sudo apt-get install docker-ce docker-ce-cli containerd.io

echo "Installing kubectl"
curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl"

echo "Installing k3d (the K3s command line tool)"
curl -s https://raw.githubusercontent.com/k3d-io/k3d/main/install.sh | bash