package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/ermes-labs/api-go/infrastructure"
)

func main() {
	// Check if the command-line argument is provided.
	if len(os.Args) < 2 {
		fmt.Println("Usage: deploy <input_file>")
		return
	}

	// The first command-line argument is the program path, the second is the
	// desired input file.
	filePath := os.Args[1]
	// Read JSON file.
	file, err := os.ReadFile(filePath)
	// Check if there was an error reading the file.
	if err != nil {
		fmt.Printf("Error reading JSON file: %s\n", err)
		return
	}

	// Unmarshal JSON file.
	infra, _, err := infrastructure.UnmarshalInfrastructure(file)
	// Check if there was an error unmarshaling the JSON file.
	if err != nil {
		fmt.Printf("Error unmarshaling JSON: %s\n", err)
		return
	}

	// Get all areas.
	areas := infra.Flatten()

	for _, area := range areas {
		// Marshal JSON node.
		jsonNode, err := infrastructure.MarshalNode(area.Node)
		// Check if there was an error marshaling the JSON node.
		if err != nil {
			fmt.Printf("Error marshaling JSON: %s\n", err)
			return
		}

		// Run Ansible playbook.
		// Define Ansible command
		cmd := exec.Command("ansible-playbook", "playbook.yml", "--extra-vars",
			fmt.Sprintf("target_node='%s' target_host='%s'", string(jsonNode), area.Host))

		// Set environment variables and execute
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err = cmd.Run()
		if err != nil {
			fmt.Printf("Error running Ansible playbook: %s\n", err)
			return
		}
	}
}
