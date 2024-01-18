package main

import (
	"flag"
	"io"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"syscall"

	"github.com/creack/pty"
	"golang.org/x/term"
)

type Client struct {
	Ip       string
	Port     int
	User     string
	Password string
}

var (
	user     string
	password string
	ip       string
	port     int
)

func init() {
	flag.StringVar(&ip, "ip", "10.45.4.37", "ip address")
	flag.IntVar(&port, "port", 80, "port")
	flag.StringVar(&user, "user", "stardust", "username")
	flag.StringVar(&password, "password", " ", "password")
}

func (cli *Client) CreateByTerm() error {
	// Create arbitrary command.
	c := exec.Command("sshpass", "-p", cli.Password, "ssh", cli.User+"@"+cli.Ip)
	log.Print(c.String())

	// Start the command with a pty.
	ptmx, err := pty.Start(c)
	if err != nil {
		return err
	}
	// Make sure to close the pty at the end.
	defer func() { _ = ptmx.Close() }() // Best effort.

	// Handle pty size.
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGWINCH)
	go func() {
		for range ch {
			if err := pty.InheritSize(os.Stdin, ptmx); err != nil {
				log.Printf("error resizing pty: %s", err)
			}
		}
	}()
	ch <- syscall.SIGWINCH                        // Initial resize.
	defer func() { signal.Stop(ch); close(ch) }() // Cleanup signals when done.

	// Set stdin in raw mode.
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}
	defer func() { _ = term.Restore(int(os.Stdin.Fd()), oldState) }() // Best effort.

	// Copy stdin to the pty and the pty to stdout.
	// NOTE: The goroutine will keep reading until the next keystroke before returning.
	go func() { _, _ = io.Copy(ptmx, os.Stdin) }()
	_, _ = io.Copy(os.Stdout, ptmx)

	return nil
}

func main() {
	flag.Parse()

	cli := &Client{
		User:     user,
		Password: password,
		Ip:       ip,
		Port:     port,
	}
	if err := cli.CreateByTerm(); err != nil {
		log.Fatal(err)
	}
}
