---
title: Overview
sidebar_icon: mountain
---


**hike** (`hk`) is a fast CLI tool for managing multi-repo development workspaces using git worktrees.

## Why hike?

When you work across multiple repositories, you need isolated workspaces per feature or task. Git worktrees are perfect for this — they let you have multiple working directories for the same repo, each on a different branch, without cloning the repo multiple times.

**hike** manages worktrees across many repos at once:

- **One command** creates a workspace with worktrees for every repo in your group
- **Each workspace** gets its own branch across all repos, a VS Code `.code-workspace` file, and lifecycle hooks
- **Parallel operations** — clones, syncs, pulls, pushes, and hooks all run concurrently
- **Clean teardown** — delete a workspace and all worktrees/branches are cleaned up

## Built for AI workflows

When multiple AI coding agents work on the same codebase, they need isolation. Without it, agents step on each other's changes, create merge conflicts, and produce inconsistent state.

Git worktrees give each agent its own working directory and branch without cloning the entire repo. hike manages this across many repos at once — create a workspace, hand it to an agent, clean up when done.

## Directory structure

```
my-projects/
├── hike.yaml                          # Your config
├── .hike/                             # Hidden — main repo clones
│   └── platform/
│       ├── my-app/
│       └── shared-lib/
├── platform-my-feature/               # A workspace
│   ├── .hike-project.json             # Metadata
│   ├── platform-my-feature.code-workspace
│   ├── my-app/                        # git worktree
│   └── shared-lib/                    # git worktree
└── platform-bold-cedar/               # Another workspace (random name)
    └── ...
```

## Quick start

```bash
# Install
curl -fsSL https://raw.githubusercontent.com/zuplo/hike/main/install.sh | sh

# Set up
mkdir my-projects && cd my-projects
hk init        # creates hike.yaml
# edit hike.yaml to add your repos
hk sync        # clones repos

# Create a workspace
hk platform my-feature    # → platform-my-feature/
hk platform               # → platform-bold-cedar/ (random)

# Work inside a project
cd platform-my-feature
hk pull        # pull all repos
hk push        # push all repos
hk status      # git status across repos
hk delete      # clean up
```
