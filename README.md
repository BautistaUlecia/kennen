# kennen

Backend for unnamed league of legends group tracker

## Prerequisites

- Go 1.25.3 or higher (for local development)
- Docker (for containerized deployment)
- Riot Games API Key

## Environment Variables

Create a `.env` file in the root directory with the following variable:

```
RIOT_API_KEY=your_riot_api_key_here
```

Ask the team to get the API key.

## Running Locally

1. **Install dependencies:**

```bash
   go mod download
```

2. **Set up environment variables:**
   Create a `.env` file with your Riot API key (see Environment Variables section above).

3. **Run the application:**

   ```bash
   go run cmd/kennen/main.go
   ```

4. **Verify the application is running:**
   The server will start on port 8080. You can check the health endpoint:
   ```bash
   curl http://localhost:8080/health
   ```

## Running with Docker

1. **Build the Docker image:**

   ```bash
   docker build -t kennen .
   ```

2. **Run the container:**

   ```bash
   docker run --name kennen -p 8080:8080 -e RIOT_API_KEY=your_riot_api_key_here kennen
   ```

   Alternatively, use a `.env` file:

   ```bash
   docker run --name kennen -p 8080:8080 --env-file .env kennen
   ```

3. **Verify the application is running:**
   ```bash
   curl http://localhost:8080/health
   ```

## API Endpoints

- `GET /health` - Health check endpoint
- `GET /groups` - List all groups
- `GET /groups/:id` - Get group by ID
- `POST /groups` - Create a new group
- `POST /groups/:id/summoners` - Add a summoner to a group
