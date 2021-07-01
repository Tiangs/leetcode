package main

import "net/http"
import "io/ioutil"
import "fmt"
func main(){
	resp,err:= http.Get("https://www.baidu.com")
	if err!=nil{
		panic(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK{
		fmt.Println("error:status code",resp.StatusCode)
	}
	all,err:=ioutil.ReadAll(resp.Body)
	if err!=nil{
		panic(err)
	}
	fmt.Printf("%s\n",all)
	

}

