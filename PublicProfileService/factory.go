package PublicProfileService

import (
	"time"

	"github.com/bluele/gcache"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func NewClient(ahost, aport string) Client {

	c := &pubprofileclient{
		host:          ahost,
		port:          aport,
		bot_chatID:    -1001469468779,
		bot_token:     "1108341214:AAEKNbFf6PO7Y6UJGK-xepDDOGKlBU2QVCg",
		botClient:     nil,
		cache:         gcache.New(1024 * 1024).LRU().Build(),
		cacheExperied: time.Duration(300) * time.Minute,
	}
	bot, err := tgbotapi.NewBotAPI(c.bot_token)
	if err == nil {
		c.botClient = bot
	}
	return c

}

func NewClientWithCache(ahost, aport string, cacheSize int64, experiedTimeMinue int64) Client {
	c := &pubprofileclient{
		host:          ahost,
		port:          aport,
		bot_chatID:    -1001469468779,
		bot_token:     "1108341214:AAEKNbFf6PO7Y6UJGK-xepDDOGKlBU2QVCg",
		botClient:     nil,
		cache:         gcache.New(int(cacheSize) * 2).LRU().Build(),
		cacheExperied: time.Duration(experiedTimeMinue) * time.Minute,
	}
	bot, err := tgbotapi.NewBotAPI(c.bot_token)
	if err == nil {
		c.botClient = bot
	}
	return c

}
