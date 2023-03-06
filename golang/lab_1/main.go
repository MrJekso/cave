package main

import (
	"fmt"
	"os"
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
		file,err := os.Create(os.Args[1]);
		if(err!=nil){panic(err);}
		fmt.Println("[*]Создал файл:",file.Name());
		//MR ROBOTS =)
	}
}
