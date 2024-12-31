package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/francisco-alonso/key-value-store/internal/kvstore"
)

func main() {
	kv := kvstore.NewKeyValueStore()
	reader := bufio.NewReader(os.Stdin)
		
	fmt.Println("Key-Value Store CLI")
	fmt.Println("Commands: SET key value, GET key, DELETE key, EXISTS key, EXIT")
	
	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		parts := strings.Split(input, " ")

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
			return
		default:
			fmt.Println("Unknown command")
		}
	}
}