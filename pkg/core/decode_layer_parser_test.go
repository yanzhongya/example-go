package core

import (
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"testing"
)

var (
	ethLayer layers.Ethernet
	ipLayer  layers.IPv4
	tcpLayer layers.TCP
)

func Test_decoderLayerParser(t *testing.T) {
	handle, err := pcap.OpenLive(device, snapshot_len, promiscuous, timeout)

	if err != nil {
		t.Log(err.Error())
	}

	defer handle.Close()

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

	for packet := range packetSource.Packets() {
		parser := gopacket.NewDecodingLayerParser(layers.LayerTypeEthernet, &ethLayer, &ipLayer, &ipLayer)
		var foundLayerType []gopacket.LayerType

		err := parser.DecodeLayers(packet.Data(), &foundLayerType)
		if err != nil {
			t.Log("不能解码：", err.Error())
		}

		for _, layerType := range foundLayerType {
			if layerType == layers.LayerTypeIPv4 {
				t.Log("IPv4 Src Ip:", ipLayer.SrcIP, "Dst IP:", ipLayer.DstIP)
			}
			if layerType == layers.LayerTypeTCP {
				t.Log("TCP Src port:", tcpLayer.SrcPort, ",Dst port", tcpLayer.DstPort)
			}

			if layerType == layers.LayerTypeEthernet {
				t.Log("Ethernet Src mac:", ethLayer.SrcMAC, ",dst mac:", ethLayer.DstMAC)
			}
		}

	}
}
