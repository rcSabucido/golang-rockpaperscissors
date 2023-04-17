package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("Rock, paper, or scissors?: ")

	var userWeapon string
	fmt.Scanln(&userWeapon)

	npcWeapon := rand.Intn(3)

	fmt.Println("")
	fmt.Println("Rock, Paper, Scissors, Shoot!")

	switch npcWeapon {
	case 0:
		fmt.Println("The opponent has chosen: Rock!")
	case 1:
		fmt.Println("The opponent has chosen: Paper!")
	case 2:
		fmt.Println("The opponent has chosen: Scissors!")
	}
	fmt.Println("")
	fmt.Println("")

	switch {
	case userWeapon == "Rock" && npcWeapon == 0:
		fmt.Println("It's a draw!")
	case userWeapon == "Rock" && npcWeapon == 1:
		fmt.Println("You Lose :(")
	case userWeapon == "Rock" && npcWeapon == 2:
		fmt.Println("You win!")
	case userWeapon == "Paper" && npcWeapon == 0:
		fmt.Println("You win!")
	case userWeapon == "Paper" && npcWeapon == 1:
		fmt.Println("It's a draw!")
	case userWeapon == "Paper" && npcWeapon == 2:
		fmt.Println("You lose :(")
	case userWeapon == "Scissors" && npcWeapon == 0:
		fmt.Println("You lose :(")
	case userWeapon == "Scissors" && npcWeapon == 1:
		fmt.Println("You win!")
	case userWeapon == "Scissors" && npcWeapon == 2:
		fmt.Println("It's a draw!")
	}
}
