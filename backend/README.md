# My Coffee Log Backend

Go + Gin + GORM + MySQL + Redis + JWT

## Quick Start

### Local Development

1. Ensure MySQL 8 and Redis are running locally
2. Copy `.env.example` to `.env` and adjust values
3. Run:

```bash
go run ./cmd/server
```

Server starts at `http://localhost:8080`

### Docker Compose

From project root:

```bash
docker-compose up
```

## API Endpoints

| Method | Path | Auth | Description |
|--------|------|------|-------------|
| POST | /api/v1/auth/register | No | Register |
| POST | /api/v1/auth/login | No | Login |
| GET | /api/v1/users/me | Yes | Current user |
| PUT | /api/v1/users/me | Yes | Update user |
| POST | /api/v1/coffee-logs | Yes | Create log |
| GET | /api/v1/coffee-logs | Yes | List logs |
| GET | /api/v1/coffee-logs/:id | Yes | Get log detail |
| PUT | /api/v1/coffee-logs/:id | Yes | Update log |
| DELETE | /api/v1/coffee-logs/:id | Yes | Delete log |
| GET | /api/v1/stats/overview | Yes | Stats overview |
| GET | /api/v1/stats/flavor-profile | Yes | Flavor profile |
| GET | /api/v1/stats/monthly | Yes | Monthly stats |
| POST | /api/v1/ai/flavor-summary | Yes | AI flavor summary |

## Architecture

```
Handler → Service → Repository → Model (GORM)
```

All user data is isolated by `user_id`.
