## Constructores

* Los constructores permiten la instanciación de una clase a un objeto, asimismo permite
definir propiedades predefinidas.
* En Go podemos utilizar funciones que puedan crear structs con propiedades que nosotros
pasamos como parámetros.


### Formas de Instanciar un objeto en go

1. Crear un struct con valores `zero` por defecto:
```
e  := Employee{}
```

2. Asignarle valores a las propiedades.
```
e2 := Employee{ id: 1, name: "Name", vacation: true }
```

3. Usando la palabra reservada `new`.
    * Regresa una referencia a la instancia creada, para acceder el valor usamos `*e3`
    ```
    e3 := new(Employee)
    ```

4. Crear una función que actúe como método constructor.
    * La función regresa con & para que regrese la referencia y no una copia, de la struct.
    ```
    func NewEmployee(id int, name string, vacation bool) *Employee {
    	return &Employee{
    		id:       id,
    		name:     name,
    		vacation: vacation,
    	}
    }
    ...
    e4 := NewEmployee(1, "Name 2", true)
    ```



















