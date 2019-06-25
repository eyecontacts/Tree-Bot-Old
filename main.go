package main

import (
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
)

func main() {

	// open auth file
	file, err := os.Open("auth.txt")
	// error check
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// read file
	token := make([]byte, 59)
	_, err = file.Read(token)
	if err != nil {
		log.Fatal(err)
	}

	discord, err := discordgo.New(string(token))
	if err != nil {
		log.Fatal(err)
	}

	// init bot
	err = discord.Open()
	if err != nil {
		log.Fatal(err)
	}
}
