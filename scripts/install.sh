#!/bin/bash

# Installation script for crossfhir
set -e

# Variables
GITHUB_REPO="TheMomentumAI/crossfhir"
BINARY_NAME="crossfhir"
INSTALL_DIR="/usr/local/bin"

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Detect OS and architecture
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

# Convert architecture names to match release artifacts
case ${ARCH} in
    x86_64)
        ARCH="x86_64"
        ;;
    arm64|aarch64)
        ARCH="arm64"
        ;;
    *)
        echo -e "${RED}Unsupported architecture: ${ARCH}${NC}"
        exit 1
        ;;
esac

# Convert OS names to match release artifacts
case ${OS} in
    darwin)
        OS="Darwin"
        ;;
    linux)
        OS="Linux"
        ;;
    *)
        echo -e "${RED}Unsupported operating system: ${OS}${NC}"
        exit 1
        ;;
esac

# Get the latest release version
echo -e "${BLUE}Fetching latest release...${NC}"
LATEST_VERSION=$(curl -s "https://api.github.com/repos/${GITHUB_REPO}/releases/latest" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')

if [ -z "$LATEST_VERSION" ]; then
    echo -e "${RED}Failed to get latest version${NC}"
    exit 1
fi

# Construct download URL
DOWNLOAD_URL="https://github.com/${GITHUB_REPO}/releases/download/${LATEST_VERSION}/${BINARY_NAME}_${LATEST_VERSION#v}_${OS}_${ARCH}.tar.gz"

# Create temporary directory
TMP_DIR=$(mktemp -d)
trap 'rm -rf -- "$TMP_DIR"' EXIT

echo -e "${BLUE}Downloading ${BINARY_NAME} ${LATEST_VERSION}...${NC}"
curl -sL "$DOWNLOAD_URL" -o "${TMP_DIR}/${BINARY_NAME}.tar.gz"

echo -e "${BLUE}Extracting...${NC}"
tar -xzf "${TMP_DIR}/${BINARY_NAME}.tar.gz" -C "${TMP_DIR}"

echo -e "${BLUE}Installing to ${INSTALL_DIR}...${NC}"
# Ensure install directory exists
sudo mkdir -p "${INSTALL_DIR}"

# Install binary
sudo mv "${TMP_DIR}/${BINARY_NAME}" "${INSTALL_DIR}/"
sudo chmod +x "${INSTALL_DIR}/${BINARY_NAME}"

echo -e "${GREEN}Successfully installed ${BINARY_NAME} ${LATEST_VERSION}${NC}"
echo -e "${BLUE}Test the installation with: ${BINARY_NAME} --version${NC}"