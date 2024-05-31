# Happy Birthday Mailer

Small CLI utility to send birthday emails to a list of recipients. Written in [Go](https://go.dev/).

It uses the [Imgflip](https://imgflip.com) api to generate funny memes by randomly selecting a template from a list of IDs.

It is built to run via a daily cron job to send out birthday emails automatically to all recipients in a given JSON file.

## First Run

On first run, the program will create a `config.json` file in the **same directory as the executable**.

```toml
title = "Happy Birthday %s 🎉🎂🎈🎁🎊🥳" # Email subject

[smtp]
host = "smtp.gmail.com"      # SMTP server    
port = 587                   # SMTP port
user = "example@gmail.com"   # SMTP username
pass = "CHANGE_ME"           # SMTP password

[imgflip]
text = "Happy Birthday!"     # Text to display on the meme
user = "example"             # Imgflip username
pass = "CHANGE_ME"           # Imgflip password
ids = [                      # Meme template IDs
    11797874,
    31643391
]
```

## Input

This is an example of the input `fellows.json` file.

```json
{
  "1": {
    "Forename": "John",
    "Surname": "Doe",
    "Birthdate": "01.06.2000",
    "Email": "john.doe@example.com"
  },
  "2": {
    "Forename": "Mary",
    "Surname": "Jane",
    "Birthdate": "31.05.2000",
    "Email": "mary.jane@example.com"
  }
}
```

## Usage

```bash
bdaymail fellows.json
```

## Build

    make

## Cross-compile

    make cross

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---
Copyright © 2023 Typomedia Foundation. All rights reserved.
