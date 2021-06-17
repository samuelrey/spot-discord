package discord

import (
	"github.com/bwmarrin/discordgo"
	"github.com/samuelrey/spot-discord/framework"
)

// NewContext takes discord specific data and produces a Context which can be
// used by commands.
func NewContext(
	session *discordgo.Session,
	channelID string,
	enrolledUsers *[]framework.User,
	actor *discordgo.User,
) *framework.Context {
	ctx := new(framework.Context)
	discordMessager := DiscordMessager{
		Session:   session,
		ChannelID: channelID,
	}
	ctx.Messager = discordMessager
	ctx.EnrolledUsers = enrolledUsers
	ctx.Actor = framework.User{
		ID:       actor.ID,
		Username: actor.Username,
	}
	return ctx
}

type DiscordMessager struct {
	Session   *discordgo.Session
	ChannelID string
}

func (d DiscordMessager) Reply(content string) error {
	_, err := d.Session.ChannelMessageSend(d.ChannelID, content)
	return err
}

func (d DiscordMessager) DirectMessage(recipient, content string) error {
	channel, err := d.Session.UserChannelCreate(recipient)
	if err != nil {
		return err
	}

	_, err = d.Session.ChannelMessageSend(channel.ID, content)
	return err
}
