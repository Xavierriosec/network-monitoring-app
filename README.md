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


