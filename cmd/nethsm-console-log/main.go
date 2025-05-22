// Package main is a console logger for NetHSM.
package main

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
	"path"
	"time"

	"github.com/alecthomas/kong"
	"github.com/tarm/serial"
)

var opt struct {
	Dir    string `kong:"help='directory for log files'"`
	Device string `kong:"help='name of serial device',default='/dev/ttyUSB0'"`
	Print  bool   `kong:"help='print log to stdout as well'"`
}

const (
	fileMode       = 0600
	filenameFormat = "hsm-console_2006-01-02-150405.log"
)

func main() {
	kong.Parse(&opt,
		kong.Description("Nitrokey NetHSM console logger"),
		kong.UsageOnError(),
		kong.ConfigureHelp(kong.HelpOptions{Summary: true, Compact: true}),
	)

	c := &serial.Config{
		Name:     opt.Device,
		Baud:     115200,
		Parity:   serial.ParityNone,
		StopBits: serial.Stop1,
	}

	// open output file
	filename := path.Join(opt.Dir, time.Now().Local().Format(filenameFormat))

	out, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, fileMode)
	if err != nil {
		slog.Error("error opening output file", "filename", filename, "err", err)
		return
	}
	defer out.Close()

	// open serial port
	in, err := serial.OpenPort(c)
	if err != nil {
		slog.Error("error opening serial port", "device", opt.Device, "err", err)
		return
	}
	defer in.Close()

	_, err = in.Write([]byte("la\n"))
	if err != nil {
		slog.Error("failed to sent 'la' (enable all log channels) command", "err", err)
		return
	}

	err = in.Flush()
	if err != nil {
		slog.Error("failed to flush 'la' (enable all log channels) command", "err", err)
	}

	_, err = in.Write([]byte("st\n"))
	if err != nil {
		slog.Error("failed to sent 'st' (print debug server status) command", "err", err)
		return
	}

	err = in.Flush()
	if err != nil {
		slog.Error("failed to flush 'st' (print debug server status) command", "err", err)
	}

	serialReader := bufio.NewReader(in)

	slog.Info("logging NetHSM console", "device", opt.Device, "logfile", filename)
	for {
		line, err := serialReader.ReadString('\n')
		if err != nil {
			slog.Error("error reading from serial port", "err", err)
			return
		}

		if opt.Print {
			fmt.Printf("> %s", line)
		}

		_, err = out.WriteString(line)
		if err != nil {
			slog.Error("write error, exiting", "err", err)
		}
	}
}
