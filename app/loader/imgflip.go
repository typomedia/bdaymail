package loader

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"github.com/typomedia/bdaymail/app/structs"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
)

func Meme(fellow structs.Fellow) string {
	api := "https://api.imgflip.com/caption_image"
	data := url.Values{}
	data.Set("template_id", strconv.Itoa(randomId()))
	data.Set("username", viper.GetString("imgflip.user"))
	data.Set("password", viper.GetString("imgflip.pass"))
	data.Set("text0", viper.GetString("imgflip.text"))
	data.Set("text1", fmt.Sprintf("%s %s", fellow.Forename, fellow.Surname))

	client := &http.Client{}
	req, err := http.NewRequest("POST", api, bytes.NewBufferString(data.Encode()))
	if err != nil {
		log.Println("Error creating request:", err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error making request:", err)
	}
	defer resp.Body.Close()

	var result structs.Imgflip
	json.NewDecoder(resp.Body).Decode(&result)

	return result.Data.URL
}

func randomId() int {
	ids := viper.GetIntSlice("imgflip.ids")

	// Get a random ID from the list
	id := rand.Intn(len(ids))

	return ids[id]
}
