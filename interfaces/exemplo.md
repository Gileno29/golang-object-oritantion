```go

type Reader interface {
    Read(p []byte) (n int, err error)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}

type ReadWriter interface{
    Reader
    Writer
}

```

```go
type Programador interface {
    Estudos
    SocialMidia
    Descansar(local string)
    Exercitar(type string) error
    Comer()
    Cafe()
    Trabalhar()
}


```