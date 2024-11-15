package helpers

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Prompt affiche une question et renvoie la réponse de l'utilisateur en tant que chaîne
func Prompt(question string) string {
	fmt.Print(question)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Error reading input: %v\n", err)
		return ""
	}
	return strings.TrimSpace(input)
}

// PromptWithDefault affiche une question avec une valeur par défaut
// Si l'utilisateur appuie sur "Entrée" sans rien saisir, la valeur par défaut est utilisée
func PromptWithDefault(question, defaultValue string) string {
	fmt.Printf("%s (default: %s): ", question, defaultValue)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Error reading input: %v\n", err)
		return defaultValue
	}
	input = strings.TrimSpace(input)
	if input == "" {
		return defaultValue
	}
	return input
}

// PromptYesNo affiche une question avec une réponse oui/non
// Retourne true pour oui (y/Y), false pour non (n/N)
func PromptYesNo(question string) bool {
	for {
		fmt.Printf("%s (y/n): ", question)
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Error reading input: %v\n", err)
			continue
		}
		input = strings.TrimSpace(strings.ToLower(input))
		if input == "y" {
			return true
		} else if input == "n" {
			return false
		} else {
			fmt.Println("Please enter 'y' or 'n'.")
		}
	}
}