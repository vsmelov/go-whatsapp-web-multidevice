package chat

// Request and Response structures for chat operations

type ListChatsRequest struct {
	Limit    int    `json:"limit" query:"limit"`
	Offset   int    `json:"offset" query:"offset"`
	Search   string `json:"search" query:"search"`
	HasMedia bool   `json:"has_media" query:"has_media"`
	Archived *bool  `json:"archived" query:"archived"`
}

type ListChatsResponse struct {
	Data       []ChatInfo         `json:"data"`
	Pagination PaginationResponse `json:"pagination"`
}

type GetChatMessagesRequest struct {
	ChatJID   string  `json:"chat_jid" uri:"chat_jid"`
	Limit     int     `json:"limit" query:"limit"`
	Offset    int     `json:"offset" query:"offset"`
	StartTime *string `json:"start_time" query:"start_time"`
	EndTime   *string `json:"end_time" query:"end_time"`
	MediaOnly bool    `json:"media_only" query:"media_only"`
	IsFromMe  *bool   `json:"is_from_me" query:"is_from_me"`
	Search    string  `json:"search" query:"search"`
}

type GetChatMessagesResponse struct {
	Data       []MessageInfo      `json:"data"`
	Pagination PaginationResponse `json:"pagination"`
	ChatInfo   ChatInfo           `json:"chat_info"`
}

// Pin Chat operations
type PinChatRequest struct {
	ChatJID string `json:"chat_jid" uri:"chat_jid"`
	Pinned  bool   `json:"pinned"`
}

type PinChatResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	ChatJID string `json:"chat_jid"`
	Pinned  bool   `json:"pinned"`
}

type ChatInfo struct {
	JID                 string `json:"jid"`
	Name                string `json:"name"`
	LastMessageTime     string `json:"last_message_time"`
	EphemeralExpiration uint32 `json:"ephemeral_expiration"`
	CreatedAt           string `json:"created_at"`
	UpdatedAt           string `json:"updated_at"`
	Archived            bool   `json:"archived"`
}

type MessageInfo struct {
	ID                string         `json:"id"`
	ChatJID           string         `json:"chat_jid"`
	SenderJID         string         `json:"sender_jid"`
	SenderDisplayName string         `json:"sender_display_name"`
	Content           string         `json:"content"`
	Timestamp         string         `json:"timestamp"`
	IsFromMe          bool           `json:"is_from_me"`
	MediaType         string         `json:"media_type"`
	Reactions         []ReactionInfo `json:"reactions,omitempty"`
	// CallMetadata is JSON when media_type is "call" (incoming call log).
	CallMetadata string `json:"call_metadata,omitempty"`
	Filename     string `json:"filename"`
	URL          string `json:"url"`
	FileLength   uint64 `json:"file_length"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}

type PaginationResponse struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
	Total  int `json:"total"`
}

// Disappearing Messages operations
type SetDisappearingTimerRequest struct {
	ChatJID      string `json:"chat_jid" uri:"chat_jid"`
	TimerSeconds uint32 `json:"timer_seconds"`
}

type SetDisappearingTimerResponse struct {
	Status       string `json:"status"`
	Message      string `json:"message"`
	ChatJID      string `json:"chat_jid"`
	TimerSeconds uint32 `json:"timer_seconds"`
}

// Archive Chat operations
type ArchiveChatRequest struct {
	ChatJID  string `json:"chat_jid" uri:"chat_jid"`
	Archived bool   `json:"archived"`
}

type ArchiveChatResponse struct {
	Status   string `json:"status"`
	Message  string `json:"message"`
	ChatJID  string `json:"chat_jid"`
	Archived bool   `json:"archived"`
}

// Backfill (on-demand history sync) operations.
// Asks the primary device (phone) for up to Count messages that precede the
// anchor message. Nothing is sent to the chat itself: the request is a peer
// protocol message to the account's own devices. The answer arrives
// asynchronously as a HistorySync event of type ON_DEMAND and is written to
// chat storage by the history sync handler.
type BackfillChatRequest struct {
	ChatJID string `json:"chat_jid" uri:"chat_jid"`
	// Count is the number of messages to request (1..50, default 50).
	Count int `json:"count"`
	// Optional explicit anchor. When AnchorMessageID is empty the oldest stored
	// message of the chat is used.
	AnchorMessageID string `json:"anchor_message_id"`
	AnchorTimestamp string `json:"anchor_timestamp"` // RFC3339, required with AnchorMessageID
	AnchorIsFromMe  bool   `json:"anchor_is_from_me"`
	// UseLID addresses the chat by its @lid JID in the request (1:1 chats only).
	UseLID bool `json:"use_lid"`
}

type BackfillAnchor struct {
	MessageID string `json:"message_id"`
	Timestamp string `json:"timestamp"`
	IsFromMe  bool   `json:"is_from_me"`
	SenderJID string `json:"sender_jid"`
}

type BackfillChatResponse struct {
	Status           string         `json:"status"`
	Message          string         `json:"message"`
	ChatJID          string         `json:"chat_jid"`
	RequestJID       string         `json:"request_jid"`
	Count            int            `json:"count"`
	Anchor           BackfillAnchor `json:"anchor"`
	RequestMessageID string         `json:"request_message_id"`
	StoredMessages   int64          `json:"stored_messages"`
}
