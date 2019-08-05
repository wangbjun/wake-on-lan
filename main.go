package main

import (
	"log"
	"net"
	"time"
	"wakeOnLan/ping"
	"wakeOnLan/wol"
)

const (
	IP   = "192.168.31.214"    //手机的ip
	HOST = "192.168.31.8"      //电脑ip
	PORT = ":9"                //wol端口，默认都是9
	MAC  = "4C:ED:FB:94:71:0F" //电脑mac地址
)

func main() {
	// 晚上8点到11点之间才触发,由于路由器时区有问题，这里手动+8小时
	log.Println("Now Time is: " + time.Now().Add(8*time.Hour).Format("2006-01-02 15:04:05"))
	hour := time.Now().Hour() + 8
	if hour < 20 || hour > 23 {
		return
	}
	// 判断当前主机是否已经在线
	addr, err := net.ResolveIPAddr("ip4", HOST)
	err = ping.Ping(1, addr)
	if err == nil {
		log.Println("target host is online")
	} else {
		// 当手机连入WiFi的时候发送唤醒数据包
		addr, err := net.ResolveIPAddr("ip4", IP)
		if err = ping.Ping(1, addr); err != nil {
			log.Println("ip unreachable,error:" + err.Error())
		} else {
			err := wol.Wol(HOST+PORT, MAC)
			if err != nil {
				log.Println("wake on lan failed, error:" + err.Error())
			}
		}
	}
}
