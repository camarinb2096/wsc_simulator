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

<li> 1. Clone el repositorio </li>
<li> 2. Abra una terminal en la raiz del directorio del repositorio </li>
<li> 3. Navege hasta la carpeta docker </li>
<li> 4. Ejecute el comando "sudo docker compose up", esto levantara los servicios incluido la base de datos </li>
<li> 5. Verifique con el comando "sudo docker ps -a" que los servivios hayan sido levantados con exito </li>
<li> 6. Si los servicios se están ejecutando satisfactoriamente puede probar las funcionalidades con la colección de postman adjunta</li>

<p>El servicio de la aplicación puede tardar unos segundos en iniciarse mientras encuentra disponible el servicio de base de datos</p>

<p>Tenga en cuenta que al no haber instalado la imagen de mysql en la primera ejecución puede tardar la descarga de dicho contenedor</p>

# Funcionalidad

<li>El programa tiene la funcionalidad de cargar equipos y jugadores a través de un CSV mediante un endpoint.</li>
<li>Para interactuar con el servicio se adjunta una colección de POSTMAN</li>

[Colección](https://documenter.getpostman.com/view/25279603/2sA3Bn5XQT)
