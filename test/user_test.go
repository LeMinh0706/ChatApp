package test

import (
	"context"
	"testing"

	"github.com/LeMinh0706/ChatApp/db"
	"github.com/LeMinh0706/ChatApp/util"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) db.User {
	hashPassword, err := util.HashPashword("kocanpass")
	require.NoError(t, err)
	arg := db.CreateUserParams{
		Username:  util.RandomString(6),
		Password:  hashPassword,
		UrlAvatar: util.RandomURL(),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, arg.Username)
	require.Equal(t, arg.UrlAvatar, arg.UrlAvatar)

	return user
}
func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}
