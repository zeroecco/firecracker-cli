package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/urfave/cli"
)

func callSocket(content) {
	c, err := net.Dial("unix", "/tmp/firecracker.socket")
	if err != nil {
		panic(err)
	}
	defer c.Close()

	_, err = c.Write([]byte(content))
	if err != nil {
		log.Fatal("write error:", err)
	}
}

func main() {
	app := cli.NewApp()
	app.Name = "Firecracker CLI"
	app.Usage = "make an explosive microVM entrance"
	app.Version = "0.0.1"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "file-system, f",
			Value: "./rootfs.ext4",
			Usage: "Root FS to use",
		},
		cli.StringFlag{
			Name:  "kernel, k",
			Value: "./kernel.bin",
			Usage: "Kernel to use",
		},
		cli.StringFlag{
			Name:  "cpu, c",
			Value: "1",
			Usage: "How many vCPU's to use",
		},
		cli.StringFlag{
			Name:  "memory, m",
			Value: "1024",
			Usage: "amount of memory to use (in MegaBytes)",
		},
	}

	app.Action = func(c *cli.Context) error {
		fmt.Println("What happened?! Do better next time.")
		fmt.Println("")
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
