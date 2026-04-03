---
title: Configuration
sidebar_icon: settings
---


hike is configured with a single `hike.yaml` file at the root of your projects directory. Run `hk init` to generate an annotated example.

## Full example

```yaml
git:
  org: your-org
  provider: github
  host: github.com
  ssh: false

hooks:
  onCreate: pnpm install

groups:
  platform:
    default: true
    aliases: [plat]
    hooks:
      onCreate: npm install
    repos:
      - my-app
      - shared-lib
      - repo: api
        name: api-service
        branch: develop
        hooks:
          onCreate: yarn install

  marketing:
    aliases: [mktg]
    repos:
      - website
      - cms

templates:
  variables:
    ORG: your-org
    TEAM: platform
```

## `git`

Configure default git provider settings. When `org` is set, repos can use short names instead of full URLs.

| Field | Type | Default | Description |
|-------|------|---------|-------------|
| `org` | string | — | Default organization/owner for short repo names |
| `provider` | string | `github` | Git provider: `github`, `gitlab`, or `bitbucket` |
| `host` | string | auto | Git host. Auto-detected from provider, or set for self-hosted |
| `ssh` | bool | `false` | Use SSH URLs instead of HTTPS |

### URL expansion examples

| Config | Short name | Expanded URL |
|--------|-----------|-------------|
| `org: acme` | `my-app` | `https://github.com/acme/my-app.git` |
| `org: acme, ssh: true` | `my-app` | `git@github.com:acme/my-app.git` |
| `org: acme, provider: gitlab` | `my-app` | `https://gitlab.com/acme/my-app.git` |
| `org: acme, host: git.internal.com` | `my-app` | `https://git.internal.com/acme/my-app.git` |

Full URLs (SSH or HTTPS) are never modified.

## `groups`

Groups organize repos that are worked on together. Each group's main clones live in `.hike/{group}/`.

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `default` | bool | no | Use this group when none is specified. Only one group can be default. |
| `aliases` | string[] | no | Shorthand names (e.g. `[plat]` for `platform`) |
| `hooks` | object | no | Group-level lifecycle hooks |
| `repos` | list | yes | List of repos |

If only one group exists, it's automatically the default.

### Repo formats

**Plain string** — just the repo name or full URL:

```yaml
repos:
  - my-app                                # Short name (requires git.org)
  - git@github.com:other-org/special.git  # Full SSH URL
  - https://github.com/org/repo.git       # Full HTTPS URL
```

**Object** — with overrides:

```yaml
repos:
  - repo: my-app          # Required: repo name or full URL
    name: app             # Optional: directory name override
    branch: develop       # Optional: default branch (default: main)
    hooks:                # Optional: repo-level hooks
      onCreate: yarn install
```

## `hooks`

Lifecycle hooks run commands after project creation. Most specific wins: **repo > group > global**.

```yaml
hooks:
  onCreate: npm install           # Global default

groups:
  platform:
    hooks:
      onCreate: pnpm install      # Overrides global for this group
    repos:
      - my-app                    # Uses pnpm install
      - repo: legacy-app
        hooks:
          onCreate: yarn          # Overrides group for this repo
```

All hooks run in parallel across repos.

| Hook | When |
|------|------|
| `onCreate` | After worktrees are created. Runs in each repo directory. |

## `templates`

Files in `.template/` at the root are processed with Go's `text/template` and copied into each new project.

```yaml
templates:
  variables:
    ORG: your-org
```

Available variables: `{{.ProjectName}}`, `{{.Group}}`, and any custom key from `templates.variables`.
