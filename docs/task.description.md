# Learning Task: Go Application Development and Kubernetes Deployment

## 1. Develop a Simple Go Application with a REST API

Create a Go application that serves a REST API with four endpoints: one for
fetching entities from https://jsonplaceholder.org/posts, parsing it and putting it into a database,
one for retrieving a specific entity, one for listing all entities, and one for deleting one or multiple entities. Also
add unit-tests for the diffent methods.

## 2. Integrate an ORM for Database Operations

Incorporate an ORM library in Go, such as GORM, to enable database
operations.

## 3. Add Prometheus Metrics to the Application

Enhance the application with Prometheus metrics, exposing performance
indicators such as request counts, durations, and error rates.

## 4. Implement Tracing

Implement tracing in the application using OpenTelemetry, tracing
API requests through the application and database operations.

## 5. Implement Logging

Implement logging, the simple Go application should create useful logs like
the endpoints that have been accessed.

## 6. Containerize the Go Application

Prepare the application for deployment by creating a Dockerfile, facilitating
the building of a Docker image for the app.

## 7. Deploy Everything on Kubernetes

Deploy the Go application, a suitable database (PostgreSQL or MySQL for a more
realistic setup), Prometheus, and a tracing solution (e.g., Jaeger) on
Kubernetes, writing Kubernetes manifests or Helm charts.

## 8. Deploy a Log collector

Deploy a collector for logs like FluentBit that fetches logs from the different
components that are deployed in the Kubernetes cluster.

## Bonus Task: Build a Frontend application

If you are done with the tasks and are still motivated you can build a frontend for the Go backend application. An easy
to use framework for building frontends is vue.js.

## Useful Documentation Links

- JSON Placeholder API: [API](https://jsonplaceholder.org/posts)
- Go programming: [A Tour of Go](https://tour.golang.org/)
- REST API design: [Microsoft REST API Guidelines](https://github.com/microsoft/api-guidelines)
- GORM (ORM library): [GORM Official Documentation](https://gorm.io/docs/)
- Prometheus: [Prometheus Documentation](https://prometheus.io/docs/introduction/overview/)
- OpenTelemetry: [OpenTelemetry Go SDK](https://github.com/open-telemetry/opentelemetry-go)
- FluentBit [FluentBit](https://fluentbit.io/)
- Docker: [Docker Documentation](https://docs.docker.com/)
- Kubernetes: [Kubernetes Documentation](https://kubernetes.io/docs/home/)
- Kubernetes
  Onboarding: [Onboarding] (https://mmmake.atlassian.net/wiki/spaces/IDCDEV/pages/2304507948/Kubernetes+Onboarding)
- vue.js: [vue.js](https://vuejs.org/)
