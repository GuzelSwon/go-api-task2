docker_build('go-api-task-guzel', '.',
    dockerfile='app/Dockerfile')
k8s_yaml('k8s/app/deployment.yaml')
k8s_yaml('k8s/mysql/deployment.yaml')
k8s_yaml('k8s/prometheus/deployment.yaml')
k8s_yaml('k8s/otel/deployment.yaml')
k8s_resource('go-api-task-guzel', port_forwards=8000)