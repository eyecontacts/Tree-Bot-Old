package main

import (
	"github.com/bwmarrin/discordgo"
	"io"
	"os"
)

func main(){
	
	// open auth file
	file, err := os.Open("auth.txt")
	// error check
    if err != nil {
        log.Fatal(err)
    }
	defer file.Close()
	
	// read file
	token, err := io.Read(file)
	if err != nil {
        log.Fatal(err)
    }

	discord, err := discordgo.New("Bot " + token)
}