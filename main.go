package main

import (
	"os"
	"log"
	"io"
	"strings"
	"github.com/uia-worker/misc/conv"
)	

func main() {
	src, err := os.Open("table.csv")
	//src, err := os.Open("/home/janisg/minyr/kjevik-temp-celsius-20220318-20230318.csv")
	if err != nil {
        	log.Fatal(err)
	}
	defer src.Close()
        log.Println(src)
        
	
	var buffer []byte
	var linebuf []byte // nil
	buffer = make([]byte, 1)
        bytesCount := 0
	for {
		_, err := src.Read(buffer)
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}

		bytesCount++
		//log.Printf("%c ", buffer[:n])
		if buffer[0] == 0x0A {
	           log.Println(string(linebuf))
		   // Her
		   elementArray := strings.Split(string(linebuf), ";")
		   if len(elementArray) > 3 {
			 celsius := elementArray[3]
			 fahr := conv.CelsiusToFahrenheit(celsius)
		         log.Println(elementArray[3])
	   	   }
                   linebuf = nil		   
		} else {
                   linebuf = append(linebuf, buffer[0])
		}	
		//log.Println(string(linebuf))
		if err == io.EOF {
			break
		}
	}

}
