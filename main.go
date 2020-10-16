package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/HexSeal/MemeMachineBot/meme"
)

// Testing a basic implementation of a bot from
// https://github.com/bwmarrin/discordgo/blob/master/examples/pingpong/main.go

// Variables used for command line parameters
var (
	Token string
)

func init() {

	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}

func main() {

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)

	// In this example, we only care about receiving message events.
	dg.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildMessages)

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}
	
	usermsg := strings.SplitAfterN(m.Content, " ", 2)

	// If they want a meme, humbly oblige
	if usermsg[0] == "#meme " {
		// Check if the bot has been summoned, this splits the #meme from the rest
		// println(usermsg[0])

		// Breaking down the format and captions after the usercommand
		userinput := strings.SplitAfterN(m.Content, "#meme ", 2)
		// println(userinput[0], userinput[1])

		// Breaks down the input to get the individual attributes
		usercommand := strings.SplitAfterN(userinput[1], ",", 3)
		format := usercommand[0]
		caption1 := usercommand[1]
		caption2 := usercommand[2]
		println(format, caption1, caption2)

		location := ""
		// Get the right meme format
		switch format {
		case "facts,":
			location = "./meme_formats/facts_meme.jpg"
		case "wonka,":
			location = "./meme_formats/willy_wonka.jpg"
		default:
			location = "./meme_formats/facts_meme.jpg"
		}
		println("Location: ", location)

		// This is where we'll eventually use a function from antoher file to combine the image with the captions
		// Eventually, this will open the new meme itself instead of the format
		meme, err := os.Open(location)
		if err != nil {
			log.Fatalln(err)
		}

		s.ChannelFileSend(m.ChannelID, "meme.jpg", meme)
		s.ChannelMessageSend(m.ChannelID, caption1)
		s.ChannelMessageSend(m.ChannelID, caption2)

		// Testing meme creation 
		meme_machine.CreateMeme(location, caption1, caption2)
	}
}