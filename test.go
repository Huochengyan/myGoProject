package test

type Books struct {
	id int
	name string
	pass string
	address string
}

func test(){

  //fmt.Println("asdfa")
  //b :=7
  //for i :=0; i<10 ;i++  {
  //  if b==i{
  //    fmt.Println("test")
  //    continue
  //  }
  // fmt.Println(i)
  //}

  //var  cmap map[string]string /* create map */
  //cmap=make(map[string]string)
  //cmap["name"]="go"
  //cmap["gender"]="f1111"
  //cmap["f1"]="f1111"

  //ok :=cmap["f1"]
  //if(ok!="") {
  // fmt.Println("f1:", ok)
  //}else{
  // fmt.Println("is null")
  //}

   //const a, b, c = 1, false, "str" //多重赋值
   //const d=10
   //println(a, b, c,d)

   //var arr=[...]string{"1","Abc"}
   //var arr=make([]string,0)
   //println(len(arr))
   //arr=append(arr, "111");
   //for  i :=0 ;i<len(arr);i++{
   //   println("e:"+arr[i])
   //}

   //for i :=0;i<10000;i++{
	//	arr=append(arr, strconv.Itoa(i));
	//}
   // fmt.Println("arr len:",len(arr))
	//var bookinfo Books
	//bookinfo.id=1;
	//bookinfo.name="10W why?"
	//bookinfo.address="addddd"
	//bookinfo.pass="11111111"
	//println(bookinfo.pass)
	//ftest(bookinfo)
    // int type default value 0
	var a int
   	println("default:",a)

}

func  ftest(books Books)  {
	println(books.pass)
	println("go....................")
    f2(0)
}

func f2(i int)  {
	i++
	println(i)
	if i<10 {
		f2(i);
	}
	return
}





