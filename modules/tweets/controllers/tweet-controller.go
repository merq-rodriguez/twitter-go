package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	HttpStatus "github.com/merq-rodriguez/twitter-clone-backend-go/common/response/http"
	. "github.com/merq-rodriguez/twitter-clone-backend-go/modules/tweets/models"
	tweetService "github.com/merq-rodriguez/twitter-clone-backend-go/modules/tweets/services"
)

func CreateTweet(w http.ResponseWriter, r *http.Request) {
	var tweet Tweet
	err := json.NewDecoder(r.Body).Decode(&tweet)

	registry := Tweet{
		UserID:    tweet.UserID,
		Message:   tweet.Message,
		Timestamp: time.Now(),
	}

	_, status, err := tweetService.CreateTweet(registry)
	if err != nil {
		http.Error(w, "Error tweet not created "+err.Error(), HttpStatus.BAD_REQUEST)
		return
	}

	if status == false {
		http.Error(w, "Tweet not created", HttpStatus.BAD_REQUEST)
		return
	}

	w.WriteHeader(HttpStatus.CREATED)
}
