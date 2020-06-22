package main

import (
	"fmt"

	"github.com/Trey2k/gostocks-discord/mysql"
	"github.com/Trey2k/gostocks-discord/utils"
	"github.com/bwmarrin/discordgo"
	"github.com/pkg/errors"
)

func chatListener(discord *discordgo.Session, message *discordgo.MessageCreate) {
	serverID := message.GuildID
	channelID := message.ChannelID

	for i := 0; i < len(utils.Config.Discord.GuildIDs); i++ {
		if serverID == utils.Config.Discord.GuildIDs[i] && channelID == utils.Config.Discord.ChannelIDs[i] {
			order := ChatParse(message.Content, *message.Author, message.ID)
			ordersChannel <- order
		}
	}
}

func editListener(discord *discordgo.Session, message *discordgo.MessageUpdate) {
	hasFailed, fail, err := mysql.HasFailed(message.ID)
	if err != nil {
		fmt.Println("Error querying db: " + errors.WithStack(err).Error())
	}
	if hasFailed {
		if fail.FailCode != 101 {
			err = mysql.DeleteFail(fail.ID)
			if err != nil {
				fmt.Println("Error querying db: " + errors.WithStack(err).Error())
			}
			order := ChatParse(message.Content, *message.Author, message.ID)
			ordersChannel <- order
		}
	}
}

func discordStatus(discord *discordgo.Session, ready *discordgo.Ready) {
	fmt.Println("Started discord client.")
	_, err := discord.UserUpdateStatus("invisible")
	utils.ErrCheck("Error connecting to discord: ", err)
}
