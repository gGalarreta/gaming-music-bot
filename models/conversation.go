package models

type Conversation struct {
    recipient_id string
	sender_id string
	status int
}

func NewConversation(recipient_id string, sender_id string, status int) *Conversation {
	return &Conversation{recipient_id, sender_id, status}
}