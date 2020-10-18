package cli

import (
	"fmt"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"

	"github.com/exidioco/modules/joint/client/common"
)

func queryAccount(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "account [identity]",
		Short: "Query an account",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.NewCLIContext().WithCodec(cdc)

			identity, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			account, err := common.QueryAccount(ctx, identity)
			if err != nil {
				return err
			}

			fmt.Println(account)
			return nil
		},
	}
}

func queryAccounts(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "accounts",
		Short: "Query accounts",
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

			accounts, err := common.QueryAccounts(ctx, page, limit)
			if err != nil {
				return err
			}

			for _, account := range accounts {
				fmt.Printf("%s\n\n", account)
			}

			return nil
		},
	}

	cmd.Flags().Int(flagPage, 1, "page")
	cmd.Flags().Int(flagLimit, 0, "limit")

	return cmd
}

func queryTransfer(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "transfer [identity]",
		Short: "Query a transfer",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.NewCLIContext().WithCodec(cdc)

			identity, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			transfer, err := common.QueryTransfer(ctx, identity)
			if err != nil {
				return err
			}

			fmt.Println(transfer)
			return nil
		},
	}
}

func queryTransfers(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "transfers",
		Short: "Query transfers",
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

			transfers, err := common.QueryTransfers(ctx, page, limit)
			if err != nil {
				return err
			}

			for _, transfer := range transfers {
				fmt.Printf("%s\n\n", transfer)
			}

			return nil
		},
	}

	cmd.Flags().Int(flagPage, 1, "page")
	cmd.Flags().Int(flagLimit, 0, "limit")

	return cmd
}
