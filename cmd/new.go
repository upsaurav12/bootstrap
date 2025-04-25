/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
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

	err = renderTemplateDir("templates/"+template, projectName, TemplateData{
		ModuleName: projectName,
	})

	if err != nil {
		fmt.Fprintf(out, "Error rendering templates", err)
	}
	fmt.Fprintf(out, "Created '%s' successfully\n", projectName)
}

type TemplateData struct {
	ModuleName string
}

func renderTemplateDir(templatePath, destinationPath string, data TemplateData) error {
	return filepath.Walk(templatePath, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relPath, _ := filepath.Rel(templatePath, path)
		targetPath := filepath.Join(destinationPath, strings.TrimSuffix(relPath, ".tmpl"))

		if info.IsDir() {
			return os.MkdirAll(targetPath, 0755)
		}

		tmpl, err := template.ParseFiles(path)
		if err != nil {
			return err
		}

		outFile, err := os.Create(targetPath)
		if err != nil {
			return err
		}
		defer outFile.Close()

		return tmpl.Execute(outFile, data)
	})
}
