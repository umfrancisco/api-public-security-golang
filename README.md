# Criminal Statistics API

This project is a RESTful API built with **Encore (Go)** that reads CSV files and exposes structured data about **criminal statistics** (e.g., homicides, thefts, robberies) for different cities and years.

The API processes raw CSV datasets and returns clean JSON responses, making it easy to integrate with dashboards, frontend apps, or data analysis tools.

---

## Features

* Reads CSV files from local or remote sources (e.g., GitHub)
* Provides structured crime statistics by city and year
* Built with Encore for rapid backend development
* Easy to integrate with React dashboards or other clients

---

## Prerequisites

### Install Encore

* **macOS:**

  ```bash
  brew install encoredev/tap/encore
  ```

* **Linux:**

  ```bash
  curl -L https://encore.dev/install.sh | bash
  ```

* **Windows:**

  ```powershell
  iwr https://encore.dev/install.ps1 | iex
  ```

---

## Run the Project

Clone the repository and run:

```bash
encore run
```

The API will be available at:

```
http://localhost:4000
```

---

## API Usage

### Get Crime Data

Example request:

```bash
curl http://localhost:4000/api/sao_paulo
```

### Example Response

```json
{
  "Content": [
    {
      "id": 1,
      "city": "São Paulo",
      "year": 2025,
      "homicides": 4.4,
      "thefts": 2194.73,
      "robberies": 862.9
    }
  ]
}
```

---

## Data Source

The API reads CSV files containing criminal statistics. Example structure:

```
year,homicides,thefts,robberies
2025,4.40,2194.73,862.9
```

These files can be:

* Stored locally
* Hosted remotely (e.g., GitHub raw URLs)

---

## How It Works

1. The API receives a request with a city parameter
2. It selects the corresponding CSV file
3. Parses the file into Go structs
4. Returns JSON data to the client

---

## Testing

Run tests with:

```bash
encore test ./...
```

---

## Local Development Dashboard

While running the app, open:

```
http://localhost:9400/
```

You can:

* View API requests and traces
* Explore your service architecture
* Test endpoints interactively

---

## Deployment

### Deploy with Encore Cloud

```bash
git add -A .
git commit -m "Deploy API"
git push encore
```

### Or Build a Docker Image

```bash
encore build docker
```

---

## Use Cases

* Crime data dashboards (React, charts)
* Data analysis pipelines
* Public security research
* Open data APIs

---

## Learn More

* https://encore.dev/docs/go
* https://encore.dev/docs/go/primitives/services
* https://encore.dev/docs/go/primitives/defining-apis



