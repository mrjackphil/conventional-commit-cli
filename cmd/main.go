package main

import (
	"conventional-emoji-in-shell/internal/controller"
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"os"
)

/*
	Types
		feat:    A new feature
		docs:    Documentation only changes
		fix:     A bug fix
		refactor: A code change that neither fixes a bug nor adds a feature
		chore:   Other changes that don't modify src or test files
		style:   Changes that do not affect the meaning of the code (white-space, formatting, missing semi-colons, etc)
		perf:    A code change that improves performance
		test:    Adding missing tests or correcting existing tests
		build:   Changes that affect the build system or external dependencies (example scopes: gulp, broccoli, npm)
		ci:      Changes to our CI configuration files and scripts (example scopes: Travis, Circle, BrowserStack, SauceLabs)
		revert:  Reverts a previous commit
*/

/*
	Gitmoji
		none
		📝  - Writing docs.
		🔥  - Removing code or files.
		🔀  - Merging branches.
		🐛  - Fixing a bug.
		🎨  - Improving structure / format of the code.
		⚡️  - Improving performance.
		🚑  - Critical hotfix.
		✨  - Introducing new features.
		🚀  - Deploying stuff.
		💄  - Updating the UI and style files.
		🎉  - Initial commit.
		✅  - Updating tests.
		🔒  - Fixing security issues.
		🍎  - Fixing something on macOS.
		🐧  - Fixing something on Linux.
		🏁  - Fixing something on Windows.
		🤖  - Fixing something on Android.
		🍏  - Fixing something on iOS.
		🔖  - Releasing / Version tags.
		🚨  - Removing linter warnings.
		🚧  - Work in progress.
		💚  - Fixing CI Build.
		⬇️  - Downgrading dependencies.
		⬆️  - Upgrading dependencies.
		📌  - Pinning dependencies to specific versions.
		👷  - Adding CI build system.
		📈  - Adding analytics or tracking code.
		♻️  - Refactoring code.
		🐳  - Work about Docker.
		➕  - Adding a dependency.
		➖  - Removing a dependency.
		🔧  - Changing configuration files.
		🌐  - Internationalization and localization.
		✏️  - Fixing typos.
		💩  - Writing bad code that needs to be improved.
		⏪  - Reverting changes.
		📦  - Updating compiled files or packages.
		👽  - Updating code due to external API changes.
		🚚  - Moving or renaming files.
		📄  - Adding or updating license.
		💥  - Introducing breaking changes.
		🍱  - Adding or updating assets.
		👌  - Updating code due to code review changes.
		♿️  - Improving accessibility.
		💡  - Documenting source code.
		🍻  - Writing code drunkenly.
		💬  - Updating text and literals.
		🗃  - Performing database related changes.
		🔊  - Adding logs.
		🔇  - Removing logs.
		👥  - Adding contributor(s).
		🚸  - Improving user experience / usability.
		🏗  - Making architectural changes.
		📱  - Working on responsive design.
		🤡  - Mocking things.
		🥚  - Adding an easter egg.
		🙈  - Adding or updating a .gitignore file
		📸  - Adding or updating snapshots
		⚗  - Experimenting new things
		🔍  - Improving SEO
		☸️  - Work about Kubernetes
		🏷️  - Adding or updating types (Flow, TypeScript)
*/

func main() {
	p := tea.NewProgram(controller.InitModel())
	if err := p.Start(); err != nil {
		fmt.Printf("There's been an error: %v", err)
		os.Exit(1)
	}

	// Select the type of change that you're committing

	// Denote the scope of this change

	// Choose a gitmoji

	// Write a short, imperative tense description of the change

	// Provide a longer description of the change

	// List any breaking changes or issues closed by this change
}
