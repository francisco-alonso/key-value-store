package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/francisco-alonso/key-value-store/internal/kvstore"
	"github.com/francisco-alonso/key-value-store/internal/server"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, proceeding with default or system environment variables.")
	}

	// Read the file path from environment variable or provide a default
	filePath := os.Getenv("KVSTORE_FILE_PATH")
	if filePath == "" {
		filePath = "./data/data.json" // Default path
	}

	// Initialize the KeyValueStore with the persistence file path
	kv, err := kvstore.NewKeyValueStore(filePath)
	if err != nil {
		log.Fatalf("Failed to initialize key-value store: %v", err)
	}

	// Start the API server in a goroutine
	go server.StartAPI(kv)

	// Setup signal handling for shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// Start the CLI loop
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Key-Value Store CLI")
	fmt.Println("Commands: SET key value, GET key, DELETE key, EXISTS key, EXIT")

	// Main CLI loop
	for {
		select {
		case <-stop:
			// Handle graceful shutdown on signal
			fmt.Println("Shutdown signal received. Exiting...")
			err := kv.Save()
			if err != nil {
				log.Printf("Failed to save data on exit: %v", err)
			}
			return
		default:
			// Read user input from the CLI
			fmt.Print("> ")
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)
			parts := strings.Split(input, " ")

			// Handle CLI commands
			switch strings.ToUpper(parts[0]) {
			case "SET":
				if len(parts) != 3 {
					fmt.Println("Usage: SET key value")
					continue
				}
				kv.Set(parts[1], parts[2])
				fmt.Println("OK")
			case "GET":
				if len(parts) != 2 {
					fmt.Println("Usage: GET key")
					continue
				}
				if value, err := kv.Get(parts[1]); err == nil {
					fmt.Printf("Value: %s\n", value)
				} else {
					fmt.Println(err)
				}
			case "DELETE":
				if len(parts) != 2 {
					fmt.Println("Usage: DELETE key")
					continue
				}
				if err := kv.Delete(parts[1]); err == nil {
					fmt.Println("Deleted")
				} else {
					fmt.Println(err)
				}
			case "EXISTS":
				if len(parts) != 2 {
					fmt.Println("Usage: EXISTS key")
					continue
				}
				if kv.Exists(parts[1]) {
					fmt.Println("Key exists")
				} else {
					fmt.Println("Key does not exist")
				}
			case "EXIT":
				fmt.Println("Exiting...")
				err := kv.Save()
				if err != nil {
					log.Printf("Failed to save data on exit: %v", err)
				}
				return
			default:
				fmt.Println("Unknown command")
			}
		}
	}
}
