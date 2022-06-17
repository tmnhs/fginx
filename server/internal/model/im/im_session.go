package model

type SessionModel struct {
	Id          int64 `json:"id" `
	UserId      int64 `json:"user_id" gorm:"column:user_id"`
	PeerId      int64 `json:"peer_id" gorm:"column:peer_id"`
	SessionType int   `json:"session_type" gorm:"column:session_type"`
	//SessionStatus  int `json:"session_status" gorm:"column:session_status"`
	IsRobotSession int `json:"is_robot_session" gorm:"column:is_robot_session"`

	CreatedTime int64 `json:"created_time" gorm:"column:created_time"`
	UpdatedTime int64 `json:"updated_time" gorm:"column:updated_time"`
}
