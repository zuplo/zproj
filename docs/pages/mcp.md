---
title: MCP Server
sidebar_icon: cpu
---


hike includes a built-in [Model Context Protocol](https://modelcontextprotocol.io) (MCP) server so AI assistants can manage workspaces autonomously.

## Setup

### Claude Code / Claude Desktop

Add to `~/.claude/claude_desktop_config.json`:

```json
{
  "mcpServers": {
    "hike": {
      "command": "hike",
      "args": ["mcp"],
      "cwd": "/path/to/your/projects"
    }
  }
}
```

:::tip
Set `cwd` to the directory containing your `hike.yaml`.
:::

### Cursor

Add the same configuration to Cursor's MCP settings.

### Any MCP client

Run `hike mcp` to start a stdio-based MCP server.

## Available tools

| Tool | Description |
|------|-------------|
| `create_project` | Create a project. Params: `group`, `name`, `color`. |
| `delete_project` | Delete a project. Requires `name`. |
| `list_projects` | List all projects. |
| `pull_project` | Pull all repos in a project. Requires `name`. |
| `push_project` | Push all repos in a project. Requires `name`. |
| `project_status` | Git status across repos. Requires `name`. |
| `sync_repos` | Clone missing repos and sync. Optional `group`. |

## Why MCP for AI workflows?

When multiple AI agents work on a codebase, they need isolation. Without it, agents create merge conflicts and inconsistent state.

**Git worktrees solve this** — each agent gets its own working directory and branch without cloning the entire repo.

**hike makes it practical at scale** — one MCP call creates a full multi-repo workspace, another tears it down.

### Typical agent workflow

1. Agent calls `create_project` → gets an isolated workspace
2. Agent works in the workspace
3. Agent calls `push_project` → pushes changes
4. Agent calls `delete_project` → cleans up
