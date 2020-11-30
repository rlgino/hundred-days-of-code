# Factory

Tipo de Patron: Creacional

### Objetivo
Esté patrón al ser una fábrica justamente, se útiliza para la creación de familias de objetos, 
simplifica la creación y la desacopla haciendola un poco más dínamica en tiempo de ejecución. 
Las bibliotecas para crear interfaces gráficas suelen utilizar este patrón y cada familia sería 
un sistema operativo distinto. Así pues, el usuario declara un Botón, pero de forma más interna 
lo que está creando es un BotónWindows o un BotónLinux, por ejemplo.

### Caso de uso de ejemplo
En nuestro caso, proponemos un cuadro de dialogo, que puede funcionar tanto en web como en mobile. 
Nuestro factory puede ser un Factory Web o un Factory Mobile, dentro de cada una se crea un nuevo 
producto(un boton en nuestro caso). Y esto hace desacoplar un poco más nuestras distintas 
implementaciones haciendo menor la repetición de código.