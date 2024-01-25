package models

import "time"

type CreateUserSubscriptionReq struct {
	Id             string
	UserId         string
	SubscriptionId string
	StartTime      time.Time
	EndTime        time.Time
	Active         bool
}
