package cli

import (
	"fmt"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"

	"github.com/exidioco/modules/escrow/client/common"
)

func queryEscrow(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "escrow [identity]",
		Short: "Query an escrow",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.NewCLIContext().WithCodec(cdc)

			identity, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			escrow, err := common.QueryEscrow(ctx, identity)
			if err != nil {
				return err
			}

			fmt.Println(escrow)
			return nil
		},
	}
}

func queryEscrows(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "escrows",
		Short: "Query escrows",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.NewCLIContext().WithCodec(cdc)

			page, err := cmd.Flags().GetInt(flagPage)
			if err != nil {
				return err
			}

			limit, err := cmd.Flags().GetInt(flagLimit)
			if err != nil {
				return err
			}

			escrows, err := common.QueryEscrows(ctx, page, limit)
			if err != nil {
				return err
			}

			for _, escrow := range escrows {
				fmt.Printf("%s\n\n", escrow)
			}

			return nil
		},
	}

	cmd.Flags().Int(flagPage, 1, "page")
	cmd.Flags().Int(flagLimit, 0, "limit")

	return cmd
}
