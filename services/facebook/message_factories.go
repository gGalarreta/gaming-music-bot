package facebook

import (
	"bot/app/utils"
	"fmt"
)

func HandleResponses(sender string, text string, postback string)  {
	Greetings(sender)
	TimeOptions(sender)
}

func Greetings(sender string)  {
	SendingText(sender, messages.GREETINGS_MESSAGE)
	SendingText(sender, messages.BOT_GOAL_MESSAGE)
}

func TimeOptions(sender string)  {
	thirty := fmt.Sprintf(`{
		"type":"postback",
		"payload": "%s",
		"title": "%s"
	  }`, messages.THIRTY_PAYLOAD, messages.THIRTY_TITLE)
	one_hour := fmt.Sprintf(`{
		"type":"postback",
		"payload": "%s",
		"title": "%s"
		}`, messages.ONE_HOUR_PAYLOAD, messages.ONE_HOUR_TITLE)
	two_hours := fmt.Sprintf(`{
		"type":"postback",
		"payload": "%s",
		"title": "%s"
		}`, messages.TWO_HOURS_PAYLOAD, messages.TWO_HOURS_TITLE)
	elements := thirty + "," + one_hour + "," + two_hours
	SendingButtons(sender, messages.ASKTIME_MESSAGE, elements)
}

