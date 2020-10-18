package cli

import (
	"bufio"
	"strconv"
	"strings"
	"time"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/spf13/cobra"

	"github.com/exidioco/modules/joint/types"
)

func txCreate(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create [consents] [holders]",
		Short: "Create a joint account",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			var (
				buffer = bufio.NewReader(cmd.InOrStdin())
				txb    = auth.NewTxBuilderFromCLI(buffer).WithTxEncoder(utils.GetTxEncoder(cdc))
				ctx    = context.NewCLIContextWithInput(buffer).WithCodec(cdc)
			)

			consents, err := strconv.Atoi(args[0])
			if err != nil {
				return err
			}

			var (
				splits  = strings.Split(args[1], ",")
				holders = make([]sdk.AccAddress, 0, len(splits))
			)

			for _, split := range splits {
				address, err := sdk.AccAddressFromBech32(split)
				if err != nil {
					return err
				}

				holders = append(holders, address)
			}

			msg := types.NewMsgCreate(ctx.FromAddress, consents, holders)
			return utils.GenerateOrBroadcastMsgs(ctx, txb, []sdk.Msg{msg})
		},
	}

	return cmd
}

func txDeposit(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "deposit [identity] [amount]",
		Short: "Make a deposit into a joint account",
		RunE: func(cmd *cobra.Command, args []string) error {
			var (
				buffer = bufio.NewReader(cmd.InOrStdin())
				txb    = auth.NewTxBuilderFromCLI(buffer).WithTxEncoder(utils.GetTxEncoder(cdc))
				ctx    = context.NewCLIContextWithInput(buffer).WithCodec(cdc)
			)

			identity, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			amount, err := sdk.ParseCoins(args[1])
			if err != nil {
				return err
			}

			msg := types.NewMsgDeposit(ctx.FromAddress, identity, amount)
			return utils.GenerateOrBroadcastMsgs(ctx, txb, []sdk.Msg{msg})
		},
	}

	return cmd
}

func txSend(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "send [from] [to] [amount] [deadline]",
		Short: "Send amount from a joint account",
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

			msg := types.NewMsgSend(ctx.FromAddress, identity, to, coins, deadline)
			return utils.GenerateOrBroadcastMsgs(ctx, txb, []sdk.Msg{msg})
		},
	}

	return cmd
}

func txApprove(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "approve [identity]",
		Short: "Approve a transfer",
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
