# ShortlyLink - URL Shortener

Shortly - Link is a URL shortening project developed using **Golang**, **Gin**, **GORM**, **MySQL**, and **React (TypeScript)**. It allows users to create short URLs and redirect to the original URLs easily.

## Features

- Generate short URLs from original URLs
- Redirect to the original URL via the short URL
- Store data in MySQL database
- Uses Gin framework and GORM ORM for the backend
- Uses React + TypeScript for the frontend
- Supports CORS

## Installation

### 1. Clone the repository

```sh
git clone https://github.com/dekbadnerd/shortlylink.git
cd shortlylink
```

### 2. Setup environment variables

Create a `.env` file and add the following values:

```env
DB_USERNAME=root
DB_PASSWORD=root
DB_HOST=127.0.0.1
DB_PORT=3306
DB_NAME=shortlylink
```

### 3. Install dependencies

#### Backend (Go)

```sh
cd backend
go mod tidy
```

#### Frontend (React)

```sh
cd frontend
npm install
```

### 4. Run the application

#### Backend

```sh
cd backend
go run main.go
```

The backend will run on `http://localhost:8000`

#### Frontend

```sh
cd frontend/shortly-link
npm run dev
```

The frontend will run on `http://localhost:3000`

## API Endpoints

### 1. Shorten URL

**POST** `/shorten`

#### Request Body (JSON)

```json
{
  "url": "https://example.com"
}
```

#### Response (JSON)

```json
{
  "short_url": "abc123"
}
```

### 2. Redirect to Original URL

**GET** `/{shortURL}`

- Example: `GET http://localhost:8000/abc123`
- If the URL exists, it will redirect to the original URL
- If not found, it will return a JSON error message

## Frontend Usage

- Enter the URL you want to shorten in the input field
- Click the "SHORTEN !!!" button to generate a short URL
- Once generated, the shortened URL will be displayed
- Click "Open URL" to open the link or "Copy URL" to copy the link to the clipboard


 
