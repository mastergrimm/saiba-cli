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

		if tui.GetStore() {
			_ = spinner.New().Title("Adding Store").Action(addStore).Run()
		}

		if tui.GetUtils() != nil {
			fmt.Printf("These are the Utils you have chosen: %v\n", tui.GetUtils())
		}
	},
}

func createSvelteKit(projectName string) error {

	cmd := exec.Command("pnpm", "create", "svelte@latest", projectName, "--silent")

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("command finished with error: %v", err)
	}

	fmt.Printf("\nSvelte Project created successfully.\n")

	return nil
}

func addUtil(util string) {

}

func addStore() {
	err := utils.CloneAndCopySubdir("https://github.com/mastergrimm/saiba-cli.git", "templates/default/svelte-store", "src/lib")

	if err != nil {
		fmt.Printf("Failed to add initial scaffold: %v\n", err)
	}

	fmt.Printf("Store added successfully.\n")
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

	cmd := exec.Command("pnpm", "install", "--silent", "--save-dev", "@iconify/svelte")
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
