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
		err := run()
		if err != nil {
			log.Printf("wake on lan failed: %s\n", err)
		}
		log.Println("nothing happened")
		<-ticker.C
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
	if strings.Contains("jwang", result) {
		return errors.New("jwang is online")
	}
	if strings.Contains("redmi-k20-pro-premium-edition", result) {
		err = wol.Wol("192.168.1.2:9", "2c:f0:5d:3b:7f:ca")
		if err != nil {
			return err
		}
	}
	return nil
}
