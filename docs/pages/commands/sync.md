---
title: sync
sidebar_icon: refresh-cw
---


Clone any missing repos and sync all `.hike/` repos to the latest `origin/HEAD`.

## Usage

```bash
hk sync [-g group]
```

## Examples

```bash
hk sync                # Sync all groups
hk sync -g backend     # Sync only the backend group
```

## What it does

For each repo in the config (or a specific group):

1. **Not cloned yet** — clones into `.hike/{group}/{repo}/`
2. **Already cloned** — runs `git fetch --all --prune` then `git reset --hard origin/{branch}`

All operations run in parallel.

:::caution
Sync performs a **hard reset** on `.hike/` repos. Any uncommitted or unpushed changes in `.hike/` directories **will be lost**. This is by design — `.hike/` repos are clean mirrors of the remote. Always work in project worktrees, never directly in `.hike/`.
:::

## When to run

- After `hk init` and editing the config
- After adding new repos to the config
- Before creating a new project to work from the latest code
