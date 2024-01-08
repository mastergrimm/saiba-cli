package tui

import (
	"errors"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

type model struct {
	form           *huh.Form
	projectName    string
	repoName       string
	features       []string
	includeSaibaUI bool
	store          bool
	utils          []string
}

var theme *huh.Theme = huh.ThemeBase16()

func NewModel() model {
	m := model{}

	m.form = huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Key("projectName").
				Title("Enter Project Name").
				Value(&m.projectName).
				Validate(func(str string) error {
					if str == "" {
						return errors.New("project name cannot be empty")
					}

					return nil
				}),
			huh.NewInput().
				Key("repoName").
				Title("Enter Repo Name").
				Value(&m.repoName).
				Validate(func(str string) error {

					return nil
				}),

			huh.NewMultiSelect[string]().
				Key("features").
				Title("Select all features you'd like (Space Bar)").
				Options(
					huh.NewOption("SASS", "SASS"),
					huh.NewOption("Lucia Auth", "Lucia Auth"),
					huh.NewOption("Iconify Icons", "Iconify Icons"),
				).
				Value(&m.features),
		),

		huh.NewGroup(
			huh.NewConfirm().
				Key("includeSaibaUI").
				Title("Would you like to add SaibaUI?").
				Affirmative("Yes").
				Negative("No").
				Value(&m.includeSaibaUI),
		),
		huh.NewGroup(
			huh.NewConfirm().
				Key("store").
				Title("Add Svelte Store (including folders)?").
				Affirmative("Yes").
				Negative("No").
				Value(&m.store),
		),
		huh.NewGroup(
			huh.NewMultiSelect[string]().
				Key("utils").
				Title("Select all utils you'd like (Space Bar)").
				Options(
					huh.NewOption("clickOutside", "clickoutside"),
					huh.NewOption("truncateText", "truncatetext"),
				).
				Value(&m.utils),
		).WithTheme(theme),
	)

	return m
}

func (m model) Init() tea.Cmd {
	return m.form.Init()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "esc", "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	var updatedForm tea.Model
	updatedForm, cmd = m.form.Update(msg)

	if updatedForm, ok := updatedForm.(*huh.Form); ok {
		m.form = updatedForm
	} else {
		fmt.Println("Updated form is not of type *huh.Form")
	}

	if m.form.State == huh.StateCompleted {

		return m, tea.Quit
	}

	return m, cmd
}

func (m model) View() string {
	if m.form.State == huh.StateCompleted {
		formOutput := "Form complete. Please wait for the SvelteKit prompt...\n"

		return formOutput
	}
	return m.form.View()
}

var FormValues = NewModel()

func RunPrompt() error {
	if _, err := tea.NewProgram(FormValues).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}

	return nil
}

func GetProjectName() string {
	return FormValues.form.GetString("projectName")
}

func GetRepoName() string {
	return FormValues.form.GetString("repoName")
}

func GetFeatures() []string {
	return FormValues.form.Get("features").([]string)
}

func GetIncludeSaibaUI() bool {
	return FormValues.form.GetBool("includeSaibaUI")
}

func GetStore() bool {
	return FormValues.form.GetBool("store")
}

func GetUtils() []string {
	return FormValues.form.Get("utils").([]string)
}
