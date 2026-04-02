package config

import (
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"gopkg.in/yaml.v3"
)

const ConfigFile = "zproj.yaml"

type Config struct {
	Groups    map[string]Group `yaml:"groups"`
	Templates *Templates       `yaml:"templates,omitempty"`
}

type Group struct {
	Repos []Repo `yaml:"-"`
}

type Repo struct {
	URL    string `yaml:"url"`
	Name   string `yaml:"name,omitempty"`
	Branch string `yaml:"branch,omitempty"`
}

type Templates struct {
	Variables map[string]string `yaml:"variables,omitempty"`
}

// RepoName returns the resolved name for a repo.
// If Name is set, returns it. Otherwise derives from the URL.
func (r Repo) RepoName() string {
	if r.Name != "" {
		return r.Name
	}
	return repoNameFromURL(r.URL)
}

// RepoBranch returns the resolved branch for a repo.
// Defaults to "main" if not set.
func (r Repo) RepoBranch() string {
	if r.Branch != "" {
		return r.Branch
	}
	return "main"
}

func repoNameFromURL(u string) string {
	base := filepath.Base(u)
	return strings.TrimSuffix(base, ".git")
}

// UnmarshalYAML supports repos as either plain strings or objects.
func (g *Group) UnmarshalYAML(value *yaml.Node) error {
	// Expect a mapping with a "repos" key
	var raw struct {
		Repos []yaml.Node `yaml:"repos"`
	}
	if err := value.Decode(&raw); err != nil {
		return err
	}

	for _, node := range raw.Repos {
		var repo Repo
		switch node.Kind {
		case yaml.ScalarNode:
			// Plain string URL
			repo = Repo{URL: node.Value}
		case yaml.MappingNode:
			if err := node.Decode(&repo); err != nil {
				return fmt.Errorf("invalid repo entry at line %d: %w", node.Line, err)
			}
		default:
			return fmt.Errorf("invalid repo entry at line %d: expected string or mapping", node.Line)
		}
		if repo.URL == "" {
			return fmt.Errorf("repo entry missing url at line %d", node.Line)
		}
		g.Repos = append(g.Repos, repo)
	}
	return nil
}

// Load reads and parses the config from the given path.
func Load(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading config: %w", err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("parsing %s: %w", filepath.Base(path), err)
	}
	if err := cfg.Validate(); err != nil {
		return nil, err
	}
	return &cfg, nil
}

// Validate checks the config for common errors and returns helpful messages.
func (c *Config) Validate() error {
	if c.Groups == nil || len(c.Groups) == 0 {
		return fmt.Errorf(`config error: no groups defined

Your %s must have at least one group with repos. Example:

  groups:
    default:
      repos:
        - git@github.com:org/repo.git`, ConfigFile)
	}

	for groupName, group := range c.Groups {
		if err := validateGroupName(groupName); err != nil {
			return fmt.Errorf("config error in group %q: %w", groupName, err)
		}
		if len(group.Repos) == 0 {
			return fmt.Errorf("config error: group %q has no repos\n\nAdd at least one repo URL to the group's repos list.", groupName)
		}

		seen := make(map[string]bool)
		for i, repo := range group.Repos {
			if err := validateRepo(repo, i, groupName); err != nil {
				return err
			}
			name := repo.RepoName()
			if seen[name] {
				return fmt.Errorf("config error: duplicate repo name %q in group %q\n\nUse the \"name\" field to give one a unique name:\n  - url: %s\n    name: %s-2", name, groupName, repo.URL, name)
			}
			seen[name] = true
		}
	}

	return nil
}

var validGroupNameRe = regexp.MustCompile(`^[a-zA-Z0-9][a-zA-Z0-9_-]*$`)

func validateGroupName(name string) error {
	if !validGroupNameRe.MatchString(name) {
		return fmt.Errorf("invalid group name %q\n\nGroup names must start with a letter or number and contain only letters, numbers, hyphens, and underscores.", name)
	}
	return nil
}

var sshURLRe = regexp.MustCompile(`^[\w.-]+@[\w.-]+:[\w./-]+$`)

func validateRepo(repo Repo, index int, group string) error {
	repoURL := repo.URL

	if repoURL == "" {
		return fmt.Errorf("config error: repo #%d in group %q is missing a URL", index+1, group)
	}

	if strings.Contains(repoURL, " ") {
		return fmt.Errorf("config error: repo URL contains spaces in group %q: %q\n\nRepo URLs should not contain spaces.", group, repoURL)
	}

	isSSH := sshURLRe.MatchString(repoURL)
	isHTTPS := false
	if !isSSH {
		if u, err := url.Parse(repoURL); err == nil && (u.Scheme == "https" || u.Scheme == "http") && u.Host != "" {
			isHTTPS = true
		}
	}

	if !isSSH && !isHTTPS {
		return fmt.Errorf(`config error: invalid repo URL in group %q: %q

Repo URLs should be either:
  SSH:   git@github.com:org/repo.git
  HTTPS: https://github.com/org/repo.git`, group, repoURL)
	}

	name := repo.RepoName()
	if name == "" {
		return fmt.Errorf("config error: could not derive repo name from URL %q in group %q\n\nSet an explicit \"name\" field for this repo.", repoURL, group)
	}

	return nil
}

// FindConfigFile returns the config file path within a directory.
func FindConfigFile(dir string) (string, bool) {
	path := filepath.Join(dir, ConfigFile)
	if _, err := os.Stat(path); err == nil {
		return path, true
	}
	return "", false
}

// FindRoot walks up from startDir looking for zproj.yaml.
// Returns the directory containing it.
func FindRoot(startDir string) (string, error) {
	dir, err := filepath.Abs(startDir)
	if err != nil {
		return "", err
	}
	for {
		if _, found := FindConfigFile(dir); found {
			return dir, nil
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return "", fmt.Errorf("could not find %s in any parent directory", ConfigFile)
		}
		dir = parent
	}
}

// GroupDir returns the directory path for a group within the root.
// "default" group uses the root directly; named groups use [groupname]/.
func GroupDir(root, group string) string {
	if group == "default" {
		return root
	}
	return filepath.Join(root, "["+group+"]")
}

// MainDir returns the .main directory for a group.
func MainDir(root, group string) string {
	return filepath.Join(GroupDir(root, group), ".main")
}
