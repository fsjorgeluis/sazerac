package main

import (
	"fmt"
	"log"

	"{{ .Module }}/cmd/{{ .ProjectName }}/di"
)

func main() {
	// Initialize dependency injection container
	container, err := di.NewContainer()
	if err != nil {
		log.Fatalf("Failed to initialize container: %v", err)
	}

	// Run the handler
	if err := container.{{ .UseCase }}Handler.Run(); err != nil {
		log.Fatalf("Handler execution failed: %v", err)
	}

	fmt.Println("✅ Execution completed successfully!")
}
