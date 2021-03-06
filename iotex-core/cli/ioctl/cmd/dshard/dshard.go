// Copyright (c) 2019 IoTeX
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

package dshard

import (
	"github.com/spf13/cobra"
)

// Flags
var (
	epochNum uint64
)

// NodeCmd represents the node command
var DshardCmd = &cobra.Command{
	Use:   "dshard",
	Short: "Find the weak shard in the system",
	Args:  cobra.ExactArgs(1),
}

func init() {
	DshardCmd.AddCommand(dnodeDelegateCmd)
	DshardCmd.AddCommand(dnodeRewardCmd)
	dnodeDelegateCmd.Flags().Uint64VarP(&epochNum, "epoch-num", "e", 0, "specify specific epoch")
}
