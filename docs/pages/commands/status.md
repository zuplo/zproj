---
title: status
sidebar_icon: activity
---


Show the branch, dirty state, and ahead/behind status for each repo in the project.

## Usage

```bash
hk status [project-name]
```

## Example

```bash
$ hk status
Project: platform-my-feature

  my-app       branch: platform-my-feature
  shared-lib   branch: platform-my-feature  [dirty]
  api          branch: platform-my-feature  (1	0)
```

## Output columns

| Column | Description |
|--------|-------------|
| Repo name | The repository directory |
| Branch | Current branch of the worktree |
| `[dirty]` | Uncommitted changes present |
| `(ahead behind)` | Commits ahead/behind `origin/{branch}` |

Auto-detects the project from cwd by looking for `.hike-project.json`.
