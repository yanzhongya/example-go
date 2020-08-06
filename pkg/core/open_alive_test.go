package core

import (
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"testing"
	"time"
)

var (
	device       string = "\\Device\\NPF_{9E9B3985-5BAC-4BC1-BFC3-0CA8C145BF10}"
	snapshot_len int32  = 1024
	promiscuous  bool   = false
	err          error
	timeout      time.Duration = 30 * time.Second
	handle       *pcap.Handle
)

func Test_openAlive(t *testing.T) {
	// OpenLive 打开设备并返回 Handle。
	// device 设备名称
	// snapshot_len 每次读取数据包大小
	// promiscuous 是否将接口至于混杂模式，见 TCP/IP 详解
	// timeout 超时时间
	handle, err = pcap.OpenLive(device, snapshot_len, promiscuous, timeout)
	if err != nil {
		fmt.Println(err.Error())
	}


	// 使用 handle 作为信息包来源处理所有信息包
	/**
	 		return &PacketSource{
				source:  source,
				decoder: decoder,
			}
	*/
	package_source := gopacket.NewPacketSource(handle,handle.LinkType())
	/**
		这种方法最简单，最容易实现，但缺乏灵活性。数据包返回“chan 的数据包”，然后将
	数据包异步写到该channel。数据包会阻塞 channel，如果基础包 PacketDATa Source 返回 EOF
	就关闭改接口。其他错误被忽略
	*/

	for packet := range package_source.Packets() {
		fmt.Println(packet)
	}
}
