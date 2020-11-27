# Bridge

Tipo de Patron: Estructural

### Objetivo
El patrón brdge es útil para poder evitar un enlace permanente entre la abstracción y su implementación. 
La idea es poder transportar diferentes implementaciones de una misma funcionalidad y poder cambiarla en 
tiempo de ejecución.

### Caso de uso de ejemplo
El caso propuesto trata sobre una calculadora con distintas operaciones. En las interfaces podemos encontrarnos 
con la interfaz del puente y de la acción. Ambos serán implementados por bridge y por cada una de las acciones. 
El puente recibe justamente la interfaz y poder dentro se ejecuta la acción de la operación.