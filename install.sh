#!/usr/bin/env bash
set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
BINARY_NAME="fn-cleaner"
INSTALL_PATH="/usr/local/bin/${BINARY_NAME}"

echo "Building ${BINARY_NAME}..."
go build -o "${BINARY_NAME}" "${SCRIPT_DIR}/main.go"

if [ ! -f "${BINARY_NAME}" ]; then
    echo "Error: Build failed. Could not find binary."
    exit 1
fi

echo "Installing to ${INSTALL_PATH}..."
mv "${BINARY_NAME}" "${INSTALL_PATH}"
chmod +x "${INSTALL_PATH}"

echo "Done! '${BINARY_NAME}' is now installed at ${INSTALL_PATH}."
