package facebook

import (
	"bot/app/utils"
)

func HandleResponses(sender string, text string, postback string)  {
	Greetings(sender, postback)
}

func Greetings(sender string, postback string)  {
	SendingText(sender, messages.GREETINGS_MESSAGE)
}
