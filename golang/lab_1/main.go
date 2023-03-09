package main

import (
	"fmt"
	"os"
	"io"
	"os/exec"
)

func main(){
	val,err := os.Getwd();
	if(err!=nil){
		panic(err);
	}else{
		fmt.Println("[*]Текущая директория:",val);
		err = os.Mkdir("bar",0776);
		if (err!=nil){
			panic(err);
		}
		err = os.Chdir("./bar");
		if(err!=nil){
			panic(err);
		}
		val,err := os.Getwd();
		if(err!=nil){panic(err);}
		fmt.Println("[*]Сменил директорию:",val);
		if(len(os.Args)<2){
			os.Remove("./../bar");
			panic("Ой ой забыл указать имя файла");
		}
		file,err := os.Create("krik");
		if(err!=nil){panic(err);}
		data := make([]byte,64);
		krik,_ := os.Open("./../krik");
		for{
			_,err:=krik.Read(data);
			file.Write(data);
			if(err==io.EOF){
				file.Close();
				krik.Close();
				break;
			}
		}
		os.Chmod("./bar/krik",0776);
		cmd:=exec.Command("./bar/krik","krik");
		err = cmd.Run();
		if(err!=nil){panic(err);}
		fmt.Println("[*]Создал файл:",file.Name());
		//MR ROBOTS =)
	}
}
