package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const ConfigFile = "zproj.config.json"

type Config struct {
	Groups    map[string]Group `json:"groups"`
	Templates *Templates       `json:"templates,omitempty"`
}

type Group struct {
	Repos []Repo `json:"-"`

	RawRepos []json.RawMessage `json:"repos"`
}

type Repo struct {
	URL    string `json:"url"`
	Name   string `json:"name,omitempty"`
	Branch string `json:"branch,omitempty"`
}

type Templates struct {
	Variables map[string]string `json:"variables,omitempty"`
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

func repoNameFromURL(url string) string {
	// Handle both SSH and HTTPS URLs
	// git@github.com:org/repo.git -> repo
	// https://github.com/org/repo.git -> repo
	base := filepath.Base(url)
	return strings.TrimSuffix(base, ".git")
}

func (g *Group) UnmarshalJSON(data []byte) error {
	type rawGroup struct {
		Repos []json.RawMessage `json:"repos"`
	}
	var rg rawGroup
	if err := json.Unmarshal(data, &rg); err != nil {
		return err
	}

	for _, raw := range rg.Repos {
		var repo Repo
		// Try string first (plain URL)
		var s string
		if err := json.Unmarshal(raw, &s); err == nil {
			repo = Repo{URL: s}
		} else {
			// Try object
			if err := json.Unmarshal(raw, &repo); err != nil {
				return fmt.Errorf("invalid repo entry: %s", string(raw))
			}
		}
		if repo.URL == "" {
			return fmt.Errorf("repo entry missing url: %s", string(raw))
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
	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("parsing config: %w", err)
	}
	if cfg.Groups == nil || len(cfg.Groups) == 0 {
		return nil, fmt.Errorf("config must define at least one group")
	}
	return &cfg, nil
}

// FindRoot walks up from startDir looking for zproj.config.json.
// Returns the directory containing it.
func FindRoot(startDir string) (string, error) {
	dir, err := filepath.Abs(startDir)
	if err != nil {
		return "", err
	}
	for {
		if _, err := os.Stat(filepath.Join(dir, ConfigFile)); err == nil {
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
