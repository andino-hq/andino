# Andino

<div align="center">

[![Go Version](https://img.shields.io/badge/Go-1.25.1-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://golang.org/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg?style=for-the-badge&logo=opensourceinitiative&logoColor=white)](https://opensource.org/licenses/MIT)
[![Build Status](https://img.shields.io/github/actions/workflow/status/andino-hq/andino/go.yml?branch=main&style=for-the-badge&logo=githubactions&logoColor=white)](https://github.com/andino-hq/andino/actions)
[![Go Report Card](https://img.shields.io/badge/Go%20Report-A+-brightgreen?style=for-the-badge&logo=go&logoColor=white)](https://goreportcard.com/report/github.com/andino-hq/andino)
[![Test Coverage](https://img.shields.io/badge/Coverage-85%25-brightgreen?style=for-the-badge&logo=codecov&logoColor=white)](https://codecov.io/gh/andino-hq/andino)

[![GitHub Issues](https://img.shields.io/github/issues/andino-hq/andino?style=for-the-badge&logo=github&logoColor=white)](https://github.com/andino-hq/andino/issues)
[![GitHub Pull Requests](https://img.shields.io/github/issues-pr/andino-hq/andino?style=for-the-badge&logo=github&logoColor=white)](https://github.com/andino-hq/andino/pulls)
[![GitHub Stars](https://img.shields.io/github/stars/andino-hq/andino?style=for-the-badge&logo=github&logoColor=white)](https://github.com/andino-hq/andino/stargazers)
[![Discord](https://img.shields.io/discord/your-server-id?style=for-the-badge&logo=discord&logoColor=white&label=Community)](https://discord.gg/your-invite-link)
[![Contributors](https://img.shields.io/github/contributors/andino-hq/andino?style=for-the-badge&logo=github&logoColor=white)](https://github.com/andino-hq/andino/graphs/contributors)

**🇨🇱 The definitive Chilean business integration library**

_Technology-agnostic API layer for fiscal compliance, payments, and business automation_

[Features](#-features) •
[Quick Start](#-quick-start) •
[API Docs](#-api-documentation) •
[Contributing](#-contributing) •
[Community](#-community)

</div>

---

## 🎯 **Mission Statement**

**Andino** is Chile's first comprehensive, open-source business integration library designed to eliminate the complexity of Chilean fiscal, tax, accounting, and digital payment systems. We bridge the gap between complex regulatory requirements and modern software development, empowering small and medium enterprises (SMEs) to automate their business processes with confidence.

### 🚀 **Why Andino Exists**

Chilean businesses face unique challenges:

- 📋 **Complex Fiscal Requirements**: SII integration, DTE management, and ever-changing tax regulations
- 💳 **Fragmented Payment Ecosystem**: Difficult-to-implement payment gateways with poor documentation
- 🏢 **SME Barriers**: Limited resources to implement and maintain compliance systems
- 🔧 **Technology Lock-in**: Vendor-specific solutions that limit flexibility
- 📊 **Manual Processes**: Lack of integrated business intelligence and reporting tools

**Andino solves these problems with a single, unified, technology-agnostic solution.**

## ✨ **Core Principles**

- **🌍 Technology Agnostic**: Works with any programming language, framework, or platform
- **🇨🇱 Chile-First**: Built specifically for Chilean regulations and business practices
- **🔓 Open Source**: Transparent, community-driven development
- **🚀 SME-Focused**: Designed for small and medium enterprises
- **📈 Future-Ready**: Extensible architecture for analytics and business intelligence
- **🛡️ Compliance-First**: Automatic adherence to Chilean legal requirements

## 📖 Table of Contents

- [Vision & Roadmap](#-vision--roadmap)
- [Features](#-features)
- [Architecture](#-architecture)
- [Tech Stack](#-tech-stack)
- [Quick Start](#-quick-start)
- [API Documentation](#-api-documentation)
- [Configuration](#-configuration)
- [Project Structure](#-project-structure)
- [Plugin System](#-plugin-system)
- [Testing](#-testing)
- [Deployment](#-deployment)
- [Contributing](#-contributing)
- [Community](#-community)
- [Roadmap](#-roadmap)
- [License](#-license)

## 🔮 **Vision & Roadmap**

### 🎯 **Phase 1: Foundation** (Current)

- ✅ Core fiscal integration (SII, DTE)
- ✅ Unified payment processing (Transbank, Flow, Khipu)
- ✅ Technology-agnostic API design
- ✅ Plugin architecture
- ✅ Basic business automation

### 📊 **Phase 2: Intelligence** (2025 Q2-Q3)

- 📈 **Business Analytics Engine**: Real-time financial insights
- 📋 **Automated Reporting**: Regulatory reports generation
- 🔍 **Compliance Monitoring**: Automatic regulation updates
- 📊 **Dashboard & Visualization**: Business intelligence tools
- 🤖 **AI-Powered Insights**: Predictive analytics for SMEs

### 🌟 **Phase 3: Ecosystem** (2025 Q4+)

- 🏪 **E-commerce Integration**: Plug-and-play solutions
- 🏦 **Banking APIs**: Direct integration with Chilean banks
- 📱 **Mobile SDKs**: Native mobile app support
- 🌐 **Marketplace**: Community plugins and extensions
- 🔐 **Enterprise Features**: Advanced security and compliance tools

## 🚀 Features

### 🏛️ **Fiscal & Tax Compliance**

- 📄 **Electronic Tax Documents (DTEs)**: Automated invoice, receipt, credit note, and debit note generation
- 🏛️ **SII Integration**: Real-time integration with Chile's tax authority (Servicio de Impuestos Internos)
- 🔐 **Digital Signatures**: Automatic document signing with valid digital certificates
- 📊 **Tax Calculation**: Real-time IVA, additional taxes, and retention calculations
- 🔍 **Audit Trail**: Complete regulatory compliance tracking and reporting
- 📋 **Legal Document Templates**: Pre-built templates compliant with Chilean regulations
- ⚡ **Automatic Updates**: Stay current with changing tax laws and regulations

### 💳 **Payment Processing Revolution**

- 🌐 **Unified Payment API**: Single interface for all major Chilean payment gateways
- 🏪 **Gateway Support**: Transbank (WebPay Plus, OnePay), Flow, Khipu, MercadoPago Chile
- 💰 **Transaction Management**: Complete lifecycle from authorization to settlement
- 🔄 **Webhook Processing**: Real-time payment notifications and status updates
- 💱 **Multi-Currency**: CLP, USD, EUR with automatic conversion rates
- 🛡️ **Fraud Protection**: Built-in security validation and risk assessment
- 📱 **Mobile Payments**: Support for mobile payment methods and wallets
- 🔁 **Recurring Payments**: Subscription and installment payment handling

### 🔧 **Developer Experience**

- 🌍 **Technology Agnostic**: REST API, GraphQL, gRPC, and WebSocket interfaces
- 🔌 **Plugin Architecture**: Extensible system for custom business logic
- 📚 **Comprehensive SDKs**: Go, Python, Node.js, PHP, .NET, Java clients
- 📖 **Interactive Documentation**: OpenAPI 3.0 with Swagger UI and code examples
- 🛠️ **Multiple Interfaces**: Web API, CLI tools, and desktop applications
- 🧪 **Testing Tools**: Sandbox environments and test data generators
- 📦 **Easy Integration**: Docker containers, Kubernetes operators, and cloud templates

### 📊 **Business Intelligence** (Coming Soon)

- 📈 **Financial Analytics**: Revenue tracking, expense analysis, and profit reports
- 📋 **Regulatory Reporting**: Automatic generation of tax returns and compliance reports
- 🎯 **KPI Dashboards**: Real-time business metrics and performance indicators
- 🔮 **Predictive Analytics**: Cash flow forecasting and business trend analysis
- 📊 **Data Visualization**: Interactive charts and graphs for business insights
- 📤 **Export Capabilities**: PDF, Excel, and CSV report generation

### ⚡ **Performance & Reliability**

- 🚀 **High Throughput**: 10,000+ transactions per second
- 🔄 **Asynchronous Processing**: Background job queues for heavy operations
- 📈 **Auto-scaling**: Kubernetes-native horizontal scaling
- 💾 **Smart Caching**: Redis-powered response and data caching
- 📊 **Monitoring**: Prometheus metrics, health checks, and alerting
- 🔄 **High Availability**: Multi-region deployment support

### 🔒 **Security & Compliance**

- 🔐 **Enterprise Security**: JWT authentication, OAuth 2.0, and API key management
- 🛡️ **Rate Limiting**: Configurable request throttling and DDoS protection
- ✅ **Data Validation**: Comprehensive input/output validation and sanitization
- 🌍 **CORS & Security Headers**: Configurable cross-origin and security policies
- 🔒 **Encryption**: End-to-end encryption with TLS 1.3 support
- 📋 **Compliance**: SOX, GDPR, and Chilean data protection law adherence
- 🔍 **Audit Logging**: Complete request/response logging for compliance

## 🏗️ Architecture

**Andino** implements **Clean Architecture** principles with **Domain-Driven Design** to ensure maintainability, testability, and technology independence:

```mermaid
graph TB
    subgraph "External Systems"
        SII[🏛️ SII<br/>Tax Authority]
        Banks[🏦 Payment<br/>Gateways]
        Email[📧 Email<br/>Services]
    end

    subgraph "Interfaces Layer"
        REST[🌐 REST API]
        GraphQL[📊 GraphQL]
        gRPC[⚡ gRPC]
        CLI[💻 CLI Tool]
        Worker[⚙️ Background<br/>Worker]
    end

    subgraph "Application Layer"
        Services[🔧 Application<br/>Services]
        UseCases[📋 Use Cases]
        DTOs[📦 Data Transfer<br/>Objects]
        Ports[🔌 Ports &<br/>Interfaces]
    end

    subgraph "Domain Layer"
        DTE[📄 DTE<br/>Domain]
        Payments[💳 Payments<br/>Domain]
        SII_Domain[🏛️ SII<br/>Domain]
        Common[🔧 Common<br/>Domain]
    end

    subgraph "Infrastructure Layer"
        DB[(🗄️ PostgreSQL<br/>Database)]
        Cache[(🔴 Redis<br/>Cache)]
        Config[⚙️ Configuration]
        Logger[📝 Logging]
        Crypto[🔐 Cryptography]
    end

    subgraph "Plugin System"
        PaymentPlugins[💳 Payment<br/>Plugins]
        SIIPlugins[🏛️ SII<br/>Plugins]
        CustomPlugins[🔧 Custom<br/>Plugins]
    end

    REST --> Services
    GraphQL --> Services
    gRPC --> Services
    CLI --> Services
    Worker --> Services

    Services --> UseCases
    Services --> DTOs
    UseCases --> Ports

    Ports --> DTE
    Ports --> Payments
    Ports --> SII_Domain
    DTE --> Common
    Payments --> Common
    SII_Domain --> Common

    Services --> DB
    Services --> Cache
    Services --> Config
    Services --> Logger
    Services --> Crypto

    Services --> PaymentPlugins
    Services --> SIIPlugins
    Services --> CustomPlugins

    PaymentPlugins --> Banks
    SIIPlugins --> SII
    Services --> Email
```

### 🎯 **Architectural Benefits**

- **🔌 Technology Independence**: Core business logic is isolated from external dependencies
- **🧪 Testability**: Each layer can be tested independently with mocks and stubs
- **🔄 Maintainability**: Clear separation of concerns makes updates and changes easier
- **📈 Scalability**: Stateless design allows for horizontal scaling
- **🔧 Extensibility**: Plugin system allows custom functionality without core changes

## 🛠️ Tech Stack

### 🚀 **Core Technologies**

![Go](https://img.shields.io/badge/Go-1.25.1-00ADD8?style=flat-square&logo=go&logoColor=white)
![Gin](https://img.shields.io/badge/Gin_Gonic-v1.9+-009688?style=flat-square&logo=go&logoColor=white)

### 🗄️ **Data & Storage**

![PostgreSQL](https://img.shields.io/badge/PostgreSQL-15+-4169E1?style=flat-square&logo=postgresql&logoColor=white)
![Redis](https://img.shields.io/badge/Redis-7+-DC382D?style=flat-square&logo=redis&logoColor=white)
![SQLite](https://img.shields.io/badge/SQLite-Testing-003B57?style=flat-square&logo=sqlite&logoColor=white)

### 🔧 **Infrastructure & DevOps**

![Docker](https://img.shields.io/badge/Docker-Ready-2496ED?style=flat-square&logo=docker&logoColor=white)
![Kubernetes](https://img.shields.io/badge/Kubernetes-Native-326CE5?style=flat-square&logo=kubernetes&logoColor=white)
![Nginx](https://img.shields.io/badge/Nginx-Proxy-009639?style=flat-square&logo=nginx&logoColor=white)
![Terraform](https://img.shields.io/badge/Terraform-IaC-7B42BC?style=flat-square&logo=terraform&logoColor=white)

### 🛠️ **Development & Testing**

![Swagger](https://img.shields.io/badge/OpenAPI_3.0-Docs-85EA2D?style=flat-square&logo=swagger&logoColor=black)
![Air](https://img.shields.io/badge/Air-Live_Reload-FF6B6B?style=flat-square&logo=go&logoColor=white)
![Testify](https://img.shields.io/badge/Testify-Testing-00D9FF?style=flat-square&logo=go&logoColor=white)
![Zap](https://img.shields.io/badge/Zap-Logging-FFA500?style=flat-square&logo=go&logoColor=white)

### 💳 **Chilean Integrations**

![SII](https://img.shields.io/badge/SII-Fiscal-1E88E5?style=flat-square&logo=government&logoColor=white)
![Transbank](https://img.shields.io/badge/Transbank-WebPay-E53935?style=flat-square&logo=creditcard&logoColor=white)
![Flow](https://img.shields.io/badge/Flow-Payments-4CAF50?style=flat-square&logo=payments&logoColor=white)
![Khipu](https://img.shields.io/badge/Khipu-Transfers-FF9800?style=flat-square&logo=payments&logoColor=white)

## 🚀 Quick Start

### ✅ Prerequisites

![Go](https://img.shields.io/badge/Go-1.25.1+-00ADD8?style=flat-square&logo=go&logoColor=white)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-15+-4169E1?style=flat-square&logo=postgresql&logoColor=white)
![Docker](https://img.shields.io/badge/Docker-Latest-2496ED?style=flat-square&logo=docker&logoColor=white)

- **Go** 1.25.1+ ([Download](https://golang.org/dl/))
- **PostgreSQL** 15+ ([Download](https://www.postgresql.org/download/))
- **Docker & Docker Compose** ([Download](https://docs.docker.com/get-docker/))
- **Make** (optional, for build automation)

### ⚡ One-Command Setup

```bash
# Clone and start everything with Docker
git clone https://github.com/andino-hq/andino.git
cd andino
make docker-up
```

### 🛠️ Manual Installation

<details>
<summary><b>👆 Click to expand manual setup steps</b></summary>

1. **📥 Clone the repository**

   ```bash
   git clone https://github.com/andino-hq/andino.git
   cd andino
   ```

2. **⚙️ Environment Setup**

   ```bash
   # Copy configuration template
   cp configs/config.yaml.example configs/config.yaml

   # Edit with your settings (database, SII credentials, etc.)
   nano configs/config.yaml
   ```

3. **🗄️ Database Setup**

   ```bash
   # Create database
   createdb andino_dev

   # Run migrations (auto-runs on first start)
   go run cmd/server/main.go --migrate
   ```

</details>

### 🐳 Docker Setup (Recommended)

The fastest way to get Andino running:

```bash
# 🚀 Start all services (API + DB + Redis + Worker)
cd deployments/docker
docker-compose up -d

# 📊 Verify services are running
docker-compose ps

# 🔍 View logs
docker-compose logs -f andino-server
```

**✅ Services Now Running:**

- 🌐 **API Server**: http://localhost:8080
- 📚 **Swagger UI**: http://localhost:8080/swagger/
- ⚙️ **Background Worker**: Processing jobs
- 💻 **CLI Tool**: `docker exec -it andino-cli andino --help`
- 🗄️ **PostgreSQL**: localhost:5432
- 🔴 **Redis**: localhost:6379

### 💻 Local Development

For development with hot-reload:

```bash
# 📦 Install dependencies
go mod download

# 🔥 Install Air for live reload
go install github.com/cosmtrek/air@latest

# 🚀 Run API server with hot reload
air -c .air.server.toml

# ⚙️ Run worker in another terminal
air -c .air.worker.toml

# 💻 Build CLI tool
go build -o bin/andino cmd/cli/main.go

# 🛠️ Available make targets
make help
```

### ✨ Quick Verification

```bash
# 🩺 Health check
curl http://localhost:8080/api/v1/health

# 🏛️ SII connection test
curl http://localhost:8080/api/v1/sii/status

# 💳 Payment gateways status
curl http://localhost:8080/api/v1/payments/gateways

# 📊 Server info
curl http://localhost:8080/api/v1/info

# 📚 Interactive API docs
open http://localhost:8080/swagger/index.html

# 💻 CLI tool help
./bin/andino --help
```

## 📚 API Documentation

### 🌐 **Multiple Interface Options**

**Andino** provides several ways to interact with the system:

| Interface       | Endpoint        | Purpose           | Documentation                                       |
| --------------- | --------------- | ----------------- | --------------------------------------------------- |
| 🌐 **REST API** | `:8080/api/v1`  | Standard HTTP API | [Swagger UI](http://localhost:8080/swagger/)        |
| 📊 **GraphQL**  | `:8080/graphql` | Flexible queries  | [GraphQL Playground](http://localhost:8080/graphql) |
| ⚡ **gRPC**     | `:9090`         | High-performance  | [Protocol Buffers](./docs/api/grpc/)                |
| 💻 **CLI**      | `andino`        | Command-line tool | `andino --help`                                     |

### 📖 **Interactive Documentation**

Visit `http://localhost:8080/swagger/index.html` for:

- 🔍 Complete API reference with examples
- 🧪 Interactive endpoint testing
- 📝 Request/response schemas
- 🔐 Authentication guides
- 💡 Integration tutorials

### ⚡ **Quick API Examples**

<details>
<summary><b>📄 Create Electronic Invoice (DTE)</b></summary>

```bash
curl -X POST http://localhost:8080/api/v1/dte/invoices \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_TOKEN" \
  -d '{
    "issuer": {
      "rut": "12345678-9",
      "company_name": "Mi Empresa SpA"
    },
    "receiver": {
      "rut": "98765432-1",
      "company_name": "Cliente Ltda"
    },
    "items": [
      {
        "description": "Consultoría de Software",
        "quantity": 1,
        "unit_price": 100000,
        "tax_rate": 19
      }
    ]
  }'
```

</details>

<details>
<summary><b>💳 Process Payment with Transbank</b></summary>

```bash
curl -X POST http://localhost:8080/api/v1/payments/transbank/webpay \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_TOKEN" \
  -d '{
    "amount": 119000,
    "currency": "CLP",
    "order_id": "ORDER-2024-001",
    "return_url": "https://myapp.com/payment/success",
    "description": "Pago por servicios de consultoría"
  }'
```

</details>

<details>
<summary><b>📊 Get Business Analytics</b></summary>

```bash
curl -X GET "http://localhost:8080/api/v1/analytics/revenue?period=monthly&year=2024" \
  -H "Authorization: Bearer YOUR_TOKEN"
```

</details>

<details>
<summary><b>💻 CLI Examples</b></summary>

```bash
# Generate invoice
andino invoice create \
  --issuer-rut "12345678-9" \
  --receiver-rut "98765432-1" \
  --amount 100000 \
  --description "Consultoría"

# Check SII status
andino sii status

# List payment gateways
andino payments list-gateways

# Generate monthly report
andino reports generate --type monthly --month 2024-01
```

</details>

## ⚙️ Configuration

Andino uses **environment-specific YAML configuration** with **environment variable overrides**:

### 📋 **Configuration Files**

| File               | Environment     | Description                    |
| ------------------ | --------------- | ------------------------------ |
| `config.yaml`      | **Default**     | Base configuration             |
| `config.dev.yaml`  | **Development** | Local development settings     |
| `config.prod.yaml` | **Production**  | Production-optimized settings  |
| `config.test.yaml` | **Testing**     | Test environment configuration |

### 🔧 **Core Configuration**

```yaml
# Server Configuration
server:
  host: "0.0.0.0"
  port: 8080
  timeout: 30s
  read_timeout: 15s
  write_timeout: 15s

# Database Configuration
database:
  url: "postgresql://user:pass@localhost/andino"
  max_connections: 25
  ssl_mode: "prefer"
  auto_migrate: true

# Redis Cache Configuration
redis:
  url: "redis://localhost:6379"
  db: 0
  password: ""
  max_retries: 3

# Logging Configuration
logging:
  level: "info" # debug, info, warn, error
  format: "json" # json, text
  output: "stdout" # stdout, file
  file_path: "./logs/andino.log"

# SII Integration
sii:
  environment: "certification" # certification, production
  certificate_path: "./certs/sii.p12"
  certificate_password: "your_cert_password"
  timeout: 30s

# Payment Gateways
payments:
  transbank:
    environment: "integration" # integration, production
    commerce_code: "your_commerce_code"
    api_key: "your_api_key"
  flow:
    api_key: "your_flow_api_key"
    secret_key: "your_flow_secret"
  khipu:
    receiver_id: "your_khipu_receiver_id"
    secret: "your_khipu_secret"

# Background Jobs
worker:
  concurrency: 10
  queues:
    - "default"
    - "sii_processing"
    - "payment_processing"
    - "email_notifications"
```

### 🌍 **Environment Variables Override**

Any configuration can be overridden using environment variables:

```bash
# Server configuration
export ANDINO_SERVER_PORT=8080
export ANDINO_SERVER_HOST="0.0.0.0"

# Database
export ANDINO_DATABASE_URL="postgresql://..."

# SII Configuration
export ANDINO_SII_CERTIFICATE_PASSWORD="your_password"
export ANDINO_SII_ENVIRONMENT="production"

# Payment Gateways
export ANDINO_PAYMENTS_TRANSBANK_COMMERCE_CODE="your_code"
export ANDINO_PAYMENTS_TRANSBANK_API_KEY="your_key"

# Logging
export ANDINO_LOG_LEVEL="debug"
```

### 🔐 **Security Configuration**

```yaml
security:
  jwt:
    secret_key: "${JWT_SECRET_KEY}"
    expiration: "24h"

  rate_limiting:
    enabled: true
    requests_per_minute: 60

  cors:
    enabled: true
    allowed_origins: ["http://localhost:3000"]
    allowed_methods: ["GET", "POST", "PUT", "DELETE"]

  tls:
    enabled: true
    cert_file: "./certs/server.crt"
    key_file: "./certs/server.key"
```

## 📁 Project Structure

Our project follows **Clean Architecture** principles with clear separation of concerns:

```
🏗️ andino/
├── 🚀 cmd/                          # Application entry points
│   ├── 💻 cli/                      # CLI tool for management and automation
│   ├── 🌐 server/                   # HTTP/REST API server
│   └── ⚙️ worker/                   # Background job processor
├── 📝 configs/                      # Configuration files by environment
│   ├── config.dev.yaml              # Development settings
│   ├── config.prod.yaml             # Production settings
│   ├── config.test.yaml             # Testing configuration
│   └── config.yaml                  # Default configuration
├── 🚢 deployments/                  # Infrastructure as Code
│   ├── 🐳 docker/                   # Container definitions
│   │   ├── docker-compose.yaml      # Multi-service setup
│   │   └── Dockerfile               # Application container
│   ├── ☸️ kubernetes/               # K8s manifests and operators
│   ├── 🔧 nginx/                    # Reverse proxy configurations
│   │   ├── conf.d/                  # Virtual host configurations
│   │   ├── html/                    # Error pages and static content
│   │   ├── nginx.conf               # Main nginx configuration
│   │   └── proxy_params             # Proxy parameters
│   └── 🏗️ terraform/                # Infrastructure provisioning
├── 📖 docs/                         # Documentation
│   ├── 🔌 api/                      # OpenAPI/Swagger specifications
│   ├── 🏛️ architecture/             # System design and architecture docs
│   ├── 🚀 deployment/               # Deployment guides and runbooks
│   └── 🔧 plugins/                  # Plugin development documentation
├── 🧠 internal/                     # Private application code
│   ├── 📋 application/              # Application layer (Clean Architecture)
│   │   ├── dto/                     # Data Transfer Objects
│   │   ├── ports/                   # Interfaces and contracts
│   │   ├── services/                # Application services
│   │   └── usecases/                # Business use cases
│   ├── 🏛️ domain/                   # Domain layer (Business Logic)
│   │   ├── common/                  # Shared domain models and value objects
│   │   ├── dte/                     # Electronic Tax Documents domain
│   │   ├── payments/                # Payment processing domain
│   │   └── sii/                     # SII integration domain models
│   ├── 🔧 infrastructure/           # Infrastructure layer
│   │   ├── config/                  # Configuration management
│   │   ├── 🔐 crypto/               # Cryptographic utilities and signing
│   │   ├── 🗄️ database/             # Database connections and migrations
│   │   ├── 🌐 external/             # External service integrations
│   │   │   ├── 📧 email/            # Email service clients
│   │   │   ├── 🏛️ sii/              # SII API client implementation
│   │   │   └── 💳 transbank/        # Transbank payment gateway client
│   │   ├── interfaces/              # Interface adapters
│   │   │   ├── 💻 cli/              # Command-line interface handlers
│   │   │   │   └── commands/        # CLI command implementations
│   │   │   ├── 🔌 grpc/             # gRPC server and handlers
│   │   │   │   ├── handler/         # gRPC service implementations
│   │   │   │   └── proto/           # Protocol buffer definitions
│   │   │   └── 🌐 http/             # HTTP/REST controllers and middleware
│   │   │       ├── controllers/     # HTTP request handlers
│   │   │       ├── middleware/      # HTTP middleware (auth, logging, etc.)
│   │   │       └── router/          # Route definitions and setup
│   │   ├── 📝 logger/               # Logging infrastructure (Zap)
│   │   └── 🗄️ repositories/         # Data persistence layer
│   │       ├── gorm/                # GORM-based implementations
│   │       └── memory/              # In-memory implementations for testing
│   └── interfaces/                  # Legacy interface directory (deprecated)
├── 📜 LICENSE                       # MIT License
├── 🛠️ Makefile                      # Build automation and common tasks
├── 📦 pkg/                          # Public/shared packages
│   ├── 🔌 api/                      # Public API client libraries
│   ├── ❌ errors/                   # Custom error types and handling
│   ├── 🛠️ utils/                    # Utility functions and helpers
│   └── ✅ validators/               # Input validation and sanitization
├── 🔌 plugins/                      # Plugin system architecture
│   ├── 💳 payments/                 # Payment gateway plugins
│   │   ├── flow/                    # Flow payment integration plugin
│   │   ├── khipu/                   # Khipu payment integration plugin
│   │   └── transbank/               # Transbank WebPay integration plugin
│   ├── 📋 README.md                 # Plugin development guide
│   └── 🏛️ sii/                      # SII service plugins
│       ├── certification/           # Digital certificate handling
│       └── reporting/               # Tax reporting and compliance
├── 📜 scripts/                      # Build and utility scripts
│   └── generate-ssl-certs.sh        # SSL certificate generation script
└── 🧪 test/                         # Comprehensive test suites
    ├── 📁 fixtures/                 # Test data and sample files
    │   ├── 🔐 certificates/         # Test SSL certificates and keys
    │   ├── 📄 dte_samples.json      # Sample Electronic Tax Documents
    │   └── 📜 xml_examples/         # XML test files and templates
    ├── 🔗 integration/              # Integration tests
    │   ├── 🌐 api/                  # API endpoint integration tests
    │   ├── 🗄️ database/             # Database integration tests
    │   └── 🌍 external/             # External service integration tests
    └── 🧩 unit/                     # Unit tests by architectural layer
        ├── 📋 application/          # Application layer unit tests
        ├── 🏛️ domain/               # Domain logic unit tests
        └── 🔧 infraestructure/      # Infrastructure layer unit tests
```

### 🏗️ **Architecture Highlights**

- **🧅 Clean Architecture**: Domain-driven design with clear layer separation
- **🔌 Plugin System**: Extensible payment and fiscal integrations
- **📦 Multi-Binary**: CLI, server, and worker applications from single codebase
- **🐳 Container-Ready**: Complete Docker and Kubernetes deployment configurations
- **📊 Observability**: Structured logging, metrics, and comprehensive health checks
- **🧪 Comprehensive Testing**: Unit, integration, and fixture-based testing at all layers

## 🔌 Plugin System

**Andino's plugin architecture** enables seamless integration of new payment gateways, fiscal services, and business logic without modifying the core system.

### 🎯 **Plugin Categories**

#### 💳 **Payment Plugins**

```
plugins/payments/
├── transbank/           # Transbank WebPay Plus, OnePay
├── flow/               # Flow.cl payment gateway
├── khipu/              # Khipu bank transfers
└── mercadopago/        # MercadoPago Chile (coming soon)
```

#### 🏛️ **SII Plugins**

```
plugins/sii/
├── certification/      # Digital certificate management
├── reporting/         # Tax reports and compliance
└── validation/        # Document validation (coming soon)
```

### 🔧 **Creating Custom Plugins**

<details>
<summary><b>📝 Payment Plugin Example</b></summary>

```go
// plugins/payments/mypayment/plugin.go
package mypayment

import (
    "context"
    "github.com/andino-hq/andino/pkg/payments"
)

type MyPaymentPlugin struct {
    apiKey    string
    secretKey string
    baseURL   string
}

func NewMyPaymentPlugin(config Config) *MyPaymentPlugin {
    return &MyPaymentPlugin{
        apiKey:    config.APIKey,
        secretKey: config.SecretKey,
        baseURL:   config.BaseURL,
    }
}

func (p *MyPaymentPlugin) ProcessPayment(ctx context.Context, req payments.PaymentRequest) (*payments.PaymentResponse, error) {
    // Implement payment processing logic
    return &payments.PaymentResponse{
        TransactionID: "txn_123456",
        Status:       payments.StatusPending,
        GatewayURL:   "https://mypayment.com/checkout/xyz",
    }, nil
}

func (p *MyPaymentPlugin) GetPaymentStatus(ctx context.Context, transactionID string) (*payments.PaymentStatus, error) {
    // Implement status checking logic
    return &payments.PaymentStatus{
        TransactionID: transactionID,
        Status:       payments.StatusCompleted,
        Amount:       100000,
    }, nil
}

// Plugin registration
func init() {
    payments.RegisterPlugin("mypayment", NewMyPaymentPlugin)
}
```

</details>

<details>
<summary><b>🏛️ SII Plugin Example</b></summary>

```go
// plugins/sii/myreport/plugin.go
package myreport

import (
    "context"
    "github.com/andino-hq/andino/pkg/sii"
)

type MyReportPlugin struct {
    certificatePath string
    environment     string
}

func (p *MyReportPlugin) GenerateReport(ctx context.Context, req sii.ReportRequest) (*sii.ReportResponse, error) {
    // Implement custom report generation
    return &sii.ReportResponse{
        ReportID: "report_789",
        Format:   "PDF",
        Data:     reportData,
    }, nil
}

// Plugin registration
func init() {
    sii.RegisterPlugin("myreport", NewMyReportPlugin)
}
```

</details>

### 📦 **Plugin Configuration**

```yaml
# configs/config.yaml
plugins:
  payments:
    mypayment:
      enabled: true
      api_key: "${MY_PAYMENT_API_KEY}"
      secret_key: "${MY_PAYMENT_SECRET}"
      base_url: "https://api.mypayment.com"

  sii:
    myreport:
      enabled: true
      certificate_path: "./certs/myreport.p12"
      environment: "production"
```

## 🧪 Testing

### 🎯 Test Coverage & Quality

![Test Coverage](https://img.shields.io/badge/Coverage-85%25-brightgreen?style=flat-square&logo=codecov)
![Tests Passing](https://img.shields.io/badge/Tests-Passing-success?style=flat-square&logo=checkmarx)
![Benchmarks](https://img.shields.io/badge/Benchmarks-Optimized-blue?style=flat-square&logo=speedtest)

### 🏃‍♂️ Running Tests

```bash
# 🧩 Unit tests
make test-unit

# 🔗 Integration tests
make test-integration

# 🌍 End-to-end tests
make test-e2e

# 📊 Test with coverage
make test-coverage

# 🌐 Generate HTML coverage report
make coverage-html

# 🚀 Benchmark tests
make benchmark

# 🏷️ Run specific test suite
go test ./test/unit/domain/... -v -tags=unit

# 🔄 Watch mode (re-run tests on file changes)
make test-watch
```

### 📁 Test Structure & Strategy

```
🧪 test/
├── 📁 fixtures/                     # Test data & samples
│   ├── 🔐 certificates/             # SSL/TLS test certificates
│   │   ├── test-cert.pem           # Test certificate
│   │   ├── test-key.pem            # Test private key
│   │   └── sii-test.p12            # SII test certificate
│   ├── 📄 dte_samples.json          # Electronic document samples
│   └── 📜 xml_examples/             # XML test files and templates
│       ├── invoice.xml             # Sample invoice XML
│       ├── receipt.xml             # Sample receipt XML
│       └── credit_note.xml         # Sample credit note XML
├── 🔗 integration/                  # End-to-end testing
│   ├── 🌐 api/                      # API endpoint integration tests
│   │   ├── dte_test.go             # DTE endpoint tests
│   │   ├── payments_test.go        # Payment endpoint tests
│   │   └── health_test.go          # Health check tests
│   ├── 🗄️ database/                 # Database integration tests
│   │   ├── migrations_test.go      # Migration tests
│   │   └── repositories_test.go    # Repository integration tests
│   └── 🌍 external/                 # External service tests
│       ├── sii_integration_test.go # SII service integration
│       └── payment_gateway_test.go # Payment gateway tests
└── 🧩 unit/                         # Unit test suites
    ├── 📋 application/              # Application layer tests
    │   ├── services_test.go        # Service layer tests
    │   └── usecases_test.go        # Use case tests
    ├── 🏛️ domain/                   # Domain logic tests
    │   ├── dte_test.go             # DTE domain tests
    │   ├── payments_test.go        # Payment domain tests
    │   └── sii_test.go             # SII domain tests
    └── 🔧 infrastructure/           # Infrastructure tests
        ├── database_test.go        # Database layer tests
        ├── external_test.go        # External service tests
        └── repositories_test.go    # Repository tests
```

### 🎯 **Testing Best Practices**

- **🧩 Unit Tests**: Fast, isolated tests for business logic
- **🔗 Integration Tests**: Test component interactions
- **🌍 E2E Tests**: Full system testing with real external services
- **📊 Test Coverage**: Maintain >80% coverage across all layers
- **🔄 Mocking**: Use interfaces for external dependencies
- **📋 Test Data**: Consistent fixtures and sample data

## 🚢 Deployment

### 🐳 **Docker (Recommended)**

![Docker](https://img.shields.io/badge/Docker-Production_Ready-2496ED?style=flat-square&logo=docker&logoColor=white)

```bash
# 🏗️ Build production image
docker build -f deployments/docker/Dockerfile -t andino:latest .

# 🚀 Run with docker-compose
cd deployments/docker && docker-compose up -d

# 📊 Check service health
docker-compose ps
curl http://localhost:8080/health

# 🔍 View logs
docker-compose logs -f andino-server

# 🛑 Stop services
docker-compose down
```

### ☸️ **Kubernetes**

![Kubernetes](https://img.shields.io/badge/Kubernetes-Cloud_Native-326CE5?style=flat-square&logo=kubernetes&logoColor=white)

```bash
# 🚀 Deploy to Kubernetes
kubectl apply -f deployments/kubernetes/

# 📊 Monitor deployment
kubectl get pods -l app=andino -w

# 🔍 Check logs
kubectl logs -f deployment/andino-server

# 📈 Scale deployment
kubectl scale deployment andino-server --replicas=3

# 🔧 Port forward for local access
kubectl port-forward svc/andino-service 8080:8080
```

### 🔧 **Nginx Reverse Proxy**

```bash
# 🔧 Copy nginx configuration
sudo cp deployments/nginx/nginx.conf /etc/nginx/
sudo cp deployments/nginx/conf.d/* /etc/nginx/conf.d/

# ✅ Test configuration
sudo nginx -t

# 🔄 Reload nginx
sudo systemctl reload nginx

# 📊 Check status
sudo systemctl status nginx
```

### ☁️ **Infrastructure as Code**

![Terraform](https://img.shields.io/badge/Terraform-Infrastructure-7B42BC?style=flat-square&logo=terraform&logoColor=white)

```bash
# 🏗️ Initialize Terraform
cd deployments/terraform
terraform init

# 📋 Plan infrastructure changes
terraform plan -var-file="environments/prod.tfvars"

# 🚀 Apply infrastructure
terraform apply -var-file="environments/prod.tfvars"

# 🗑️ Destroy infrastructure (if needed)
terraform destroy -var-file="environments/prod.tfvars"
```

### 🎛️ **Available Services**

| Service            | Port   | Description          | Health Check       | Scaling    |
| ------------------ | ------ | -------------------- | ------------------ | ---------- |
| 🌐 **API Server**  | `8080` | Main REST API        | `/health`          | Horizontal |
| 💻 **CLI Tool**    | -      | Management commands  | `andino --version` | N/A        |
| ⚙️ **Worker**      | `8081` | Background jobs      | `/health`          | Horizontal |
| 🔌 **gRPC Server** | `9090` | High-performance API | Health service     | Horizontal |
| 📊 **Metrics**     | `9091` | Prometheus metrics   | `/metrics`         | Single     |

### 🔒 **Production Checklist**

- [ ] **🔐 Security**: Configure JWT secrets and API keys
- [ ] **🗄️ Database**: Set up PostgreSQL with proper backups
- [ ] **🔴 Cache**: Configure Redis for optimal performance
- [ ] **📊 Monitoring**: Set up Prometheus and Grafana dashboards
- [ ] **📝 Logging**: Configure log aggregation (ELK/Fluentd)
- [ ] **🔒 TLS**: Enable HTTPS with valid SSL certificates
- [ ] **🛡️ Network**: Configure firewalls and security groups
- [ ] **📋 Compliance**: Ensure data protection and audit logging
- [ ] **⚡ Performance**: Configure resource limits and auto-scaling
- [ ] **🔄 Backup**: Set up automated database and configuration backups

## ⚡ Performance

### 📊 **Benchmarks & Metrics**

| Metric                 | Value           | Notes                        |
| ---------------------- | --------------- | ---------------------------- |
| 🚀 **Throughput**      | 10,000+ req/sec | Under optimal conditions     |
| ⚡ **Response Time**   | <50ms average   | 95th percentile <200ms       |
| 💾 **Memory Usage**    | ~64MB baseline  | ~200MB under load            |
| 🗄️ **DB Connections**  | 25 max pool     | Configurable per environment |
| 🔴 **Cache Hit Ratio** | >90%            | Redis-powered caching        |
| 📈 **CPU Usage**       | <30% average    | During normal operations     |

### 🎯 **Performance Optimizations**

- **🔄 Connection Pooling**: Efficient database connection management
- **💾 Smart Caching**: Redis-powered response and query caching
- **📦 Request Batching**: Bulk processing for high-volume operations
- **⚡ Async Processing**: Background jobs for heavy operations
- **🗜️ Response Compression**: Gzip compression for API responses
- **📊 Query Optimization**: Indexed database queries and efficient JOINs

### 🚀 **Scaling Strategies**

```yaml
# kubernetes/hpa.yaml - Horizontal Pod Autoscaler
apiVersion: autoscaling/v2
kind: HorizontalPodAutoscaler
metadata:
  name: andino-server-hpa
spec:
  scaleTargetRef:
    apiVersion: apps/v1
    kind: Deployment
    name: andino-server
  minReplicas: 2
  maxReplicas: 10
  metrics:
    - type: Resource
      resource:
        name: cpu
        target:
          type: Utilization
          averageUtilization: 70
    - type: Resource
      resource:
        name: memory
        target:
          type: Utilization
          averageUtilization: 80
```

## 🔒 Security

### 🛡️ **Security Features**

- **🔐 JWT Authentication**: Secure token-based authentication with configurable expiration
- **🔑 API Key Management**: Multiple API key support with rate limiting per key
- **🛡️ Rate Limiting**: Configurable request throttling and DDoS protection
- **✅ Input Validation**: Comprehensive request/response validation and sanitization
- **🌍 CORS Support**: Configurable cross-origin resource sharing policies
- **🔒 TLS Encryption**: End-to-end encryption with TLS 1.3 support
- **📋 Audit Logging**: Complete request/response logging for compliance
- **🔐 Data Encryption**: Database field encryption for sensitive data

### 🔐 **Authentication & Authorization**

```bash
# Get JWT token
curl -X POST http://localhost:8080/auth/login \
  -H "Content-Type: application/json" \
  -d '{"username":"admin","password":"secure_password"}'

# Use token for API calls
curl -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  http://localhost:8080/api/v1/protected-endpoint
```

### 🛡️ **Security Configuration**

```yaml
# Security settings
security:
  jwt:
    secret_key: "${JWT_SECRET_KEY}" # Use strong secret
    expiration: "24h" # Token expiration

  rate_limiting:
    enabled: true
    global_limit: 1000 # Requests per hour
    per_user_limit: 100 # Requests per user per hour

  cors:
    enabled: true
    allowed_origins:
      - "https://yourdomain.com"
      - "http://localhost:3000" # Development only
    allowed_methods: ["GET", "POST", "PUT", "DELETE"]
    allowed_headers: ["Authorization", "Content-Type"]

  encryption:
    enabled: true
    key: "${ENCRYPTION_KEY}" # AES-256 key

  audit:
    enabled: true
    log_requests: true
    log_responses: false # Don't log sensitive data
    retention_days: 90
```

## 🤝 Contributing

We welcome contributions from the Chilean developer community! **Andino** is built by and for Chilean businesses and developers.

### 🌟 **How to Contribute**

1. **🍴 Fork** the repository on GitHub
2. **📥 Clone** your fork: `git clone https://github.com/YOUR_USERNAME/andino.git`
3. **🌿 Create** a feature branch: `git checkout -b feature/amazing-feature`
4. **💻 Make** your changes and add tests
5. **✅ Ensure** tests pass: `make test`
6. **📝 Commit** your changes: `git commit -m 'Add amazing feature'`
7. **🚀 Push** to your branch: `git push origin feature/amazing-feature`
8. **📋 Open** a Pull Request with a clear description

### 📋 **Development Guidelines**

- **🧪 Testing**: Write tests for new functionality and bug fixes
- **📝 Documentation**: Update docs for API changes and new features
- **🎨 Code Style**: Follow Go conventions and use `gofmt`
- **✅ CI/CD**: Ensure GitHub Actions pass before submitting PR
- **🔒 Security**: Follow security best practices for sensitive code
- **🌍 Internationalization**: Support both Spanish and English documentation

### 🎯 **Contribution Areas**

| Area                      | Description                            | Skill Level  | Impact |
| ------------------------- | -------------------------------------- | ------------ | ------ |
| 🔌 **Plugin Development** | New payment gateways, SII integrations | Intermediate | High   |
| 📊 **Analytics Features** | Business intelligence, reporting       | Advanced     | High   |
| 🧪 **Testing**            | Unit tests, integration tests          | Beginner     | Medium |
| 📚 **Documentation**      | API docs, tutorials, guides            | Beginner     | High   |
| 🎨 **UI/Dashboard**       | Admin dashboard, monitoring            | Intermediate | Medium |
| 🔧 **DevOps**             | CI/CD, deployment, monitoring          | Advanced     | Medium |
| 🐛 **Bug Fixes**          | Issues and bug reports                 | All levels   | High   |

### 📚 **Resources for Contributors**

- 📖 [Contributing Guide](CONTRIBUTING.md) - Detailed contribution guidelines
- 🏗️ [Architecture Docs](docs/architecture/) - System design and patterns
- 🔌 [Plugin Development](docs/plugins/) - How to create plugins
- 💻 [Development Setup](docs/development/) - Local development environment
- 🧪 [Testing Guide](docs/testing/) - Testing strategies and best practices

## 🌟 Community

### 💬 **Join the Chilean Developer Community**

- **💬 Discord**: [Join our community](https://discord.gg/your-invite-link) - Chat with other Chilean developers
- **📋 GitHub Discussions**: [Ask questions and share ideas](https://github.com/andino-hq/andino/discussions)
- **🐛 Issues**: [Report bugs and request features](https://github.com/andino-hq/andino/issues)
- **📧 Newsletter**: Subscribe to our monthly updates and Chilean fintech news

### 🎯 **Community Initiatives**

- **🎓 Workshops**: Monthly online workshops about Chilean fintech integration
- **🏆 Hackathons**: Quarterly hackathons focused on Chilean business automation
- **📚 Knowledge Base**: Community-contributed examples and tutorials
- **🤝 Mentorship**: Connect experienced developers with newcomers
- **🌍 User Groups**: Local meetups in Santiago, Valparaíso, and Concepción

### 🏅 **Contributors**

Thanks to all our amazing contributors who make Andino possible:

[![Contributors](https://contrib.rocks/image?repo=andino-hq/andino)](https://github.com/andino-hq/andino/graphs/contributors)

## 🗺️ Roadmap

### 🎯 **2024 Q4** - Foundation ✅

- [x] Core SII integration (DTE generation, validation)
- [x] Payment gateway unification (Transbank, Flow, Khipu)
- [x] Clean Architecture implementation
- [x] Plugin system architecture
- [x] Comprehensive testing suite
- [x] Docker and Kubernetes deployment

### 📊 **2025 Q1** - Intelligence & Analytics

- [ ] **Business Intelligence Dashboard**
  - Real-time financial metrics and KPIs
  - Interactive charts and visualizations
  - Custom report builder
- [ ] **Automated Compliance Reporting**
  - Monthly and annual tax reports
  - SII report automation
  - Regulatory change notifications
- [ ] **Advanced Analytics**
  - Cash flow forecasting
  - Revenue trend analysis
  - Expense categorization and insights

### 🚀 **2025 Q2** - Expansion & Integration

- [ ] **E-commerce Integration**
  - WooCommerce plugin
  - Shopify app
  - PrestaShop module
  - Custom e-commerce API
- [ ] **Banking APIs**
  - Direct bank account integration
  - Automatic reconciliation
  - Bank transfer automation
- [ ] **Mobile SDKs**
  - iOS Swift SDK
  - Android Kotlin SDK
  - React Native package
  - Flutter plugin

### 🌟 **2025 Q3** - AI & Automation

- [ ] **AI-Powered Features**
  - Intelligent expense categorization
  - Fraud detection algorithms
  - Predictive cash flow analysis
  - Automated tax optimization suggestions
- [ ] **Advanced Automation**
  - Workflow automation builder
  - Conditional business rules
  - Multi-step approval processes
  - Integration with CRM systems

### 🌍 **2025 Q4** - Ecosystem & Enterprise

- [ ] **Marketplace & Extensions**
  - Community plugin marketplace
  - Third-party integrations
  - Custom business logic extensions
- [ ] **Enterprise Features**
  - Multi-tenant architecture
  - Advanced security controls
  - Dedicated support channels
  - SLA guarantees and monitoring
- [ ] **Regional Expansion**
  - Argentina integration support
  - Peru fiscal system support
  - Colombia payment gateways

### 🔮 **Future Vision** (2026+)

- [ ] **Blockchain Integration**: Cryptocurrency payments and NFT invoicing
- [ ] **International Expansion**: Support for other Latin American countries
- [ ] **AI Business Assistant**: Natural language business query interface
- [ ] **IoT Integration**: Smart device payment and inventory tracking

## 📋 Changelog

See [CHANGELOG.md](CHANGELOG.md) for detailed release notes, breaking changes, and migration guides.

### 🚀 Latest Release Highlights

**v0.1.0** - Foundation Release

- ✨ Complete SII integration with DTE support
- 💳 Unified payment gateway API
- 🔌 Plugin architecture implementation
- 🐳 Production-ready Docker deployment
- 📚 Comprehensive API documentation

## 📄 License

This project is licensed under the **MIT License** - see the [LICENSE](LICENSE) file for details.

### 📝 License Summary

- ✅ **Commercial Use**: Use Andino in commercial projects
- ✅ **Modification**: Modify and adapt the source code
- ✅ **Distribution**: Distribute original or modified versions
- ✅ **Private Use**: Use Andino for private projects
- ❌ **Liability**: Authors not liable for damages
- ❌ **Warranty**: No warranty provided

## 🙏 Acknowledgments

- 🇨🇱 **Chilean Fintech Community**: For feedback, contributions, and support
- 🏛️ **SII (Servicio de Impuestos Internos)**: For comprehensive API documentation
- 🏦 **Chilean Payment Providers**: Transbank, Flow, Khipu for developer support
- 🏢 **Chilean SMEs**: For real-world testing and feedback
- 💻 **Open Source Community**: For the amazing tools and libraries we build upon
- 👥 All our [contributors](https://github.com/andino-hq/andino/graphs/contributors) and community members

---

<div align="center">

**🇨🇱 Made with ❤️ in Chile**

_Empowering Chilean businesses through technology_

[⭐ Star us on GitHub](https://github.com/andino-hq/andino) • [🐛 Report Bug](https://github.com/andino-hq/andino/issues) • [💡 Request Feature](https://github.com/andino-hq/andino/issues) • [💬 Join Discord](https://discord.gg/your-invite-link)

**"We simplify Chilean tax technology so you can focus on growing your business."**

</div>
