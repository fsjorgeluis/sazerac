package main

import (
	"log"

	"{{ .Module }}/cmd/{{ .ProjectName }}/di"
)

func main() {
	// Initialize dependencies
	container, err := di.NewContainer()
	if err != nil {
		log.Fatalf("Failed to initialize dependencies: %v", err)
	}
	defer container.Close()

	// Execute handler to demonstrate the full flow
	// This runs: Handler -> UseCase -> Repository
	if err := container.{{ .UseCase }}Handler.Run(); err != nil {
		log.Fatalf("Failed to execute handler: %v", err)
	}
}
