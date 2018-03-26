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
	"regexp"
	"net/http"
	"os"
	"math/rand"
	"strconv"
	"strings"
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
	imageURL := app.appBaseURL + "/static/buttons/1040.jpg"
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
				
				//抓取字串至第一個空格
				var firsttime = regexp.MustCompile(`^\S+`)
				fstword := firsttime.FindString(message.Text)
				
				//輸出運算結果
				wordtitle ,retype := titleread (fstword)
				
				//判斷輸出類別
				switch retype {
				case "dice" :
					//負責傳出訊息
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(wordtitle)).Do()
					        err != nil {
						log.Print(err)
					} 
					//負責傳出訊息
				case "box" :
					if _, err := app.bot.ReplyMessage(
									replyToken,
							linebot.NewTemplateMessage(wordtitle, template),
									).Do()
					err != nil {
									return err
										}
					
				}
				
			}
		}
	}
}


//輸入比較符號與兩數字給出真偽
func camepareto (wordin string,num1 int,num2 int)(string){
	word := ""
	switch wordin {
			case ">=":
		if num1 >= num2 {
			word = "true"
		}else{
			word = "false"
		}
			case "<=":
		if num1 <= num2 {
			word = "true"
		}else{
			word = "false"
		}
			case "=":
		if num1 == num2 {
			word = "true"
		}else{
			word = "false"
		}
			case ">":
		if num1 > num2 {
			word = "true"
		}else{
			word = "false"
		}
			case "<":
		if num1 < num2 {
			word = "true"
		}else{
			word = "false"
		}
	}
	return word
}


//執行多次普通擲骰
func ddtitle(wordin string) string {
	var reg = regexp.MustCompile(`\d+(?i:d)\d+`)
	compare, _ := regexp.MatchString("[0-9]+[>=<]{1,2}[0-9]+$", wordin)
	word := "基本擲骰:\n"+"("+wordin+")\n→"
	times := strings.Count(wordin, "+") + 1
	totleresult := 0
	ttresolt := reg.FindAllString(wordin, times)
	for i:= 0 ; i < times ; i++ {
		word1 , number1 := ddone(ttresolt[i])
		totleresult = totleresult + number1
		if i == times-1 {
			word = word + word1 + ""
		}else{
			word = word + word1 + "+"
		}
	}
	
	if compare {
		var comeparetype = regexp.MustCompile(`[>=<]{1,2}`)
		var numbercompare = regexp.MustCompile("[0-9]+$")
		ase := numbercompare.FindString(wordin)
		moon1 := comeparetype.FindString(wordin)
		int11,_ :=strconv.Atoi(ase)  
		word = word+moon1+ase+"\n→"+camepareto (moon1,totleresult,int11)	
	}else{
		word = word + "\n→" + strconv.Itoa(totleresult)
	}
	return word
	
}

//普通擲骰正則表達式提取數字
func cutmath(wordin string) (int, int) {
        var mdm = regexp.MustCompile(`\d+(?i:d)\d+`)
	var md = regexp.MustCompile(`^\d+`)
	var dm = regexp.MustCompile(`\d+$`)
	word := mdm.FindString(wordin)
	math1 ,_ :=strconv.Atoi(md.FindString(word))
	math2 ,_ :=strconv.Atoi(dm.FindString(word))
	return math1 , math2
}

//普通擲骰
func ddone(wordin string) ( string , int ) {
	word := "("
	dicenumber, diceside := cutmath(wordin)
	diceresult := make([]int, dicenumber)
	number := 0
	
	for i:=0; i < dicenumber ; i++ { 
		tmemath := diceroll(diceside) 
		diceresult[i] = tmemath
	}
	for i:=0; i < dicenumber ; i++ { 
		word1 := strconv.Itoa(diceresult[i])
		number = number + diceresult[i]
		if i == 0 {
			word = word + word1
		}else{
		word = word +"+"+ word1
		}
	}
	word = word + ")"
	return word , number
	
}

//測試
func tetitle() string {
	word := "測試輸出:"
	var a,b,c,d,e,f int =0,0,0,0,0,0
	for i := 0 ; i <420 ; i++ {
		abc := diceroll(6)
		switch abc {
			case 1 :
			a++
			case 2 :
			b++
			case 3 :
			c++
			case 4 :
			d++
			case 5 :
			e++
			case 6 :
			f++
		}
		
	}
	aa:=strconv.Itoa(a)  
		bb:=strconv.Itoa(b)  
		cc:=strconv.Itoa(c)  
		dd:=strconv.Itoa(d)  
		ff:=strconv.Itoa(e)  
		ee:=strconv.Itoa(f)  
		word = word +"\n"+aa+"\n"+bb+"\n"+cc+"\n"+dd+"\n"+ee+"\n"+ff
	return word
}

