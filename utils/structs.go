package utils

import (
	"time"

	"github.com/bwmarrin/discordgo"
)

//OrderStruct order struct
type OrderStruct struct {
	Buy          bool
	Risky        bool
	Ticker       string
	ExpDate      time.Time
	StrikPrice   float64
	ContractType string
	Price        float64
	StopLoss     float64
	Sender       discordgo.User
	MessageID    string
}
