package internal

import (
	"fmt"
	"github.com/jaypipes/ghw"
	"strings"
	"github.com/safchain/ethtool"
)

type NIC struct {
	NAME	string
	MAC	string
	PCI	string
	FLAG	bool
}

func (n NIC) Check100GNIC() NIC {

	pci, err := ghw.PCI()

	if err != nil {
		fmt.Printf("Error getting PCI info: %v", err)
	}

	n.FLAG = false

	for _, device := range pci.Devices {
		vendor := device.Vendor
		vendorName := vendor.Name

		if len(vendor.Name) > 20 {
			vendorName = string([]byte(vendorName)[0:17]) + "..."
		}

		product := device.Product
		productName := product.Name
		if len(product.Name) > 40 {
			productName = string([]byte(productName)[0:37]) + "..."
		}

		if strings.Contains(productName,`E810-C`) {
			//fmt.Printf("%-12s\t%-20s\t%-40s\n", device.Address, vendorName, productName)
			if device.Address == n.PCI {
				n.FLAG = true
			}
		}
	}

	return n
}

func (n NIC) GetEthInfo() NIC {

	ethHandle, err := ethtool.NewEthtool()

	if err != nil {
		panic(err.Error())
	}

	defer ethHandle.Close()

	stats, err := ethHandle.DriverInfo(n.NAME)

	if err != nil {
		panic(err.Error())
	}

	n.PCI = stats.BusInfo

	return n
}

func ListAllNIC() []NIC {

	NICs := []NIC{}
	net, err := ghw.Network()

	if err != nil {

		fmt.Printf("Error getting network info: %v", err)
        }

	for _, nic := range net.NICs {

		if nic.IsVirtual == false {
			newnic := NIC {
				NAME:	nic.Name,
				MAC:	nic.MacAddress,
			}

			NICs = append(NICs,newnic)
		}

	}

	return NICs
}
