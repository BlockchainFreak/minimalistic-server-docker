## Overview: Lightweight Go Web Server Docker Image

### Description
This Docker image contains a lightweight web server built with Go (Golang). The image is optimized for minimal size, coming in at just 5.76 MB. The server is designed to be simple and efficient, making it ideal for lightweight web applications or microservices.

### Features
- **Small Size:** The Docker image is only 5.76 MB, ensuring quick downloads and minimal storage usage.
- **Configurable Port:** The server listens on a configurable port, which can be set using the `PORT` environment variable. By default, it listens on port 8000.
- **Port Flexibility:** The server can operate on any port greater than 1024, providing flexibility for deployment in various environments.

### Usage

#### Pulling the Image
To pull the image from Docker Hub, use the following command:
```sh
docker pull umernaeem/go-web-server:latest
```

#### Running the Container
To run the container with the default port (8000), use:
```sh
docker run -d -p 8000:8000 umernaeem/go-web-server:latest
```

To run the container with a custom port, use the `-e` flag to set the `PORT` environment variable:
```sh
docker run -d -p <host-port>:<container-port> -e PORT=<container-port> umernaeem/go-web-server:latest
```
For example, to run the server on port 8080:
```sh
docker run -d -p 8080:8080 -e PORT=8080 umernaeem/go-web-server:latest
```

### Example
Here is an example of running the container on port 8080:
```sh
docker run -d -p 8080:8080 -e PORT=8080 umernaeem/go-web-server:latest
```
You can then access the web server at `http://localhost:8080`.

### Use Cases

#### Infrastructure as Code (IaC) with Cloud Services

1. **Google Cloud Run:**
   - **Port Customizability:** When deploying applications on Google Cloud Run, the ability to customize the port via the `PORT` environment variable allows seamless integration with Cloud Run's flexible port configuration.
   - **Pre-built and Lightweight:** The small size of the Docker image (5.76 MB) ensures quick deployment and scaling, reducing startup times and improving overall performance.

2. **AWS App Runner:**
   - **Port Customizability:** AWS App Runner allows you to deploy containerized applications with ease. The customizable port feature of this image ensures compatibility with App Runner's networking configurations.
   - **Pre-built and Lightweight:** The lightweight nature of the image makes it ideal for rapid deployment and scaling, which is crucial for applications that require high availability and quick response times.

3. **Kubernetes Deployments:**
   - **Port Customizability:** In Kubernetes, services often need to run on specific ports. The ability to set the port via an environment variable simplifies the configuration of Kubernetes services and deployments.
   - **Pre-built and Lightweight:** The small image size reduces the time required to pull the image from the container registry, speeding up pod creation and scaling operations.

4. **CI/CD Pipelines:**
   - **Port Customizability:** During continuous integration and continuous deployment (CI/CD) processes, different stages might require the application to run on different ports. The customizable port feature allows for flexible testing and deployment scenarios.
   - **Pre-built and Lightweight:** The minimal size of the image ensures that CI/CD pipelines run faster, as the image can be pulled and started quickly, reducing overall build and deployment times.

5. **Serverless Architectures:**
   - **Port Customizability:** In serverless architectures where functions or services might need to run on specific ports, the customizable port feature provides the necessary flexibility.
   - **Pre-built and Lightweight:** The lightweight image is ideal for serverless environments where resources are provisioned on-demand, ensuring that functions start quickly and efficiently.

6. **Edge Computing:**
   - **Port Customizability:** Edge computing often involves deploying services on devices with specific network configurations. The ability to set the port via an environment variable allows for easy adaptation to various network environments.
   - **Pre-built and Lightweight:** The small size of the image makes it suitable for edge devices with limited storage and processing power, ensuring efficient use of resources.

By leveraging the customizable port feature and the lightweight nature of this Docker image, you can achieve greater flexibility and efficiency in various cloud and containerized environments.

### Environment Variables
- `PORT`: The port on which the web server will listen. Must be greater than 1024. Defaults to 8000.

### Docker Hub Repository
You can find the Docker image on Docker Hub at:
[Docker Hub Repository Link](https://hub.docker.com/r/umernaeem/go-web-server)

### License
This project is licensed under the MIT License.