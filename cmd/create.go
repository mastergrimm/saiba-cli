package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/mastergrimm/saiba-cli/internal/utils"

	"github.com/spf13/cobra"

	"github.com/charmbracelet/huh/spinner"
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
		if err := tui.RunPrompt(); err != nil {
			fmt.Printf("Could not initialize prompt: %v\n", err)
			return
		}

		if err := createSvelteKit(tui.GetProjectName()); err != nil {
			fmt.Printf("Could not create SvelteKit project: %v\n", err)
			return
		}

		for _, feature := range tui.GetFeatures() {
			switch feature {
			case "SASS":
				_ = spinner.New().Title("Adding Sass folders").Action(addSASS).Run()
			case "Lucia Auth":
				_ = spinner.New().Title("Adding Lucia Auth").Action(addLuciaAuth).Run()

			case "Iconify Icons":
				_ = spinner.New().Title("Adding Iconify Icons").Action(addIconifyIcons).Run()
			}
		}

		if tui.GetIncludeSaibaUI() {
			_ = spinner.New().Title("Adding SaibaUI").Action(addSaibaUI).Run()
		}
	},
}

func createSvelteKit(projectName string) error {
	cmd := exec.Command("npm", "create", "svelte@latest", projectName)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("command finished with error: %v", err)
	}

	fmt.Print("\n")

	fmt.Println("Svelte Project created successfully.")

	fmt.Print("\n")

	return nil
}
func addSASS() {
	utils.GotoDir(tui.GetProjectName())

	cmd := exec.Command("npx", "-q", "svelte-add@latest", "sass")

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Printf("Could not add SASS: %v\n", err)
	} else {
		fmt.Printf("SASS installed!\n")
	}

	err := utils.CloneAndCopySubdir("https://github.com/mastergrimm/saiba-cli.git", "templates/sass", "src/lib")

	if err != nil {
		fmt.Printf("Failed to add SASS: %v\n", err)
	}
}

func addLuciaAuth() {
	utils.GotoDir(tui.GetProjectName())

	err := utils.CloneAndCopySubdir("https://github.com/mastergrimm/saiba-cli.git", "templates/lucia", "src/lib")

	if err != nil {
		fmt.Printf("Failed to add Lucia Auth: %v\n", err)
	}

}

func addIconifyIcons() {
	utils.GotoDir(tui.GetProjectName())

	cmd := exec.Command("npm", "install", "--silent", "--no-progress", "--save-dev", "@iconify/svelte")
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
	utils.GotoDir(tui.GetProjectName())

	err := utils.CloneAndCopySubdir("https://github.com/mastergrimm/saiba-cli.git", "templates/saibaUI", "src/lib")
	if err != nil {
		fmt.Printf("Failed to add SaibaUI: %v\n", err)
	}
}

func init() {
	rootCmd.AddCommand(createCmd)
}
