# TnP Portal System Design Document

## Table of Contents

1. [System Overview](#1-system-overview)
2. [Architecture](#2-architecture)
3. [Backend Design](#3-backend-design)
4. [Frontend Design](#4-frontend-design)
5. [Database Design](#5-database-design)
6. [API Design](#6-api-design)
7. [Security](#7-security)
8. [Deployment](#8-deployment)
9. [Monitoring and Logging](#9-monitoring-and-logging)
10. [Future Improvements](#10-future-improvements)

## 1. System Overview

### 1.1 Purpose

The TnP (Training and Placement) Portal is a comprehensive platform designed to streamline the placement process for educational institutions. It facilitates interaction between students, companies, and placement officers.

### 1.2 Key Features

- Student Profile Management
- Job Portal
- Company Management
- Application Tracking
- Interview Scheduling
- Document Management
- Analytics and Reporting
- Notification System

### 1.3 System Requirements

- High Availability
- Scalability
- Security
- Real-time Updates
- Mobile Responsiveness
- Multi-tenant Support

## 2. Architecture

### 2.1 High-Level Architecture

```
┌─────────────────┐     ┌─────────────────┐     ┌─────────────────┐
│                 │     │                 │     │                 │
│  Frontend       │◄───►│  API Gateway    │◄───►│  Backend        │
│  (NextJS)       │     │  (Nginx)        │     │  (Go)           │
│                 │     │                 │     │                 │
└─────────────────┘     └─────────────────┘     └────────┬────────┘
                                                         │
                                                         ▼
                                                 ┌─────────────────┐
                                                 │                 │
                                                 │  Database       │
                                                 │  (PostgreSQL)   │
                                                 │                 │
                                                 └─────────────────┘
```

### 2.2 Technology Stack

#### Backend

- **Language**: Go 1.22+
- **Framework**: gofr
- **Database**: PostgreSQL 15+
- **Cache**: Redis
- **Message Queue**: RabbitMQ
- **Search**: Elasticsearch
- **Object Storage**: MinIO/S3

#### Frontend

- **Framework**: React 18+
- **State Management**: Redux Toolkit
- **UI Library**: Material-UI
- **Form Handling**: React Hook Form
- **API Client**: Axios
- **Real-time**: Socket.io

#### Infrastructure

- **Containerization**: Docker
- **Orchestration**: Kubernetes
- **CI/CD**: GitHub Actions
- **Monitoring**: Prometheus + Grafana
- **Logging**: ELK Stack
- **CDN**: Cloudflare

## 3. Backend Design

### 3.1 Project Structure

```
backend/
├── cmd/
│   ├── api/              # Main API service
│   ├── worker/           # Background worker service
│   ├── analytics/        # Analytics service
│   ├── notification/     # Notification service
│   └── document/         # Document processing service
├── internal/
│   ├── config/          # Configuration management
│   ├── models/          # Data models
│   ├── repository/      # Database operations
│   ├── service/         # Business logic
│   ├── handler/         # HTTP handlers
│   ├── middleware/      # Middleware functions
│   └── utils/           # Utility functions
├── pkg/                 # Reusable packages
├── migrations/          # Database migrations
└── static/             # Static files
```

### 3.2 Microservices Architecture

#### 3.2.1 Service Overview

```
┌─────────────────┐     ┌─────────────────┐     ┌─────────────────┐
│                 │     │                 │     │                 │
│  API Service    │◄───►│  Worker Service │◄───►│  Analytics      │
│  (Go)           │     │  (Go)           │     │  Service (Go)   │
│                 │     │                 │     │                 │
└────────┬────────┘     └────────┬────────┘     └────────┬────────┘
         │                       │                       │
         ▼                       ▼                       ▼
┌─────────────────┐     ┌─────────────────┐     ┌─────────────────┐
│                 │     │                 │     │                 │
│  Notification   │◄───►│  Document       │◄───►│  Database       │
│  Service (Go)   │     │  Service (Go)   │     │  (PostgreSQL)   │
│                 │     │                 │     │                 │
└─────────────────┘     └─────────────────┘     └─────────────────┘
```

#### 3.2.2 Service Communication

```
┌─────────────────────────────────────────────────────────────────┐
│                         NATS Message Queue                      │
└───────────────────────────────┬─────────────────────────────────┘
                                │
        ┌───────────────────────┴───────────────────────┐
        │                                               │
┌───────▼───────┐                               ┌───────▼───────┐
│               │                               │               │
│  API Service  │                               │  Worker       │
│               │                               │  Service      │
└───────┬───────┘                               └───────┬───────┘
        │                                               │
        ▼                                               ▼
┌───────────────┐     ┌───────────────┐     ┌───────────────┐
│               │     │               │     │               │
│  Notification │◄───►│  Document     │◄───►│  Analytics    │
│  Service      │     │  Service      │     │  Service      │
│               │     │               │     │               │
└───────────────┘     └───────────────┘     └───────────────┘
```

#### 3.2.3 Data Flow

```
┌─────────────────────────────────────────────────────────────────┐
│                         API Service                             │
└───────────────────────────────┬─────────────────────────────────┘
                                │
        ┌───────────────────────┴───────────────────────┐
        │                                               │
┌───────▼───────┐                               ┌───────▼───────┐
│               │                               │               │
│  Job          │                               │  User         │
│  Processing   │                               │  Management   │
│               │                               │               │
└───────┬───────┘                               └───────┬───────┘
        │                                               │
        ▼                                               ▼
┌───────────────┐     ┌───────────────┐     ┌───────────────┐
│               │     │               │     │               │
│  Document     │◄───►│  Notification │◄───►│  Analytics    │
│  Processing   │     │  Service      │     │  Processing   │
│               │     │               │     │               │
└───────────────┘     └───────────────┘     └───────────────┘
```

#### 3.2.4 Service Responsibilities

1. **API Service**
   - HTTP request handling
   - Authentication/Authorization
   - Request validation
   - Response formatting
   - Service orchestration

2. **Worker Service**
   - Background job processing
   - Queue management
   - Task scheduling
   - Error handling
   - Retry mechanisms

3. **Analytics Service**
   - Data aggregation
   - Report generation
   - Statistical analysis
   - Trend detection
   - Performance metrics

4. **Notification Service**
   - Email notifications
   - Push notifications
   - SMS notifications
   - Notification templates
   - Delivery tracking

5. **Document Service**
   - File upload/download
   - Document processing
   - OCR processing
   - Format conversion
   - Storage management

#### 3.2.5 Service Dependencies

```
┌─────────────────────────────────────────────────────────────────┐
│                         Shared Dependencies                     │
└───────────────────────────────┬─────────────────────────────────┘
                                │
        ┌───────────────────────┴───────────────────────┐
        │                                               │
┌───────▼───────┐                               ┌───────▼───────┐
│               │                               │               │
│  Database     │                               │  Cache        │
│  (PostgreSQL) │                               │  (Redis)      │
│               │                               │               │
└───────┬───────┘                               └───────┬───────┘
        │                                               │
        ▼                                               ▼
┌───────────────┐     ┌───────────────┐     ┌───────────────┐
│               │     │               │     │               │
│  Message      │◄───►│  Object       │◄───►│  Search       │
│  Queue        │     │  Storage      │     │  Engine       │
│  (NATS)       │     │  (MinIO/S3)   │     │  (Elastic)    │
│               │     │               │     │               │
└───────────────┘     └───────────────┘     └───────────────┘
```

## 4. Frontend Design

### 4.1 Project Structure

```
frontend/
├── src/
│   ├── components/         # Reusable UI components
│   │   ├── common/        # Shared components
│   │   ├── layout/        # Layout components
│   │   └── features/      # Feature-specific components
│   ├── pages/             # Page components
│   ├── services/          # API service calls
│   ├── store/             # State management
│   ├── hooks/             # Custom React hooks
│   ├── utils/             # Utility functions
│   └── types/             # TypeScript types
├── public/                # Static assets
└── package.json
```

## 5. Database Design

### 5.1 Schema Design

```sql
-- Users table
CREATE TABLE users (
    id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash BYTEA NOT NULL,
    role VARCHAR(50) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

-- Students table
CREATE TABLE students (
    id UUID PRIMARY KEY,
    user_id UUID REFERENCES users(id),
    degree VARCHAR(100) NOT NULL,
    department VARCHAR(100) NOT NULL,
    major VARCHAR(100),
    graduation_year INTEGER NOT NULL,
    contact_info VARCHAR(255),
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

-- Companies table
CREATE TABLE companies (
    id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    website VARCHAR(255),
    industry VARCHAR(100),
    size VARCHAR(50),
    location VARCHAR(255),
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

-- Jobs table
CREATE TABLE jobs (
    id UUID PRIMARY KEY,
    company_id UUID REFERENCES companies(id),
    title VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    requirements TEXT[] NOT NULL,
    location VARCHAR(255) NOT NULL,
    salary_min DECIMAL(10,2),
    salary_max DECIMAL(10,2),
    deadline TIMESTAMP NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

-- Applications table
CREATE TABLE applications (
    id UUID PRIMARY KEY,
    job_id UUID REFERENCES jobs(id),
    student_id UUID REFERENCES students(id),
    status VARCHAR(50) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

-- Interviews table
CREATE TABLE interviews (
    id UUID PRIMARY KEY,
    application_id UUID REFERENCES applications(id),
    date_time TIMESTAMP NOT NULL,
    type VARCHAR(50) NOT NULL,
    location VARCHAR(255),
    notes TEXT,
    status VARCHAR(50) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);
```

### 5.2 Indexes

```sql
-- Performance optimization
CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_jobs_company ON jobs(company_id);
CREATE INDEX idx_applications_job ON applications(job_id);
CREATE INDEX idx_applications_student ON applications(student_id);
CREATE INDEX idx_interviews_application ON interviews(application_id);
```

## 6. API Design

### 6.1 RESTful Endpoints

```
/v1/
├── auth/
│   ├── register
│   └── login
├── users/
│   ├── profile
│   └── {userId}
├── students/
│   ├── profile
│   └── {studentId}
├── companies/
│   ├── profile
│   └── {companyId}
├── jobs/
│   ├── {jobId}
│   └── {jobId}/apply
├── applications/
│   ├── {applicationId}
│   └── {applicationId}/status
├── interviews/
│   ├── schedule
│   └── {interviewId}
└── analytics/
    ├── placements
    ├── students/performance
    └── companies/trends
```

### 6.2 WebSocket Endpoints

```
/ws/
├── notifications
└── interviews
```

## 7. Security

### 7.1 Authentication

- JWT-based authentication
- Refresh token rotation
- Secure password hashing (bcrypt)
- Rate limiting
- IP blocking for suspicious activity

### 7.2 Authorization

- Role-based access control (RBAC)
- Resource-level permissions
- API key management for companies

### 7.3 Data Protection

- Encryption at rest
- Encryption in transit (TLS)
- Secure file uploads
- Data masking for sensitive information

## 8. Deployment

### 8.1 Infrastructure

```
[Cloudflare (Edge)]
    │
    ├── [CDN]
    ├── [SSL/TLS]
    ├── [DDoS Protection]
    └── [Load Balancer]
            │
            ├── [Oracle Cloud (Primary)]
            │       │
            │       ├── [API Server]
            │       ├── [PostgreSQL]
            │       ├── [NATS Server]
            │       ├── [Redis Cache]
            │       └── [MinIO/S3]
            │
            └── [Render (Secondary)]
                    │
                    ├── [Worker Services]
                    │       ├── [Job Processing]
                    │       ├── [Notifications]
                    │       └── [Analytics]
                    │
                    ├── [Document Processing]
                    └── [Backup Services]
```

### 8.2 Deployment Strategy

#### 8.2.1 Oracle Cloud (Primary Infrastructure)
- **Compute**: Always Free AMD-based VMs
- **Database**: Always Free Autonomous Database
- **Storage**: Always Free Object Storage
- **Networking**: Always Free Load Balancer
- **Services**:
  - Main API Server
  - PostgreSQL Database
  - NATS Message Queue
  - Redis Cache
  - MinIO/S3 Storage

#### 8.2.2 Render (Background Processing)
- **Free Tier**:
  - 750 hours/month runtime
  - Automatic HTTPS
  - Custom domains
  - Automatic deployments
- **Services**:
  - Worker Services
  - Analytics Processing
  - Document Processing
  - Notification Service
  - Backup Services

#### 8.2.3 Cloudflare (Edge Services)
- **Free Tier**:
  - Global CDN
  - SSL/TLS certificates
  - DDoS protection
  - DNS management
  - Workers (serverless)
- **Features**:
  - Edge caching
  - Load balancing
  - WAF protection
  - Rate limiting
  - Analytics

### 8.3 Deployment Process

#### 8.3.1 Oracle Cloud Setup
```bash
# Create VM instance
# Install Docker and Docker Compose
sudo dnf install docker-engine docker-compose

# Clone repository
git clone https://github.com/your-username/tnp-portal.git
cd tnp-portal/backend

# Create .env file
cp .env.example .env
# Edit .env with your values

# Start services
docker-compose -f docker-compose.prod.yml up -d
```

#### 8.3.2 Render Setup
```bash
# Connect GitHub repository
# Deploy worker service
render.yaml deploy

# Deploy backup service
render.yaml deploy
```

#### 8.3.3 Cloudflare Setup
```bash
# Add DNS records
# A record: tnp-api.yourdomain.com -> Oracle VM IP
# CNAME: www.tnp-api.yourdomain.com -> tnp-api.yourdomain.com

# Configure SSL/TLS
# Set encryption mode to "Full"
# Create Origin Certificate
# Install certificate on Oracle VM

# Configure caching rules
# Set up load balancing
# Configure WAF rules
```

### 8.4 Monitoring and Scaling

#### 8.4.1 Oracle Cloud
- VM monitoring
- Database monitoring
- Resource utilization
- Auto-scaling rules

#### 8.4.2 Render
- Worker monitoring
- Service health checks
- Log aggregation
- Performance metrics

#### 8.4.3 Cloudflare
- CDN analytics
- Security events
- Performance metrics
- Traffic patterns

### 8.5 Cost Optimization

#### 8.5.1 Oracle Cloud
- Use Always Free resources
- Optimize VM sizing
- Implement auto-scaling
- Use reserved instances

#### 8.5.2 Render
- Monitor runtime hours
- Optimize worker processes
- Use efficient scheduling
- Implement caching

#### 8.5.3 Cloudflare
- Use free tier features
- Optimize caching rules
- Implement rate limiting
- Use workers for edge computing

### 8.6 Security Considerations

#### 8.6.1 Oracle Cloud
- Network security groups
- VCN configuration
- Database security
- Access control

#### 8.6.2 Render
- Environment variables
- Service isolation
- Access control
- Secure communication

#### 8.6.3 Cloudflare
- WAF rules
- DDoS protection
- SSL/TLS configuration
- Security headers

## 9. Monitoring and Logging

### 9.1 Metrics

- Application metrics (Prometheus)
- Business metrics (custom)
- Infrastructure metrics (Cloud provider)

### 9.2 Logging

- Structured logging (ELK Stack)
- Error tracking (Sentry)
- Audit logging

### 9.3 Alerts

- System health alerts
- Business metric alerts
- Security alerts

## 10. Future Improvements

### 10.1 Technical Improvements

1. **Microservices Architecture**
   - Split into smaller, focused services
   - Service mesh implementation
   - Event-driven architecture

2. **Performance Optimization**
   - Implement GraphQL
   - Add CDN for static assets
   - Implement edge computing

3. **Scalability**
   - Implement sharding
   - Add read replicas
   - Implement caching strategies

### 10.2 Feature Improvements

1. **AI/ML Integration**
   - Job matching algorithm
   - Resume analysis
   - Interview preparation

2. **Enhanced Analytics**
   - Predictive analytics
   - Real-time dashboards
   - Custom reports

3. **Mobile Application**
   - Native mobile apps
   - Offline capabilities
   - Push notifications

### 10.3 Security Improvements

1. **Advanced Security**
   - Implement 2FA
   - Add biometric authentication
   - Implement zero-trust architecture

2. **Compliance**
   - GDPR compliance
   - Data residency
   - Audit trails

### 10.4 User Experience

1. **Accessibility**
   - WCAG compliance
   - Screen reader support
   - Keyboard navigation

2. **Internationalization**
   - Multi-language support
   - Localization
   - Timezone handling

## Conclusion

This system design provides a robust foundation for the TnP Portal, with clear separation of concerns, secure communication, and scalable architecture. The design is modular and can be extended to accommodate future requirements and improvements.
