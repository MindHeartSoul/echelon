package cli

import (
	"strconv"
	"encoding/hex"

	"github.com/spf13/cobra"
	"github.com/coniks-sys/coniks-go/crypto/vrf"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/echelonfoundation/echelon/v3/x/vrf/types"
)

var _ = strconv.Itoa(0)

func CmdCreateRandom() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-random [multiplier]",
		Short: "create-random",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argMultiplier, erri := strconv.ParseUint(args[0], 10, 64)
			if erri != nil {
				return erri
			}
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			// directly generate vrv + proof + pubkey
			sender := clientCtx.GetFromAddress().String()
//			userval, isFound := k.GetUserval(clientCtx, sender)

//			var user_key_count int64 = 1
//			if isFound {
//				user_key_count = userval.Count + 1
//			}

			sk, err := vrf.GenerateKey(nil)
			if err != nil {
				return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Secret Key is not generated")
			}

			random_val_key := sender //+ "," + strconv.FormatInt(user_key_count, 10)
			a_message := []byte(random_val_key)

			vrv, proof := sk.Prove(a_message) // Generate vrv (verifiable random value) and proof
			pub_key, ok_bool := sk.Public()   // public key creation

			if ok_bool == false {
				return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Public Key is not generated")
			}

			// continue msg generate
			msg := types.NewMsgCreateRandom(
				clientCtx.GetFromAddress().String(),
				argMultiplier,
				hex.EncodeToString(vrv),
				hex.EncodeToString(proof),
				hex.EncodeToString(pub_key),
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
