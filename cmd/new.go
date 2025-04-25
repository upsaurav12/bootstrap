/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"text/template"

	"github.com/spf13/cobra"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "command for creating a new project.",
	Long:  `command for creating a new project.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Check if the project name is provided
		if len(args) < 1 {
			fmt.Fprintln(cmd.OutOrStdout(), "Error: project name is required")
			return
		}

		// Get the template flag value from the command context
		tmpl, _ := cmd.Flags().GetString("template")

		// Get the project name (first argument)
		dirName := args[0]

		// Create the new project
		createNewProject(dirName, tmpl, cmd.OutOrStdout())
	},
}

func init() {
	// Add the new command to the rootCmd
	rootCmd.AddCommand(newCmd)

	// Define the --template flag for this command
	newCmd.Flags().StringP("template", "t", "", "Specify a project template (eg, go, node, python)")
}

// Function to create the project
func createNewProject(projectName string, template string, out io.Writer) {
	// Attempt to create the project directory
	err := os.Mkdir(projectName, 0755)
	if err != nil {
		fmt.Fprintf(out, "Error creating directory: %v\n", err)
		return
	}

	// Print the template that was passed

	createGoTemplate(projectName, out)

	// Print success message
	fmt.Fprintf(out, "Created '%s' successfully\n", projectName)
}

func createGoTemplate(projectName string, out io.Writer) {
	// Define the go.mod template content
	goModTemplate := `module {{.ModuleName}}

go 1.18
`

	// Define the main.go content
	mainGoFile := filepath.Join(projectName, "main.go")
	mainGoContent := `package main

import "fmt"

func main() {
	fmt.Println("Hello, world!")
}
`

	// Create go.mod file and apply the template
	goModFile := filepath.Join(projectName, "go.mod")
	tmpl, err := template.New("go.mod").Parse(goModTemplate)
	if err != nil {
		fmt.Fprintf(out, "Error creating go.mod template: %v\n", err)
		return
	}

	// Create a file to write the go.mod content
	goModFileContent := struct {
		ModuleName string
	}{
		ModuleName: projectName, // Use project name as module name
	}

	// Write go.mod file with parsed template
	goModFileHandle, err := os.Create(goModFile)
	if err != nil {
		fmt.Fprintf(out, "Error creating go.mod file: %v\n", err)
		return
	}
	defer goModFileHandle.Close()

	err = tmpl.Execute(goModFileHandle, goModFileContent)
	if err != nil {
		fmt.Fprintf(out, "Error executing go.mod template: %v\n", err)
		return
	}

	// Create main.go file
	err = os.WriteFile(mainGoFile, []byte(mainGoContent), 0644)
	if err != nil {
		fmt.Fprintf(out, "Error creating main.go: %v\n", err)
		return
	}

	// Print the result
	fmt.Fprintf(out, "Go project created in '%s'\n", projectName)
}
