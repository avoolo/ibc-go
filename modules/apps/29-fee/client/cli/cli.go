package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"
)

// GetQueryCmd returns the query commands for 29-fee
func GetQueryCmd() *cobra.Command {
	queryCmd := &cobra.Command{
		Use:                        "ibc-fee",
		Short:                      "IBC relayer incentivization query subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
	}

	queryCmd.AddCommand(
		GetCmdIncentivizedPacket(),
		GetCmdIncentivizedPackets(),
		GetCmdTotalRecvFees(),
		GetCmdTotalAckFees(),
		GetCmdTotalTimeoutFees(),
		GetCmdIncentivizedPacketsForChannel(),
		GetCmdCounterpartyAddress(),
		GetCmdFeeEnabledChannel(),
		GetCmdFeeEnabledChannels(),
	)

	return queryCmd
}

// NewTxCmd returns the transaction commands for 29-fee
func NewTxCmd() *cobra.Command {
	txCmd := &cobra.Command{
		Use:                        "ibc-fee",
		Short:                      "IBC relayer incentivization transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	txCmd.AddCommand(
		NewRegisterPayeeCmd(),
		NewPayPacketFeeAsyncTxCmd(),
		NewRegisterCounterpartyAddressCmd(),
	)

	return txCmd
}
