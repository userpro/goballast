# goballast

golang ballast library, reduce GC frequency

# Usage

## Run example

~~~bash
cd example
go mod tidy
go run main.go
~~~

open `127.0.0.1:6060/debug/charts/` show dashboard

## Demo

~~~go
import (
	"github.com/userpro/goballast"
)

func main() {
    // Set GC trigger memory usage target 4G
    goballast.New(2 << 30)
    // goballast.NewWithDebug(2 << 30) // print runtime debug info

    // TODO something
}
~~~

# Reference

* [Go memory ballast: How I learnt to stop worrying and love the heap](https://blog.twitch.tv/en/2019/04/10/go-memory-ballast-how-i-learnt-to-stop-worrying-and-love-the-heap/) 
* [proposal: runtime: add a mechanism for specifying a minimum target heap size](https://github.com/golang/go/issues/23044)

