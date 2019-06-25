package main

import (
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
)

func main() {

	// open auth file
	file, err := os.Open("Auth.txt")
	// error check
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// read file
	token := make([]byte, 100)
	_, err = file.Read(token)
	if err != nil {
		log.Fatal(err)
	}

	discord, err := discordgo.New("Bot " + string(token))

	// init bot
	err = discord.Open()
}
