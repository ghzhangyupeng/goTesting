package main

import (
	"os"
	"bufio"
	"encoding/gob"
	"fmt"

)

func main()  {
	ref := make(map[int32][]float32)
	f, err:= os.Open("./iclick_custom_rfid.gob")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	r := bufio.NewReader(f)
	dec := gob.NewDecoder(r)
	if err := dec.Decode(&ref); err != nil {
		panic(err)
	}
	fmt.Println(ref)
}





