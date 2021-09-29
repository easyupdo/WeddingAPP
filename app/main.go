package main

/*
#cgo CFLAGS: -I /Users/jay/Personal/Study/go/src/app
#cgo LDFLAGS: -L /Users/jay/Personal/Study/go/src/app -lhello
#include "hello.h"
#include <stdio.h>
//#include "hello.h"
//void Hello(){
//    printf("Hello This is hello.c\n");
//}
*/
import "C"
import ("fmt"
         "time"
         "net/http"
         "log"
         "io/ioutil"
)

var b int =100;

func Hello1(pa int)int {
	fmt.Println("Hello func");
	fmt.Println("pa:",pa);
	if true {
		fmt.Println("if true");
	}
	return 999;
}

func Ret(pa1 string,pa2 int,f func())(string,int)  {
	f();
	return pa1,pa2
}

func callback()  {
	fmt.Println("This is callback func!");
}

type ST struct{
	sa int;
	strb string;
}


type GG interface{
   ggtest()
}

func (st ST) ggtest()  {
   fmt.Println("GG interface ggtest");
}

func Say(str string)  {
   for i:=0;i<100;i++{
      fmt.Println("str:",str);
      time.Sleep(time.Duration(5)*time.Millisecond)
   }
}

//say hello to the world
func sayHello(w http.ResponseWriter, r *http.Request) {
   //n, err := fmt.Fprintln(w, "hello world")
   
   data1 , err2:= ioutil.ReadFile("./index.html");
   if err2 == nil{
      log.Println(err2);
   }
   fmt.Println(string(data1))
   fmt.Println("Method",r.Method);

   _, _ = w.Write(data1)
   
}

//say hello to the world
func sayHello1(w http.ResponseWriter, r *http.Request) {
   //n, err := fmt.Fprintln(w, "hello world")
   
   resp,err3 :=http.Get("http://www.baidu.com");
   if err3 == nil{
      log.Println(err3);
   }
   //fmt.Println(resp);
   body,berr:=ioutil.ReadAll(resp.Body);
   if berr==nil{
      log.Println(berr);
   }
   fmt.Println(string(body));

	_, _ = w.Write(body)
}


//say hello to the world
func sayHello2(w http.ResponseWriter, r *http.Request) {
   //n, err := fmt.Fprintln(w, "hello world")
   
   fmt.Println("sayHello2 Method:",r.Method);
   
   r.ParseForm();
   fmt.Println("Form:",r.Form);
   fmt.Println("username:",r.Form["the_heart_name"]);
   fmt.Println("content:",r.Form["the_heart_content"]);

	//_, _ = w.Write(body)
}


func main() {
  test();
   fmt.Println("Hello, World!");
   ri := Hello1(888);
   fmt.Println(ri);
   a:=100;
   fmt.Println("println");
   fmt.Println(a);
   fmt.Println("a:",a);
   fmt.Println(1,2);

   var iArr [5]int = [5]int{1, 2, 3, 4}
   for key,value := range iArr{
	   fmt.Println(key,value);
	   //fmt.Println(value);
   }

   for i := 0;i<len(iArr);i++{
	   fmt.Println("i",i);
   }
   ra1,ra2:=Ret("abc",222,callback);
   fmt.Println(ra1,ra2);

   i_ptr := &ri;
   fmt.Println("i_ptr:",*i_ptr);

   sti := ST{12345,"abcdef"}
   fmt.Println(sti.sa);
   fmt.Println(sti.strb);

   gst := new(ST);
   gst.ggtest();

   //go Say("Hello");
   //Say("World");

   //chan

   ch := make(chan int)
   go func ()  {
      ch <- 100;
   }()
   rch := <- ch;
   fmt.Println("rch:",rch);
   log.Printf("Log")
   //log.Fatal("Log Fatal")

 

   http.Handle("/", http.FileServer(http.Dir("/Users/jay/Personal/Study/nodejs/HTML/wedding/")))
   //http.Handle("/", http.FileServer(http.Dir("/Users/jay/Personal/Study/nodejs/HTML/gallrey")))
  //1.注册一个处理器函数
	//http.HandleFunc("/", sayHello)
   http.HandleFunc("/a", sayHello1)
   http.HandleFunc("/form", sayHello2)

   C.Hello();

   //var a1,b1 int;
   //fmt.Scanln(&a1,&b1);
   //fmt.Println("a1:",a1,"b1:",b1);
  

	//2.设置监听的TCP地址并启动服务
	//参数1:TCP地址(IP+Port)
	//参数2:handler handler参数一般会设为nil，此时会使用DefaultServeMux。
	err := http.ListenAndServe("192.168.40.104:5555", nil)
	if err != nil {
		fmt.Printf("http.ListenAndServe()函数执行错误,错误为:%v\n", err)
		return
	}
}
