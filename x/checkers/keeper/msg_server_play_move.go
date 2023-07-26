package keeper

import (
	"context"
	"github.com/alice/checkers/x/checkers/rules"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"strconv"

	"github.com/alice/checkers/x/checkers/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) PlayMove(goCtx context.Context, msg *types.MsgPlayMove) (*types.MsgPlayMoveResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	storedGame, found := k.Keeper.GetStoredGame(ctx, msg.GameIndex)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrGameNotFound, "%s", msg.GameIndex)
	}

	if storedGame.Winner != rules.PieceStrings[rules.NO_PLAYER] {
		return nil, sdkerrors.Wrapf(types.ErrGameFinished, "%s", msg.GameIndex)
	}

	isBlack := storedGame.Black == msg.Creator
	isRed := storedGame.Red == msg.Creator

	var player rules.Player

	if !isBlack && !isRed {
		return nil, sdkerrors.Wrapf(types.ErrPlayerNotInGame, "%s", msg.Creator)
	} else if isBlack && isRed {
		player = rules.StringPieces[storedGame.Turn].Player
	} else if isBlack {
		player = rules.BLACK_PLAYER
	} else {
		player = rules.RED_PLAYER
	}

	game, err := storedGame.ParseGame()
	if err != nil {
		panic(err.Error())
	}

	if !game.TurnIs(player) {
		return nil, sdkerrors.Wrapf(types.ErrNotPlayerTurn, "%s", player)
	}

	err = k.Keeper.CollectWager(ctx, &storedGame)
	if err != nil {
		return nil, err
	}

	captured, moveErr := game.Move(
		rules.Pos{
			X: int(msg.FromX),
			Y: int(msg.FromY),
		},
		rules.Pos{
			X: int(msg.ToX),
			Y: int(msg.ToY),
		},
	)
	if moveErr != nil {
		return nil, sdkerrors.Wrapf(types.ErrWrongMove, moveErr.Error())
	}

	storedGame.Winner = rules.PieceStrings[game.Winner()]
	lastBoard := game.String()
	systemInfo, found := k.Keeper.GetSystemInfo(ctx)
	if !found {
		panic("system info not found")
	}

	if storedGame.Winner == rules.PieceStrings[rules.NO_PLAYER] {
		storedGame.Board = lastBoard
		k.Keeper.SendToFifoTail(ctx, &storedGame, &systemInfo)
	} else {
		storedGame.Board = ""
		k.Keeper.RemoveFromFifo(ctx, &storedGame, &systemInfo)
		k.Keeper.MustPayWinnings(ctx, &storedGame)

		// register winner and loser
		k.Keeper.MustRegisterPlayerWin(ctx, &storedGame)
	}

	storedGame.Deadline = types.FormatDeadline(types.GetNextDeadline(ctx)) // not emitted event for this
	storedGame.MoveCount++                                                 // not emitted event for this
	storedGame.Turn = rules.PieceStrings[game.Turn]
	k.Keeper.SetStoredGame(ctx, storedGame)
	k.Keeper.SetSystemInfo(ctx, systemInfo)
	ctx.GasMeter().ConsumeGas(types.PlayMoveGas, "Play a move")

	// emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.MovePlayedEventType,
			sdk.NewAttribute(types.MovePlayedEventCreator, msg.Creator),
			sdk.NewAttribute(types.MovePlayedEventGameIndex, msg.GameIndex),
			sdk.NewAttribute(types.MovePlayedCaptureX, strconv.FormatInt(int64(captured.X), 10)),
			sdk.NewAttribute(types.MovePlayedCaptureY, strconv.FormatInt(int64(captured.Y), 10)),
			sdk.NewAttribute(types.MovePlayedWinner, rules.PieceStrings[game.Winner()]),
			sdk.NewAttribute(types.MovePlayedEventBoard, lastBoard),
		),
	)

	return &types.MsgPlayMoveResponse{
		CapturedX: int32(captured.X),
		CapturedY: int32(captured.Y),
		Winner:    rules.PieceStrings[game.Winner()],
	}, nil
}
