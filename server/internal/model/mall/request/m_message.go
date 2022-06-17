package request

type Message struct {
	MessageType int   `json:"messageType" form:"messageType"`
	FromUid     int64 `json:"from_uid" form:"from_uid"`
	ToUid       int64 `json:"to_uid" form:"to_uid"`
}
