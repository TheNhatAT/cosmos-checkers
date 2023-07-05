package keeper_test

import (
	"github.com/alice/checkers/x/checkers/testutil"
	"github.com/alice/checkers/x/checkers/types"
	"github.com/stretchr/testify/require"
	"testing"
)

const (
	alice = testutil.Alice
	bob   = testutil.Bob
	carol = testutil.Carol
)

func TestCreateGame(t *testing.T) {
	msgServer, context := setupMsgServer(t)
	createResponse, err := msgServer.CreateGame(context, &types.MsgCreateGame{
		Creator: alice,
		Black:   bob,
		Red:     carol,
	})
	require.Nil(t, err)
	require.EqualValues(t, types.MsgCreateGameResponse{
		GameIndex: "", // TODO: update with a proper value when updated
	}, *createResponse)
}
