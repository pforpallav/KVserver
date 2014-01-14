package main

import (
        "fmt";
        "net";
        "os";
        "io";
)

func main() {
        var (
                host = "127.0.0.1";
                port = "9090";
                remote = host + ":" + port;
                data = make([]byte, 1024);
        )

        fmt.Println("Starting Key server...");

        lis, error := net.Listen("tcp", remote);

        defer lis.Close();

        if error != nil {
                fmt.Printf("Error creating listener: %s\n", error );
                os.Exit(1);
        }

        for {
                var response string;
                var arg1 string;
                var read = true;
                //var count = 1;

                //Accepting connection
                con, error := lis.Accept();
                if error != nil {
                        fmt.Printf("Error: Accepting data: %s\n", error);
                        os.Exit(2);
                }
                
                fmt.Printf("New Connection received from: %s \n", con.RemoteAddr());

                for read {
                        n, error := con.Read(data);
                        switch error {
                                case io.EOF:
                                        //fmt.Printf("Warning: End of data reached: %s \n", error);
                                        read = false;
                                case nil:
                                        //fmt.Println(string(data[0:n])); // Debug
                                        response = string(data[0:n]);
                                        if(response == string("GET")){
                                                n, error = con.Read(data);
                                                arg1 = string(data[0:n]);
                                                fmt.Printf("%s", arg1);
                                        } else if(response == string("PUT")){
                                                
                                        } else if(response == string("DEL")){
                                                
                                        }
                                        //count+
                                default:
                                        fmt.Printf("Error: Reading data : %s \n", error);
                                        read = false;
                        }
                }

                con.Close();
        }
}