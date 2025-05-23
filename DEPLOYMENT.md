# Deployment Guide

This application can be deployed to various cloud platforms. Here are the recommended options:

## 🚀 Quick Deploy Options

### Option 1: Railway (Recommended - Easiest)

[![Deploy on Railway](https://railway.app/button.svg)](https://railway.app/new/template?template=https://github.com/MaoDaGreith/SoftwareEngineerChallenge)

1. **Fork this repository**
2. **Sign up at [Railway.app](https://railway.app)**
3. **Connect your GitHub account**
4. **Deploy from GitHub repository**
5. **Railway will automatically detect the Dockerfile and deploy**

### Option 2: Render

[![Deploy to Render](https://render.com/images/deploy-to-render-button.svg)](https://render.com/deploy)

1. **Fork this repository**
2. **Sign up at [Render.com](https://render.com)**
3. **Create new Web Service**
4. **Connect your GitHub repository**
5. **Render will build using the Dockerfile**

### Option 3: Fly.io

```bash
# Install Fly CLI
curl -L https://fly.io/install.sh | sh

# Login and deploy
fly auth login
fly launch
fly deploy
```

### Option 4: Google Cloud Run

```bash
# Build and push to Google Container Registry
gcloud builds submit --tag gcr.io/PROJECT_ID/order-packs-calculator

# Deploy to Cloud Run
gcloud run deploy --image gcr.io/PROJECT_ID/order-packs-calculator --platform managed
```

## 🔧 Manual Deployment

### Prerequisites
- Docker installed
- Port 8080 available

### Build and Run

```bash
# Clone repository
git clone https://github.com/MaoDaGreith/SoftwareEngineerChallenge.git
cd SoftwareEngineerChallenge

# Build Docker image
docker build -t order-packs-calculator .

# Run container
docker run -p 8080:8080 order-packs-calculator
```

## 🌍 Environment Variables

The application supports these optional environment variables:

- `PORT`: Server port (default: 8080)
- `PACK_SIZES`: Comma-separated default pack sizes (e.g., "250,500,1000,2000,5000")

Example:
```bash
docker run -p 8080:8080 -e PACK_SIZES="100,250,500,1000" order-packs-calculator
```

## 🧪 GitHub Actions CI/CD

The repository includes GitHub Actions workflow that:
- Runs tests on every push/PR
- Builds Docker image
- Can be extended to deploy automatically

## 📱 API Endpoints

Once deployed, your application will have:

- **Frontend**: `https://your-app.railway.app/`
- **API**: `https://your-app.railway.app/api/packs/calculate`

## 🔍 Health Check

The application includes a health check endpoint at `/` that can be used by load balancers and monitoring services.

## 📊 Example Request

```bash
curl -X POST https://your-app.railway.app/api/packs/calculate \
  -H "Content-Type: application/json" \
  -d '{"pack_sizes": [250, 500, 1000], "order_amount": 1250}'
```

## 🎯 Zero-Downtime Updates

For production deployments:
1. Use health checks for rolling updates
2. Consider using multiple instances
3. Implement proper logging and monitoring 