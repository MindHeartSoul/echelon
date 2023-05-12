package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/echelonfoundation/echelon/v3/x/vrf/types"
)

func (k msgServer) BroadcastRandom(goCtx context.Context, msg *types.MsgBroadcastRandom) (*types.MsgBroadcastRandomResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := k.BroadcastRandomNumber(ctx, msg)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgBroadcastRandomResponse{}, err
}
