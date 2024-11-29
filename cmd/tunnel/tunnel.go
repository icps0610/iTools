package main

import (
	"flag"
	"fmt"
	"io"
	"net"

	"tools/pkg/cmdutils"
	"tools/pkg/utils"
)

func main() {
	appVersion := `1.0`
	argsCount := 0
	description := `This program enables customizable port forwarding, 
allowing you to specify the local port and remote IP/port.`
	useage := `Usage: tunnel -lp [port] -ip [ip] -p [port]

 -lp, --localPort          localPort (the local port the application will listen on)
      --ip                 ip (the remote IP address to connect to).
  -p, --port               port (the remote port to connect to).

Example: 
    tunnel -lp 50 -ip 192.168.0.1 -p 80`

	// var args []string
	var localPort, ip, port string

	// 設定參數
	{
		cmdutils.Set(useage, description)

		flag.StringVar(&localPort, "lp", "", "localPort")
		flag.StringVar(&localPort, "lport", "", "localPort")

		flag.StringVar(&ip, "ip", "", "ip")

		flag.StringVar(&port, "p", "", "port")
		flag.StringVar(&port, "port", "", "port")

		cmdutils.SetArgs(argsCount, appVersion) // args =

	}

	// 主程式
	{
		// required
		if localPort == "" || ip == "" || port == "" {
			cmdutils.ShowHelp()
		}

		ln, err := net.Listen("tcp", ":"+localPort)
		utils.PrintError(err)

		defer ln.Close()

		for {
			conn, err := ln.Accept()
			utils.PrintError(err)

			fmt.Println("New connection from:", conn.RemoteAddr().String())

			remote, err := net.Dial("tcp", ip+":"+port)
			if err != nil {
				utils.PrintError(err)
				conn.Close()
				continue
			}
			// remote
			go func() {
				defer conn.Close()
				defer remote.Close()

				_, err := io.Copy(remote, conn)
				utils.PrintError(err)
			}()

			// local
			go func() {
				defer conn.Close()
				defer remote.Close()

				_, err := io.Copy(conn, remote)
				utils.PrintError(err)
			}()

		}
	}
}
