package model

type Subscription struct {
	UserID               string   `json:"user_id" bson:"user_id"`
	Topics               []string `json:"topics" bson:"topics"`
	NotificationChannels struct {
		Email            string `json:"email" bson:"email"`
		SMS              string `json:"sms" bson:"sms"`
		PushNotification bool   `json:"push_notifications" bson:"push_notifications"`
	} `json:"notification_channels" bson:"notification_channels"`
}

type Notification struct {
	Topic string `json:"topic" bson:"topic"`
	Event struct {
		EventID   string `json:"event_id" bson:"event_id"`
		Timestamp string `json:"timestamp" bson:"timestamp"`
		Details   struct {
			UserID   string `json:"user_id" bson:"user_id"`
			Email    string `json:"email" bson:"email"`
			Username string `json:"username" bson:"username"`
		} `json:"details" bson:"details"`
	} `json:"event" bson:"event"`
	Message struct {
		Title string `json:"title" bson:"title"`
		Body  string `json:"body" bson:"body"`
	} `json:"message" bson:"message"`
}
