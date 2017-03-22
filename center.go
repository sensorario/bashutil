package bashutil

import (
	"fmt"
	"github.com/nsf/termbox-go"
)

func Center(s string) {
	if err := termbox.Init(); err != nil {
		panic(err)
	}
	w, _ := termbox.Size()
	termbox.Close()
	fmt.Printf(
		fmt.Sprintf("%%-%ds", w/2),
		fmt.Sprintf(fmt.Sprintf("%%%ds\n", w/2+len(s)/2), s),
	)
}
