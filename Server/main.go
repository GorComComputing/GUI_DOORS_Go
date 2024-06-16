package main

import (
	"fmt"
	"net/http"
	"log"
	"strings"
	"github.com/Syfaro/telegram-bot-api"
	"flag"
	//"encoding/gob"
    "os"
    //"bufio"
    //tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
    "io/ioutil"
)


var telegramBotToken string					// глобальная переменная в которой храним токен
var chatID     		int64					// chatid чата для нотификаций
var tgBotChan	 	chan string = make(chan string, 10)

var numericKeyboardSTART = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonURL("gorcom.online", "http://www.gorcom.online"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonURL("Камеры", "http://www.gorcom.online:8081"),
	),
)

var numericKeyboardCOFFEE = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Red ON", "Red ON COFFEE"),
		tgbotapi.NewInlineKeyboardButtonData("Red OFF", "Red OFF COFFEE"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Гирлянда ON", "Relay ON COFFEE"),
		tgbotapi.NewInlineKeyboardButtonData("Гирлянда OFF", "Relay OFF COFFEE"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Вентилятор ON", "Fun ON COFFEE"),
		tgbotapi.NewInlineKeyboardButtonData("Вентилятор OFF", "Fun OFF COFFEE"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Насос ON", "Pump ON COFFEE"),
		tgbotapi.NewInlineKeyboardButtonData("Насос OFF", "Pump OFF COFFEE"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Конвейер ON", "Track ON COFFEE"),
		tgbotapi.NewInlineKeyboardButtonData("Конвейер OFF", "Track OFF COFFEE"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonURL("gorcom.online", "http://www.gorcom.online"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonURL("Камеры", "http://www.gorcom.online:8081"),
	),
)

var numericKeyboardPICO = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Red ON", "Red ON PICO"),
		tgbotapi.NewInlineKeyboardButtonData("Red OFF", "Red OFF PICO"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Yellow ON", "Yellow ON PICO"),
		tgbotapi.NewInlineKeyboardButtonData("Yellow OFF", "Yellow OFF PICO"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Green ON", "Green ON PICO"),
		tgbotapi.NewInlineKeyboardButtonData("Green OFF", "Green OFF PICO"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Led ON", "Led ON PICO"),
		tgbotapi.NewInlineKeyboardButtonData("Led OFF", "Led OFF PICO"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Laser ON", "Laser ON PICO"),
		tgbotapi.NewInlineKeyboardButtonData("Laser OFF", "Laser OFF PICO"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonURL("gorcom.online", "http://www.gorcom.online"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonURL("Камеры", "http://www.gorcom.online:8081"),
	),
)


// Инимциализация для подключения к Телеграм-боту
func init() {
	// https://api.telegram.org/bot6741341050:AAH5uDOLQHpMnmZxGvLBwiXljavoaA3OOms/getUpdates
	// принимаем на входе флаг -telegrambottoken
	flag.StringVar(&telegramBotToken, "telegrambottoken", "6741341050:AAH5uDOLQHpMnmZxGvLBwiXljavoaA3OOms", "Telegram Bot Token")
	flag.Int64Var(&chatID, "chatid", -4265825342, "chatId to send messages")

	flag.Parse()

	// без него не запускаемся
	if telegramBotToken == "" {
		log.Print("-telegrambottoken is required")
		os.Exit(1)
	}
	
	if chatID == 0 {
		log.Print("-chatid is required")
		os.Exit(1)
	}
}


func tgBotChanSend(c chan string, bot *tgbotapi.BotAPI) {
    bot.Send(tgbotapi.NewMessage(chatID, <-c))
	for mes := range c {
		bot.Send(tgbotapi.NewMessage(chatID, mes))
	}

}


type Handler struct {
	fileServer http.Handler
}


func main() {	

	// start web server
	fmt.Println("Server DOORS started OK")
	fmt.Println("Port: 8081")
	
	go http.ListenAndServe(":8081", &Handler{
		fileServer: http.FileServer(http.Dir("www")),
	})

	
	// Подключение к Телеграм-боту
	// используя токен создаем новый инстанс бота
	bot, err := tgbotapi.NewBotAPI(telegramBotToken)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println("Authorized on Telegram: ", bot.Self.UserName)

	// u - структура с конфигом для получения апдейтов
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	// используя конфиг u создаем канал в который будут прилетать новые сообщения
	updates, err := bot.GetUpdatesChan(u)
	
	go tgBotChanSend(tgBotChan, bot)		// Обработчик канала Телеграм-бота

	// в канал updates прилетают структуры типа Update
	// вычитываем их и обрабатываем
	for update := range updates {
		reply := "Неизвестная команда"
		if update.CallbackQuery != nil {
    		switch update.CallbackQuery.Data {
			case "Red ON COFFEE":
				req, err := http.NewRequest("GET", "http://192.168.0.12/get?cmd=red%20on", nil)		////////////// GET
				if err != nil {
					reply = "Request FAIL"
				}
				resp, err := http.DefaultClient.Do(req)
				if err != nil {
					reply = "Request FAIL"
				}		
				defer resp.Body.Close()
				body_resp, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					reply = "Request FAIL\n"
				}
				reply = update.CallbackQuery.From.UserName + ": " + string(body_resp)
				msg := tgbotapi.NewMessage(chatID, reply)
				bot.Send(msg)
			case "Red OFF COFFEE":			
				req, err := http.NewRequest("GET", "http://192.168.0.12/get?cmd=red%20off", nil)		////////////// GET
				if err != nil {
					reply = "Request FAIL"
				}
				resp, err := http.DefaultClient.Do(req)
				if err != nil {
					reply = "Request FAIL"
				}		
				defer resp.Body.Close()
				body_resp, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					reply = "Request FAIL\n"
				}
				reply = update.CallbackQuery.From.UserName + ": " + string(body_resp)
				msg := tgbotapi.NewMessage(chatID, reply)
				bot.Send(msg)
			case "Relay ON COFFEE":
				req, err := http.NewRequest("GET", "http://192.168.0.12/get?cmd=relay%20on", nil)		////////////// GET
				if err != nil {
					reply = "Request FAIL"
				}
				resp, err := http.DefaultClient.Do(req)
				if err != nil {
					reply = "Request FAIL"
				}		
				defer resp.Body.Close()
				body_resp, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					reply = "Request FAIL\n"
				}
				reply = update.CallbackQuery.From.UserName + ": " + string(body_resp)
				msg := tgbotapi.NewMessage(chatID, reply)
				bot.Send(msg)
			case "Relay OFF COFFEE":	
				req, err := http.NewRequest("GET", "http://192.168.0.12/get?cmd=relay%20off", nil)		////////////// GET
				if err != nil {
					reply = "Request FAIL"
				}
				resp, err := http.DefaultClient.Do(req)
				if err != nil {
					reply = "Request FAIL"
				}		
				defer resp.Body.Close()
				body_resp, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					reply = "Request FAIL\n"
				}
				reply = update.CallbackQuery.From.UserName + ": " + string(body_resp)
				msg := tgbotapi.NewMessage(chatID, reply)
				bot.Send(msg)
			case "Fun ON COFFEE":
				req, err := http.NewRequest("GET", "http://192.168.0.12/get?cmd=fun%20on", nil)		////////////// GET
				if err != nil {
					reply = "Request FAIL"
				}
				resp, err := http.DefaultClient.Do(req)
				if err != nil {
					reply = "Request FAIL"
				}		
				defer resp.Body.Close()
				body_resp, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					reply = "Request FAIL\n"
				}
				reply = update.CallbackQuery.From.UserName + ": " + string(body_resp)
				msg := tgbotapi.NewMessage(chatID, reply)
				bot.Send(msg)
			case "Fun OFF COFFEE":			
				req, err := http.NewRequest("GET", "http://192.168.0.12/get?cmd=fun%20off", nil)		////////////// GET
				if err != nil {
					reply = "Request FAIL"
				}
				resp, err := http.DefaultClient.Do(req)
				if err != nil {
					reply = "Request FAIL"
				}		
				defer resp.Body.Close()
				body_resp, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					reply = "Request FAIL\n"
				}
				reply = update.CallbackQuery.From.UserName + ": " + string(body_resp)
				msg := tgbotapi.NewMessage(chatID, reply)
				bot.Send(msg)
			case "Pump ON COFFEE":
				req, err := http.NewRequest("GET", "http://192.168.0.12/get?cmd=pump%20on", nil)		////////////// GET
				if err != nil {
					reply = "Request FAIL"
				}
				resp, err := http.DefaultClient.Do(req)
				if err != nil {
					reply = "Request FAIL"
				}		
				defer resp.Body.Close()
				body_resp, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					reply = "Request FAIL\n"
				}
				reply = update.CallbackQuery.From.UserName + ": " + string(body_resp)
				msg := tgbotapi.NewMessage(chatID, reply)
				bot.Send(msg)
			case "Pump OFF COFFEE":			
				req, err := http.NewRequest("GET", "http://192.168.0.12/get?cmd=pump%20off", nil)		////////////// GET
				if err != nil {
					reply = "Request FAIL"
				}
				resp, err := http.DefaultClient.Do(req)
				if err != nil {
					reply = "Request FAIL"
				}		
				defer resp.Body.Close()
				body_resp, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					reply = "Request FAIL\n"
				}
				reply = update.CallbackQuery.From.UserName + ": " + string(body_resp)
				msg := tgbotapi.NewMessage(chatID, reply)
				bot.Send(msg)
			case "Track ON COFFEE":
				req, err := http.NewRequest("GET", "http://192.168.0.12/get?cmd=track%20on", nil)		////////////// GET
				if err != nil {
					reply = "Request FAIL"
				}
				resp, err := http.DefaultClient.Do(req)
				if err != nil {
					reply = "Request FAIL"
				}		
				defer resp.Body.Close()
				body_resp, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					reply = "Request FAIL\n"
				}
				reply = update.CallbackQuery.From.UserName + ": " + string(body_resp)
				msg := tgbotapi.NewMessage(chatID, reply)
				bot.Send(msg)
			case "Track OFF COFFEE":			
				req, err := http.NewRequest("GET", "http://192.168.0.12/get?cmd=track%20off", nil)		////////////// GET
				if err != nil {
					reply = "Request FAIL"
				}
				resp, err := http.DefaultClient.Do(req)
				if err != nil {
					reply = "Request FAIL"
				}		
				defer resp.Body.Close()
				body_resp, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					reply = "Request FAIL\n"
				}
				reply = update.CallbackQuery.From.UserName + ": " + string(body_resp)
				msg := tgbotapi.NewMessage(chatID, reply)
				bot.Send(msg)
			/////////////////////////////////////////////////////////////////
			case "Red ON PICO":
				req, err := http.NewRequest("GET", "http://192.168.0.14/get?cmd=red%20on", nil)		////////////// GET
				if err != nil {
					reply = "Request FAIL"
				}
				resp, err := http.DefaultClient.Do(req)
				if err != nil {
					reply = "Request FAIL"
				}		
				defer resp.Body.Close()
				body_resp, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					reply = "Request FAIL\n"
				}
				reply = update.CallbackQuery.From.UserName + ": " + string(body_resp)
				msg := tgbotapi.NewMessage(chatID, reply)
				bot.Send(msg)
			case "Red OFF PICO":			
				req, err := http.NewRequest("GET", "http://192.168.0.14/get?cmd=red%20off", nil)		////////////// GET
				if err != nil {
					reply = "Request FAIL"
				}
				resp, err := http.DefaultClient.Do(req)
				if err != nil {
					reply = "Request FAIL"
				}		
				defer resp.Body.Close()
				body_resp, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					reply = "Request FAIL\n"
				}
				reply = update.CallbackQuery.From.UserName + ": " + string(body_resp)
				msg := tgbotapi.NewMessage(chatID, reply)
				bot.Send(msg)
			case "Yellow ON PICO":
				req, err := http.NewRequest("GET", "http://192.168.0.14/get?cmd=yellow%20on", nil)		////////////// GET
				if err != nil {
					reply = "Request FAIL"
				}
				resp, err := http.DefaultClient.Do(req)
				if err != nil {
					reply = "Request FAIL"
				}		
				defer resp.Body.Close()
				body_resp, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					reply = "Request FAIL\n"
				}
				reply = update.CallbackQuery.From.UserName + ": " + string(body_resp)
				msg := tgbotapi.NewMessage(chatID, reply)
				bot.Send(msg)
			case "Yellow OFF PICO":			
				req, err := http.NewRequest("GET", "http://192.168.0.14/get?cmd=yellow%20off", nil)		////////////// GET
				if err != nil {
					reply = "Request FAIL"
				}
				resp, err := http.DefaultClient.Do(req)
				if err != nil {
					reply = "Request FAIL"
				}		
				defer resp.Body.Close()
				body_resp, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					reply = "Request FAIL\n"
				}
				reply = update.CallbackQuery.From.UserName + ": " + string(body_resp)
				msg := tgbotapi.NewMessage(chatID, reply)
				bot.Send(msg)
			case "Green ON PICO":
				req, err := http.NewRequest("GET", "http://192.168.0.14/get?cmd=green%20on", nil)		////////////// GET
				if err != nil {
					reply = "Request FAIL"
				}
				resp, err := http.DefaultClient.Do(req)
				if err != nil {
					reply = "Request FAIL"
				}		
				defer resp.Body.Close()
				body_resp, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					reply = "Request FAIL\n"
				}
				reply = update.CallbackQuery.From.UserName + ": " + string(body_resp)
				msg := tgbotapi.NewMessage(chatID, reply)
				bot.Send(msg)
			case "Green OFF PICO":			
				req, err := http.NewRequest("GET", "http://192.168.0.14/get?cmd=green%20off", nil)		////////////// GET
				if err != nil {
					reply = "Request FAIL"
				}
				resp, err := http.DefaultClient.Do(req)
				if err != nil {
					reply = "Request FAIL"
				}		
				defer resp.Body.Close()
				body_resp, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					reply = "Request FAIL\n"
				}
				reply = update.CallbackQuery.From.UserName + ": " + string(body_resp)
				msg := tgbotapi.NewMessage(chatID, reply)
				bot.Send(msg)
			case "Led ON PICO":
				req, err := http.NewRequest("GET", "http://192.168.0.14/get?cmd=led%20on", nil)		////////////// GET
				if err != nil {
					reply = "Request FAIL"
				}
				resp, err := http.DefaultClient.Do(req)
				if err != nil {
					reply = "Request FAIL"
				}		
				defer resp.Body.Close()
				body_resp, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					reply = "Request FAIL\n"
				}
				reply = update.CallbackQuery.From.UserName + ": " + string(body_resp)
				msg := tgbotapi.NewMessage(chatID, reply)
				bot.Send(msg)
			case "Led OFF PICO":			
				req, err := http.NewRequest("GET", "http://192.168.0.14/get?cmd=led%20off", nil)		////////////// GET
				if err != nil {
					reply = "Request FAIL"
				}
				resp, err := http.DefaultClient.Do(req)
				if err != nil {
					reply = "Request FAIL"
				}		
				defer resp.Body.Close()
				body_resp, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					reply = "Request FAIL\n"
				}
				reply = update.CallbackQuery.From.UserName + ": " + string(body_resp)
				msg := tgbotapi.NewMessage(chatID, reply)
				bot.Send(msg)
			case "Laser ON PICO":
				req, err := http.NewRequest("GET", "http://192.168.0.14/get?cmd=laser%20on", nil)		////////////// GET
				if err != nil {
					reply = "Request FAIL"
				}
				resp, err := http.DefaultClient.Do(req)
				if err != nil {
					reply = "Request FAIL"
				}		
				defer resp.Body.Close()
				body_resp, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					reply = "Request FAIL\n"
				}
				reply = update.CallbackQuery.From.UserName + ": " + string(body_resp)
				msg := tgbotapi.NewMessage(chatID, reply)
				bot.Send(msg)
			case "Laser OFF PICO":			
				req, err := http.NewRequest("GET", "http://192.168.0.14/get?cmd=laser%20off", nil)		////////////// GET
				if err != nil {
					reply = "Request FAIL"
				}
				resp, err := http.DefaultClient.Do(req)
				if err != nil {
					reply = "Request FAIL"
				}		
				defer resp.Body.Close()
				body_resp, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					reply = "Request FAIL\n"
				}
				reply = update.CallbackQuery.From.UserName + ": " + string(body_resp)
				msg := tgbotapi.NewMessage(chatID, reply)
				bot.Send(msg)
			}
    		continue
		}
		if update.Message == nil {
			continue
		}
		
		text := update.Message.Text      			// Текст сообщения
		//chatID := update.Message.Chat.ID 			//  ID чата
		//userID := update.Message.From.ID 			// ID пользователя
		userName := update.Message.From.UserName	// Имя пользователя

		// логируем от кого какое сообщение пришло
		log.Printf("[%s] %s", userName, text)

		// свитч на обработку комманд
		// комманда - сообщение, начинающееся с "/"
		switch update.Message.Command() {
		case "start":
			reply := "Выбери робота в меню или:"
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)
			msg.ReplyMarkup = numericKeyboardSTART
			bot.Send(msg)
			reply = update.Message.From.UserName + ": открыл бот"
			msg = tgbotapi.NewMessage(chatID, reply)
			bot.Send(msg)
			continue
		case "coffee":
			reply := "Выбери действие COFFEE:"
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)
			msg.ReplyMarkup = numericKeyboardCOFFEE
			bot.Send(msg)
			reply = update.Message.From.UserName + ": открыл COFFEE"
			msg = tgbotapi.NewMessage(chatID, reply)
			bot.Send(msg)
			continue
		case "pico":
			reply := "Выбери действие PICO:"
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)
			msg.ReplyMarkup = numericKeyboardPICO
			bot.Send(msg)
			reply = update.Message.From.UserName + ": открыл PICO"
			msg = tgbotapi.NewMessage(chatID, reply)
			bot.Send(msg)
			continue
		case "statec":
			req, err := http.NewRequest("GET", "http://192.168.0.12/get?cmd=state", nil)		////////////// GET
			if err != nil {
				reply = "Request FAIL"
			}
			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				reply = "Request FAIL"
			}		
			defer resp.Body.Close()
			body_resp, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				reply = "Request FAIL\n"
			}
			reply = update.Message.From.UserName + ": запросил состояние"
			msg := tgbotapi.NewMessage(chatID, reply)
			bot.Send(msg)
			reply = string(body_resp)
		case "statep":
			req, err := http.NewRequest("GET", "http://192.168.0.14/get?cmd=state", nil)		////////////// GET
			if err != nil {
				reply = "Request FAIL"
			}
			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				reply = "Request FAIL"
			}		
			defer resp.Body.Close()
			body_resp, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				reply = "Request FAIL\n"
			}
			reply = update.Message.From.UserName + ": запросил состояние"
			msg := tgbotapi.NewMessage(chatID, reply)
			bot.Send(msg)
			reply = string(body_resp)
		case "redon":
			req, err := http.NewRequest("GET", "http://192.168.0.12/get?cmd=red%20on", nil)		////////////// GET
			if err != nil {
				reply = "Request FAIL"
			}
			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				reply = "Request FAIL"
			}		
			defer resp.Body.Close()
			body_resp, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				reply = "Request FAIL\n"
			}
			reply = string(body_resp)
		case "redoff":
			req, err := http.NewRequest("GET", "http://192.168.0.12/get?cmd=red%20off", nil)		////////////// GET
			if err != nil {
				reply = "Request FAIL"
			}
			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				reply = "Request FAIL"
			}		
			defer resp.Body.Close()
			body_resp, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				reply = "Request FAIL\n"
			}
			reply = string(body_resp)
		case "relayon":
			req, err := http.NewRequest("GET", "http://192.168.0.12/get?cmd=relay%20on", nil)		////////////// GET
			if err != nil {
				reply = "Request FAIL"
			}
			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				reply = "Request FAIL"
			}		
			defer resp.Body.Close()
			body_resp, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				reply = "Request FAIL\n"
			}
			reply = string(body_resp)
		case "relayoff":
			req, err := http.NewRequest("GET", "http://192.168.0.12/get?cmd=relay%20off", nil)		////////////// GET
			if err != nil {
				reply = "Request FAIL"
			}
			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				reply = "Request FAIL"
			}		
			defer resp.Body.Close()
			body_resp, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				reply = "Request FAIL\n"
			}
			reply = string(body_resp)
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)
		bot.Send(msg)
	}
}


// Роутер
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("%v %v", r.Method, r.URL.Path)
	
	if strings.Trim(r.URL.Path, "/") == "" {
		tgBotChan <- strings.Split(r.RemoteAddr, ":")[0] + ": зашел на сайт"
		index_page(w, r)
		return
	}
	// API from HTTP
	if strings.Trim(r.URL.Path, "/") == "api" {
		http_pars(w, r)
		return
	}
	// API from JSON 
	if strings.Trim(r.URL.Path, "/") == "json" {
		json_pars(w, r)
		return
	}
	// Long poll
	if strings.Trim(r.URL.Path, "/") == "poll" {
		PollResponse(w, r)
		return
	}

	if strings.Trim(r.URL.Path, "/") == "savebyte" {
		savebyte(w, r)
		return
	}
	
	if strings.Trim(r.URL.Path, "/") == "saveutf8" {
		saveutf8(w, r)
		return
	}
	
	// serve static assets from 'static' dir:
	h.fileServer.ServeHTTP(w, r)
}
