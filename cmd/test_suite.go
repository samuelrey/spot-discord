package cmd

import (
	"github.com/samuelrey/spot-the-bot/framework"
	"github.com/samuelrey/spot-the-bot/message"
	"github.com/stretchr/testify/suite"
)

type CommandTestSuite struct {
	suite.Suite
	Actor           message.User
	Ctx             CommandContext
	UserQueue       framework.UserQueue
	Messager        framework.MockMessager
	PlaylistCreator framework.MockPlaylistCreator
}

func (suite *CommandTestSuite) SetupTest() {
	suite.Actor = message.User{ID: "amethyst#4422", Username: "amethyst"}
	suite.Messager = framework.MockMessager{}
	suite.PlaylistCreator = framework.MockPlaylistCreator{}

	suite.UserQueue = framework.NewUserQueue([]message.User{})

	suite.Ctx = CommandContext{
		Messager:        &suite.Messager,
		PlaylistCreator: &suite.PlaylistCreator,
		PlaylistName:    "Einstok",
		UserQueue:       &suite.UserQueue,
		Actor:           suite.Actor,
	}
}
