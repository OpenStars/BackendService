package SimpleSessionService

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func NewClient(etcdServers []string, sid, defaultHost, defaultPort string) Client {

	sessionsv := &simpleSessionClient{
		host: defaultHost,
		port: defaultPort,
		sid:  sid,
		// etcdManager: EndpointsManager.GetEtcdBackendEndpointManagerSingleton(etcdServers),
		bot_chatID: -1001469468779,
		bot_token:  "1108341214:AAEKNbFf6PO7Y6UJGK-xepDDOGKlBU2QVCg",
		botClient:  nil,
	}
	bot, err := tgbotapi.NewBotAPI(sessionsv.bot_token)
	if err == nil {
		sessionsv.botClient = bot
	}
	// if sessionsv.etcdManager == nil {
	// 	return sessionsv
	// }
	// err = sessionsv.etcdManager.SetDefaultEntpoint(sid, defaultHost, defaultPort)
	// if err != nil {
	// 	log.Println("SetDefaultEndpoint sid", sid, "err", err)
	// 	return nil
	// }

	return sessionsv

}
