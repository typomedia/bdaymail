package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"github.com/XotoX1337/tinymail"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"github.com/typomedia/bdaymail/app/loader"
	"github.com/typomedia/bdaymail/app/reader"
	"github.com/typomedia/bdaymail/app/structs"
	"log"
	"time"
)

//go:embed app/assets/*
var fs embed.FS
var err error

func main() {
	pflag.Parse()
	input := pflag.Arg(0)

	config, _ := fs.ReadFile("app/assets/config.toml")
	loader.Config(config)

	if input == "" {
		log.Fatal("Please provide a file path or URL using the -input flag")
	}

	var data []byte
	data, err = reader.File(input)

	if err != nil {
		log.Fatal(err)
	}

	var fellows structs.Fellows
	err = json.Unmarshal(data, &fellows)
	if err != nil {
		log.Fatal(err)
	}

	host := viper.GetString("smtp.host")
	user := viper.GetString("smtp.user")
	password := viper.GetString("smtp.pass")
	text := viper.GetString("imgflip.text")

	file, _ := fs.ReadFile("app/assets/email.html")
	mailer := tinymail.New(user, password, host)

	for _, fellow := range fellows {
		birthdate, _ := time.Parse("02.01.2006", fellow.Birthdate)
		birthday := birthdate.Format("02-01")
		today := time.Now().Format("02-01")

		if birthday == today {
			log.Println(text, fellow.Forename, fellow.Surname)

			fellow.Meme = loader.Meme(fellow)

			message, _ := tinymail.FromTemplateString(fellow, string(file))
			message.SetFrom(user)
			message.SetTo(fellows.Emails()...)
			message.SetSubject(fmt.Sprintf(viper.GetString("title"), fellow.Forename))
			err = mailer.SetMessage(message).Send()
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}
