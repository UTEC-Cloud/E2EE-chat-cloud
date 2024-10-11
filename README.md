# E2EE-chat

## Integrantes
* Miguel Yurivilca
* Rodrigo Salazar

## Aplicación

Sistema de mensajería segura con cifrado E2EE basado en X3HD.

## Funcionalidades

|  ID |  CATEGORIA | REQUERIMIENTO |
|---|---|---|
| RF01 | Auth | Usuario puede crear una cuenta |
| RF02 | Auth | Usuario puede iniciar sesión a su cuenta |
| RF03 | Chat | Usuario es capaz de enviar mensaje de forma asincrónica |
| RF04 | Chat | Usuario recibe notificación en tiempo real al recibir un mensaje |
| RF05 | E2EE | Usuario es capaz de generar un par de llaves de identidad (privada y publica) |
| RF06 | E2EE | Usuario es capaz de compartir su llave de identidad pública |

## Características

### Protocolo de encriptación X3DH

El protocolo X3DH es un protocolo de acuerdo de claves que permite a dos partes establecer una clave secreta compartida a través de un canal inseguro. El protocolo X3DH se basa en el protocolo de Diffie-Hellman y utiliza una combinación de curvas elípticas y funciones hash criptográficas para garantizar la seguridad de la comunicación. 

![image](images/flujo_x3dh.png "Flujo X3DH")

La clave secreta compartida se utiliza para cifrar y descifrar mensajes, estableciendo un canal de comunicación E2EE seguro.




## Arquitectura actual

La arquitectura actual es una arquitectura monolítica que consta en unico servidor responsable de la gestión de usuarios y mensajes. La comunicación entre el cliente y el servidor se realiza a través de una conexión HTTPS (autenticación) y Secure WebSockets (mensajes). 

La base de datos utilizada es una base de documentos: MongoDB. Almacena la información de los usuarios, sus paquetes de llaves y los mensajes enviados.

![image](images/current_architecture.png "Current architecture")

## Arquitectura propuesta

El escalamiento de servicios con websockets es un desafío en arquitecturas monolíticas. Para resolver este problema, se propone una arquitectura serverless que permite escalar de forma horizontal los servicios de chat y autenticación.

![image](images/new_architecture.png "New architecture")

## Pasos para ejecutar la aplicación

### Inicializar BD
```bash
docker-compose up -d
```
### Ejecutar Servidor
```bash
cd e2ee_server
go run .
```

### Ejecutar Cliente
```bash
cd e2ee_client
go run .
```



## Tópicos de cloud

* Cloud Serverless
    * Migración de la Arquitectura Monolítica a Serverless
    * Escalabilidad de WebSockets

* Cloud Databases
    * Base de datos de documentos para usar User Key Bundles
    * Base de datos de llave-valor para gestión de WebSockets

* Dockerización
    * Despliegue de lambdas con imágenes docker

* Cloud & DevOps
    * CI/CD pipeline de despliegue de aplicación

* Cloud Monitoring:
    * Logging, alertas y escalabilidad

* Cloud Security
    * Autenticación con servicios cloud (Amazon Cognito)

## Referencias

1) The X3DH Key Agreement Protocol. (2016). Signal Messenger. https://signal.org/docs/specifications/x3dh/
2) API Gateway WebSocket APIs - Amazon API Gateway. (n.d.). https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api.html
3) Vishnevskiy, S. (2021, August 25). How Discord Scaled Elixir to 5,000,000 Concurrent Users. Discord. https://discord.com/blog/how-discord-scaled-elixir-to-5-000-000-concurrent-users
4) Monitor your Amazon EC2 Auto Scaling groups - Amazon EC2 Auto Scaling. (n.d.). https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-monitoring-features.html