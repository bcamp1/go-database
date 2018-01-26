package main

import (
	"fmt"
	"iofile"
)

func main() {
	menu()
}

// Main menu
func menu() {
	fmt.Print("\nMenu\n(1) Log In\n(2) Create Account\n")
	switch iofile.Get() {
	case "1":
		login()
		break
	case "2":
		createAccount()
		break
	default:
		menu()
	}
}

func login() {
	fmt.Println("=====Returning User Login=====Type q to go to back to menu=====")
	n := iofile.Ask("Username: ")
	if n == "q" {
		menu()
	}
	p := iofile.Ask("Password: ")
	if p == "q" {
		menu()
	}
	line := whichUser(n)
	if line != 0 {
		fmt.Println("username found, checking password...")
		if iofile.ReadLine("users.txt", line+1) == p {
			fmt.Println("Logged in successfully.")
		} else {
			fmt.Print("Password incorrect. Please try again\n\n")
		}
	} else {
		fmt.Print("username not found. Please try again\n\n")
	}
}

func whichUser(username string) int {
	for i := 1; i < 99; i += 2 {
		if iofile.ReadLine("users.txt", i) == username {
			return i
		}
	}
	return 0
}

func createAccount() {
	fmt.Println("=====New User Login=====Type q to go to back to menu=====")
	n := iofile.Ask("Please enter your username: ")
	if n == "q" {
		menu()
	}
	p := iofile.Ask("Welcome " + n + ". Enter new password: ")
	if p == "q" {
		menu()
	}
	if iofile.Ask("Confirm Password: ") == p {
		fmt.Print("Setting up your account... \nYou can now log in\n\n")
		iofile.Append("users.txt", n+"\n")
		iofile.Append("users.txt", p+"\n")
		login()
	} else {
		fmt.Println("Passwords do not match. Try again.")
		createAccount()
	}
}
