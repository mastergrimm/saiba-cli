package cmd

import (
	"fmt"
	"os"
	"os/exec"

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
				fmt.Println("Lucia Auth")
			case "Iconify Icons":
				fmt.Println("Iconify Icons")
			}
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
	if err := os.Chdir(tui.GetProjectName()); err != nil {
		fmt.Printf("Could not change directory: %v\n", err)
	} else {
		fmt.Printf("Changed directory to %s\n", tui.GetProjectName())
	}

	cmd := exec.Command("npx", "svelte-add@latest", "sass")

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Printf("Could not add SASS: %v\n", err)
	} else {
		fmt.Printf("SASS installed correctly!\n")
	}
}

func init() {
	rootCmd.AddCommand(createCmd)
}
