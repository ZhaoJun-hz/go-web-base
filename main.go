package main

import "github.com/ZhaoJun-hz/go-web-base/cmd"

func main() {
	defer cmd.Clean()
	cmd.Start()
}
