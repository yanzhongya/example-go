package core

import (
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"net"
	"testing"
)

/**
创建和发送packet
演示如何使用ethernet、IP和TCP层创建数据包。所有的东西都是默认的和空的，所以它实际上不做任何事情
 */
var (
	buffer       gopacket.SerializeBuffer
	options      gopacket.SerializeOptions
)
func Test_createOrSendPacket(t *testing.T) {
	handle, err := pcap.OpenLive(device,snapshot_len,promiscuous,timeout)
	if err != nil{
		t.Log(err.Error())
	}

	// 发送的数据
	rawBytes := []byte{10,20,30}
	// 将给定的数据注入pcap句柄。
	err = handle.WritePacketData(rawBytes)
	if err != nil {
		t.Log(err.Error())
	}

	// 创建一个正确格式的数据包
	// 只需要携带空的数据，应该填充 MAC 地址、IP 地址、等等
	buffer = gopacket.NewSerializeBuffer()
	gopacket.SerializeLayers(buffer,options,
		&layers.Ethernet{},
		&layers.IPv4{},
		&layers.TCP{},
		gopacket.Payload(rawBytes),
	)

	outgoingPacket := buffer.Bytes()
	t.Log(string(outgoingPacket))

	// 写出数据
	err = handle.WritePacketData(outgoingPacket)
	if err != nil {
		t.Log(err.Error())
	}

	ipLayer := &layers.IPv4{
		SrcIP: net.IP{127,0,0,1},
		DstIP: net.IP{8,8,8,8},
	}

	ethernetLayer := &layers.Ethernet{
		SrcMAC: net.HardwareAddr{0xFF, 0xAA, 0xFA, 0xAA, 0xFF, 0xAA},
		DstMAC: net.HardwareAddr{0xBD, 0xBD, 0xBD, 0xBD, 0xBD, 0xBD},
	}
	tcpLayer := &layers.TCP{
		SrcPort: layers.TCPPort(4321),
		DstPort: layers.TCPPort(80),
	}

	buffer = gopacket.NewSerializeBuffer()
	err = gopacket.SerializeLayers(buffer,options,
		ethernetLayer,
		ipLayer,
		tcpLayer,
		gopacket.Payload(rawBytes))
	if err != nil {
		t.Log(err.Error())
	}

	outgoingPacket = buffer.Bytes()
	t.Log(string(outgoingPacket))
}
