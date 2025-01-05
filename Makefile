IMAGES = cart-service product-service user-service auth-service
DOCKER_COMPOSE_FILE = ./deployments/compose/docker-compose.yaml
DEPLOYMENTS_DIR = ./deployments
KUBECTL_APPLY_FILES = $(DEPLOYMENTS_DIR)/*

# Сборка всех сервисов
build:
	@for service in $(IMAGES); do \
		echo "Building $$service..."; \
		docker build -t $$service:latest ./$$service; \
	done

# Загрузка образов в Minikube
load-images:
	@for service in $(IMAGES); do \
		echo "Loading $$service into Minikube..."; \
		minikube image load $$service:latest; \
	done

# Конвертация docker-compose в манифесты k8s
convert:
	kompose convert -f $(DOCKER_COMPOSE_FILE) -o $(DEPLOYMENTS_DIR)

# Запуск Minikube
start-minikube:
	minikube start --driver=docker

# Включение Ingress
enable-ingress:
	minikube addons enable ingress

# Применение конфигураций k8s
apply:
	kubectl apply -f $(KUBECTL_APPLY_FILES)

# Запуск туннеля Minikube
tunnel:
	minikube tunnel

# Полный цикл деплоя
deploy: build start-minikube load-images enable-ingress convert apply tunnel

# Удаление деплоя
delete:
	kubectl delete -f $(KUBECTL_APPLY_FILES)
	minikube stop
	minikube delete

# Открытие интерфейса мониторинга
grafana:
	minikube service grafana

# Проверка
check:
	minikube image list
	kubectl get pods                          
	kubectl get services