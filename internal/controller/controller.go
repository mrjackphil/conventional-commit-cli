package controller

import (
	"conventional-emoji-in-shell/internal/input"
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"os/exec"
	"strings"
)

type model struct {
	step    int
	steps   []step
	options map[string][]option
	cursor  int
	input   input.Model
	result  result
	done    bool
}

type result struct {
	name    string
	scope   string
	gitmoji string
	summary string
	desc    string
	changes string
}

type step struct {
	name string
	msg  string
}

type option struct {
	text  string
	desc  string
	value string
}

type Err struct{ error error }

func InitModel() tea.Model {
	return model{
		step: 0,
		steps: []step{
			{name: "type", msg: "Choose type of change:"},
			{name: "scope", msg: "Enter scope of change:"},
			{name: "gitmoji", msg: "Choose gitmoji:"},
			{name: "summary", msg: "Enter short summary:"},
			{name: "description", msg: "Enter full description:"},
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
				{
					text:  "🎨",
					value: ":art:",
					desc:  "Improving structure / format of the code.",
				},
				{
					text:  "⚡️",
					value: ":zap:",
					desc:  "Improving performance.",
				},
				{
					text:  "🔥",
					value: ":fire:",
					desc:  "Removing code or files.",
				},
				{
					text:  "🐛",
					value: ":bug:",
					desc:  "Fixing a bug.",
				},
				{
					text:  "🚑",
					value: ":ambulance:",
					desc:  "Critical hotfix.",
				},
				{
					text:  "✨",
					value: ":sparkles:",
					desc:  "Introducing new features.",
				},
				{
					text:  "📝",
					value: ":pencil:",
					desc:  "Writing docs.",
				},
				{
					text:  "🚀",
					value: ":rocket:",
					desc:  "Deploying stuff.",
				},
				{
					text:  "💄",
					value: ":lipstick:",
					desc:  "Updating the UI and style files.",
				},
				{
					text:  "🎉",
					value: ":tada:",
					desc:  "Initial commit.",
				},
				{
					text:  "✅",
					value: ":white_check_mark:",
					desc:  "Adding tests.",
				},
				{
					text:  "🔒",
					value: ":lock:",
					desc:  "Fixing security issues.",
				},
				{
					text:  "🍎",
					value: ":apple:",
					desc:  "Fixing something on macOS.",
				},
				{
					text:  "🐧",
					value: ":penguin:",
					desc:  "Fixing something on Linux.",
				},
				{
					text:  "🏁",
					value: ":checkered_flag:",
					desc:  "Fixing something on Windows.",
				},
				{
					text:  "🤖",
					value: ":robot:",
					desc:  "Fixing something on Android.",
				},
				{
					text:  "🍏",
					value: ":green_apple:",
					desc:  "Fixing something on iOS.",
				},
				{
					text:  "🔖",
					value: ":bookmark:",
					desc:  "Releasing / Version tags.",
				},
				{
					text:  "🚨",
					value: ":rotating_light:",
					desc:  "Removing linter warnings.",
				},
				{
					text:  "🚧",
					value: ":construction:",
					desc:  "Work in progress.",
				},
				{
					text:  "💚",
					value: ":green_heart:",
					desc:  "Fixing CI Build.",
				},
				{
					text:  "⬇️",
					value: ":arrow_down:",
					desc:  "Downgrading dependencies.",
				},
				{
					text:  "⬆️",
					value: ":arrow_up:",
					desc:  "Upgrading dependencies.",
				},
				{
					text:  "📌",
					value: ":pushpin:",
					desc:  "Pinning dependencies to specific versions.",
				},
				{
					text:  "👷",
					value: ":construction_worker:",
					desc:  "Adding CI build system.",
				},
				{
					text:  "📈",
					value: ":chart_with_upwards_trend:",
					desc:  "Adding analytics or tracking code.",
				},
				{
					text:  "♻️",
					value: ":recycle:",
					desc:  "Refactoring code.",
				},
				{
					text:  "🐳",
					value: ":whale:",
					desc:  "Work about Docker.",
				},
				{
					text:  "➕",
					value: ":heavy_plus_sign:",
					desc:  "Adding a dependency.",
				},
				{
					text:  "➖",
					value: ":heavy_minus_sign:",
					desc:  "Removing a dependency.",
				},
				{
					text:  "🔧",
					value: ":wrench:",
					desc:  "Changing configuration files.",
				},
				{
					text:  "🌐",
					value: ":globe_with_meridians:",
					desc:  "Internationalization and localization.",
				},
				{
					text:  "✏️",
					value: ":pencil2:",
					desc:  "Fixing typos.",
				},
				{
					text:  "💩",
					value: ":poop:",
					desc:  "Writing bad code that needs to be improved.",
				},
				{
					text:  "⏪",
					value: ":rewind:",
					desc:  "Reverting changes.",
				},
				{
					text:  "🔀",
					value: ":twisted_rightwards_arrows:",
					desc:  "Merging branches.",
				},
				{
					text:  "📦",
					value: ":package:",
					desc:  "Updating compiled files or packages.",
				},
				{
					text:  "👽",
					value: ":alien:",
					desc:  "Updating code due to external API changes.",
				},
				{
					text:  "🚚",
					value: ":truck:",
					desc:  "Moving or renaming files.",
				},
				{
					text:  "📄",
					value: ":page_facing_up:",
					desc:  "Adding or updating license.",
				},
				{
					text:  "💥",
					value: ":boom:",
					desc:  "Introducing breaking changes.",
				},
				{
					text:  "🍱",
					value: ":bento:",
					desc:  "Adding or updating assets.",
				},
				{
					text:  "👌",
					value: ":ok_hand:",
					desc:  "Updating code due to code review changes.",
				},
				{
					text:  "♿️",
					value: ":wheelchair:",
					desc:  "Improving accessibility.",
				},
				{
					text:  "💡",
					value: ":bulb:",
					desc:  "Documenting source code.",
				},
				{
					text:  "🍻",
					value: ":beers:",
					desc:  "Writing code drunkenly.",
				},
				{
					text:  "💬",
					value: ":speech_balloon:",
					desc:  "Updating text and literals.",
				},
				{
					text:  "🗃",
					value: ":card_file_box:",
					desc:  "Performing database related changes.",
				},
				{
					text:  "🔊",
					value: ":loud_sound:",
					desc:  "Adding logs.",
				},
				{
					text:  "🔇",
					value: ":mute:",
					desc:  "Removing logs.",
				},
				{
					text:  "👥",
					value: ":busts_in_silhouette:",
					desc:  "Adding contributor(s).",
				},
				{
					text:  "🚸",
					value: ":children_crossing:",
					desc:  "Improving user experience / usability.",
				},
				{
					text:  "🏗",
					value: ":building_construction:",
					desc:  "Making architectural changes.",
				},
				{
					text:  "📱",
					value: ":iphone:",
					desc:  "Working on responsive design.",
				},
				{
					text:  "🤡",
					value: ":clown_face:",
					desc:  "Mocking things.",
				},
				{
					text:  "🥚",
					value: ":egg:",
					desc:  "Adding an easter egg.",
				},
				{
					text:  "🙈",
					value: ":see_no_evil:",
					desc:  "Adding or updating a .gitignore file.",
				},
				{
					text:  "📸",
					value: ":camera_flash:",
					desc:  "Adding or updating snapshots.",
				},
				{
					text:  "⚗",
					value: ":alembic:",
					desc:  "Experimenting new things.",
				},
				{
					text:  "🔍",
					value: ":mag:",
					desc:  "Improving SEO.",
				},
				{
					text:  "☸️",
					value: ":wheel_of_dharma:",
					desc:  "Work about Kubernetes.",
				},
				{
					text:  "🏷️",
					value: ":label:",
					desc:  "Adding or updating types (Flow, TypeScript).",
				},
				{
					text:  "🌱",
					value: ":seedling:",
					desc:  "Adding or updating seed files.",
				},
				{
					text:  "🚩",
					value: ":triangular_flag_on_post:",
					desc:  "Adding, updating, or removing feature flags.",
				},
				{
					text:  "💫",
					value: ":dizzy:",
					desc:  "Adding or updating animations and transitions.",
				}},
		},
		input: input.Init(),
	}
}

