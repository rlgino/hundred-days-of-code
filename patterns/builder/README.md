# Builder

Tipo de Patron: Creacional

### Objetivo
El patron builder viene a solucionar el hecho de tener un constructor con muchos 
parametros, proporcionandonos una interfaz con métodos para poder crearlos. Justamente 
este patrón consta de una interfaz builder que va a ser la que construya el producto. 
Este builder tiene que tener metodos que reciban los parametros del constructor 
respectivamente.
Además tiene un método Build que devolverá el producto creado con su estado formado.

### Caso de uso de ejemplo
Nuestro ejemplo queremos construir una estructura Persona. Para ello vamos a crear 
un builder de persona justamente. Este builder va a tener varios métodos para crear 
nuestra persona con nombre, edad y dirección. Notese que si alguno de esos valores no 
son pasados, nuestro producto tendrá los valores default.