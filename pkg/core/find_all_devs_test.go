package core

import (
	"fmt"
	"github.com/google/gopacket/pcap"
	"testing"
)

/**
查找所有设备
 */
func Test_findAllDevs(t *testing.T) {
	devices, err := pcap.FindAllDevs()
	if err != nil{
		t.Error(err.Error())
		return
	}
	for _, device := range devices {
		t.Log("设备名称：" , device.Name)
		t.Log("设备状态：", device.Flags)
		t.Log("设备描述",device.Description)
		t.Log("设备地址:")
		for _, address := range device.Addresses {
			t.Log("\t IP 地址",address.IP)
			t.Log("\t P2P 地址",address.P2P)
			t.Log("\t 广播地址",address.Broadaddr)
			t.Log("\t 子网掩码",address.Netmask.String())
		}
		fmt.Println()
	}
}