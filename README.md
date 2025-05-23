# Order Packs Calculator

A Go application to calculate optimal pack combinations for fulfilling item orders with a modern web interface.

## 🌐 Live Demo

**🚀 Try it now:** [https://past-birthday-production.up.railway.app/](https://past-birthday-production.up.railway.app/)

The application is live and fully functional - you can test the pack calculator immediately!

## 🚀 Quick Deploy

[![Deploy on Railway](https://railway.app/button.svg)](https://railway.app/new/template?template=https://github.com/MaoDaGreith/SoftwareEngineerChallenge)

**One-click deployment to Railway** - Fork this repo and click the button above!

## ✨ Features

- **Smart Algorithm**: Minimizes over-shipping and pack count
- **Web Interface**: Modern, responsive UI
- **REST API**: Clean JSON API for integration
- **Configurable**: Environment variables, JSON config, or API parameters
- **Docker Ready**: Multi-stage builds for production
- **CI/CD**: GitHub Actions for automated testing and deployment

## 🛠 Setup

### Quick Start with Docker Compose

```bash
git clone https://github.com/MaoDaGreith/SoftwareEngineerChallenge.git
cd SoftwareEngineerChallenge
docker-compose up
```

Visit `http://localhost:8080` to use the web interface.

### Local Development

```bash
go run cmd/server/main.go
```

### Docker

Build and run:
```bash
docker build -t order-packs-calculator .
docker run -p 8080:8080 order-packs-calculator
```

With custom pack sizes:
```bash
docker run -p 8080:8080 -e PACK_SIZES=250,500,1000,2000,5000 order-packs-calculator
```

## 🌐 Deployment

See [DEPLOYMENT.md](DEPLOYMENT.md) for detailed deployment instructions including:

- **Railway** (recommended)
- **Render**
- **Fly.io**
- **Google Cloud Run**
- **Heroku**

## 📱 Usage

### Web Interface
Visit the [live demo](https://past-birthday-production.up.railway.app/) to use the interactive web interface.

### API

POST to `/api/packs/calculate`:

```bash
curl -X POST https://past-birthday-production.up.railway.app/api/packs/calculate \
  -H "Content-Type: application/json" \
  -d '{"pack_sizes":[250,500,1000],"order_amount":1250}'
```

**Request:**
```json
{
  "pack_sizes": [250, 500, 1000],
  "order_amount": 1250
}
```

**Response:**
```json
{
  "packs": { "1000": 1, "250": 1 }
}
```

**Error:**
```json
{
  "error": "Cannot fulfill order exactly with given pack sizes."
}
```

## ⚙️ Configuration

### Environment Variables
- `PORT`: Server port (default: 8080)
- `PACK_SIZES`: Default pack sizes (e.g., "250,500,1000,2000,5000")

### Priority Order
1. API request `pack_sizes` parameter
2. `PACK_SIZES` environment variable  
3. Hardcoded defaults: [250, 500, 1000, 2000, 5000]

## 🧪 Testing

```bash
go test ./...
```

## 🏗 Architecture

```
├── cmd/server/          # Application entrypoint
├── internal/
│   ├── api/            # HTTP handlers and routing
│   ├── config/         # Configuration management
│   └── packs/          # Core calculation logic
├── web/                # Frontend assets
├── .github/workflows/  # CI/CD
└── Dockerfile         # Container definition
```

## 🎯 Algorithm

The pack calculator uses dynamic programming to:

1. **Minimize over-shipping** (fewest extra items)
2. **Minimize pack count** (fewest total packs)
3. **Prefer larger packs** (tie-breaking)

## 📈 Example

For order of 1250 items with packs [250, 500, 1000]:
- ✅ Returns: `{1000: 1, 250: 1}` (1250 items, 2 packs)
- ❌ Not: `{500: 3}` (1500 items, 3 packs, more over-shipping)
