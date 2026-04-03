---
title: pull
sidebar_icon: arrow-down
---


Run `git pull --ff-only` on every repo in the project.

## Usage

```bash
hk pull [project-name]
```

## Examples

```bash
cd platform-my-feature
hk pull                          # Auto-detect from cwd

hk pull platform-my-feature      # Explicit name
```

## Details

Runs fast-forward only pulls in parallel across all repos. If a fast-forward isn't possible, the pull fails for that repo and you'll need to resolve it manually.

Auto-detects the project from cwd by looking for `.hike-project.json`.
