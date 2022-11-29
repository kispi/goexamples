package main

import (
	"goexamples/localpackage" // 이처럼 "루트패키지이름/상대패키지경로" 순으로 import해야함. (./상대패키지경로 (X))
)

func main() {
	localpackage.Hello()
	localpackage.World()
}
