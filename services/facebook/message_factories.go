package facebook

func HandleResponses(sender string, text string, postback string)  {
	Greetings(sender, postback)
}

func Greetings(sender string, postback string)  {
	message := "Bienvenido al bot"
	SendingText(sender, message)
}
