package process

import (
	"fmt"
	"time"
)

type Process struct {
	Id uint64
	Count uint64
}

func (p *Process) RunProcess(c chan uint64) {
	for {
		select {
			case msg := <- c:
				if msg == p.Id {
					return
				} else {
					c <- msg
				}
			default:
				fmt.Println("default")
				fmt.Printf("%d : %d\n",p.Id, p.Count)
				(*p).Count = (*p).Count + 1
				time.Sleep(time.Millisecond * 500)
		}
	}
}