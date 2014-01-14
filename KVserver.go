package main

import (
        "fmt";
        "net";
        "os";
        "io";
        //"string";
)

func main() {
        var (
                host = "127.0.0.1";
                port = "9090";
                remote = host + ":" + port;
                data = make([]byte, 1024);
        )

        var m map[string]string;
        m = make(map[string]string);

        fmt.Println("Starting Key Value Server ...");

        lis, error := net.Listen("tcp", remote);

        defer lis.Close();

        if error != nil {
                fmt.Printf("Error creating listener: %s\n", error );
                os.Exit(1);
        }

        for {
                var response string;
                var arg1 string;
                var arg2 string;
                var val string;
                var read = true;
                //var count = 1;

                //Accepting connection
                con, error := lis.Accept();
                if error != nil {
                        fmt.Printf("Error accepting connection: %s\n", error);
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

                                                if(error != nil){
                                                       if error != nil { fmt.Printf("Error recieving data from server: %s\n", error);}
                                                }
                                                arg1 = string(data[0:n]);

                                                val = m[arg1];
                                                if(val == ""){
                                                        val = string("UNDEFINED");
                                                }
                                                con.Write([]byte(val));
                                        } else if(response == string("PUT")){
                                                n, error = con.Read(data);
                                                if(error != nil){
                                                        if error != nil { fmt.Printf("Error recieving data from server: %s\n", error);}
                                                }
                                                arg1 = string(data[0:n]);

                                                n, error = con.Read(data);
                                                if(error != nil){
                                                        if error != nil { fmt.Printf("Error recieving data from server: %s\n", error);}
                                                }
                                                arg2 = string(data[0:n]);

                                                m[arg1] = arg2;
                                        } else if(response == string("DEL")){
                                                n, error = con.Read(data);

                                                if(error != nil){
                                                       if error != nil { fmt.Printf("Error recieving data from server: %s\n", error);}
                                                }
                                                arg1 = string(data[0:n]);

                                                delete(m, arg1);
                                        } else {
                                                fmt.Printf("Error accepting command: %s\n", response);
                                        }
                                        //count+
                                default:
                                        read = false;
                        }
                }

                con.Close();
        }
}