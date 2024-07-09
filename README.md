# proxy-server

The **Proxy Server** project is an HTTP server designed to proxy requests to third-party services.


## Installation

1. **Clone the repository:**
   ```bash
   git clone https://github.com/xurzfuufzu/proxy-server.git
   cd proxy-server
   ```
2. **Build the Docker images:**
   ```bash
   make build
   ```
3. **Start the Docker containers:**
   ```bash
   make up
   ```

## Request Format

The server expects a JSON request with the following fields:

```
{
"method": "GET",
"url": "http://example.com",
"headers": { "Authorization": "Bearer your_access_token" }
}
```

## Response Format

The response to the client should be in JSON format with the following fields:

```
{
"id": "requestId",
"status": <HTTP status code from the third-party service response>,
"headers": { "array of headers from the third-party service response" },
"length": <length of the response content>
}
```

## Endpoints

### `POST /proxy`

**Description:** Sends a proxy request and stores information about the request and response.

**Example Request:**
```json
{
    "method": "GET",
    "url": "http://google.com",
    "headers": { 
        "Authentication": "Basic bG9naW46cGFzc3dvcmQ="
    }   
}
```
**Example Response:**
```json
{
    "id": "1",
    "status": 200,
    "headers": {
        "Cache-Control": "private, max-age=0",
        "Content-Security-Policy-Report-Only": "object-src 'none';base-uri 'self';script-src 'nonce-XO0aV7UEKgS_meH_D2AKWg' 'strict-dynamic' 'report-sample' 'unsafe-eval' 'unsafe-inline' https: http:;report-uri https://csp.withgoogle.com/csp/gws/other-hp",
        "Content-Type": "text/html; charset=ISO-8859-1",
        "Date": "Tue, 09 Jul 2024 13:27:39 GMT",
        "Expires": "-1",
        "P3p": "CP=\"This is not a P3P policy! See g.co/p3phelp for more info.\"",
        "Server": "gws",
        "Set-Cookie": "AEC=AVYB7crFDBfTKBANtWViki9kwSlCf7QyoE7-Lwdqxi31w73vAmV_JrITNSQ; expires=Sun, 05-Jan-2025 13:27:39 GMT; path=/; domain=.google.com; Secure; HttpOnly; SameSite=lax",
        "X-Frame-Options": "SAMEORIGIN",
        "X-Xss-Protection": "0"
    },
    "length": 21201
}
```

### `GET /proxy/:id`
**Description:** Returns request by id.

**Example Response:**
```json
{
    "request": {
        "method": "GET",
        "url": "http://google.com",
        "headers": {
            "Authentication": "Basic bG9naW46cGFzc3dvcmQ="
        }
    },
    "response": {
        "id": "1",
        "status": 200,
        "headers": {
            "Cache-Control": "private, max-age=0",
            "Content-Security-Policy-Report-Only": "object-src 'none';base-uri 'self';script-src 'nonce-XO0aV7UEKgS_meH_D2AKWg' 'strict-dynamic' 'report-sample' 'unsafe-eval' 'unsafe-inline' https: http:;report-uri https://csp.withgoogle.com/csp/gws/other-hp",
            "Content-Type": "text/html; charset=ISO-8859-1",
            "Date": "Tue, 09 Jul 2024 13:27:39 GMT",
            "Expires": "-1",
            "P3p": "CP=\"This is not a P3P policy! See g.co/p3phelp for more info.\"",
            "Server": "gws",
            "Set-Cookie": "AEC=AVYB7crFDBfTKBANtWViki9kwSlCf7QyoE7-Lwdqxi31w73vAmV_JrITNSQ; expires=Sun, 05-Jan-2025 13:27:39 GMT; path=/; domain=.google.com; Secure; HttpOnly; SameSite=lax",
            "X-Frame-Options": "SAMEORIGIN",
            "X-Xss-Protection": "0"
        },
        "length": 21201
    }
}
```