package command

import (
        "github.com/spf13/cobra"
	"auto/internal"
	"fmt"
)

var cmddpdk = &cobra.Command {
        Use: "dpdk",
        Short: "dpdk commandline",
}

func init() {

	cmdRoot.AddCommand(cmddpdk)
	
	actions := []string{"enable","disable"}
	for _, action := range(actions) {
                 cmdApi := &cobra.Command {
                        Use: action,
                        Short:  action + " action for dpdk",
                        Run: func(cmd *cobra.Command, args []string) {
                                switch v:=cmd.Use;v {
					case "enable":
						//wf := internal.Workflow{}
					        //wf = wf.WithScripts("step0","whoami")
						//wf.RunScript("step0")
						//internal.List100GNIC()
						//internal.GetEthInfo("eth0")
						//internal.ListAllNIC()
						//nics := internal.ListAllNIC()
						//for _, nic := range(nics) {
						//	fmt.Printf("%+v\n",nic)
						//}
						
						NICS100G := []internal.NIC{}
						NICS := internal.ListAllNIC()
						for _, nic := range(NICS) {
							nic = nic.GetEthInfo()	
							nic = nic.Check100GNIC()
							if nic.FLAG == true {
								NICS100G = append(NICS100G,nic)
							}
						}
	
						for _, nic := range(NICS100G) {
							fmt.Printf("%+v\n",nic)
						}
	
				}
                        },

                 }

        cmddpdk.AddCommand(cmdApi)
        }
}
