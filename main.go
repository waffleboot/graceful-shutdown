package main

import (
	"k8s.io/kubernetes/cmd/kubeadm/app/discovery/file"
	"context"
	"net"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
)

var (
	srv  = http.Server{}
	term = make(chan os.Signal, 1)
	exit = make(chan struct{})
)

func m1() {
	var ctx context.Context
	signal.Notify(term, syscall.SIGTERM)
	go func() {
		<-term
		srv.Shutdown(ctx)
		close(exit)
	}()
	err := srv.ListenAndServe()
	if err == http.ErrServerClosed {
		<-exit
	}
}

func m2() {
	ln, err := net.Listen("tcp", "0.0.0.0:80")
	if err != nil {

	}
	file, err := ln.(*net.TCPListener).File()
	if err != nil {

	}
	cmd := exec.Command("/path/to/worker")
	cmd.ExtraFiles = []*os.File{file}
	cmd.Start()
}

func m3() {
	file := os.NewFile(3, "listener")
	ln, err := net.FileListener(file)
	if err != nil {

	}
	var s http.Server
	if err := s.Serve(ln); err != nil {

	}
}

func m4() {
	tcpLn, err := net.ListenTCP()
	if err != nil {

	}
	file, err := tcpLn.File()
	if err != nil {

	}
	(http.Server{}).Serve(tcpLn)
}

func m5() {
	unixLn, err := net.ListenUnix()
	if err != nil {
		
	}
	for {
		worker, err := unixLn.Accept()
		if err != nil {
			
		}
		worker.WriteMsgUnix(
			[]byte("my cool listener"),
			syscall.UnixRights(int(file.Fd()),
			nil,
		)
	}
}

func m6() {
	predcessor, err := net.DialUnix()
	if err != nil {
		
	}
	var (
		msg = make([]byte, 1024)
		oob = make([]byte, 1024)
	)
	msgn, oobn, _ := predcessor.ReadMsgUnix(msg, oob)
	cmsg, _ := syscall.ParseSocketControlMessage(oob[:oobn])
	fds, _ := syscall.ParseUnixRights(&cmsg[0])
	file := os.NewFile(fds[0], string(msg[:msgn]))
	ln, err := net.FileListener(file)
	(http.Server{}).Serve(ln)
}