package main

import (
	"encoding/json"
	"fmt"
)

func main()  {
	//const jsonStreanm = `
	//{"name":"Ed", "Text":"knock knock"}
	//{ "Name" : "Sam" , "Text" : "Who's there?" }
	//{ "Name" : "Ed" , "Text" : "Go fmt." }
	//{ "Name" : "Sam" , "Text" : "Go fmt who?" }
	//{ "Name" : "Ed" , "Text" : "Go fmt yourself!" }
	//`
	//type Message struct {
	//	Name , Text string
	//}
	//dec := json.NewDecoder(strings.NewReader(jsonStreanm))
	//for {
	//	var m Message
	//	if err := dec.Decode(&m); err == io.EOF{
	//
	//		break
	//	} else {
	//		log.Fatal(err)
	//	}
	//	fmt.Printf("%s:%s\n", m.Name, m.Text)
	//}

	var jsonBlob = [] byte ( ` [ 
        { "Name" : "Platypus" , "Order" : "Monotremata" } , 
        { "Name" : "Quoll" ,     "Order" : "Dasyuromorphia" } 
    ] ` )
	type Animal struct {
		Name  string
		Order string
	}
	var animals [] Animal
	err := json. Unmarshal ( jsonBlob , & animals )
	if err != nil {
		fmt. Println ( "error:" , err )
	}
	fmt. Printf ( "%+v" , animals )

}