---
title: update
sidebar_icon: refresh-cw
---


Self-update to the latest release.

## Usage

```bash
hk update
```

## What it does

1. Checks the latest release via the GitHub API
2. Downloads the archive for your OS/architecture
3. Replaces `~/.hike/bin/hike`

No sudo needed — the symlinks in `/usr/local/bin` point to `~/.hike/bin/hike` and automatically pick up the new version.

## Automatic update check

The CLI checks once per day in the background (2-second timeout). If outdated:

```
A new version of hike is available: 0.10.0 → 0.11.0
Run 'hike update' to upgrade.
```

State stored in `~/.hike/update-check.json`.
