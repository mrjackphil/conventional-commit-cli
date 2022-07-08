package controller

import (
	"conventional-emoji-in-shell/internal/input"
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"strings"
)

type model struct {
	step    int
	steps   []step
	options map[string][]option
	cursor  int // which to-do list item our cursor is pointing at
	input   tea.Model
}

type step struct {
	name string
	msg  string
}

type option struct {
	text string
	desc string
}

func InitModel() model {
	return model{
		step: 0,
		steps: []step{
			{name: "type", msg: "Choose type of change:"},
			{name: "scope", msg: "Choose scope of change:"},
			{name: "gitmoji", msg: "Choose gitmoji:"},
			{name: "description", msg: "Enter description:"},
			{name: "breaking changes", msg: "Enter breaking changes:"},
		},
		options: map[string][]option{
			"type": {
				{text: "feat", desc: "A new feature"},
				{text: "docs", desc: "Documentation only changes"},
				{text: "fix", desc: "A bug fix"},
				{text: "refactor", desc: "A code change that neither fixes a bug nor adds a feature"},
				{text: "chore", desc: "Other changes that don't modify src or test files"},
				{text: "style", desc: "Changes that do not affect the meaning of the code (white-space, formatting, missing semi-colons, etc)"},
				{text: "perf", desc: "A code change that improves performance"},
				{text: "test", desc: "Adding missing tests or correcting existing tests"},
				{text: "build", desc: "Changes that affect the build system or external dependencies (example scopes: gulp, broccoli, npm)"},
				{text: "ci", desc: "Changes to our CI configuration files and scripts (example scopes: Travis, Circle, BrowserStack, SauceLabs)"},
				{text: "revert", desc: "Reverts a previous commit"},
			},
			"gitmoji": {
				{text: "none", desc: "No gitmoji"},
				{text: "📝", desc: "Writing docs."},
				{text: "🔥", desc: "Removing code or files."},
				{text: "🔀", desc: "Merging branches."},
				{text: "🐛", desc: "Fixing a bug."},
				{text: "🎨", desc: "Improving structure / format of the code."},
				{text: "⚡️", desc: "Improving performance."},
				{text: "🚑", desc: "Critical hotfix."},
				{text: "✨", desc: "Introducing new features."},
				{text: "🚀", desc: "Deploying stuff."},
				{text: "💄", desc: "Updating the UI and style files."},
				{text: "🎉", desc: "Initial commit."},
				{text: "✅", desc: "Updating tests."},
				{text: "🔒", desc: "Fixing security issues."},
				{text: "🍎", desc: "Fixing something on macOS."},
				{text: "🐧", desc: "Fixing something on Linux."},
				{text: "🏁", desc: "Fixing something on Windows."},
				{text: "🤖", desc: "Fixing something on Android."},
				{text: "🍏", desc: "Fixing something on iOS."},
				{text: "🔖", desc: "Releasing / Version tags."},
				{text: "🚨", desc: "Removing linter warnings."},
				{text: "🚧", desc: "Work in progress."},
				{text: "💚", desc: "Fixing CI Build."},
				{text: "⬇️", desc: "Downgrading dependencies."},
				{text: "⬆️", desc: "Upgrading dependencies."},
				{text: "📌", desc: "Pinning dependencies to specific versions."},
				{text: "📦", desc: "Updating packages."},
				{text: "📄", desc: "Adding or updating license."},
				{text: "💥", desc: "Introducing breaking changes."},
				{text: "🍱", desc: "Adding or updating assets."},
				{text: "👌", desc: "Updating code due to code review changes."},
				{text: "♿️", desc: "Improving accessibility."},
				{text: "💡", desc: "Documenting source code."},
				{text: "🍻", desc: "Writing code drunkenly."},
				{text: "💬", desc: "Updating text and literals."},
				{text: "🗃", desc: "Performing database related changes."},
				{text: "🔊", desc: "Adding logs."},
				{text: "🔇", desc: "Removing logs."},
				{text: "👥", desc: "Adding contributor(s)."},
				{text: "🚸", desc: "Improving user experience / usability."},
				{text: "🏗", desc: "Making architectural changes."},
				{text: "📱", desc: "Working on responsive design."},
				{text: "🤡", desc: "Mocking things."},
				{text: "🥚", desc: "Adding an easter egg."},
				{text: "⚗", desc: "Experimenting new things"},
				{text: "🔍", desc: "Improving SEO"},
				{text: "☸️", desc: "Work about Kubernetes"},
				{text: "🏷️", desc: "Adding or updating types (Flow, TypeScript)"},
				{text: "🙈", desc: "Adding or updating a .gitignore file"},
				{text: "📸", desc: "Adding or updating snapshots"},
				{text: "📦", desc: "Adding or updating a dependency"},
				{text: "📁", desc: "Adding or updating a file"},
				{text: "📂", desc: "Adding or updating a directory"},
				{text: "📅", desc: "Adding or updating a timestamp"},
			},
		},
		input: input.Init(),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	if m.currentStep().name != "gitmoji" && m.currentStep().name != "feat" {
		m.input, cmd = m.input.Update(msg)
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl-c", "q", "esc": // quit
			return m, tea.Quit
		case "enter":
			m.cursor = 0
			m.step++
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.currentOptions())-1 {
				m.cursor++
			}
		}
	}

	return m, cmd
}

func (m model) View() string {
	s := "MrJackphil's git commitzen:\n\n"

	if (m.step < 0) || (m.step >= len(m.steps)) {
		s += "End of the job.\nPress Ctrl-C to quit."
		return s
	}

	step := m.currentStep()

	s += step.msg

	options := m.currentOptions()

	// Max length of the option
	longest := 0
	for _, feat := range options {

		if len(feat.text) > longest {
			longest = len(feat.text)
		}
	}

	// Setup viewport boundaries
	boundary := 15
	toffset := m.cursor - boundary
	boffset := m.cursor + boundary

	if toffset < 0 {
		toffset = 0
	}

	if boffset > len(options) {
		boffset = len(options)
	}

	// Print the options
	for feat, option := range options {
		cursor := current(m.cursor, feat)
		desc := option.desc

		if desc != "" {
			desc = "- " + desc
		}

		space := longest - len(option.text)
		if space < 0 {
			space = 0
		}

		if feat >= toffset && feat <= boffset {
			s += fmt.Sprintf("\n %s %s%s %s", cursor, option.text, strings.Repeat(" ", space), desc)
		}
	}

	if step.name != "type" && step.name != "gitmoji" {
		s += "\n" + m.input.View()
	}

	return s
}

func (m model) currentStep() *step {
	if (m.step < 0) || (m.step >= len(m.steps)) {
		return nil
	}

	return &m.steps[m.step]
}

func (m model) currentOptions() []option {
	return m.options[m.currentStep().name]
}

func current(c int, f int) string {
	if c == f {
		return "*"
	}
	return " "
}
