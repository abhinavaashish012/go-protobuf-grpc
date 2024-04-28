package main

/*
To generate .pb files
//protoc -I=$SRC_DIR --go_out=$DST_DIR $SRC_DIR/addressbook.proto
//protoc -I=proto --go_out=pb proto/processor_message.proto
https://protobuf.dev/getting-started/gotutorial/

go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
*/

// github.com/golang/protobuf/proto v1.5.4
import (
	"fmt"
	"os"
	"pcbook/sample"
	"pcbook/serializer"
)

func main() {
	fmt.Println("Hello form PCBook.")

	laptop := sample.NewLaptop()

	err := serializer.WriteProtobufToBinaryFile(laptop, "./tmp/laptop.bin")
	if err != nil {
		fmt.Println("Error writing laptop.bin === ", err.Error())
		return
	}

	err = serializer.WriteProtobufToJSONFile(laptop, "./tmp/my_laptop.json")
	if err != nil {
		fmt.Println("Error writing my_laptop.json === ", err.Error())
		return
	}
	stat1, _ := os.Stat("./tmp/laptop.bin")
	fmt.Println("Binary Size : ", stat1.Size())

	stat2, _ := os.Stat("./tmp/my_laptop.json")
	fmt.Println("Json Size : ", stat2.Size())
	err = serializer.ReadProtobufFromBinaryFile("./tmp/my_laptop.json", laptop)
}
