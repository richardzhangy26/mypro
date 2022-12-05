package main

import "fmt"

type Usb interface {
	Start()
	Stop()
}
type Phone struct {
	name string
}

func (p Phone) Start() {
	fmt.Println("手机开始工作")
}
func (p Phone) Stop() {
	fmt.Println("手机停止工作")
}
func (p Phone) Call() {
	fmt.Println("开始打电话")
}

type Camera struct {
	name string
}

func (c Camera) Start() {
	fmt.Println("相机开始工作")
}
func (c Camera) Stop() {
	fmt.Println("相机停止工作")
}

type Computer struct {
}

func (c Computer) Working(usb Usb) {
	usb.Start()
	if phone, ok := usb.(Phone); ok {
		phone.Call()
	}

	usb.Stop()
}
func main() {
	var usbArray [3]Usb
	usbArray[0] = Phone{"诺基亚"}
	usbArray[1] = Camera{"索尼"}
	usbArray[2] = Phone{"苹果"}
	var computer Computer
	for _, device := range usbArray {
		computer.Working(device)
		fmt.Println()
	}
}
