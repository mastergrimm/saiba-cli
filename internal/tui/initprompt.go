package tui

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

// Define your struct to hold form data
type model struct {
	form           *huh.Form
	projectName    string
	repoName       string
	features       []string
	includeSaibaUI bool
}

func NewModel() model {
	m := model{}

	m.form = huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Key("projectName").
				Title("Enter Project Name").
				Value(&m.projectName),
			huh.NewInput().
				Key("repoName").
				Title("Enter Repo Name").
				Value(&m.repoName),
			huh.NewMultiSelect[string]().
				Key("features").
				Title("Select all features you'd like").
				Options(
					huh.NewOption("SASS", "SASS"),
					huh.NewOption("Lucia Auth", "Lucia Auth"),
					huh.NewOption("Iconify Icons", "Iconify Icons"),
				).
				Value(&m.features),
		),

		huh.NewGroup(
			huh.NewConfirm().
				Title("Would you like to add SaibaUI?").
				Value(&m.includeSaibaUI),
		),
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
		formOutput := "Form completed. Project is being created...\n"

		return formOutput
	}
	return m.form.View()
}

var FormValues = NewModel()

func RunPrompt() {
	if _, err := tea.NewProgram(FormValues).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}

func GetProjectName() string {
	return FormValues.form.GetString("projectName")
}
