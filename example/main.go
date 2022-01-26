package main

import (
	"log"
	"net/http"
	"time"

	goballast "github.com/userpro/goballast"

	_ "github.com/mkevac/debugcharts"
)

func businessAllocMemory(a []int) {
	for i := range a {
		a[i]++
	}
}

func allocTest() {
	for {
		time.Sleep(time.Millisecond * 200)
		a := make([]int, 0)
		for i := 0; i < 1000000; i++ {
			a = append(a, i)
		}
		businessAllocMemory(a)
	}
}

func main() {
	goballast.NewWithDebug(2 << 30)

	go allocTest()

	log.Println(http.ListenAndServe("0.0.0.0:6060", nil))
}
