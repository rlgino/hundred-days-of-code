# Singleton

Tipo de Patron: Creacional

### Objetivo
Este patrón de diseño viene a solucionar el tener muchas instancias de un mismo componente. 
Con esto nos ahorrariamos estar instanciando, y por lo tanto estar ocupando memoria 
en cosas que no cambian.

### Caso de uso de ejemplo
En este ejemplo, vamos a estar utilizando una conección a una base de datos. Esta conección será 
instanciada una vez y no tendrá más instancias, porque no es necesario, debido a que la conección 
debe permanecer abierta en todo momento y no va a variar.