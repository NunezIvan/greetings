#Saludos en Go

Ester paquete proporciona una forma simple de obtener saludos personalizados

##Instalacion
Ejecuta el siguiente comando para instalar el paquete

```bash
go get -u github.com/NunezIvan/greetings
```
##Uso
Aqui tienes un ejemplo de como utilizar el paquete en tu codigo

```go
    package main
    import(
        "fmt"
        "github/NunezIvan/greetings"
    )
    func main(){
        message, err := greetings.Hello("Alex")
        if err != nil{
            fmt.Println("Ocurrio un error: ", err)
            return
        }
        fmt.Println(message)
    }
```
