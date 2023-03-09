package main

import (
	"fmt"
	"os"
	"os/exec"
	"io"
)

func main(){
	fmt.Println("[*]Ну поехали");
	file,err := os.Open("./../main");
	if(err!=nil){panic(err);os.Exit(1);}
	fileNew,err := os.Create("./main");
	if(err!=nil){panic(err);os.Exit(1)};
	data := make([]byte,64);
	for{
		_,err:=file.Read(data);
		fileNew.Write(data);
		if(err==io.EOF){
			file.Close();
			fileNew.Close();
			break;
		}
	}
	os.Chmod("main",0776);
	cmd:=exec.Command("./main","krik");
	err = cmd.Run();
	if(err!=nil){panic(err);}
	fmt.Println("[*]Еле добрались");
}
