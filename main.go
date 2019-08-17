package main

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"net"
	"os"
	"time"
	"wakeOnLan/ping"
	"wakeOnLan/wol"
)

func main() {
	app := cli.NewApp()
	app.Name = "autoWol"
	app.Version = "0.0.1"
	app.EnableBashCompletion = true
	app.Usage = "一个根据条件定时唤醒电脑的小工具,可以运行在路由器上面"
	app.Description = "这个工具主要的用途是运行在我的路由器上面（也可以运行在其它平台），不断的ping我的手机IP，" +
		"如果我下班回家手机一旦连入路由器就会自动唤醒我的电脑，懒人必备！" +
		"\n\t 使用方式，举个例子: auto-wol -r 192.168.31.214 -i 192.168.31.8 -m 4C:ED:FB:94:71:0F"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:     "replyIp,r",
			Required: true,
			Usage:    "根据哪个设备IP的状态去唤醒电脑，比如手机的IP",
		},
		cli.StringFlag{
			Name:     "wolIp,i",
			Required: true,
			Usage:    "需要唤醒的设备的IP，比如电脑IP",
		},
		cli.StringFlag{
			Name:     "wolMac,m",
			Required: true,
			Usage:    "需要唤醒的设备的MAC地址，比如电脑MAC地址",
		},
		cli.StringFlag{
			Name:  "wolPort,p",
			Value: "9",
			Usage: "需要唤醒的设备的PORT，一般电脑主板默认都是9端口",
		},
		cli.IntFlag{
			Name:  "start,s",
			Value: 20,
			Usage: "唤醒生效开始时间，比如晚上8点",
		},
		cli.IntFlag{
			Name:  "end,e",
			Value: 22,
			Usage: "唤醒生效结束时间，比如晚上10点",
		},
		cli.IntFlag{
			Name:  "interval,it",
			Value: 30,
			Usage: "检查状态间隔时间，也就是Ping的频率，默认30s",
		},
	}
	app.Action = func(c *cli.Context) error {
		return wakeUp(c)
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func wakeUp(c *cli.Context) error {
	var (
		relyIp  = c.String("relyIp")
		wolIp   = c.String("wolIp")
		wolPort = c.String("wolPort")
		wolMac  = c.String("wolMac")
		start   = c.Int("start")
		end     = c.Int("end")
		it      = c.Int("interval")
	)
	for {
		ticker := time.NewTicker(time.Second * time.Duration(it))
		select {
		case <-ticker.C:
			// 由于路由器时区有问题，默认UTC时区，这里手动+8小时
			fmt.Printf("Now Time is: " + time.Now().Add(8*time.Hour).Format("2006-01-02 15:04:05") + "\n")
			hour := time.Now().Hour() + 8
			if hour < start || hour > end {
				return nil
			}
			// 判断当前主机是否已经在线
			addr, err := net.ResolveIPAddr("ip4", wolIp)
			err = ping.Ping(1, addr)
			if err == nil {
				fmt.Printf("target host is online\n")
			} else {
				// 当手机连入WiFi的时候发送唤醒数据包
				addr, err := net.ResolveIPAddr("ip4", relyIp)
				if err = ping.Ping(1, addr); err == nil {
					err := wol.Wol(wolIp+wolPort, wolMac)
					if err != nil {
						fmt.Println("wake on lan failed, error:" + err.Error())
					}
				}
			}
			return nil
		}
	}
}
