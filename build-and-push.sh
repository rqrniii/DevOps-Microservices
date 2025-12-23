#!/bin/bash
# Build and Push Docker Images to ECR (AMD64 for EC2)
# Usage: ./build-and-push.sh [version-tag]

set -e

# Configuration
export AWS_REGION="me-south-1"
export AWS_ACCOUNT_ID="405894839882"
export ECR_REGISTRY="${AWS_ACCOUNT_ID}.dkr.ecr.${AWS_REGION}.amazonaws.com"

# Version tag (default to latest if not provided)
VERSION_TAG="${1:-latest}"

# Colors
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m'

echo -e "${YELLOW}🐳 TaskGenius - Build and Push Docker Images (AMD64)${NC}"
echo "======================================================="
echo "ECR Registry: $ECR_REGISTRY"
echo "Version Tag: $VERSION_TAG"
echo "Platform: linux/amd64 (for EC2 x86_64)"
echo ""

# Check if docker is running
if ! docker info > /dev/null 2>&1; then
    echo -e "${RED}❌ Docker is not running. Please start Docker.${NC}"
    exit 1
fi

# Setup buildx for multi-platform
echo -e "${YELLOW}Setting up Docker Buildx${NC}"
docker buildx create --use --name multiarch 2>/dev/null || docker buildx use multiarch
docker buildx inspect --bootstrap
echo ""

# Login to ECR
echo -e "${YELLOW}Step 1: Login to ECR${NC}"
echo "-----------------------------------"
aws ecr get-login-password --region $AWS_REGION | \
    docker login --username AWS --password-stdin $ECR_REGISTRY

if [ $? -eq 0 ]; then
    echo -e "${GREEN}✅ Successfully logged in to ECR${NC}"
else
    echo -e "${RED}❌ Failed to login to ECR${NC}"
    exit 1
fi
echo ""

# Build and Push Auth Service
echo -e "${YELLOW}Step 2: Building Auth Service (AMD64)${NC}"
echo "-----------------------------------"
docker buildx build --platform linux/amd64 \
    -f services/auth-service/Dockerfile \
    -t $ECR_REGISTRY/auth-service:$VERSION_TAG \
    -t $ECR_REGISTRY/auth-service:latest \
    --push .

if [ $? -eq 0 ]; then
    echo -e "${GREEN}✅ Auth service built and pushed${NC}"
else
    echo -e "${RED}❌ Failed to build auth service${NC}"
    exit 1
fi
echo ""

# Build and Push Todo Service
echo -e "${YELLOW}Step 3: Building Todo Service (AMD64)${NC}"
echo "-----------------------------------"
docker buildx build --platform linux/amd64 \
    -f services/todo-service/Dockerfile \
    -t $ECR_REGISTRY/todo-service:$VERSION_TAG \
    -t $ECR_REGISTRY/todo-service:latest \
    --push .

if [ $? -eq 0 ]; then
    echo -e "${GREEN}✅ Todo service built and pushed${NC}"
else
    echo -e "${RED}❌ Failed to build todo service${NC}"
    exit 1
fi
echo ""

# Build and Push AI Service
echo -e "${YELLOW}Step 4: Building AI Service (AMD64)${NC}"
echo "-----------------------------------"
docker buildx build --platform linux/amd64 \
    -f services/ai-service/Dockerfile \
    -t $ECR_REGISTRY/ai-service:$VERSION_TAG \
    -t $ECR_REGISTRY/ai-service:latest \
    --push .

if [ $? -eq 0 ]; then
    echo -e "${GREEN}✅ AI service built and pushed${NC}"
else
    echo -e "${RED}❌ Failed to build AI service${NC}"
    exit 1
fi
echo ""

# Build and Push Frontend
echo -e "${YELLOW}Step 5: Building Frontend (AMD64)${NC}"
echo "-----------------------------------"
docker buildx build --platform linux/amd64 \
    -f frontend/Dockerfile \
    -t $ECR_REGISTRY/frontend:$VERSION_TAG \
    -t $ECR_REGISTRY/frontend:latest \
    --push \
    ./frontend

if [ $? -eq 0 ]; then
    echo -e "${GREEN}✅ Frontend built and pushed${NC}"
else
    echo -e "${RED}❌ Failed to build frontend${NC}"
    exit 1
fi
echo ""

# Summary
echo -e "${GREEN}🎉 All images built and pushed successfully (AMD64)!${NC}"
echo "====================================================="
echo ""
echo "📦 Images in ECR (version: $VERSION_TAG):"
echo "  - $ECR_REGISTRY/auth-service:$VERSION_TAG"
echo "  - $ECR_REGISTRY/todo-service:$VERSION_TAG"
echo "  - $ECR_REGISTRY/ai-service:$VERSION_TAG"
echo "  - $ECR_REGISTRY/frontend:$VERSION_TAG"
echo ""
echo "🔍 Verify images:"
echo "  aws ecr describe-images --repository-name auth-service --region $AWS_REGION"
echo ""
echo -e "${YELLOW}Next steps:${NC}"
echo "1. Update Kubernetes secrets:"
echo "   kubectl create secret generic taskgenius-secrets \\"
echo "     --from-literal=jwt-secret=YOUR_JWT_SECRET \\"
echo "     --from-literal=openai-api-key=YOUR_OPENAI_KEY \\"
echo "     -n taskgenius --dry-run=client -o yaml | kubectl apply -f -"
echo ""
echo "2. Deploy to Kubernetes:"
echo "   kubectl apply -f k8s/deployments.yaml"
echo ""
echo "3. Restart deployments (if updating existing):"
echo "   kubectl rollout restart deployment -n taskgenius"
echo ""
echo "4. Check deployment status:"
echo "   kubectl get pods -n taskgenius"
echo "   kubectl get svc -n taskgenius"
