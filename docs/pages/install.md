---
title: Installation
sidebar_icon: download
---


hike runs on macOS (Apple Silicon, Intel) and Linux (x86_64, ARM64). Both `hike` and `hk` commands are installed automatically.

## Quick install

```bash
curl -fsSL https://raw.githubusercontent.com/zuplo/hike/main/install.sh | sh
```

This script:
1. Detects your OS and architecture
2. Downloads the latest release from GitHub
3. Installs the binary to `~/.hike/bin/hike`
4. Creates symlinks for `hike` and `hk` in `/usr/local/bin` (requires sudo once)

## Manual install

Download from [GitHub Releases](https://github.com/zuplo/hike/releases):

```bash
mkdir -p ~/.hike/bin
tar -xzf hike_*_darwin_arm64.tar.gz -C ~/.hike/bin/
sudo ln -sf ~/.hike/bin/hike /usr/local/bin/hike
sudo ln -sf ~/.hike/bin/hike /usr/local/bin/hk
```

## Updating

No sudo needed after the initial install:

```bash
hk update
```

The CLI checks for updates once per day in the background and notifies you when a new version is available.

## File locations

| Path | Purpose |
|------|---------|
| `~/.hike/bin/hike` | Binary |
| `~/.hike/update-check.json` | Update check state |
| `/usr/local/bin/hike` | Symlink |
| `/usr/local/bin/hk` | Shorthand symlink |

## Uninstall

```bash
sudo rm /usr/local/bin/hike /usr/local/bin/hk
rm -rf ~/.hike
```
