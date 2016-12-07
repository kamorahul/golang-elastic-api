package Utils

import (
	"net"
	"fmt"
	//"encoding/json"
	"strconv"
	"runtime"
	"github.com/fatih/structs"
	"strings"
)

func GetRecordSrv(service string) string{
	cName,addrs,err := net.LookupSRV("","",service)
	if(err != nil){
		return ""
	}
	if cName != ""{
		fmt.Println(cName)
	}
	dat1 := structs.Map(addrs[0])
	if err != nil{
		panic(err)
	}
	return "http://"+ strings.Trim(dat1["Target"].(string),".") + ":" + strconv.Itoa(int(dat1["Port"].(uint16)))

}


func GetNumCpu() int{
	num  := runtime.NumCPU()
	fmt.Println(num)
	return num
}
