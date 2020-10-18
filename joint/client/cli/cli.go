package cli

import (
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"
)

func GetQueryCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "joint",
		Short: "Joint account module sub-commands",
	}

	cmd.AddCommand(flags.GetCommands(
		queryAccount(cdc),
		queryAccounts(cdc),
		queryTransfer(cdc),
		queryTransfers(cdc),
	)...)

	return cmd
}

func GetTxCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "joint",
		Short: "Joint account module sub-commands",
	}

	cmd.AddCommand(flags.PostCommands(
		txCreate(cdc),
		txDeposit(cdc),
		txSend(cdc),
		txApprove(cdc),
	)...)

	return cmd
}
