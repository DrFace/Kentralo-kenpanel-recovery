package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	disasterkit "github.com/kentralo/kenpanel-recovery/disaster-kit"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "generate-kit":
		cmd := flag.NewFlagSet("generate-kit", flag.ExitOnError)
		host := cmd.String("host", "srv01.example.com", "Server hostname")
		cmd.Parse(os.Args[2:])

		kit := disasterkit.GenerateKit(*host)
		data, _ := json.MarshalIndent(kit, "", "  ")
		fmt.Println(string(data))

	case "version":
		fmt.Println("kenpanel-recovery v2.3.0 (standalone engine)")

	default:
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("Usage: kenpanel-recovery <command> [options]")
	fmt.Println("Commands:")
	fmt.Println("  generate-kit   Create a self-contained disaster runbook manifest")
	fmt.Println("  docker-rescue  Inspect and recover Docker containers offline")
	fmt.Println("  version        Show version")
}
