
## Test Caso 1

Lectura y escritura de un fichero empleando dos tareas distintas, 
una para la lectura y otra para la escritura. 

Ambas tareas son independientes, 
y se comunican entre ellas mediante una estructua FIFO que soporte escritura y lectura de manera concurrente.

En este caso debido a que se implementa en el lenguaje Go, se emplea un Buffered Channel 
para la comunicación de las tareas.
 
En la tarea de escritura se incluye un delay aleatorio simulando retardo en la escritura, para comprobar que la tarea de lectura finaliza antes y cierra el fichero al terminar.

Al invocar el programa, se deben pasar dos parámetros, el nombre del fichero de entrada y el nombre del fichero de salida.
 