package models

// Conversation contains the ordered messages exchanged by a user and an agent.
type Conversation struct {
	Messages []Message
}

// Message represents one message in a conversation.
type Message struct {
	ID      string
	Author  MessageAuthor
	Content string
	State   MessageState
}

// MessageAuthor identifies who produced a message.
type MessageAuthor string

const (
	MessageAuthorUser  MessageAuthor = "user"
	MessageAuthorAgent MessageAuthor = "agent"
)

// MessageState describes the delivery lifecycle of a message.
type MessageState string

const (
	MessageStatePending   MessageState = "pending"
	MessageStateStreaming MessageState = "streaming"
	MessageStateCompleted MessageState = "completed"
	MessageStateCancelled MessageState = "cancelled"
	MessageStateFailed    MessageState = "failed"
)
