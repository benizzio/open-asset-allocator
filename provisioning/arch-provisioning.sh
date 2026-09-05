#!/bin/zsh
# Provisions Arch Linux with the project-pinned Node.js runtime and container tools.
# Co-authored by: OpenCode and Igor Benicio de Mesquita

script_directory=${0:A:h}
project_root=${script_directory:h}
node_version_file="$project_root/.nvmrc"

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
    if sudo usermod -aG docker "$USER"; then
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

# Install NVM when absent, then load it in this process before installing Node.js.
export NVM_DIR="${NVM_DIR:-$HOME/.nvm}"
if [[ ! -s "$NVM_DIR/nvm.sh" ]]; then
    curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.40.2/install.sh | bash
fi

if [[ ! -s "$NVM_DIR/nvm.sh" ]]; then
    echo "Failed to install NVM."
    exit 1
fi

# shellcheck source=/dev/null
\. "$NVM_DIR/nvm.sh"

if [[ ! -f "$node_version_file" ]]; then
    echo "Node.js version file not found: $node_version_file"
    exit 1
fi

project_node_version=$(tr -d '[:space:]' < "$node_version_file")
nvm install "$project_node_version"
nvm alias default "$project_node_version"
nvm use "$project_node_version"
