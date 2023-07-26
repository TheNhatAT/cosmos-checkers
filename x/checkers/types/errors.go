package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrInvalidBlack            = sdkerrors.Register(ModuleName, 1100, "black address is invalid: %s")
	ErrInvalidRed              = sdkerrors.Register(ModuleName, 1101, "red address is invalid: %s")
	ErrGameNotParseable        = sdkerrors.Register(ModuleName, 1102, "game cannot be parsed")
	ErrInvalidGameIndex        = sdkerrors.Register(ModuleName, 1103, "game index is invalid")
	ErrInvalidPositionIndex    = sdkerrors.Register(ModuleName, 1104, "position index is invalid")
	ErrMoveAbsent              = sdkerrors.Register(ModuleName, 1105, "there is no move")
	ErrGameNotFound            = sdkerrors.Register(ModuleName, 1106, "game by id not found")
	ErrCreatorNotPlayer        = sdkerrors.Register(ModuleName, 1107, "message creator is not a player")
	ErrNotPlayerTurn           = sdkerrors.Register(ModuleName, 1108, "player tried to play out of turn")
	ErrWrongMove               = sdkerrors.Register(ModuleName, 1109, "wrong move")
	ErrPlayerNotInGame         = sdkerrors.Register(ModuleName, 1110, "player is not in game")
	ErrGameFinished            = sdkerrors.Register(ModuleName, 1111, "game is already finished")
	ErrInvalidDeadline         = sdkerrors.Register(ModuleName, 1112, "deadline cannot be parsed: %s")
	ErrCannotFindWinnerByColor = sdkerrors.Register(ModuleName, 1113, "cannot find winner by color: %s")
	ErrBlackCannotPay          = sdkerrors.Register(ModuleName, 1114, "black cannot pay")
	ErrRedCannotPay            = sdkerrors.Register(ModuleName, 1115, "red cannot pay")
	ErrNothingToPay            = sdkerrors.Register(ModuleName, 1116, "nothing to pay")
	ErrCannotRefundWager       = sdkerrors.Register(ModuleName, 1117, "cannot refund wager")
	ErrCannotPayWinnings       = sdkerrors.Register(ModuleName, 1118, "cannot pay winnings")
	ErrNotInRefundState        = sdkerrors.Register(ModuleName, 1119, "game is not in refund state")
	ErrWinnerNotParseable      = sdkerrors.Register(ModuleName, 1120, "winner cannot be parsed: %s")
	ErrThereIsNoWinner         = sdkerrors.Register(ModuleName, 1121, "there is no winner")
	ErrInvalidDateAdded        = sdkerrors.Register(ModuleName, 1122, "date added is invalid")
	ErrCannotAddToLeaderboard  = sdkerrors.Register(ModuleName, 1123, "cannot add to leaderboard")
)
