package main
import "fmt"
/*
type book struct{
title string
author string
price int
}
*/
func main(){
/*
var b1 book
b1.title = "SALATUR RASOOL (SM)"
b1.author = "Dr Muhammad Asadullah Al Ghalib"
b1.price = 200

fmt.Println(b1)
fmt.Println(b1.author)
*/

                                //annonymous struct
b1 := struct{
title string
author string
price int
}{
title: "SALATUR RASOOL (SM)",
author: "Dr Muhammad Asadullah Al Ghalib",
price: 200,
}
fmt.Println(b1)
}