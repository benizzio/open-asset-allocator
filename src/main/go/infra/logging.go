package infra

import (
	"flag"
	"fmt"
	"github.com/golang/glog"
)

func ConfigLogger() bool {

	err := flag.Set("logtostderr", "true")
	if err != nil {
		fmt.Println("Error setting logtostderr flag: ", err)
		return true
	}
	err = flag.Set("alsologtostderr", "true")
	if err != nil {
		fmt.Println("Error setting alsologtostderr flag: ", err)
		return true
	}
	err = flag.Set("stderrthreshold", "INFO")
	if err != nil {
		fmt.Println("Error setting stderrthreshold flag: ", err)
		return true
	}

	flag.Parse()
	defer glog.Flush()

	return false
}
