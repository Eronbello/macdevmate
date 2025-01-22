package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/AlecAivazis/survey/v2"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// Register the "install" subcommand
func init() {
	rootCmd.AddCommand(installCmd)
}

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Interactively choose software to install on macOS",
	Run: func(cmd *cobra.Command, args []string) {
		runInstall()
	},
}

// brewPackages: CLI tools installed with 'brew install'
var brewPackages = map[string]string{
	"Git":            "git",
	"Docker Compose": "docker-compose",
	"Golang":         "go",
	"Node.js":        "node", // includes npm
}

// caskPackages: GUI apps installed with 'brew install --cask'
var caskPackages = map[string]string{
	"Docker Desktop":     "docker",
	"Visual Studio Code": "visual-studio-code",
	"Insomnia":           "insomnia",
}

// runInstall: main logic for the install command
func runInstall() {
	color.Green("Starting installation process...")

	// 1. Check if Homebrew is installed; if not, install it
	if !commandExists("brew") {
		color.Yellow("Homebrew not found. Installing Homebrew...")
		if err := installHomebrew(); err != nil {
			color.Red(fmt.Sprintf("Error installing Homebrew: %v", err))
			return
		}
	} else {
		color.Green("Homebrew is already installed.")
	}

	// 2. Prompt the user to select which software to install
	chosenSoftware := promptUserForSoftware()
	if len(chosenSoftware) == 0 {
		color.Yellow("No software selected. Exiting.")
		return
	}

	// 3. Install each chosen software
	for _, softwareName := range chosenSoftware {
		if pkg, ok := brewPackages[softwareName]; ok {
			installBrewPackage(softwareName, pkg)
		} else if cask, ok := caskPackages[softwareName]; ok {
			installBrewCask(softwareName, cask)
		} else {
			color.Red(fmt.Sprintf("Unknown software: %s (skipping)", softwareName))
		}
	}

	color.Green("Installation process finished.")
}

// promptUserForSoftware displays a multi-select prompt of available software
func promptUserForSoftware() []string {
	// Gather all software names into a single slice
	var softwareList []string
	for name := range brewPackages {
		softwareList = append(softwareList, name)
	}
	for name := range caskPackages {
		softwareList = append(softwareList, name)
	}

	var selected []string
	prompt := &survey.MultiSelect{
		Message: "Select the software you want to install (use space to toggle, enter to confirm):",
		Options: softwareList,
	}
	err := survey.AskOne(prompt, &selected)
	if err != nil {
		color.Red(fmt.Sprintf("Prompt failed: %v", err))
		return nil
	}
	return selected
}

// commandExists checks if a command is available in PATH
func commandExists(name string) bool {
	_, err := exec.LookPath(name)
	return err == nil
}

// installHomebrew runs the official Homebrew install script
func installHomebrew() error {
	cmd := exec.Command("/bin/bash", "-c",
		"curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh | bash",
	)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// installBrewPackage installs a CLI package via 'brew install'
func installBrewPackage(softwareName, brewPackage string) {
	if commandExists(brewPackage) {
		color.Yellow(fmt.Sprintf("%s is already installed.", softwareName))
		return
	}
	color.Cyan(fmt.Sprintf("Installing %s...", softwareName))
	cmd := exec.Command("brew", "install", brewPackage)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		color.Red(fmt.Sprintf("Error installing %s: %v", softwareName, err))
	} else {
		color.Green(fmt.Sprintf("%s installed successfully!", softwareName))
	}
}

// installBrewCask installs a GUI app via 'brew install --cask'
func installBrewCask(softwareName, caskPackage string) {
	color.Cyan(fmt.Sprintf("Installing %s (cask)...", softwareName))
	cmd := exec.Command("brew", "install", "--cask", caskPackage)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		color.Red(fmt.Sprintf("Error installing %s: %v", softwareName, err))
	} else {
		color.Green(fmt.Sprintf("%s installed successfully!", softwareName))
	}
}
