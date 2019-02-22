package facebook

func HandleResponses(sender string, text string, postback string)  {
	Greetings(sender, text, postback)
}

func Greetings(sender string, text string, postback string)  {
	message := "Bienvenido al bot"
	SendingText(sender, message)
}
