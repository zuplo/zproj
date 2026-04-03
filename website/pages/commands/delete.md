---
title: delete
sidebar_icon: trash-2
---


Remove a project and clean up all git worktrees and branches.

## Usage

```bash
hk delete [project-name]
```

## Examples

```bash
cd platform-my-feature
hk delete                        # Auto-detect from cwd

hk delete platform-my-feature    # Explicit name
```

## What it does

1. Reads `.hike-project.json` to determine the group
2. Removes git worktrees for each repo (parallel)
3. Deletes local branches
4. Removes the project directory

Auto-detects the project from cwd by looking for `.hike-project.json`.
