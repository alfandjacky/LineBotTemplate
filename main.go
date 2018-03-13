// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"math/rand"
	"strconv"
	"time"
	
	"github.com/line/line-bot-sdk-go/linebot"
)
var Str1 string

var bot *linebot.Client

func main() {
	var err error
	rand.Seed(time.Now().UnixNano())
	bot, err = linebot.New(os.Getenv("ChannelSecret"), os.Getenv("ChannelAccessToken"))
	log.Println("Bot:", bot, " err:", err)
	http.HandleFunc("/callback", callbackHandler)
	port := os.Getenv("PORT")
	addr := fmt.Sprintf(":%s", port)
	http.ListenAndServe(addr, nil)
}


func callbackHandler(w http.ResponseWriter, r *http.Request) {
	events, err := bot.ParseRequest(r)
	 
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				//以上已經篩選好訊息 純文字
				switch message.Text{
				case "D66":
					Str1 = strconv.Itoa(d66())
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(Str1)).Do()
				err != nil {
					log.Print(err)
				} 
				default :
					Str1 = "沒東西"
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(Str1)).Do()
				err != nil {
					log.Print(err)
				} 
				}
				
			}
		}
	}
}

//產生隨機數
func diceroll(diceside int) int {
	san := rand.Intn(diceside)+1
	return san
}
//D66判定
func d66() int {
	var dice1 = diceroll(6)
	var dice2 = diceroll(6)
	diceresult := dice1*10 + dice2
	return diceresult
}
//if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.ID+":"+message.Text+ Str1)).Do(); err != nil {
					log.Print(err)
