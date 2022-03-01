#!/usr/bin/env bash

set -e

if [ "$EUID" -eq 0 ]; then
    echo "Error: this script should not be run as root"
    exit
fi


VERSION="v0.1"
UMLFS_VERSION="v0"
UMLKERN_VERSION="v0"

OS=`uname`
if [ "$OS" = "Linux" ]; then
    GOOS="linux"
else
    echo "Unknown operating system $OS"
    exit 1
fi

ARCH=`uname -m`
if [ "$ARCH" = "x86_64" ]; then 
    ARCH="amd64"
else
    echo "Unsupported cpu arch $ARCH"
    exit 1
fi

KOBLE_URL="https://github.com/b177y/koble/releases/download/${VERSION}/koble_${GOOS}_${ARCH}"
DIR="${HOME}/.local/bin"
mkdir -p "$DIR"

# Install podman dependencies
sudo apt install ca-certificates
wget https://packagecloud.io/shiftkey/desktop/gpgkey -O - | sudo apt-key add -
source /etc/os-release
echo "deb https://download.opensuse.org/repositories/devel:/kubic:/libcontainers:/stable/xUbuntu_${VERSION_ID}/ /" | sudo tee /etc/apt/sources.list.d/devel:kubic:libcontainers:stable.list
wget -qO - "https://download.opensuse.org/repositories/devel:/kubic:/libcontainers:/stable/xUbuntu_${VERSION_ID}/Release.key" | sudo apt-key add -
sudo apt update
sudo apt -y install podman
sudo sysctl kernel.unprivileged_userns_clone=1
sudo touch /etc/subuid /etc/subgid
sudo usermod --add-subuids 100000-165535 --add-subgids 100000-165535 $USER
systemctl --user enable --now podman.socket
podman pull docker.io/b177y/uml-runner
podman pull docker.io/b177y/koble-deb

# Install UML dependencies
# wget "https://sourceforge.net/projects/koble-fs/files/v0/koble-fs.tar.bz2/download" -O /tmp/koble-fs.tar.bz2
# wget "https://sourceforge.net/projects/koble-kernel/files/v0/koble-kernel.tar.bz2/download" -O /tmp/koble-kernel.tar.bz2
# mkdir -p ~/.local/share/uml/images
# mkdir -p ~/.local/share/uml/kernel
# tar -C ~/.local/share/uml/images -xjvf /tmp/koble-fs.tar.bz2
# tar -C ~/.local/share/uml/kernel -xjvf /tmp/koble-kernel.tar.bz2
# mv ~/.local/share/uml/kernel/linux ~/.local/share/uml/kernel/koble-kernel

echo "Downloading Koble ${VERSION}..."
wget "$KOBLE_URL" -O "${DIR}/koble"
chmod +x "${DIR}/koble"

# work out which shell (bash/zsh/fish)
USER_SHELL=$(basename "$SHELL")

if [ "$USER_SHELL" = "zsh" ]; then
    _PATH=$(zsh -c echo \$PATH)
    if [[ ":$PATH:" != *":$HOME/.local/bin:"* ]]; then
        echo "export PATH=\"\$PATH:\$HOME/.local/bin\" # Koble PATH" >> ~/.zshrc
        echo "Added koble to PATH. Run source ~/.zshrc for it to take effect in this shell session."
    fi
    grep "# Koble completion" ~/.zshrc || \
        echo "source <($HOME/.local/bin/koble completion zsh) # Koble completion" >> ~/.zshrc
elif [ "$USER_SHELL" = "bash" ]; then
    _PATH=$(bash -c echo \$PATH)
    if [[ ":$PATH:" != *":$HOME/.local/bin:"* ]]; then
        echo "export PATH=\"\$PATH:\$HOME/.local/bin\" # Koble PATH" >> ~/.bashrc
        echo "Added koble to PATH. Run source ~/.bashrc for it to take effect in this shell session."
    fi
    grep "# Koble completion" ~/.bashrc || \
        echo "source <($HOME/.local/bin/koble completion bash) # Koble completion" >> ~/.bashrc
elif [ "$USER_SHELL" = "fish" ]; then
    echo "fish shell is not supported for automatic setup."
else
    echo "Shell $USER_SHELL is not currently supported."
fi
