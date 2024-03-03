package main

import (
	"fmt"
	"io"
	"os"
	"sync"

	"time"

	"github.com/gookit/color"
)



func main() {
	buf := make([]byte, 1)
	var prev byte

	linenr := 0

	printmutex := sync.Mutex{}

	timer := time.NewTimer(time.Second * 2)

	go func() {
		for {
			<- timer.C
			printmutex.Lock()
			if prev != '\n' {
				fmt.Print("\n...")
			} else {
				fmt.Print("...\n")
			}
			printmutex.Unlock()
			timer.Reset(time.Second * 2)
		}
	}()

	boldred := color.Style{color.Bold, color.Red}
	boldyellow := color.Style{color.Bold, color.Yellow}
	boldblue := color.Style{color.Bold, color.Blue}
	boldgreen := color.Style{color.Bold, color.Green}

	colors := []color.Style{boldred, boldyellow, boldblue, boldgreen}

	fmt.Print(colors[linenr%len(colors)].Renderln(fmt.Sprintf("%03d: ", linenr)))
	for {
		_, err := io.ReadFull(os.Stdin, buf)
		if err != nil {
			if err == io.EOF {
				os.Exit(0)
			}
			panic(err)
		}
		timer.Reset(time.Second * 2)
		printmutex.Lock()
		if prev == '\n' {
			fmt.Print(colors[linenr%len(colors)].Renderln(fmt.Sprintf("%03d: ", linenr)))
		}
		fmt.Print(string(buf))
		if buf[0] == '\n' {
			linenr += 1
		}
		prev = buf[0]
		printmutex.Unlock()
	}
}