func (m model) Init() tea.Cmd {
	InitModel()
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	if m.done {
		return m, tea.Quit
	}

	if m.step > len(m.steps)-1 {
		cmd := exec.Command("git", "commit", "-a", "-m", m.getResult())
		proc := tea.ExecProcess(cmd, func(err error) tea.Msg {
			return Err{error: err}
		})
		m.done = true
		return m, proc
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl-c", "q", "esc": // quit
			return m, tea.Quit
		case "enter":
			switch m.steps[m.step].name {
			case "type":
				m.result.name = m.currentOptions()[m.cursor].text
			case "scope":
				m.result.scope = m.input.GetText()
			case "gitmoji":
				m.result.gitmoji = m.currentOptions()[m.cursor].value
			case "description":
				m.result.desc = m.input.GetText()
			case "breaking changes":
				m.result.changes = m.input.GetText()
			}

			m.cursor = 0
			if m.step < len(m.steps) {
				m.step++
			}
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

	step := m.currentStep()

	if step == nil {
		return m, cmd
	}

	if step.name != "gitmoji" && step.name != "type" {
		m.input, cmd = m.input.Update(msg)
	}

	return m, cmd
}

func (m model) View() string {
	s := "MrJackphil's git commitzen:\n\n"

	if (m.step < 0) || (m.step >= len(m.steps)) {
		s += m.getResult()
		s += "\n\n"
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
	boundary := 5
	t := m.cursor - boundary
	b := m.cursor + boundary

	if t < 0 {
		b += -t
	}

	if b > len(options) {
		t -= b - len(options)
	}

	t = clamp(t, 0, len(options))
	b = clamp(b, 0, len(options))

	// Print the options
	for feat, option := range options {
		// Print the selected option
		cursor := " "
		if m.cursor == feat {
			cursor = ">"
		}

		// Print description
		desc := option.desc
		if desc != "" {
			desc = "- " + desc
		}

		// Add space between option and description
		space := max(0, longest-len(option.text))

		if feat >= t && feat <= b {
			s += fmt.Sprintf("\n %s %s%s %s", cursor, option.text, strings.Repeat(" ", space), desc)
		}
	}

	// Add input view content
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

func (m model) getResult() string {
	changes := m.result.changes
	if m.result.changes != "" {
		changes = "\n\nBREAKING CHANGE: " + m.result.changes
	}

	return fmt.Sprintf("%s(%s): %s %s%s", m.result.name, m.result.scope, m.result.gitmoji, m.result.desc, changes)
}

func clamp(v, low, high int) int {
	if high < low {
		low, high = high, low
	}
	return min(high, max(low, v))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
