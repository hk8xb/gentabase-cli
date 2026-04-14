#!/bin/sh
# Gentabase CLI installer
# Usage: curl -fsSL https://gentabase.dev/install.sh | sh
set -e

VERSION="0.1.2"
REPO="hk8xb/gentabase-cli"
INSTALL_DIR="${GENTABASE_INSTALL_DIR:-/usr/local/bin}"

# Detect platform
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

case "$OS" in
  darwin) PLATFORM="darwin" ;;
  linux)  PLATFORM="linux" ;;
  *)      echo "Unsupported OS: $OS"; exit 1 ;;
esac

case "$ARCH" in
  x86_64|amd64) ARCH="amd64" ;;
  arm64|aarch64) ARCH="arm64" ;;
  *)             echo "Unsupported architecture: $ARCH"; exit 1 ;;
esac

ASSET="gentabase_${PLATFORM}_${ARCH}.tar.gz"
URL="https://github.com/${REPO}/releases/download/v${VERSION}/${ASSET}"

echo "Installing Gentabase CLI v${VERSION} (${PLATFORM}/${ARCH})..."

TMPDIR=$(mktemp -d)
trap 'rm -rf "$TMPDIR"' EXIT

curl -fsSL "$URL" -o "${TMPDIR}/${ASSET}"
tar xzf "${TMPDIR}/${ASSET}" -C "$TMPDIR" gentabase

if [ -w "$INSTALL_DIR" ]; then
  mv "${TMPDIR}/gentabase" "${INSTALL_DIR}/gentabase"
else
  echo "Installing to ${INSTALL_DIR} (requires sudo)..."
  sudo mv "${TMPDIR}/gentabase" "${INSTALL_DIR}/gentabase"
fi

chmod +x "${INSTALL_DIR}/gentabase"

echo ""
echo "Gentabase CLI installed successfully!"
echo ""
echo "  gentabase --help     Show available commands"
echo "  gentabase login      Authenticate with your account"
echo "  gentabase init       Initialize a new project"
echo ""
