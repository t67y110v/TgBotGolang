package main

import (
	"fmt"
)

func main() {
	botUrl := botApi + botToken
	offset := 0

	/*db, err := bolt.Open("bot.db", 0600, nil)
	if err != nil {
		fmt.Println("Error")
	}
	TokenReprository := NewTokenRepository(db)
	*/

	for {
		updates, err := GetUpdates(botUrl, offset)
		if err != nil {
			break
		}
		fmt.Println(updates)
		for _, update := range updates {
			err = Respond(botUrl, update)
			offset = update.UpdateId + 1
		}
	}
}
