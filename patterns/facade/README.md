# Facade

Tipo de Patron: Estructural

### Objetivo
Esté patrón tiene por objetivo poder proporcionar una interfaz sencilla que por detras
la implementación oculta una complejidad superior. Esto nos permite poder manejar 
métodos complejos y exponerlos de una manera más simple para los clientes que la utilicen

### Caso de uso de ejemplo
Nuestro caso trata de un finalizador de ventas. Nuestra fachada recibe justamente una
venta y por detrás ejecuta un listado de funciones para está venta. Lo que vemos es que
la interfaz recibe solo una venta y devuelve un error, pero, por dentro vemos que la 
complejidad es alta.