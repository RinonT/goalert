package message

import (
	"time"

	"github.com/target/goalert/notification"
)

// Message represents the data for an outgoing message.
type Message struct {
	ID         string
	Type       notification.MessageType
	Dest       notification.Dest
	AlertID    int
	AlertLogID int
	VerifyID   string

	UserID    string
	ServiceID string
	CreatedAt time.Time
	SentAt    time.Time

	ScheduleID string

	StatusAlertIDs []int
}
