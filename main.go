package main

import (
	"bytes"
	"errors"
	"log"
	"os/exec"
	"strings"
	"time"
	"wakeOnLan/wol"
)

func main() {
	ticker := time.NewTicker(time.Second * 20)
	for {
		<-ticker.C
		hour := time.Now().Hour()
		// 早上10点到晚上23点之间生效
		if hour < 10 && hour > 23 {
			continue
		}
		err := run()
		if err != nil {
			log.Printf("wake on lan failed: %s\n", err)
			continue
		}
		log.Println("nothing happened")
	}
}

func run() error {
	cmd := exec.Command("nmap", "-sL", "192.168.1.1-10")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return err
	}
	result := out.String()
	if strings.Contains(result, "jwang") {
		return errors.New("jwang is online")
	}
	if strings.Contains(result, "redmi-k20-pro-premium-edition") {
		return wol.Wol("2c:f0:5d:3b:7f:ca")
	}
	return nil
}
