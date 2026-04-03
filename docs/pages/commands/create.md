---
title: create (default)
sidebar_icon: plus
---


Create a new project with git worktrees for all repos in a group. This is the **default command** — you don't need to type `create`.

## Usage

```bash
hk [group] [name] [-c color] [-g group] [-n name]
```

## Examples

```bash
hk platform my-feature    # Creates platform-my-feature/
hk platform               # Random name: platform-bold-cedar/
hk my-feature             # Uses default group: platform-my-feature/
hk platform -c purple     # With VS Code title bar color
hk platform -c            # Random color
```

## Argument resolution

1. The first argument is matched against known groups and aliases
2. If it matches a group, the second argument (or a random name) is the project name
3. If it doesn't match a group, it's treated as the project name with the default group
4. The project directory is named `{group}-{name}`

## What it does

For each repo in the group:

1. Creates a git worktree from `.hike/{group}/{repo}`
2. Names the worktree branch after the project (e.g. `platform-my-feature`)
3. Generates a VS Code `.code-workspace` file
4. Processes templates from `.template/`
5. Runs `onCreate` hooks in parallel

## Flags

| Flag | Short | Description |
|------|-------|-------------|
| `--color` | `-c` | VS Code title bar color. Use a name, or `-c` alone for random. |
| `--group` | `-g` | Explicitly set the group |
| `--name` | `-n` | Explicitly set the project name |

## Colors

`blue`, `cyan`, `green`, `indigo`, `lime`, `orange`, `pink`, `purple`, `red`, `rose`, `sky`, `slate`, `teal`, `yellow`
