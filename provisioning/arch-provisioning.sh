#!/bin/zsh

# Update packages
sudo pacman -Syu

# install docker if needed
sudo pacman -S --needed docker
sudo pacman -S --needed docker-compose

# Check if the user is in the docker group
if id -nG "$USER" | grep -qw "docker"; then
    echo "User $USER is already in the docker group."
else
    echo "User $USER is not in the docker group. Adding now..."
    sudo usermod -aG docker "$USER"

    # Check if the command succeeded
    if [ $? -eq 0 ]; then
        echo "User $USER was successfully added to the docker group."
    else
        echo "Failed to add user $USER to the docker group."
        exit 1
    fi
fi

# Config docker deamon on demand
sudo systemctl disable docker.service
sudo systemctl stop docker.service
sudo systemctl enable docker.socket
sudo systemctl start docker.socket

# install nvm and node
curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.40.2/install.sh | bash
nvm install