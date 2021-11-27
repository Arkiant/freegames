package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
)

type Game struct {
	Name        string    `json:"name"`
	Photo       string    `json:"photo"`
	URL         string    `json:"url"`
	Platform    string    `json:"platform"`
	AvailableTo time.Time `bson:"available_to" json:"available_to"`
}

func main() {

	var (
		Token string
		Host  string
	)

	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.StringVar(&Host, "h", "", "Host")
	flag.Parse()

	// 1- AUTHENTICATION
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		log.Fatalf("Invalid bot parameters: %v", err)
	}

	// 2- CONFIGURATION

	// 3- HANDLER COMMANDS

	dg.AddHandler(handlerCommands(Host))
	dg.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildMessages)

	// 4- BOT CONNECTION
	err = dg.Open()
	if err != nil {
		log.Fatalf("Failed open bot: %v", err)
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

// SendFreeGames send all free games to discord client
func sendFreeGames(dg *discordgo.Session, host string, channel string) error {
	fmt.Println("Sending message into freegames channel...")
	sendFreeGamesToChannel(dg, host, channel)
	return nil
}

func getGames(host string) ([]Game, error) {
	cHTTP := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("http://%s/freegames", host), nil)
	if err != nil {
		return nil, errors.New("can't create a request")
	}

	res, err := cHTTP.Do(req)
	defer func() {
		if res != nil && res.Body != nil {
			res.Body.Close()
		}
	}()
	if err != nil {
		return nil, errors.New("can do the request")
	}

	games := []Game{}
	respText, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.New("cant' read the response")
	}

	err = json.Unmarshal(respText, &games)
	if err != nil {
		return nil, errors.New("can't unmarshal the response")
	}

	return games, nil
}

// sendFreeGamesToChannel this method send all games into a specific channel
func sendFreeGamesToChannel(dg *discordgo.Session, host, channelID string) error {

	games, err := getGames(host)
	if err != nil {
		return err
	}

	for _, v := range games {
		dg.ChannelMessageSend(channelID, v.URL)
	}

	return nil
}

// handlerCommands execute freegames command
func handlerCommands(host string) func(s *discordgo.Session, m *discordgo.MessageCreate) {
	return func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == s.State.User.ID {
			return
		}

		log.Printf("Command %s received from %s", m.Content, m.ChannelID)

		switch m.Content {
		case "!freegames":
			sendFreeGames(s, host, m.ChannelID)
		default:
			log.Printf("command not found %s", m.Content)

		}
	}
}
