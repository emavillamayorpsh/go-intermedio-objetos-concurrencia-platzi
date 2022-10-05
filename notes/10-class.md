## Interfaces

* En lenguajes como Typescript o Java es necesario definir los métodos de la interfaz para poder implementarla y la implementación es explícita.
* En cambio en Go, es suficiente definir los métodos para implementar la interfaz y la implementación es implícita.
* En el polimorfismo se utiliza una interfaz (o una clase base) para determinar en tiempo de ejecución el método a utilizar, accediéndolo por medio de la interfaz en vez de hacerlo a través de un objeto en particular, porque en este último caso el método se determina en tiempo de compilación (esto no es polimorfismo).

```
// no se conoce p hasta el momento en que 
// se ejecuta esta función
// podría ser un FullTimeEmployee
// o un TemporaryEmployee
func getMessage(p PrintInfo) {
  fmt.Println(p.getMessage)
}

// aquí ya se conoce el método a usar
// antes de compilar
ftEmployee.getMessage()
```

Entonces con polimorfismo, podrías por ejemplo, estar recibiendo desde una BD o un JSON desde un front, los datos de un empleado en tiempo de ejecución y determinar en ese momento cual es el método apropiado para ese empleado en particular.




