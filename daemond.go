package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"
)

func init() {
	if os.Getppid() != 1 {
		cmd := exec.Command(os.Args[0], os.Args[1:]...)
		cmd.Start()
		os.Exit(0)
	}

	// 监听系统信号
	go func() {
		_c := make(chan os.Signal, 1)
		signal.Notify(_c, os.Interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGKILL)
		msg := <-_c
		log.Println(msg)
		os.Exit(0)
	}()
}

func main() {

	go func() {
		fp, _ := os.OpenFile("log", os.O_WRONLY|os.O_CREATE, 0644)
		log.SetOutput(fp)
		log.Println(fmt.Sprint(os.Getpid()))
		for {
			// log.Println(fmt.Sprint("hello ", os.Getppid()))
			time.Sleep(time.Second * 5)
		}
	}()

	for {
		time.Sleep(time.Second * 1000)
	}
}
