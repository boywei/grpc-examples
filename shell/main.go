package main

import (
	"io"
	"log"
	"os"
	"os/exec"

	"github.com/creack/pty"
	"golang.org/x/term"
)

func createByPty() error {
	// 生成一个命令对象
	cmd := exec.Command("bash")

	// 获得pty
	ptyMaster, tty, err := pty.Open()
	if err != nil {
		return err
	}
	defer ptyMaster.Close()

	// 设置cmd的输入输出为pty
	cmd.Stdin = tty
	cmd.Stdout = tty
	cmd.Stderr = tty

	// 开始执行命令
	if err := cmd.Start(); err != nil {
		return err
	}

	// 准备读写线程
	go func() {
		// 读取ptyMaster并输出
		if _, err := io.Copy(os.Stdout, ptyMaster); err != nil {
			panic(err)
		}
	}()
	go func() {
		// 从stdin读取并输出到ptyMaster
		if _, err := io.Copy(ptyMaster, os.Stdin); err != nil {
			panic(err)
		}
	}()

	// 等待命令结束
	if err := cmd.Wait(); err != nil {
		return err
	}

	return nil
}

func createByTerm() error {
	// Create arbitrary command.
	c := exec.Command("bash")

	// Start the command with a pty.
	ptmx, err := pty.Start(c)
	if err != nil {
		return err
	}
	// Make sure to close the pty at the end.
	defer func() { _ = ptmx.Close() }() // Best effort.

	// Set stdin in raw mode.
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		return err
	}
	defer func() { _ = term.Restore(int(os.Stdin.Fd()), oldState) }() // Best effort.

	// Copy stdin to the pty and the pty to stdout.
	// NOTE: The goroutine will keep reading until the next keystroke before returning.
	go func() { _, _ = io.Copy(ptmx, os.Stdin) }()
	_, _ = io.Copy(os.Stdout, ptmx)

	return nil
}

func main() {
	if err := createByTerm(); err != nil {
		log.Fatal(err)
	}
}
