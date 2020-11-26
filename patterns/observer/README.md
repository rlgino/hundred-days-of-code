# Observer

Tipo de Patron: De Comportamiento

### Objetivo
El objetivo de este patrón es tener una dependencia con muy bajo acoplamiento. Por lo que es útil 
para poder ejecutar distintas acciones en un punto particular. Este patrón tiene varios actores:

* Sujeto(Contenedor): Este actor contendrá el listado de observadores y tendrá, en nuestro caso, 
3 funciones importantes: subscribir un patron, desuscribir un patron y ejecutar las acciones definidas.
* Observador: Es la interfaz que expondremos para adaptar nuestros patrones y mediante las cuales 
se definiran estos.
* Observador concreto: Es la implementación particular del observador. En donde se va a definir 
el comportamiento de las funciones definidas en el observador.

La principal ventaja de esto, es el bajo acoplamiento en sus pares y la dependencia para hacer las cosas.

### Caso de uso de ejemplo
El ejemplo consta de las acciones en una venta, en nuestro caso van a ser:
* Calculo de monto
* Cambio de estado
* Facturación y cambio de ID
Cada observador es justamente la acción de una venta y estos se ejecutan a la hora de hacer esta venta.