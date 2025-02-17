package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

//Hello devuelve un saludo para una persona especifica

func Hello(nombre string) (string,error) {
	if nombre == ""{
		return "", errors.New("Nombre no puede ser vacio")
	}
	message := fmt.Sprintf(randomFormat(), nombre)
	return message, nil
}


func randomFormat() string {
	formats := [] string {"Hola, %v! !Bienvenido!","¡Que bueno verte, %v!","¡Saludos, %s encantado de conocerte!"}
	return formats[rand.Intn(len(formats))]
}

func Hellos (nombres [] string)(map[string]string ,error){
	messages := make(map[string]string)
	for _,name := range nombres{
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}

		messages[name] = message
	}

	return messages, nil
}