package main

import (
        "fmt";
        "net";
        "os";
        //"strings";
)

func main() {
        var (
                host = "127.0.0.1";
                port = "9090";
                remote = host + ":" + port;
                msg string;
                data = make([]byte, 1024);
                val string;
        )

        con, error := net.Dial("tcp", remote);
        defer con.Close();
        if error != nil { fmt.Printf("Host not found: %s\n", error ); os.Exit(1); }

        n, error := fmt.Scanf("%s", &msg);
        
        for(n > 0){
                if(msg == string("GET")){
                        in, error := con.Write([]byte(msg));
                        if error != nil { fmt.Printf("Error sending data: %s, in: %d\n", error, in ); os.Exit(2); }
                        
                        n, error = fmt.Scanf("%s", &msg);
                        in, error = con.Write([]byte(msg));
                        if error != nil { fmt.Printf("Error sending data: %s, in: %d\n", error, in ); os.Exit(2); }
                        
                        n, error = con.Read(data);
                        if error != nil { fmt.Printf("Error recieving data from server: %s\n", error); os.Exit(2); }
                        val = string(data[0:n]);

                        fmt.Printf("Value for key %s is: %s\n", msg, val);
                } else if(msg == string("PUT")){
                        in, error := con.Write([]byte(msg));
                        if error != nil { fmt.Printf("Error sending data: %s, in: %d\n", error, in ); os.Exit(2); }
                        
                        n, error = fmt.Scanf("%s", &msg);
                        in, error = con.Write([]byte(msg));
                        if error != nil { fmt.Printf("Error sending data: %s, in: %d\n", error, in ); os.Exit(2); }
                        
                        n, error = fmt.Scanf("%s", &msg);
                        in, error = con.Write([]byte(msg));
                        if error != nil { fmt.Printf("Error sending data: %s, in: %d\n", error, in ); os.Exit(2); }

                        /*
                        n, error = con.Read(data);
                        if error != nil { fmt.Printf("Error recieving data from server: %s\n", error); os.Exit(2); }
                        val = string(data[0:n]);

                        fmt.Printf("Value for key %s is: %s", msg, val);
                        */
                } else if(msg == string("DEL")){
                        in, error := con.Write([]byte(msg));
                        if error != nil { fmt.Printf("Error sending data: %s, in: %d\n", error, in ); os.Exit(2); } 

                        n, error = fmt.Scanf("%s", &msg);
                        in, error = con.Write([]byte(msg));
                        if error != nil { fmt.Printf("Error sending data: %s, in: %d\n", error, in ); os.Exit(2); }
                } else {
                        fmt.Printf("Command not recognized!\n");
                }
                n, error = fmt.Scanf("%s", &msg);
        }

        //fmt.Println("Connection OK");

}