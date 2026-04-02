package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/ntotten/zproj/internal/config"
	"github.com/ntotten/zproj/internal/git"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize: clone all repos into .main directories",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := requireConfig(); err != nil {
			return err
		}
		return runInit()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

func runInit() error {
	for groupName, group := range cfg.Groups {
		mainDir := config.MainDir(rootDir, groupName)
		if err := os.MkdirAll(mainDir, 0755); err != nil {
			return fmt.Errorf("creating .main for group %q: %w", groupName, err)
		}

		fmt.Printf("Initializing group %q (%d repos)...\n", groupName, len(group.Repos))

		results := git.RunParallel(group.Repos, func(repo config.Repo) git.Result {
			repoName := repo.RepoName()
			dest := filepath.Join(mainDir, repoName)

			if _, err := os.Stat(dest); err == nil {
				return git.Result{Repo: repoName, Output: "already exists, skipping"}
			}

			fmt.Printf("  Cloning %s...\n", repoName)
			if err := git.Clone(repo.URL, dest, repo.RepoBranch()); err != nil {
				return git.Result{Repo: repoName, Err: err}
			}
			return git.Result{Repo: repoName, Output: "cloned"}
		})

		var errs []string
		for _, r := range results {
			if r.Err != nil {
				errs = append(errs, fmt.Sprintf("  %s: %v", r.Repo, r.Err))
			} else {
				fmt.Printf("  %s: %s\n", r.Repo, r.Output)
			}
		}
		if len(errs) > 0 {
			fmt.Fprintf(os.Stderr, "Errors in group %q:\n%s\n", groupName, strings.Join(errs, "\n"))
		}
	}

	fmt.Println("Init complete.")
	return nil
}
