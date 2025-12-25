# DevOps-Microservices

TaskGenius - AI-Powered Task Management

📋 Overview
TaskGenius is a full-stack application with AI-powered task generation, deployed on a self-managed Kubernetes cluster on AWS. The application demonstrates modern cloud-native architecture with microservices, container orchestration, and CI/CD automation.
Key Features

- User Authentication - Secure JWT-based authentication
- Task Management - Full CRUD operations for todos
- AI Integration - OpenAI GPT-4 powered task generation
- Cloud Native - Deployed on Kubernetes with high availability
- Auto-scaling - Multiple replicas with load balancing

🔧 Local Development 

Backend Services

# Auth Service

cd services/auth-service

go run main.go

# Todo Service

cd services/todo-service

go run main.go

# AI Service

cd services/ai-service

go run main.go

# Frontend

cd services/frontend

bun install

bun run dev


# Frontend

VITE_AUTH_SERVICE_URL=http://localhost:8080

VITE_TODO_SERVICE_URL=http://localhost:8081

VITE_AI_SERVICE_URL=http://localhost:8082


# 📡 API Endpoints

# Authentication

POST /api/auth/register

POST /api/auth/login

GET  /api/auth/me

# Todo Management

GET    /api/todos

POST   /api/todos

PUT    /api/todos/:id/toggle

DELETE /api/todos/:id


# AI Generation

POST /api/ai/generate
