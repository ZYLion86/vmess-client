package main

import (
	"log"
	"vmess-client/src/vmesshub"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Ltime | log.Lshortfile)
	socks5ListenUri := "127.0.0.1:10808"

	log.Println("Local Socks5 :", socks5ListenUri)
	uuid := "137a472d-16f4-11ed-9650-005056bf05aa"
	remoteHost := "103.44.238.218:22223"
	securityType := "none"
	alterId := 0

	vh, err := vmesshub.CreateVmessHub(socks5ListenUri, remoteHost, uuid, securityType, alterId)

	if err != nil {
		log.Println(err)
		return
	} else {
		log.Println("连接成功")
		vh.StartSocks5Listen()
	}
}
