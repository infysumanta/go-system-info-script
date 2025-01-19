package main

import (
	"fmt"

	"github.com/infysumanta/go-system-info-script/internal/sysinfo"
)


func main(){
	info := sysinfo.GetSystemInfo()
	fmt.Println(info)
}