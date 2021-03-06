package tun2socks

import (
	"fmt"
	"log"
	"net"

	"github.com/FlowerWrong/tun2socks/util"
	"github.com/FlowerWrong/water"
)

func Ifconfig(tunName, network string, mtu uint32) {
	var ip, ipv4Net, _ = net.ParseCIDR(network)
	ipStr := ip.To4().String()
	sargs := fmt.Sprintf("%s %s %s mtu %d netmask %s up", tunName, ipStr, ipStr, mtu, util.Ipv4MaskString(ipv4Net.Mask))
	if err := util.ExecCommand("ifconfig", sargs); err != nil {
		log.Fatal("execCommand failed", err)
	}
}

func NewTun(app *App) {
	var err error
	app.Ifce, err = water.New(water.Config{
		DeviceType: water.TUN,
	})
	if err != nil {
		log.Fatal("Create tun interface failed", err)
	}
	log.Println("[tun] interface name is", app.Ifce.Name())
	Ifconfig(app.Ifce.Name(), app.Cfg.General.Network, app.Cfg.General.Mtu)
}
