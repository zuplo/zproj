---
title: push
sidebar_icon: arrow-up
---


Run `git push` on every repo in the project.

## Usage

```bash
hk push [project-name]
```

## Examples

```bash
cd platform-my-feature
hk push                          # Auto-detect from cwd

hk push platform-my-feature      # Explicit name
```

## Details

Pushes the current branch of each worktree to its remote, in parallel.

Auto-detects the project from cwd by looking for `.hike-project.json`.
