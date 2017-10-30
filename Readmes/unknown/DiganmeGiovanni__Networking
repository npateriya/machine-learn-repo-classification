
#Conmutación y enrutamiento de redes

##Contenido

 * Capitulo 1
 * Capitulo 2
 * Capitulo 3


## Capitulo 1 - Introducción al escalamiento de redes

#### Necesidades de crecimiento de la red
 * Aplicaciones criticas
 * Trafico de redes convergentes
 * Diversas necesidades de negocio
 * Control administrativo centralizado

#### Equipo de clase empresarial
Se instala para proporcionar una red *Altamente confiable*.

#### Diseño jerarquico de la red
Divide la funcionalidad de la red en tres capas.

##### Capa de acceso
Proporciona conectividad a los usuarios.

##### Capa de distribución
Envia el trafico de una red local a otra.

##### Capa de nucleo
Representa una capa troncal de alta velocidad entre las redes dispersas.

#### Arquitectura empresarial de CISCO
Divide la red en componentes funcionales mientras mantiene las capas de nucleo, acceso y distribución.

Los modulos de la arquitectura empresarial de cisco son.

 * Campus empresarial
 * Perimetro empresarial
 * Perimetro del proveedor de servicios
 * Remoto

![Arquitectura empresarial](http://i.imgur.com/ZPHFLci.png)

##### Campus empresarial
Toda la infraestructura del campues e incluye las capas de acceso, distribución y nucleo.

El módulo de capa de acceso incluye switches de capa 2 o de capa 3 para proporcionar la densidad de puestos requerida.En este módulo se produce la implementación de las VLANs y los enlaces troncales a la capa de distribucion del edificio. La redundancia a los switches de distribución del edificio es importante.

El módulo de distriución agrega acceso al edificio mediante dispositivos de capa 3. En este modulo se llevan a cabo el routing, el control de acceso y el QoS.

El módulo de nucleo proporciona conectividad de alta velocidad entre los módulos de la capa de distribución, las granjas de servidores de los datacenters y el perimetro empresarial.

En este módulo el eje central del diseño es la redundancia, la convergencia rápida y la tolerancia a fallas.

##### Perimetro empresarial
Esta compuesto por los módulos de internet, VPN y WAN que conectan la red de la empresa a la red del proveedor de servicios. Este módulo extiende los servicios de la empresa a sitios remotos y permite que la empresa utilice recursos de internet y de socios.

Proporciona QoS, refuerzo de politicas, niveles de servicio y seguridad.

##### Perimetro del proveedor de servicios
Proporciona servicios de internet, de red publica de telefonia conmutada (PSTN) y WAN.

El modelo de red empresarial compuesta (ECNM) pasa a través de un dispositivo de extremo. Este es el momento en el que los paquetes se pueden analizar y se puede tomar una decisión sobre si se debe permitir el ingreso de estos a la red empresarial. Los sistemas de detección de intrusiones (IDS) y los sistemas de prevención de intrusiones (IPS) también se pueden configurar en el perimetro empresarial para brindar protección contra actividades malintencionadas.

####Dominio de fallas
Un dominio de fallas es el area de la red que se ve afectada cuando un dispositivo o servicio escencial experimenta problemas.

Una red bien diseñada, ademas de controlar el trafico, limita el tamaño de los dominios de fallas

#### Expansión de la red.

#####Diseño que admita escalabilidad
Diseñar una estrategia que permita que la red este disponible y se pueda escalar eficazmente. Recomentaciones:

 * Utilice equipo modular extensible o de dispositivos agrupados que puedan actualizarce facilmente para incrementar las capacidades. Algunos equipos se pueden integrar en un cluster para que funcionen como uno solo.
 * Diseñe la red jerarquica para que incluya modulos que se puedan agregar, quitar o actualizar segun sea necesario sin afectar el diséño de otras areas funcionales de la red. Por ejemplo, una capa de enlace que se pueda expandir sin afectar las capas de distribución y nucleo.
 * Cree una buena estrategia de direccionamiento IPv4

 * La implementación de enlaces redundantes entre los dispositivos escenciales en capa de acceso y capa nucleo.
 * La implementación de varios enlaces entre los equipos, ya sea con agregacion de enlaces mediante etherchannel o con balanceo de cargas para aumentar el ancho de banda. La combinación de varios enlaces Ethernet en una sola configuración con balanceo de carga Etherchannel aumenta el ancho de banda.
 * La implementación de conexión inalambrica para permitir movilidad y expansión.
 * El uso de un protocolo de routing escalable y la implementación de caracteristicas de ese protocolo para aislar las actualizaciones de routing y minimizar el tamaño de la tabla de ruteo.

##### Redundancia
La redundancia es una parte importante del diseño de una red para prevenir interrupciones de servicio y minimizar la posibilidad de un punto unico de falla.
Un metodo para implementar redundancia consiste en instalar equipos duplicados y proporcionar servicios de conmutación por falla para los dispositivos.

Otro método para implementar redundancia es mediante rutas redundantes. Estas ofrecen rutas fisicas alternativas para atravesar la red. En una red conmutada, las rutas redundantes admiten una alta disponibilidad, sin embargo, debido al funcionamiento de los switches, estas pueden generar bucles logicos infinitos, es por ello que se necesita el protocolo de expansión de arbol (STP).

El protocolo STP permite la redundancia necesaria para proporcionar confiabilidad pero elimina los bucles de switching. Para hacerlo proporciona un mecanismo para deshabilitar rutas conmutadas hasta que la ruta se vuleva necesari, por ejemplo, cuando ocurre una falla. Es un protocolo de estandares abiertos que se utiliza en un entorno de conmutación para crear una topologia lógica sin bucles.

##### Implementación de etherchannel
En el diseño de red jerarquico es posible que algunos enlaces entre los dispositivos de acceso y de enlace necesiten procesa una mayor cantidad de trafico que los enlaces entre otros dispositivos. A medida que el trafico de varios enlaces converge en un unico punto de salida, es posible que se genere un cuello de botella.

La agregación de enlaces permiten que el administrador aumente el ancho de banda de un enlace logico mediante la combinación de varios enlaces fisicos. Etherchannel es una forma de agregación de enlaces fisicos en redes conmutadas.

Etherchannel utiliza los puertos de switch existentes, por lo tanto no es necesario incurrir en gastos mayores para actualizar el enlaces a una conexión mas veloz y costosa. El enlace Etherchannel se ve como un enlace lögico que utiliza una interfaz Etherchannel, la mayoria de las tareas de configuración y adminsitración se realizar sobre la interfaz Etherchannel de manera que se asegure la coherencia de configuración en todos los enlaces. Por último, Etherchannel aprovecha el balanceo de carga entre los enlaces que forma parte del Etherchannel. Segun la plataforma de hardware se pueden implementar diferentes mecanismos de balanceo de carga.

##### Administración de la red enrutada.

###### Administración in band.
Requiere de almenos una interfaz conectada y configurada y el uso de Telnet, ssh o http para acceder al dispositivo.

###### Administración Out of band.
Requiere conexión directa a un puerto de consola o AUX y un cliente de emulación de terminal para acceder al dispositivo.


## Redundancia de LAN

#### Tormenta de broadcast.
Se produce cuanbdo en una red con enlaces redundantes el ancho de banda es agotado debido a bucles lógicos generados durante el envio de paquetes de difusión.