# network-monitoring-app
La aplicación mostraría un dashboard donde podrías ver gráficas de datos de flujo y métricas en tiempo real. Grafana presentará gráficos personalizados en base a tus métricas y los flujos de red, mientras que el frontend en React permitiría agregar y gestionar dispositivos.

1. Configuración del Entorno
Sistema Operativo: Debian es una muy buena opción para este proyecto, ya que es estable y cuenta con soporte para todas las herramientas que necesitamos. Las versiones recientes de Debian incluyen soporte para Docker y Kubernetes, lo cual facilita el despliegue.

2. Paso a Paso para Implementación
Paso 1: Instalar Docker y Kubernetes en Debian
Comienza instalando Docker y configurando Kubernetes (k8s) en tu sistema. Aquí los comandos básicos:
# Actualiza el sistema
sudo apt update && sudo apt upgrade -y

# Instalar Docker

sudo apt install -y docker.io

sudo systemctl start docker

sudo systemctl enable docker

Pasos para Ejecutar Todo con Docker Compose
Asegúrate de que Docker y Docker Compose estén instalados en tu sistema. Puedes verificar la instalación de Docker y Docker Compose con estos comandos:

docker --version

docker-compose --version

Posiciónate en la carpeta raíz del proyecto. En este caso, asegúrate de estar en la carpeta /network-monitoring-app donde está el archivo docker-compose.yml.

Construye y levanta los servicios con Docker Compose. Ejecuta el siguiente comando para construir las imágenes y levantar los contenedores:

docker-compose up -d
La opción -d ejecuta los contenedores en segundo plano (modo "detached").
Verifica que los contenedores estén en ejecución. Puedes usar este comando para ver el estado de cada contenedor:

docker-compose ps
Deberías ver todos los servicios (backend, frontend, influxdb, elasticsearch, y grafana) listados y en estado Up.

Accede a la aplicación web en el navegador. Como configuraste el frontend para que esté en el puerto 80, ahora puedes abrir el navegador y acceder a:

http://localhost
Esto debería mostrarte la interfaz React de administración y visualización de métricas.

Accede a Grafana (si deseas configurarlo). Grafana está configurado en el puerto 3000, así que puedes acceder a:

http://localhost:3000
Desde aquí, puedes crear dashboards adicionales y embebidos en tu aplicación React si deseas personalizar la visualización de métricas.

Comandos Adicionales Útiles
Para detener todos los contenedores:

docker-compose down
Para ver los logs de un servicio específico, como el backend:


docker-compose logs backend
