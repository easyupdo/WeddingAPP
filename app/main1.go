package main

import ("fmt"
         "net/http"
         "log"
         "io/ioutil"
         "reflect"
         "strconv"
         "encoding/json"
         "io"
)


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


// post renturn json
type Ret struct{
    Status bool;
    Data string;
}


//say hello to the world
func sayHello2(w http.ResponseWriter, r *http.Request) {
   //n, err := fmt.Fprintln(w, "hello world")
   
   fmt.Println("sayHello2 Method:",r.Method);
   
   r.ParseForm();
   fmt.Println("Form:",r.Form);

   var name string ;
   fmt.Println("Type:",reflect.TypeOf(name));
   for _,str := range r.Form["name"] {
    name += str;
   }
   fmt.Println("name:",name);
   var content string ;
   for _,str := range r.Form["notes"]{
      content += str;
   }
   var num string ;
   for _,str := range r.Form["guest"]{
      num += str;
   }

   fmt.Println("guest_name:",name);
   fmt.Println("content:",content);
   fmt.Println("guest_num:",num);

   num_int, err:= strconv.Atoi(num);
   if err == nil{
    log.Println(err);
   }
    
   SaveSQL(name,content,num_int);
   
   w.Header().Set("content-type","text/json");
   ret := Ret{Status:true,Data:"sql exec ok!"};
   ret_json,jerr := json.Marshal(ret);
    
   fmt.Println("Jerr:",jerr);

   fmt.Println("Ret:",ret);
   fmt.Println("Ret_Json:",string(ret_json));

   io.WriteString(w,string(ret_json));

	//_, _ = w.Write(body)
}




func TestAjax(w http.ResponseWriter, r *http.Request) {
  
   fmt.Println("TextAjax Method:",r.Method);
   fmt.Println("TestAjax Body:",r.Body);
   fmt.Println("TextAjax R:",r);
   r.ParseForm();
   fmt.Println("TestAjax Body:",r.Form);
}




func main() {
  test();

   //http.Handle("/", http.FileServer(http.Dir("/Users/jay/Personal/Study/emsdk")))
   //http.Handle("/", http.FileServer(http.Dir("/Users/jay/Personal/Study/nodejs/Electron")))
   http.Handle("/", http.FileServer(http.Dir("/home/jay/Work/wedding_app/wedding_app_data")))
  //1.注册一个处理器函数
	//http.HandleFunc("/", sayHello)
   http.HandleFunc("/a", sayHello1)
   http.HandleFunc("/rsvp", sayHello2)

   http.HandleFunc("/ajax", TestAjax)

   //var a1,b1 int;
   //fmt.Scanln(&a1,&b1);
   //fmt.Println("a1:",a1,"b1:",b1);
  

	//2.设置监听的TCP地址并启动服务
	//参数1:TCP地址(IP+Port)
	//参数2:handler handler参数一般会设为nil，此时会使用DefaultServeMux。
  fmt.Println("Server Start And Listern:192.168.40.104:5555");
	err := http.ListenAndServe("172.20.1.60:5555", nil)
	if err != nil {
		fmt.Printf("http.ListenAndServe()函数执行错误,错误为:%v\n", err)
		return
	}
}
