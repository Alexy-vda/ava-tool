package helpers

import (
	"fmt"
	"os"
	"os/exec"
)

// execGoModInit initialise un module Go et exécute go mod tidy
func ExecGoModInit(baseDir string, name string) {
	if _, err := os.Stat(baseDir); os.IsNotExist(err) {
		fmt.Printf("Directory %s does not exist\n", baseDir)
		return
	}

	// Exécuter go mod init
	cmd := exec.Command("go", "mod", "init", name)
	cmd.Dir = "./" + baseDir
	if err := cmd.Run(); err != nil {
		fmt.Printf("Failed to execute go mod init: %v\n", err)
	}

	// Exécuter go mod tidy dans le répertoire
	cmd = exec.Command("go", "mod", "tidy")
	cmd.Dir = "./" + baseDir
	if err := cmd.Run(); err != nil {
		fmt.Printf("Failed to execute go mod tidy: %v\n", err)
	}
}
