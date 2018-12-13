package main

import (
    "fmt"
    "reflect"
)

type handler struct{
    function interface{}
    args interface{}
}

func exeHandler(h handler){
    f := reflect.ValueOf(h.function)
    if f.Kind() != reflect.Func {
        panic("f must be funnc.")
    }
    f.Call([]reflect.Value{reflect.ValueOf(h.args)})
}

func main(){
    f := func(str string){
        fmt.Println(str)
    }
    h := handler{function: f, args: "hello"}
    exeHandler(h)
}