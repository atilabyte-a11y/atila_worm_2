package main 


import (
 
"fmt"
"time"

)



func main(){


port :=  []int {22, 2222, 8022}


ip  := [] string { 


"192.168",

"104.236",


}


for  ind    ,_  :=   range  ip{






for  _,  p :=  range port{


for ii :=  1 ;  ii  < 255 ; ii++{

for  i :=  1 ; i <  255; i++{


ips := fmt.Sprintf(ip[ind]  +  ".%d.%d:%d"  , i  , ii , p)



go fun(ips)


}
}
}
}


for {

time.Sleep( 1000  * time.Second)

}


}

