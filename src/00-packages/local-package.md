# How create packages

- To use multiple packages you need create a module!
```
cd <folder>
go mod init <module-name>
```

- Create file "main.go"

```
> Compile all project on single project
go build 

> Execute project
./<module-name>

```

- Create another folder and file ./auxiliar/auxiliar.go
```
> On : ./<module-name>/auxiliar/auxiliar.go

package auxiliar
import "fmt"

func Escrever(){
    fmt.Println("Write line on Auxiliar")
}

> On: ./<module-name>/main.go

package main
import(
    "<module-name>/auxiliar"
    "fmt"
    )

func main(){
    fmt.Println("Write line Main)
    auxiliar.Escrever()
}


> ON: Cmd
go run main.go

```

- Important Notes
Upper letter on functions in packages to indicate public and use  therefore.

First letter upper its public and we can use everywhere.<br>
First letter lower her scope is only inside file.<br>
