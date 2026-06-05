#!/usr/bin/env bash
set -euo pipefail

BINARY_NAME="fn-cleaner"
INSTALL_PATH="/usr/local/bin/${BINARY_NAME}"

echo "Removing ${INSTALL_PATH}..."
rm -f "${INSTALL_PATH}"

if [ ! -f "${INSTALL_PATH}" ]; then
    echo "Done! '${BINARY_NAME}' has been uninstalled."
else
    echo "Error: Failed to remove '${INSTALL_PATH}'."
    exit 1
fi
