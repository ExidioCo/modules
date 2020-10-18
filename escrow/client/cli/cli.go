package cli

import (
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"
)

func GetQueryCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "escrow",
		Short: "Escrow module sub-commands",
	}

	cmd.AddCommand(flags.GetCommands(
		queryEscrow(cdc),
		queryEscrows(cdc),
	)...)

	return cmd
}

func GetTxCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "escrow",
		Short: "Escrow module sub-commands",
	}

	cmd.AddCommand(flags.PostCommands(
		txCreate(cdc),
		txApprove(cdc),
	)...)

	return cmd
}
