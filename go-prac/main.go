package main

import (
	"fmt"
	"reflect"
)

type MyEvent struct{
	Name string `json:"What is your name?"`
	Age int `json:"How old are you?"`
}

type MyResponse struct {
	Message string `json:"Answer"`
}

func MyHandler(event MyEvent)(MyResponse, error){
	return MyResponse{Message: fmt.Sprintf("%s is %d years old!", event.Name, event.Age)}, nil
}

func main(){
	Start(MyHandler)
}

func Receive() string {
	return `{"What is your name?":"Raj", "How old are you?":16}`
}

func Start(handler interface{}){
	t := reflect.TypeOf(handler)
	var inparams = make(map[string]reflect.Type)
	inpfields := []reflect.StructField{}
	ifield := make(map[string]string)


	if t.Kind() == reflect.Func{
		for i := 0; i<t.NumIn(); i++{
			inparams[t.In(i).Name()] = t.In(i)
			fmt.Println()
		}
	}
	for _, v := range inparams{
		for i := 0; i < v.NumField(); i++{
			inpfields = append(inpfields, v.Field(i))
			ifield[v.Field(i).Name] = v.Field(i).Tag.Get("json")
		}
	}

	// st := reflect.StructOf(inpfields);

	// stt := reflect.ValueOf(st).Interface();
	// fmt.Println(inparams)
	// fmt.Println(inpfields)
	// fmt.Println(st)
	// fmt.Println(ifield)


	// r := []byte(Receive());
	// err := json.Unmarshal(r, &stt)
	// if err != nil{
	// 	fmt.Println(err)
	// }

}








// type Person struct {
// 	Name    string
// 	Age     int
// 	Address string
// 	Email   string
// }

// type Info interface{
// 	GetInfo()
// 	GetName() string
// }

// func (p Person) GetInfo(){
// 	fmt.Printf("Name is %v and Age is %v", p.Name, p.Age)
// }

// func (p Person) GetName() string {
// 	return p.Name
// }

// func main() {
// 	r := Person{
// 		Name:    "Raj",
// 		Age:     16,
// 		Address: "Saidpur",
// 		Email:   "raj@r.com",
// 	}

// 	defer (func ()  {
// 		fmt.Println("THIS WILL RUN EVENT IF THERES A PANIC!!!")
// 	})()

// 	giveMe := make(chan string)

// 	go func (giveMe chan string)  {
// 		time.Sleep(time.Second)
// 		fmt.Println("this will run after 1 second")
// 		giveMe <- "GOLANG TO CHAN"
// 	}(giveMe)

// 	select{
// 		case g:= <- giveMe:
// 			fmt.Println(g)
// 		default :
// 			fmt.Println("DONE")
// 	}
	
// 	ReceiveS(r)
// 	fmt.Println("THIS SHOULD RUN LAST")

// 	time.Sleep(4 * time.Second)
// }

// func ReceiveS(x interface{}){
// 	res, ok := x.(Person)
// 	if !ok {
// 		return
// 	}

// 	// panic("THIS WILL TERMINATE THE PROGRAM")
// 	fmt.Println(res.Name)
// }