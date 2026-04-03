package cmd

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/zuplo/hike/internal/config"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Create a new hike.yaml configuration file and templates",
	RunE: func(cmd *cobra.Command, args []string) error {
		cwd, err := os.Getwd()
		if err != nil {
			return err
		}

		cfgPath := filepath.Join(cwd, config.ConfigFile)

		if _, err := os.Stat(cfgPath); err == nil {
			if !promptYesNo(fmt.Sprintf("%s already exists. Overwrite?", config.ConfigFile)) {
				fmt.Println("Aborted.")
				return nil
			}
		}

		if err := os.WriteFile(cfgPath, []byte(exampleConfig), 0644); err != nil {
			return fmt.Errorf("writing config: %w", err)
		}
		fmt.Printf("Created %s\n", cfgPath)

		// Create .template directory with example templates
		tmplDir := filepath.Join(cwd, ".template")
		if err := os.MkdirAll(filepath.Join(tmplDir, ".claude"), 0755); err != nil {
			return fmt.Errorf("creating template dir: %w", err)
		}

		// VS Code workspace template
		wsPath := filepath.Join(tmplDir, "{{.ProjectName}}.code-workspace")
		if _, err := os.Stat(wsPath); os.IsNotExist(err) {
			if err := os.WriteFile(wsPath, []byte(workspaceTemplate), 0644); err != nil {
				return fmt.Errorf("writing workspace template: %w", err)
			}
			fmt.Println("Created .template/{{.ProjectName}}.code-workspace")
		}

		// Claude settings template
		claudePath := filepath.Join(tmplDir, ".claude", "settings.local.json")
		if _, err := os.Stat(claudePath); os.IsNotExist(err) {
			if err := os.WriteFile(claudePath, []byte(claudeTemplate), 0644); err != nil {
				return fmt.Errorf("writing claude template: %w", err)
			}
			fmt.Println("Created .template/.claude/settings.local.json")
		}

		fmt.Println("\nEdit hike.yaml to add your repos, then run 'hk sync' to clone them.")
		fmt.Println("Customize templates in .template/ — they are processed for each new project.")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

const exampleConfig = `# Git provider defaults — repos can use just the name (e.g. "my-app")
# instead of full URLs when org is set.
git:
  # provider: github  # github (default), gitlab, bitbucket, or any host
  # host: github.com  # auto-detected from provider, or set for self-hosted
  org: your-org
  # ssh: true         # use SSH URLs (default: false, uses HTTPS)

# Each group contains a set of repos that are cloned and managed together.
# Every group gets its own [groupname]/ directory.
# Set 'default: true' on a group to use it when --group is not specified.
groups:
  platform:
    default: true
    # aliases: [plat]  # Optional: use 'plat' as shorthand for 'platform'
    repos:
      # With git.org set, just use the repo name:
      - my-app
      - shared-lib

      # Or use a full URL when needed:
      # - git@github.com:other-org/special-repo.git

      # Object form lets you override name or branch:
      # - repo: api
      #   name: api
      #   branch: develop

  # Add more groups to organize repos separately:
  # backend:
  #   aliases: [be]
  #   repos:
  #     - api-service
  #     - worker

# Lifecycle hooks — run after creating a project.
# Repo-level overrides group-level, which overrides global.
# hooks:
#   onCreate: npm install

# Optional: variables available in templates
# templates:
#   variables:
#     ORG: your-org
`

// Go template for VS Code workspace file.
// The filename itself uses {{.ProjectName}} so it's named after the project.
const workspaceTemplate = `{
  "folders": [
{{- range $i, $repo := .Repos}}
    {{- if $i}},{{end}}
    { "path": "{{$repo.Name}}" }
{{- end}}
  ]{{if .Color}},
  "settings": {
    "workbench.colorCustomizations": {
      "titleBar.activeBackground": "{{.Color}}",
      "titleBar.activeForeground": "#ffffff",
      "titleBar.inactiveBackground": "{{.Color}}",
      "titleBar.inactiveForeground": "#cccccc"
    }
  }{{end}}
}
`

// Go template for Claude Code settings.
// Grants file access to all repo directories in the project.
const claudeTemplate = `{
  "additionalDirectories": [
{{- range $i, $repo := .Repos}}
    {{- if $i}},{{end}}
    "{{$repo.Name}}"
{{- end}}
  ]
}
`

func promptYesNo(question string) bool {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%s [y/N] ", question)
	answer, _ := reader.ReadString('\n')
	answer = strings.TrimSpace(strings.ToLower(answer))
	return answer == "y" || answer == "yes"
}
