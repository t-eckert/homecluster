#!/bin/bash

echo "Installing dependencies"
cat << EOM
This script uses snapd to install the following packages:
  - docker
  - kubectl
EOM

sudo snap install core
sudo snap install docker