package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	//cli()
	//WebServ()
}

func WebServ() {
	http.HandleFunc("/", WebServiceHandler)
	http.ListenAndServe(":3000", nil)
}

func WebServiceHandler(w http.ResponseWriter, r *http.Request) {
	f, _ := os.Open("./menu.txt")
	io.Copy(w, f)
}

func cli() {
	fmt.Println("what would u like me to scream")
	in := bufio.NewReader(os.Stdin)
	s, _ := in.ReadString('\n')
	s = strings.TrimSpace(s)
	s = strings.ToUpper(s)
	fmt.Println(s + "!")
}
