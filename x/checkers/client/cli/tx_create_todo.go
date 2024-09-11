package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/faridanangs/checkers/x/checkers/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdCreateTodo() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-todo [title] [text]",
		Short: "Broadcast message createTodo",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argTitle := args[0]
			argText := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateTodo(
				clientCtx.GetFromAddress().String(),
				argTitle,
				argText,
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
