# URL Shortener with Redis and GoFiber

This project provides a RESTful API for shortening URLs using Redis for storage and GoFiber as the web framework.
Docker and Docker-compose is used for containerization and deployment, ensuring consistent behavior across different environments.

## Features

- Shorten long URLs into shorter, more manageable links.
- Resolve shortened URLs to their original long form.
- Rate limiting to prevent abuse and ensure fair usage.
- Docker used for containerization and deployment.

## Prerequisites

Make sure you have Go and Redis installed on your machine if you want to run it without docker. You can download and install go from [here](https://golang.org/dl/) and redis from [here](https://redis.io/downloads/).

## Dependencies

- [GoFiber](https://github.com/gofiber/fiber): A fast and lightweight web framework for Go.
- [Redis](https://redis.io/): An open-source, in-memory data structure store used as a database, cache, and message broker.
- [go-redis/redis](https://github.com/go-redis/redis): Go client for Redis.
- [joho/godotenv](https://github.com/joho/godotenv): A Go library for loading environment variables from a .env file.
- [asaskevich/govalidator](https://github.com/asaskevich/govalidator): A package of validators and sanitizers for strings, structs, and collections.
- [google/uuid](https://github.com/google/uuid): A Go package for UUIDs (Universally Unique Identifiers) generation and parsing.

## Installation

1. Clone the repository:

   ```bash
   https://github.com/Yash-sudo-web/urlShortnerRedisGolang.git

2. Navigate to the project directory:

   ```bash
   cd urlshortnerredisgolang

3. Add ".env" file to configure enviromental variables, Example .env file:
   
   ```bash
   REDIS_URL="db:6379"
   REDIS_PASSWORD=""
   APP_PORT=":3000"
   DOMAIN="http://localhost:3000"
   API_QUOTA=10

5. Build and run the application using Docker Compose:

   ```bash
   docker-compose up --build

## API Endpoints
- POST /shorten: Shorten a long URL.
- GET /{url}: Resolve a shortened URL.

## Usage
You can use tools like cURL or Postman to interact with the API endpoints. Here are some examples:

- Send a POST request to /shorten with a JSON payload containing the URL to be shortened:

  ```bash
  curl -X POST -H "Content-Type: application/json" -d '{"url": "https://example.com"}' http://localhost:3000/shorten
  
- To resolve a shortened URL, simply append the shortened code to the base URL:
  
  ```bash
  curl http://localhost:3000/{shortened_code}

## Contributing
Contributions are welcome! Feel free to open issues or submit pull requests.

## License

This project is licensed under the [MIT License](https://github.com/Yash-sudo-web/urlShortnerRedisGolang/blob/main/LICENSE).
