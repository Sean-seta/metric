# Go Metrics with Prometheus & Grafana

A containerized Go (Gin) web application instrumented for Prometheus monitoring with a pre-configured Grafana dashboard. The entire stack is managed via Docker Compose for easy setup.

## 🚀 Getting Started

Follow these instructions to get the project up and running.

### Prerequisites

* Docker
* Docker Compose

### Installation & Running

1.  **Start the entire application stack** using the `Makefile`:
    ```bash
    make run
    ```
    This command builds the Go application and starts the `app`, `prometheus`, and `grafana` services.

2.  **(Optional) Generate traffic** to the application endpoints:
    ```bash
    make test
    ```
    This runs a test script that continuously sends requests to the application.

### Accessing Services

Once the containers are running, the services will be available at the following ports:

* **📍 Go Application**: [http://localhost:8080](http://localhost:8080)
* **📊 Prometheus**: [http://localhost:9090](http://localhost:9090)
* **📈 Grafana**: [http://localhost:3000](http://localhost:3000)
    * **Login**: `admin` / `admin`

Navigate to the Grafana URL and import using `grafana.json` to view the pre-loaded dashboard visualizing the application's metrics.
