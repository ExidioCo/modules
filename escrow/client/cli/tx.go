package cli

import (
	"bufio"
	"strconv"
	"time"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/spf13/cobra"

	"github.com/exidioco/modules/escrow/types"
)

func txCreate(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create [from] [to] [amount] [deadline]",
		Short: "Create an escrow",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			var (
				buffer = bufio.NewReader(cmd.InOrStdin())
				txb    = auth.NewTxBuilderFromCLI(buffer).WithTxEncoder(utils.GetTxEncoder(cdc))
				ctx    = context.NewCLIContextWithInput(buffer).WithCodec(cdc)
			)

			identity, err := strconv.ParseUint(args[0], 0, 64)
			if err != nil {
				return err
			}

			to, err := sdk.AccAddressFromBech32(args[1])
			if err != nil {
				return err
			}

			coins, err := sdk.ParseCoins(args[2])
			if err != nil {
				return err
			}

			deadline, err := time.Parse(time.RFC3339, args[3])
			if err != nil {
				return err
			}

			msg := types.NewMsgCreate(ctx.FromAddress, identity, to, coins, deadline)
			return utils.GenerateOrBroadcastMsgs(ctx, txb, []sdk.Msg{msg})
		},
	}

	return cmd
}

func txApprove(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "approve [identity]",
		Short: "Approve an escrow",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			var (
				buffer = bufio.NewReader(cmd.InOrStdin())
				txb    = auth.NewTxBuilderFromCLI(buffer).WithTxEncoder(utils.GetTxEncoder(cdc))
				ctx    = context.NewCLIContextWithInput(buffer).WithCodec(cdc)
			)

			identity, err := strconv.ParseUint(args[0], 0, 64)
			if err != nil {
				return err
			}

			msg := types.NewMsgApprove(ctx.FromAddress, identity)
			return utils.GenerateOrBroadcastMsgs(ctx, txb, []sdk.Msg{msg})
		},
	}

	return cmd
}
