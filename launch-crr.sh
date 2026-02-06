#!/bin/bash
# Cool Retro Radio Launcher for macOS
# Launches crr inside Cool Retro Term with retro profile

CRR_PATH="$(dirname "$0")/crr"
CRT_APP="/Applications/Cool Retro Term.app"

# Check if Cool Retro Term is installed
if [ ! -d "$CRT_APP" ]; then
    echo "Cool Retro Term not found!"
    echo "Please install it from: https://github.com/Swordfish90/cool-retro-term"
    echo ""
    echo "On macOS with Homebrew:"
    echo "  brew install --cask cool-retro-term"
    echo ""
    echo "Running in current terminal instead..."
    exec "$CRR_PATH"
fi

# Check if crr binary exists
if [ ! -f "$CRR_PATH" ]; then
    echo "crr binary not found at $CRR_PATH"
    echo "Please build it first: go build"
    exit 1
fi

# Launch Cool Retro Term with crr
# Using default profile, you can customize in CRT settings
open -a "Cool Retro Term" --args -e "$CRR_PATH"
