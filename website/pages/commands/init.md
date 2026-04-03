---
title: init
sidebar_icon: file-plus
---

Create a new `hike.yaml` configuration file and `.template/` directory in the current directory.

## Usage

```bash
hk init
```

## What it creates

| File | Purpose |
|------|---------|
| `hike.yaml` | Annotated config with git provider, groups, hooks, and template examples |
| `.template/{{.ProjectName}}.code-workspace` | VS Code workspace template — creates a workspace file for each project with all repos as folders |
| `.template/.claude/settings.local.json` | Claude Code settings template — grants file access to all repos in the project |

Prompts before overwriting an existing `hike.yaml`. Templates are only created if they don't already exist.

## Typical workflow

```bash
mkdir my-projects && cd my-projects
hk init                    # Create config + templates
# Edit hike.yaml
hk sync                    # Clone repos
hk platform my-feature     # Start working
```

## Customizing templates

After `hk init`, edit the files in `.template/` to match your workflow. Templates are processed with Go's `text/template` and support looping over repos, conditional blocks, and custom variables. See the [Configuration](/docs/configuration) page for full template documentation.
