package cmd

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"

	"github.com/mastergrimm/saiba-cli/internal/tui"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates a sveltekit project",
	Long: `Creates a sveltekit project with the following options: 
	- SASS
	- Lucia Auth 
	- Iconify Icons
	- SaibaUI`,

	Run: func(cmd *cobra.Command, args []string) {
		tui.RunPrompt()
		createSvelteKit(tui.GetProjectName())

		for _, feature := range tui.GetFeatures() {
			switch feature {
			case "SASS":
				addSASS()
			case "Lucia Auth":
				addLuciaAuth()
			case "Iconify Icons":
				addIconifyIcons()
			}
		}

		if tui.GetIncludeSaibaUI() {
			addSaibaUI()
		}

	},
}

func createSvelteKit(projectName string) {
	cmd := exec.Command("npm", "create", "svelte@latest", projectName)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Printf("Command finished with error: %v\n", err)
	} else {
		fmt.Printf("Svelte Project created!.\n")
	}
}

func addSASS() {
	gotoDir(tui.GetProjectName())

	cmd := exec.Command("npx", "svelte-add@latest", "sass")

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Printf("Could not add SASS: %v\n", err)
	} else {
		fmt.Printf("SASS installed correctly!\n")
	}

	cloneAndCopySubdir("https://github.com/mastergrimm/saiba-cli.git", "templates/sass", "src/lib")
}

func addLuciaAuth() {
	gotoDir(tui.GetProjectName())

	cloneAndCopySubdir("https://github.com/mastergrimm/saiba-cli.git", "templates/lucia", "src/lib")
}

func addIconifyIcons() {
	gotoDir(tui.GetProjectName())

	cmd := exec.Command("npm", "install", "--save-dev", "@iconify/svelte")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Println("Installing Iconify icons...")
	if err := cmd.Run(); err != nil {
		fmt.Printf("Failed to install Iconify icons: %v\n", err)
	} else {
		fmt.Println("Iconify icons installed successfully.")
	}
}

func addSaibaUI() {
	gotoDir(tui.GetProjectName())

	cloneAndCopySubdir("https://github.com/mastergrimm/saiba-cli.git", "templates/saibaUI", "src/lib")
}

func gotoDir(dir string) {

	if err := os.Chdir(dir); err != nil {
		fmt.Printf("Could not change directory: %v\n", err)
	} else {
		fmt.Printf("Changed directory to %s\n", tui.GetProjectName())
	}
}

func cloneAndCopySubdir(repo, subdir, destination string) error {
	tempDir, err := os.MkdirTemp("", "repo-")
	if err != nil {
		return fmt.Errorf("failed to create temp directory: %w", err)
	}
	defer os.RemoveAll(tempDir)

	cloneCmd := exec.Command("git", "clone", repo, tempDir)
	if err := runCommand(cloneCmd); err != nil {
		return fmt.Errorf("failed to clone repository: %w", err)
	}

	srcPath := filepath.Join(tempDir, subdir)
	if err := copyDir(srcPath, destination); err != nil {
		return fmt.Errorf("failed to copy subdirectory: %w", err)
	}

	return nil
}

func runCommand(cmd *exec.Cmd) error {
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func copyDir(src, dst string) error {
	var err error
	var fds []os.DirEntry
	var srcinfo os.FileInfo

	if srcinfo, err = os.Stat(src); err != nil {
		return err
	}

	if err = os.MkdirAll(dst, srcinfo.Mode()); err != nil {
		return err
	}

	if fds, err = os.ReadDir(src); err != nil {
		return err
	}
	for _, fd := range fds {
		srcfp := filepath.Join(src, fd.Name())
		dstfp := filepath.Join(dst, fd.Name())

		if fd.IsDir() {
			if err = copyDir(srcfp, dstfp); err != nil {
				fmt.Println(err)
			}
		} else {
			if err = copyFile(srcfp, dstfp); err != nil {
				fmt.Println(err)
			}
		}
	}
	return nil
}

func copyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	if _, err = io.Copy(out, in); err != nil {
		return err
	}
	return out.Close()
}

func init() {
	rootCmd.AddCommand(createCmd)
}
