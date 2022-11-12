# Arquitectura Hexagonal
Este es un patrón de arquitectura (arquitectura de software aplicado al diseño de software) que nos permite crear aplicaciones débilmente acopladas a sus componentes.

Nos permite dividir una aplicación en varios componentes intercambiables como:

* El core de la aplicación (la lógica de negocio)
* La interfaz de usuario
* El almacenamiento
* Los test

Los componentes son conectados entre sí a través de `Puertos` , los `puertos` son las interfaces públicas que deben usar los `adaptadores` para cumplir con el objetivo del puerto.

Los `Adaptadores` son la especialización de un contexto concreto.

![Hexagonal Architecture](client/hexagonal.png)

Nota: Imagen hecha por el creador Alistair Cockburn.

