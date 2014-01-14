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
        )

        con, error := net.Dial("tcp", remote);
        defer con.Close();
        if error != nil { fmt.Printf("Host not found: %s\n", error ); os.Exit(1); }

        n, error := fmt.Scanf("%s", &msg);
        
        for(n > 0){
                in, error := con.Write([]byte(msg));
                if error != nil { fmt.Printf("Error sending data: %s, in: %d\n", error, in ); os.Exit(2); }
                n, error = fmt.Scanf("%s", &msg);
        }

        //fmt.Println("Connection OK");

}