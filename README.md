
# Minio Storage Wrapper

## Features

- Minio Storage: Interacts with Minio for file storage.
- HTTP and gRPC Support: Serves requests in both HTTP and gRPC formats.
- Redis Cache: Stores metadata for files in a Redis cache.
- Load Balancing: Nginx configuration for load balancing between client and API servers.

## Installation

### Prerequisites

- Go installed on your machine
- Minio Server
- Redis Server
- Nginx

### Steps

1. Clone the repository:

```bash
git clone git@github.com:MikeMwita/minio_storage.git
```

2. Install dependencies:

```bash
go mod download
```

3. Configure environment variables:

Copy the example environment file:

```bash
cp .env.example .env
```

Modify the values in the .env file to match your setup.

4. Run the application:

```bash
go run main.go
```

5. Configure Nginx:

Update your Nginx configuration with the provided upstream configuration.

6. Access the application:

The HTTP server is accessible at http://yourdomain/.
The gRPC server is accessible at yourdomain:9000.

## Configuration

### Minio Configuration:

Update Minio server details in .env file.

### Redis Configuration:

Update Redis server details in .env file.

### Nginx Configuration:

Modify Nginx configuration for load balancing as per your requirements.

## Usage

### HTTP Server:

Access the HTTP server at http://yourdomain/.

### gRPC Server:

Access the gRPC server at yourdomain:9000.

## License

This project is licensed under the MIT License - see the LICENSE file for details.