//克蘇魯7th擲骰
func coc7thtitle() string {
	word := "CoC7th"
	return word
}

//絕對奴隸擲骰
func astitle( wordin string ) string {
	word := "絕對隸奴擲骰:\n→"
	var cutas = regexp.MustCompile("[^(?i:^as)]+")
	ase := cutas.FindString(wordin)
	var num1 = regexp.MustCompile("^[0-9]+")
	number1 := num1.FindString(ase)
	word1 ,_ :=strconv.Atoi(number1)
	aresult,agreat,aword := asd66 (word1)
	awresult :=strconv.Itoa(aresult)
	word = word + aword
	if aword == "(大失敗)" {
	}else{
	a, _ := regexp.MatchString("[^0-9]+", ase)
	if a {
		var num2 = regexp.MustCompile("[0-9]+$")
		number2 := num2.FindString(ase)
		word2 ,_ :=strconv.Atoi(number2)
		var cutnum = regexp.MustCompile("[^0-9]+")
		compare := cutnum.FindString(ase)
		if compare == "v"{
			bresult,_,bword := asd66 (word2)
			bwresult :=strconv.Itoa(bresult)
			word = word + "VS" + bword + "\n→" + "(" + awresult+ ")" + "VS" + "(" + bwresult + ")" +"\n→"
			if aresult >= bresult {
				word = word + "true"
			}else{
				word = word + "false"
			}
			word = word +"\n→DK增加" + agreat
			
		}else if compare == ">="{
			if aresult >= word2 {
				word = word + ">=" + number2 + "\n→" + awresult+ ">=" + number2+ "\n→" + "true" + "\n→DK增加" + agreat
			}else{
				word = word + ">=" + number2 + "\n→" + awresult+ ">=" + number2+ "\n→" + "false" + "\n→DK增加" + agreat
			}
			
		}
	}else{
		word = word + "\n→" + awresult + "\n→DK增加" + agreat
	}
	}
	return word
}
//絕對隸奴D66
func asd66 (numin int) (int , string , string){
	a , b , word66 := d66title()
	sixtime := strings.Count(word66, "6")
	fivetime := strings.Count(word66, "5")
	resort := a + b - (6*sixtime)
	sixtime2 := sixtime*sixtime
	sixword := strconv.Itoa(sixtime2)
	numout := numin - resort
	if fivetime == 2 {
		word66 = "(大失敗)"
	}
	return numout ,sixword ,word66
}
//D66擲骰
func d66title() (int,int,string) {
	var dice1 = diceroll(6)
	var dice2 = diceroll(6)
	word1 := strconv.Itoa(dice1)
	word2 := strconv.Itoa(dice2)
	word := "(" + word1 +"," + word2 +")" 
	return dice1,dice2,word
}
//女僕醬判斷
func madogo {
	template := linebot.NewCarouselTemplate(
	linebot.NewCarouselColumn(
		imageURL, "hoge", "fuga",
		linebot.NewURITemplateAction("Go to line.me", "https://line.me"),
		linebot.NewPostbackTemplateAction("Say hello1", "hello こんにちは", ""),
	),
	linebot.NewCarouselColumn(
		imageURL, "hoge", "fuga",
		linebot.NewPostbackTemplateAction("言 hello2", "hello こんにちは", "hello こんにちは"),
		linebot.NewMessageTemplateAction("Say message", "Rice=米"),
	),
)
}
//句首判斷要做什麼
func titleread(wordin string) string {
	var word string
	aa := true
	a, _ := regexp.MatchString("(?i:^cc)", wordin)
	b, _ := regexp.MatchString("(?i:^AS)[0-9]", wordin)
	c, _ := regexp.MatchString("(?i:^D66)", wordin)
	d, _ := regexp.MatchString("(?i:^te)", wordin)
	e, _ := regexp.MatchString("^[0-9]+(?i:d)[0-9]+", wordin)
	f, _ := regexp.MatchString("^[女][僕][醬]", wordin)
	switch aa{
		case a :
		wordout = coc7thtitle()
		retype = "dice"
		
		case b :
		wordout = astitle(wordin)
		retype = "dice"
		
		case c :
		_, _, word66 := d66title()
		wordout = "D66擲骰:\n" + word66
		retype = "dice"
		
		case d :
		wordout = tetitle()
		retype = "dice"
		
		case e :
		wordout = ddtitle(wordin)
		retype = "dice"
		
		csae f :
		madogo ()
		wordout = "女僕醬"
		retype = "box"
	} 
	return wordout,retype
}
	
//產生隨機數
func diceroll(diceside int) int {
	san := rand.Intn(diceside)+1
	return san
}




	

//if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.ID+":"+message.Text+ Str1)).Do(); err != nil {log.Print(err)}
