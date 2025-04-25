package cmd

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewProject_Success(t *testing.T) {
	tempDir := t.TempDir()
	projectName := "test-project"
	fullPath := filepath.Join(tempDir, projectName)

	// Change to temp directory to create projectName relative to it
	oldDir, err := os.Getwd()
	assert.NoError(t, err, "Failed to get working directory")
	defer os.Chdir(oldDir)
	err = os.Chdir(tempDir)
	assert.NoError(t, err, "Failed to change to temp directory")

	var out bytes.Buffer
	createNewProject(projectName, "go", &out)

	// Check if directory was created
	_, err = os.Stat(fullPath)
	assert.NoError(t, err, "Expected project directory to be created")

	// Check output
	expected := fmt.Sprintf("Created '%s' successfully\n", projectName)
	fmt.Println(out.String())
	assert.Equal(t, expected, out.String(), "Unexpected output")
}

/*

func TestCreateNewProject_DirectoryAlreadyExists(t *testing.T) {
	tempDir := t.TempDir()
	projectName := "test-project"
	fullPath := filepath.Join(tempDir, projectName)

	err := os.Mkdir(fullPath, 0755)
	assert.NoError(t, err, "Failed to set-up directory")
	var out bytes.Buffer
	createNewProject(projectName, "go", &out)

	_, err = os.Stat(fullPath)
	assert.NoError(t, err, "Expected directory to still exists")

	assert.Contains(t, out.String(), "Error creating directory", "Expected error message")
}*/

func TestCreateNewProject_InvalidPath(t *testing.T) {
	tempDir := t.TempDir()
	projectName := "invalid\000name"
	invalidPath := filepath.Join(tempDir, projectName)

	var out bytes.Buffer
	createNewProject(projectName, "go ", &out)

	_, err := os.Stat(invalidPath)
	assert.Error(t, err, "Expected no directory to be created")

	assert.Contains(t, out.String(), "Error creating directory", "Expected error message")
}
