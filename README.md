<h1>WorldSoccerChampionship Simulator</h1>

# Descripción

<p> De acuerdo al ejercicio propuesto, se define la implementación de un API REST, utilizando una arquitectura por capas con un enfoque DDD. En el que se le asigna un paquete completo a cada entidad que se vería involucrada en el simulador de torneo de futbol. Ofreciendo la portabilidad de cada entidad y una clara separación de responsabilidades. </p>

# Paquetes implementados

<li> ORM: GORM </li>
<li> WEB: GIN </li>

<p> Se define la implementación del ORM para implementar practicas de code first, lo que facilita la manipulación, la integridad y la seguridad de los datos. Adicional de la eficiencia al modelar la base de datos desde el código. GIN como web framework para agilidad en la configuración de rutas HTTP y en el desarrollo. </p>

# Requisitos

<li> Docker (con compose habilitado) </li>

# Instalación y puesta en marcha

<li> Abra una terminal en la raiz del directorio del repositorio </li>
<li> Navege hasta la carpeta docker </li>
<li> Ejecute el comando "sudo docker compose up", esto levantara los servicios incluido la base de datos </li>
<li> Verifique con el comando "sudo docker ps -a" que los servivios hayan sido levantados con exito </li>
<li> Si los servicios se están ejecutando satisfactoriamente puede probar las funcionalidades con la colección de postman adjunta</li>
