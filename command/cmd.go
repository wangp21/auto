package command

import (
        "github.com/spf13/cobra"
)

func Execute() {
        if err := cmdRoot.Execute(); err != nil {
                //panic(err)
        }
}

var cmdRoot = &cobra.Command{
        Use:   "auto provision toolkit",
        Short: "auto provision toolkit",
}
