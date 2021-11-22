package main

//Import
import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Main
func main() {
	loadEnv()

	secret := os.Getenv("aSECRET")

	fmt.Print("a varialbe es: ", secret)
	fmt.Printf("\n----------------------------")

	getUser()
	stan()
	fmt.Printf("\n----------------------------")
	CreateFile()
	fmt.Printf("\n----------------------------")

	ReadFile()
}

//Variables de ambiente
func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("erros fatal")
	}
}

// Argumentos de linea de comando
func getUser() {
	nombre := flag.String("nombre", "", "El nombre de la persona")
	edad := flag.Int("edad", 18, "La edad de la persona")
	flag.Parse()
	fmt.Println("\nTu nombre es:", *nombre)
	fmt.Println("\nTu edad es:", *edad)
}

// Standard
func stan() {

	var reader = bufio.NewReader(os.Stdin)
	fmt.Println("Your Name: ")
	message, _ := reader.ReadString('\n')

	fmt.Println("Welcome: ", message)
	fmt.Println("Standard Output")
	os.Stderr.WriteString("\n Standard Error")
	fmt.Printf("\n----------------------------")

}

// File I/O
func CreateFile() {

	fmt.Printf("\nWriting a file in Go")

	//Create File(Nombre)
	file, err := os.Create("test.txt")

	//Error, err: variable error
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	defer file.Close()

	// length(contenido)
	len, err := file.WriteString("File in GO (Golang), Victor Calderon")

	//Error
	if err != nil {
		log.Fatalf("failed writing to file: %s", err)
	}

	fmt.Printf("\nFile Name: %s", file.Name())
	fmt.Printf("\nLength: %d bytes", len)
}

func ReadFile() {

	fmt.Printf("\n\nReading a file in Go\n")
	fileName := "test.txt"

	data, err := ioutil.ReadFile("test.txt")

	if err != nil {
		log.Panicf("failed reading data from file: %s", err)
	}

	fmt.Printf("\nFile Name: %s", fileName)
	fmt.Printf("\nData: %s", data)
	// fmt.Println("\nData: ", data)

}
