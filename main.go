package main

import (
	"github.com/bwmarrin/discordgo"
	"io"
	"os"
	"log"
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
	token := make([]byte, 100)
	_, _ := file.Read(token)
	if err != nil {
		log.Fatal(err)
	}

	discord, err := discordgo.New("Bot " + token)
}