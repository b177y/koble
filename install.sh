#!/usr/bin/env bash

set -e

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

echo "Downloading Koble ${VERSION}..."
wget "$KOBLE_URL" -O "${DIR}/koble"
# chmod +x "${DIR}/koble"

# work out which shell (bash/zsh/fish)
USER_SHELL=$(basename "$SHELL")

if [ "$USER_SHELL" -eq "zsh" ]; then
    _PATH=$(zsh -c echo \$PATH)
    if [[ ":$PATH:" != *":$HOME/.local/bin:"* ]]; then
        echo "export PATH=\"\$PATH:\$HOME\.local/bin # Koble PATH" >> ~/.zprofile
        echo "Added koble to PATH. Run source ~/.zshrc for it to take effect in this shell session."
    fi
    grep "# Koble completion" ~/.zshrc || \
        echo "source <(koble completion zsh) # Koble completion" >> ~/.zshrc
elif [ "$USER_SHELL" -eq "bash" ]; then
    _PATH=$(bash -c echo \$PATH)
    if [[ ":$PATH:" != *":$HOME/.local/bin:"* ]]; then
        echo "export PATH=\"\$PATH:\$HOME\.local/bin # Koble PATH" >> ~/.profile
        echo "Added koble to PATH. Run source ~/.bashrc for it to take effect in this shell session."
    fi
    grep "# Koble completion" ~/.bashrc || \
        echo "source <(koble completion bash) # Koble completion" >> ~/.bashrc
elif [ "$USER_SHELL" -eq "fish" ]; then
    echo "fish shell is not supported for automatic setup."
else
    echo "Shell $USER_SHELL is not currently supported."
fi

# Install podman dependencies
source /etc/os-release
echo "deb https://download.opensuse.org/repositories/devel:/kubic:/libcontainers:/stable/xUbuntu_${VERSION_ID}/ /" | sudo tee /etc/apt/sources.list.d/devel:kubic:libcontainers:stable.list
curl -L "https://download.opensuse.org/repositories/devel:/kubic:/libcontainers:/stable/xUbuntu_${VERSION_ID}/Release.key" | sudo apt-key add -
sudo apt update
sudo apt -y install podman
sudo sysctl kernel.unprivileged_userns_clone=1
sudo touch /etc/subuid /etc/subgid
sudo usermod --add-subuids 100000-165535 --add-subgids 100000-165535 $USER

# Install UML dependencies
wget "https://github.com/b177y/koble-fs/releases/download/${UMLFS_VERSION}/koble-fs.tar.bz2" -O /tmp/koble-fs.tar.bz2
wget "https://github.com/b177y/koble-kernel/releases/download/${UMLKERN_VERSION}/koble-kernel.tar.bz2" -O /tmp/koble-kernel.tar.bz2
mkdir -p ~/.local/share/uml/images
mkdir -p ~/.local/share/uml/kernel
tar -C ~/.local/share/uml/images -xjvf /tmp/koble-fs.tar.bz2
tar -C ~/.local/share/uml/kernel -xjvf /tmp/koble-kernel.tar.bz2
mv ~/.local/share/uml/kernel/linux ~/.local/share/uml/kernel/koble-kernel
