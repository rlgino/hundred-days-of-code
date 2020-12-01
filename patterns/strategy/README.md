# Strategy

Tipo de Patron: Comportamiento

### Objetivo
El objetivo de este patron es el hecho de cambiar el comportamiento, en base a un contexto que 
posiblemente tenga un estado, todo esto en tiempo de ejecución. Lo que hace posible, escalar 
estos comportamientos en un futuro y que sea trasparente para los clientes que las ejecuten.

### Caso de uso de ejemplo
En el ejemplo vamos a usar un contexto que tiene un arreglo de numeros, y un algoritmo de 
busqueda. Como sabemos hay diferentes algoritmos de busqueda y en base a un parametro 
ingresado por el usuario, nosotros elegimos si usar un algoritmo de busqueda u otro. Fijarse 
que no tenemos las implementaciones de estos algoritmos porque basicamente nos importa como 
funciona el patron de diseño.