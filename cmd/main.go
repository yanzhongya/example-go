package main

import (
	"fmt"
	"github.com/google/gopacket/pcap"
	"log"
)

func main() {
	// 得到所有的(网络)设备
	devices, err := pcap.FindAllDevs()
	if err != nil {
		log.Fatal(err)
	}
	// 打印设备信息
	fmt.Println("设备发现:")
	for _, device := range devices {
		fmt.Println("\n名称: ", device.Name)
		fmt.Println("描述: ", device.Description)
		fmt.Println("设备描述: ", device.Description)
		for _, address := range device.Addresses {
			fmt.Println("- IP 地址: ", address.IP)
			fmt.Println("- 子网掩码: ", address.Netmask)
		}
	}
}
