package entities

// Conversation represents one persisted conversation.
type Conversation struct {
	ID string
}

// Message represents one persisted message and its position in a conversation.
type Message struct {
	ID             string
	ConversationID string
	Author         string
	Content        string
	State          string
	Position       int
}
