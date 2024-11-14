# Receipt Processor

## Introduction

Thank you for the opportunity to complete this assignment. I approached this project as both a technical challenge and a learning experience, taking the time to dive into Go and Docker—two technologies new to me. While it required a substantial effort to learn these tools from scratch, I believe the experience has been invaluable and has broadened my skills in backend development and containerization.

I am a firm believer that a software engineer's journey is about continuous learning. We’re often faced with the unknown, and embracing challenges like this assignment allowed me to test my adaptability and resilience in an unfamiliar environment. I strive to be proactive about growing my skills, knowing that in real-world settings, it’s unlikely to have complete mastery over every tool or technology. Completing this project has equipped me with new competencies and the confidence to quickly adapt, which I look forward to bringing to any team I join.

## Project Overview

The `Receipt Processor` is a lightweight web service built to handle receipt data and calculate points based on a set of predefined rules. This repository contains:

1. **Endpoints** for processing receipts and retrieving calculated points.
2. **Scalable Logic** implemented in Go for reliable and efficient computation.
3. **Dockerized Setup** for ease of deployment in containerized environments.

### Technologies Used

- **Go**: Chosen as a powerful, statically typed language ideal for creating efficient, high-performance backend applications.
- **Docker**: Used to containerize the application, ensuring that it runs consistently in any environment.

## API Endpoints

1. **`POST /receipts/process`**: Accepts a JSON receipt and returns a unique ID. Points for the receipt are computed based on rules provided in the assignment and are stored in memory.
2. **`GET /receipts/{id}/points`**: Returns the points associated with a receipt ID.

### Points Calculation

The points system is determined by various rules, such as counting alphanumeric characters in retailer names, evaluating the total amount, assessing item descriptions, and applying specific multipliers. Details on these rules can be found in the main application logic within the `main.go` file.

## Running the Application

### Prerequisites

- Docker (if not using Go directly)
- Go (if building manually)

### Running with Docker

To build and run the Docker container:

```bash
docker build -t receipt-processor .
docker run -p 8080:8080 receipt-processor
```

The service will be accessible at http://localhost:8080.

### Running Locally with Go

If Docker is unavailable, the application can also be run locally by following these steps:

```bash
go build -o receipt_service
./receipt_service
```

## Closing Note

I appreciate the opportunity to engage with this technical assessment. Simulations like these not only test my current skills but also give me a great chance to demonstrate my eagerness and capability to grow, even in unfamiliar territory. Thank you for considering my work, and I hope this project reflects both my technical aptitude and commitment to continuous learning.

Sincerely,
\
Christian Knox-Phillips
