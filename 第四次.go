package main

import (
    "fmt"
    "time"
)

func main()  {
    var t int64
     m:=9
    a :=make([]int64,0,1)
    l1:    fmt.Print("请输入时间戳：")
        fmt.Scanf("%d",&t)
        a=append(a,t)
        fmt.Println("imput ok!")
        fmt.Printf("是否继续输入\n(1 for yes,0 for no)")
        fmt.Scanf("%d",&m)
        if(m==1) {
            goto l1
        } else{
            for _,q:=range a {
                fmt.Println("the results are:\n",time.Unix(int64(q), 0))
            }
        }

}