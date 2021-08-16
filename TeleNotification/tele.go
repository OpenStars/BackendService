package telenotification

import (
	"log"
	"net"
	"os"
	"path/filepath"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// default config
var (
	id                = int64(-1001469468779)
	token             = "1108341214:AAEKNbFf6PO7Y6UJGK-xepDDOGKlBU2QVCg"
	bot               *tgbotapi.BotAPI
	msgChan           = make(chan string, 1000)
	callerServiceName = os.Args[0]
	ipCaller          = getLocalIP()
)

func getLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}
func init() {
	InitBotAgent(0, "", 1000)
}
func InitBotAgent(chatID int64, chatToken string, chanLength int) error {
	if chatID != 0 {
		id = chatID
	}
	if chatToken != "" {
		token = chatToken
	}
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err == nil {
		callerServiceName = dir + callerServiceName
	}
	log.Println("[INIT TELEBOT] id", id, "token", token, "chan", chanLength)
	bot, err = tgbotapi.NewBotAPI(token)
	if err != nil {
		return err
	}
	msgChan = make(chan string, chanLength)
	go worker()
	return nil
}

func Notify(msg string) {
	select {
	case msgChan <- msg:
	default:
		log.Println("[TELE ERROR] full chan")
	}
}

func NotifyServiceError(sid, host, port string, err error) {

	msg := "[TELE SERVICE ERROR] caller file " + callerServiceName + " caller ip " + ipCaller + " to service sid " + sid + " address " + host + ":" + port
	if err != nil {
		msg = msg + " err " + err.Error()
	}
	Notify(msg)
}

func worker() {
	for msg := range msgChan {
		if bot != nil {
			msgTele := tgbotapi.NewMessage(id, msg)
			_, err := bot.Send(msgTele)
			if err != nil {
				log.Println("[SUCCESS] send tele err", err)
			}
		}
	}
}
