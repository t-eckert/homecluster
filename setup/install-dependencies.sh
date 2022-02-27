#!/bin/bash

echo "Installing dependencies"
cat << EOM
This script installs the following packages:
  - docker
  - kubectl
  - k3d
EOM

echo "Installing snapd core"
sudo snap install core

echo "Installing docker"
sudo snap install docker

echo "Installing kubectl"
sudo snap install kubectl

echo "Installing k3d"
curl -s https://raw.githubusercontent.com/k3d-io/k3d/main/install.sh | bash