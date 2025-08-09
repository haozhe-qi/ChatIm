package sdk

const (
	MsgTypeText = "text"
)

type Chat struct {
	Nick      string
	UserID    string
	SessionID string
	conn      *connect
}

type Message struct {
	Type       string
	Name       string
	FromUserID string
	ToUserID   string
	Content    string
	Session    string
}

func NewChat(serverAddr, nick, userID, sessionID string) *Chat {
	return &Chat{
		Nick:      nick,
		UserID:    userID,
		SessionID: sessionID,
		conn:      newConnect(serverAddr),
	}
}

func (chat *Chat) Send(msg *Message) {
	chat.conn.send(msg)
}
func (chat *Chat) Close() {
	chat.conn.close()
}
func (chat *Chat) Recv() <-chan *Message {
	return chat.conn.recv()
}
