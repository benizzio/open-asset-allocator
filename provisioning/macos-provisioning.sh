#!/bin/zsh
# This script prepares the environment for the open-asset-allocator project on Mac,
# in an equivalent way to the original provisioning script for Linux (which used pacman).
# Co-authored by: OpenCode and Igor Benicio de Mesquita
# Features:
#  - Checks if Docker Desktop is installed; if not, guides the user to install it.
#  - Installs and activates the project-pinned Node.js version using NVM.
#  - Ensures that npm (Node package manager) is working properly.
#  - Installs necessary global npm packages (e.g., rimraf) if they're not already installed.
# Clear messages inform the user about each step being executed.
# The script is idempotent (can be run multiple times without repeating actions already completed).
# Note: This script assumes the default system shell is Zsh (standard on modern macOS).
#
# Differences compared to Linux:
#  - On Linux (e.g., Arch Linux with pacman), the original script installed packages directly via package manager.
#    On macOS we use Homebrew (cask) or specific installers; here we chose to use NVM to manage Node.js installation.
#  - On macOS it's not necessary to use systemctl to start/activate Docker; Docker Desktop manages the daemon internally.
#  - It's also not necessary to configure user groups for Docker. On Linux distributions, the user is typically added to the "docker" group
#    to allow Docker usage without sudo; on macOS, Docker Desktop doesn't require this adjustment because it runs in user mode transparently.
#

script_directory=${0:A:h}
project_root=${script_directory:h}
node_version_file="$project_root/.nvmrc"

# Safety check: ensure it's running on macOS (Darwin)
if [[ "$(uname)" != "Darwin" ]]; then
    echo "This provisioning script is designed to run on macOS (Darwin)."
    exit 1
fi

echo "Starting provisioning for the open-asset-allocator project on macOS..."
echo

# 1. Check Docker Desktop
echo "Checking if Docker Desktop is installed..."
# Uses the -v command to check if the "docker" binary is accessible in PATH
if command -v docker &> /dev/null; then
    echo "Docker Desktop is already installed on the system."
    echo "Make sure the Docker Desktop application is running to use the Docker CLI."
else
    echo "Docker Desktop is NOT installed."
    echo "Please install Docker Desktop for Mac."
    echo "You can download it from: https://www.docker.com/products/docker-desktop"
    echo "Or, if you prefer via Homebrew Cask (if Homebrew is installed):"
    echo "    brew install --cask docker"
    # Note: We don't perform automatic installation of Docker Desktop because it's a GUI application
    # that requires manual steps from the user (download, drag to Applications, grant permissions, etc.).
fi
echo

# 2. Install and select the project Node.js version through NVM
echo "Installing the project Node.js version using NVM..."
export NVM_DIR="${NVM_DIR:-$HOME/.nvm}"
if [[ ! -s "$NVM_DIR/nvm.sh" ]]; then
    echo "NVM is not installed. Starting installation of NVM (Node Version Manager)..."
    if command -v curl &> /dev/null; then
        curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.40.2/install.sh | bash
    elif command -v wget &> /dev/null; then
        wget -qO- https://raw.githubusercontent.com/nvm-sh/nvm/v0.40.2/install.sh | bash
    else
        echo "Error: curl or wget not found to download NVM. Install one of these utilities and run the script again."
        exit 1
    fi
fi

if [[ ! -s "$NVM_DIR/nvm.sh" ]]; then
    echo "Error: Failed to install NVM."
    exit 1
fi

# shellcheck source=/dev/null
\. "$NVM_DIR/nvm.sh"
# shellcheck source=/dev/null
[ -s "$NVM_DIR/bash_completion" ] && \. "$NVM_DIR/bash_completion"

if [[ ! -f "$node_version_file" ]]; then
    echo "Error: Node.js version file not found: $node_version_file"
    exit 1
fi

PROJECT_NODE_VERSION=$(tr -d '[:space:]' < "$node_version_file")
nvm install "$PROJECT_NODE_VERSION"
nvm alias default "$PROJECT_NODE_VERSION"
nvm use "$PROJECT_NODE_VERSION"

if [[ "$(node -v)" != "v$PROJECT_NODE_VERSION" ]]; then
    echo "Error: Failed to activate Node.js $PROJECT_NODE_VERSION."
    exit 1
fi
echo "Node.js $PROJECT_NODE_VERSION successfully installed and activated via NVM."
echo

# 3. Check if npm is functional
echo "Checking if npm (Node package manager) is accessible..."
# The npm command is installed with Node.js; here we confirm its availability
if command -v npm &> /dev/null; then
    NPM_VERSION=$(npm -v)
    echo "npm is installed and accessible (version $NPM_VERSION)."
else
    echo "Error: the npm command is not available. Something went wrong with the Node.js/NVM installation."
    echo "Please check the Node.js installation steps above and try again."
    exit 1
fi
echo
# Ensure npm uses a directory accessible by the user (without sudo)
echo "Configuring npm to use a global packages directory within home..."

# Define alternative directory for global npm installations
NPM_GLOBAL_DIR="$HOME/.npm-global"

# Create the directory if it doesn't exist yet
mkdir -p "$NPM_GLOBAL_DIR"

# Configure npm to use it
npm config set prefix "$NPM_GLOBAL_DIR"

# Ensure the directory is in the user's PATH
# Keep shell variables literal in the generated Zsh configuration.
# shellcheck disable=SC2016
if ! grep -q 'export PATH="$HOME/.npm-global/bin:$PATH"' "$HOME/.zshrc"; then
    echo 'export PATH="$HOME/.npm-global/bin:$PATH"' >> "$HOME/.zshrc"
    echo "Line added to ~/.zshrc to include the global npm packages directory in PATH."
else
    echo "The global npm packages directory is already in the PATH in ~/.zshrc."
fi

# Apply the changes immediately in the current session
export PATH="$HOME/.npm-global/bin:$PATH"

echo "npm will now use the directory $NPM_GLOBAL_DIR for global packages (without needing sudo)."

# 4. Install necessary global npm packages (if not already installed)
echo "Checking and installing necessary global npm packages..."
GLOBAL_PACKAGES=("rimraf")  # list here the global packages needed for the project
for pkg in "${GLOBAL_PACKAGES[@]}"; do
    # Check if the package executable already exists in PATH (indicating global installation)
    if command -v "$pkg" &> /dev/null; then
        echo "Global package '$pkg' is already installed."
    else
        echo "Installing global package '$pkg'..."
        npm install -g "$pkg"
    fi
done
echo

# 5. Final messages to the user
echo "Provisioning successfully completed!"
echo "Final tips:"
echo " - If you just installed Docker Desktop, open the Docker application to complete the setup (you need to agree to the terms of service on first run)."
echo " - The Node.js and npm environment is ready to use."
echo " - Note: if this was the first installation of Node.js via NVM, open a new terminal or run 'source ~/.zshrc' "
echo "   to load the NVM definitions and have access to the 'node' and 'npm' commands."
echo
echo "All set! You can now proceed with using the open-asset-allocator project."
