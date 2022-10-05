## Go modules

* En `Node` podemos usar `npm` para generar un árbol de dependencias de nuestros proyectos, en `Golang` existen los módulos con los cuales podremos definir en que paquetes tenemos dependencias para poder instalarlas:

Algunos comando del cli de Golang:

* Inicializar un modulo
```
go mod init github.com/username/module 
```

* Descargar una dependencia
```
go get github.com/donvito/hellomod
```

* Limpiar dependencias sin utilizar
```
go mod tidy
```

* Informacion de los modulos cacheados
```
 go mod download -json
```

* Podemos utilizar “Alias” al momento de importar un paquete de go si asi lo necesitáramos
```
import (
	"github.com/donvito/hellomod"
	hellomod2 "github.com/donvito/hellomod/v2"
)
```

