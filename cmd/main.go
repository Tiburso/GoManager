package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/urfave/cli/v2"
)

func ReadLine(scanner *bufio.Scanner) string {
	if scanner.Scan() {
		return strings.TrimSpace(scanner.Text())
	}

	return ""
}

func ShowMenu() {
	fmt.Println("1. Create application")
	fmt.Println("2. Delete application")
	fmt.Println("3. Update application")
	fmt.Println("4. Show applications")
	fmt.Println("5. Show companies")
	fmt.Println("6. Show company applications")
	fmt.Println("7. Exit")
	fmt.Print("Enter your choice: ")
}

// func CLITool(db *gorm.DB) error {
// 	scanner := bufio.NewScanner(os.Stdin)

// 	for {
// 		// show menu
// 		ShowMenu()

// 		// read user input from stdin
// 		input := ReadLine(scanner)

// 		// switch on user input
// 		var err error
// 		switch input {
// 		case "1":
// 			err = CreateApplication(db)
// 		case "2":
// 			err = DeleteApplication(db)
// 		case "3":
// 			err = UpdateApplication(db)
// 		case "4":
// 			ShowApplications(db)
// 		case "5":
// 			ShowCompanies(db)
// 		case "6":
// 			err = ShowCompanyApplications(db)
// 		case "7":
// 			return nil
// 		default:
// 			fmt.Println("Invalid input")
// 		}

// 		if err != nil {
// 			return err
// 		}
// 	}
// }

func main() {
	app := &cli.App{
		EnableBashCompletion: true,
		Name:                 "GoManager",
		Usage:                "GoManager CLI",
		Commands: []*cli.Command{
			ServerCommand(),
			ClientCommand(),
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
