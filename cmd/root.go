package cmd

import (
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "macdevmate",
	Short: "MacDevMate: A CLI to install common dev tools on macOS",
	Long: `MacDevMate is a Go-based command-line tool that helps you
install development tools on macOS via Homebrew.`,
}

// An ASCII art banner. Feel free to customize or replace with your own!
var asciiBanner = `
::::    ::::      :::      ::::::::  :::::::::  :::::::::: :::     ::: ::::    ::::      ::: ::::::::::: ::::::::::
+:+:+: :+:+:+   :+: :+:   :+:    :+: :+:    :+: :+:        :+:     :+: +:+:+: :+:+:+   :+: :+:   :+:     :+:
+:+ +:+:+ +:+  +:+   +:+  +:+        +:+    +:+ +:+        +:+     +:+ +:+ +:+:+ +:+  +:+   +:+  +:+     +:+
+#+  +:+  +#+ +#++:++#++: +#+        +#+    +:+ +#++:++#   +#+     +:+ +#+  +:+  +#+ +#++:++#++: +#+     +#++:++#
+#+       +#+ +#+     +#+ +#+        +#+    +#+ +#+         +#+   +#+  +#+       +#+ +#+     +#+ +#+     +#+
#+#       #+# #+#     #+# #+#    #+# #+#    #+# #+#          #+#+#+#   #+#       #+# #+#     #+# #+#     #+#
###       ### ###     ###  ########  #########  ##########     ###     ###       ### ###     ### ###     ##########
`


func init() {
	// Print the ASCII banner before any subcommand executes
	rootCmd.PersistentPreRun = func(cmd *cobra.Command, args []string) {
		color.Cyan(asciiBanner)
	}

	// If you prefer to show a usage message at the root level without a subcommand
	rootCmd.SetHelpTemplate(
		`{{.Long}}

Usage:
  {{.Use}} [command]

Available Commands:
  install       Interactively choose software to install on macOS

Use "{{.Use}} [command] --help" for more information about a command.
`,
	)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		color.Red(err.Error())
		os.Exit(1)
	}
}
