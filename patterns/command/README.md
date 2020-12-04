# Command

Tipo de Patron: Comportamiento

### Objetivo
Esté patrón tiene por objetivo cumplir la dependencia entre funciones o métodos. Mediante
una interfaz `command` que tiene un método `execute()`, el cual se va a ejecutar cuando
nosotros decidamos. Es muy común ver este patrón cuando se necesita cierta dependencia entre
tareas. Suele verse mucho para pila de tareas, o mejor dicho pasar callbacks entre métodos.

### Caso de uso de ejemplo
Nuestro caso de ejemplo consta de una lista de tareas que iremos apilando. Una vez apiladas 
estas tareas se ejecutan y vemos el resultado. En Java, por ejemplo podemos ver este tipo de 
patrón en las clases Runneable, y así mismo lo podemos ver más fácilmente en la programación 
funcional como lo son los streams.