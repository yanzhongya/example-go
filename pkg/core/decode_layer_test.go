package core

import (
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"strings"
	"testing"
)

/**
解码包数据
*/
func Test_decodeLayer(t *testing.T) {
	handle, err := pcap.OpenLive(device, snapshot_len, promiscuous, timeout)
	if err != nil {
		t.Log()
	}

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

	for packet := range packetSource.Packets() {

		t.Log()
		ethernetLayer := packet.Layer(layers.LayerTypeEthernet)
		if ethernetLayer != nil {
			// 可以解析 Ethernet 信息
			t.Log("Ethernet 层解码")
			ethernetPacket, _ := ethernetLayer.(*layers.Ethernet)
			t.Log("来源 Mac 地址", ethernetPacket.SrcMAC)
			t.Log("目标 Mac 地址", ethernetPacket.DstMAC)
			t.Log("Ethernet 类型", ethernetPacket.EthernetType)
		}
		t.Log()
		layerTypeIPv4 := packet.Layer(layers.LayerTypeIPv4)
		if layerTypeIPv4 != nil {
			// 可以解析 Ethernet 信息
			t.Log("IPv4 层解码")
			ipv4, _ := layerTypeIPv4.(*layers.IPv4)
			t.Log("来源 IP 地址", ipv4.SrcIP)
			t.Log("目标 IP 地址", ipv4.DstIP)
			t.Log("协议", ipv4.Protocol)
		}
		t.Log()

		layerTypeTCP := packet.Layer(layers.LayerTypeTCP)
		if layerTypeIPv4 != nil {
			// 可以解析 TCP 信息
			tcp, _ := layerTypeTCP.(*layers.TCP)
			if tcp != nil {
				t.Log("TCP 层解码")
				t.Log("来源 Port 地址", tcp.SrcPort)
				t.Log("目标 Port 地址", tcp.DstPort)
				t.Log("TCP 序列号", tcp.Seq)
			}
		}
		t.Log()

		// 输出所有的接收到的包类型
		t.Log("All Packet")
		for _, layer := range packet.Layers() {
			t.Log("-",layer.LayerType())
		}
		t.Log()

		// 返回数据包中的第一个应用层
		applicationLayer := packet.ApplicationLayer()
		if applicationLayer != nil{
			t.Log("应用层或则负载发现")
			if strings.Contains(string(applicationLayer.Payload()),"HTTP") {
				t.Log("这是一个 HTTP 请求")
				t.Log(string(applicationLayer.Payload()))
			}
		}
		// 检查是否有错误
		if err := packet.ErrorLayer(); err != nil {
			t.Log("解码数据包的某些部分时出错:", err)
		}


	}

}
