package main

import (
	"log"
	"net"
	"time"
	"wakeOnLan/ping"
	"wakeOnLan/wol"
)

const (
	IP   = "192.168.31.214"
	HOST = "192.168.31.8"
	PORT = "9"
	MAC  = "4C:ED:FB:94:71:0F"
)

func main() {
	log.Println("Now Time is: " + time.Now().Add(8*time.Hour).Format("2006-01-02 15:04:05"))
	// 晚上8点到11点之间才触发
	hour := time.Now().Hour() + 8
	if hour < 20 || hour > 23 {
		return
	}
	// 当主机不在线的时候发送唤醒数据包
	addr, err := net.ResolveIPAddr("ip4", HOST)
	err = ping.Ping(1, addr)
	if err == nil {
		log.Println("target host is online")
	} else {
		addr, err := net.ResolveIPAddr("ip4", IP)
		if err != nil {
			log.Println("resolve ip addr error:" + err.Error())
		}
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
