# Adapter

Tipo de Patron: Estructural

### Objetivo
El patrón de diseño Adapter, como su nombre lo índica funciona para adaptar ciertas funcionalidades 
en lugares donde las necesitamos, pero su interfaz no es compatible de la misma forma. El concepto es 
hacer un emboltorio (wrapper o adaptador) para poder utilizar la interfaz que necesitamos.

### Caso de uso de ejemplo
En nuestro caso planteamos la utilización de una caja de ahorro que por dentro ahorra dolares, pero 
nosotros necesitamos pesos argentinos. Para poder utilizar esos dolares, vamos a hacer un adaptador 
que en el medio trasforme esos dolares a pesos argentinos.