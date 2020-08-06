package core

import (
	"github.com/google/gopacket"
	"testing"
)

/**
自定义层
相当于自己定义了一层协议，甚至可以不使用 TCP/IP 或者 Ethernet
*/
// 自定义层数据结构
type CustomLayer struct {
	SomeByte    byte
	AnotherByte byte
	restOfData  []byte
}

//RegisterLayerType创建一个新的 Layer 类型并在全局进行注册。
//传入的数字必须唯一，否则将发生运行时恐慌。 编号0-999保留用于gopacket库。
//数字1000-1999应该用于常见的特定于应用程序的类型，并且速度非常快。
//任何其他数字（负数或> = 2000）都可以用于不常见的应用程序特定类型，
//并且速度稍慢（它们需要在数组索引上进行映射查找）。
var CustomLayerType = gopacket.RegisterLayerType(
	20001,
	gopacket.LayerTypeMetadata{
		"CustomTypeData",
		gopacket.DecodeFunc(decodeCustomLayer),
	},
)

// 查询的时候，返回类型
func (l CustomLayer) LayerType() gopacket.LayerType {
	return CustomLayerType
}

// 返回 layer 提供的信息。
func (l CustomLayer) LayerContents() []byte {
	return []byte{l.SomeByte, l.AnotherByte}
}

// LayerPayload返回在我们的图层或原始有效负载之上构建的后续图层
func (l CustomLayer) LayerPayload() []byte {
	return l.restOfData
}

// 自定义解码功能
func decodeCustomLayer(data []byte, p gopacket.PacketBuilder) error {
	p.AddLayer(&CustomLayer{data[0], data[1], data[2:]})
	return p.NextDecoder(gopacket.LayerTypePayload)
}

func Test_mainCustomDecoder(t *testing.T) {
	//	如果您创建自己的编码和解码，则实际上可以创建自己的协议或实现layers包中尚未定义的协议。
	//	在我们的示例中，我们只是将普通的以太网数据包包装到我们自己的层中。
	//	如果您想创建一些难以理解的二进制数据类型，则创建自己的协议是很好的

	rawData := []byte{0xF0, 0x0F, 65, 65, 66, 67, 68}

	packet := gopacket.NewPacket(rawData, CustomLayerType, gopacket.Default)
	t.Log("创建包数据")
	t.Log(packet)

	customLayer := packet.Layer(CustomLayerType)
	if customLayer != nil{
		t.Log("packet 已成功解码,自定义层解码器")
		customLayerContent,_ := customLayer.(*CustomLayer)

		t.Log("Payload: ", customLayerContent.LayerPayload())
		t.Log("SomeByte element:", customLayerContent.SomeByte)
		t.Log("AnotherByte element:", customLayerContent.AnotherByte)
	}
}
