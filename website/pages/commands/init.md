---
title: init
sidebar_icon: file-plus
---


Create a new `hike.yaml` configuration file in the current directory.

## Usage

```bash
hk init
```

## What it does

Writes an annotated example `hike.yaml` with git provider config, an example group with repos, and commented examples for hooks, aliases, and templates.

Prompts before overwriting an existing file.

## Typical workflow

```bash
mkdir my-projects && cd my-projects
hk init                    # Create config
# Edit hike.yaml
hk sync                    # Clone repos
hk platform my-feature     # Start working
```
