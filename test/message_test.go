package test

import (
	"context"
	"testing"

	"github.com/LeMinh0706/ChatApp/db"
	"github.com/LeMinh0706/ChatApp/util"
	"github.com/stretchr/testify/require"
)

func createRandomMessage(t *testing.T, user_id_1, user_id_2 int64) db.Message {
	arg := db.CreateMessageParams{
		FromUser: user_id_1,
		ToUser:   user_id_2,
		Content:  util.RandomDescription(),
	}

	message, err := testQueries.CreateMessage(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, message)

	require.Equal(t, user_id_1, message.FromUser)
	require.Equal(t, user_id_2, message.ToUser)
	require.Equal(t, arg.Content, message.Content)
	return message
}
func TestCreateMessage(t *testing.T) {
	user1 := createRandomUser(t)
	user2 := createRandomUser(t)
	createRandomMessage(t, user1.ID, user2.ID)
}

func TestGetMessage(t *testing.T) {
	user1 := createRandomUser(t)
	user2 := createRandomUser(t)
	for i := 0; i < 10; i++ {
		createRandomMessage(t, user1.ID, user2.ID)
		createRandomMessage(t, user2.ID, user1.ID)
	}
	messages, err := testQueries.GetMessages(context.Background(), db.GetMessagesParams{
		FromUser: user1.ID,
		ToUser:   user2.ID,
	})
	require.NoError(t, err)
	require.NotEmpty(t, messages)
	require.Len(t, messages, 20)

}
