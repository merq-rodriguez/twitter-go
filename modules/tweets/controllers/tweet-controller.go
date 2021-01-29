package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	HttpStatus "github.com/merq-rodriguez/twitter-clone-backend-go/common/response/http"
	. "github.com/merq-rodriguez/twitter-clone-backend-go/modules/tweets/models"
	tweetService "github.com/merq-rodriguez/twitter-clone-backend-go/modules/tweets/services"
)

/*
GetTweetsByUserID function controller
*/
func GetTweetsByUserID(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	page := r.URL.Query().Get("page")

	if len(ID) < 1 {
		http.Error(w, "Id parameter is required", HttpStatus.NOT_ACCEPTABLE)
		return
	}
	if len(page) < 1 {
		http.Error(w, "Page parameter is required", HttpStatus.NOT_ACCEPTABLE)
		return
	}

	_page, err := strconv.Atoi(page)
	if err != nil {
		http.Error(w, "The page parameter must be greater than 0", HttpStatus.NOT_ACCEPTABLE)
		return
	}

	pag := int64(_page)
	results, success := tweetService.GetTweetsByUserID(ID, pag)

	if success == false {
		http.Error(w, "Error read tweets", HttpStatus.NOT_ACCEPTABLE)
		return
	}

	w.WriteHeader(HttpStatus.OK)
	json.NewEncoder(w).Encode(results)
}

/*
CreateTweet function controller
*/
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
