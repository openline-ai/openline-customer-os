#!/bin/bash

echo "  β³ Openline dependency check..."

# Xcode
xcode-select -p
if [ $? -eq 0 ]; then
    echo "  β Xcode"
else
    echo "  π¦¦ Installing Xcode.  This may take awhile, please let the script do it's thing.  It will prompt when completed."
    xcode-select --install
    if [ $? -eq 0 ]; then
        echo "  β Xcode"
    else
        echo "  β Xcode installation failed"
    fi
fi

# Homebrew
if [[ $(brew --version) == *"Homebrew"* ]];
    then
        echo "  β Homebrew"
    else
        echo "  π¦¦ Installing Homebrew..."
        /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
        if [ $? -eq 0 ]; then
            echo "  β Homebrew"
        else
            echo "  β Homebrew installation failed."
        fi
fi

# Docker
if [[ $(docker --version) == *"Docker version"* ]];
    then
        echo "  β Docker"
    else
        echo "  π¦¦ Installing Docker..."
        echo "  βοΈ This can take a while, please let the script do it's thing.  It will prompt when completed."
        softwareupdate --install-rosetta
        
        if [[ $(arch) == 'arm64' ]]; 
            then
                echo "  Installing Apple silicon version..."
                curl -L https://desktop.docker.com/mac/main/arm64/Docker.dmg --output openline-setup/Docker.dmg
            else
                echo "  Installing Intel version..."
                curl -L https://desktop.docker.com/mac/main/amd64/Docker.dmg --output openline-setup/Docker.dmg
        fi

        sudo hdiutil attach openline-setup/Docker.dmg
        sudo /Volumes/Docker/Docker.app/Contents/MacOS/install
        sudo hdiutil detach /Volumes/Docker

        echo "  β Docker"
        echo "  βοΈPlease open Docker desktop via the GUI to initialize the application before proceeding."
        rm -r openline-setup/Docker.dmg
        echo "  Attempting to open Docker..."
        open -a Docker.app
        read -p "  => Press enter to continue once Docker GUI has opened..."
fi

# Minikube
if [[ $(minikube version) == *"minikube version"* ]];
    then
        echo "  β Minikube"
    else
        echo "  π¦¦ Installing Minikube..."
        brew install minikube
        if [ $? -eq 0 ]; then
            echo "  β Minikube"
        else
            echo "  β Minikube installation failed."
        fi
fi

# Helm
if [[ $(helm version) == *"version.BuildInfo"* ]];
    then
        echo "  β Helm"
    else
        echo "  π¦¦ Installing Helm..."
        brew install helm
        if [ $? -eq 0 ]; then
            echo "  β Helm"
        else
            echo "  β Helm installation failed."
        fi
fi