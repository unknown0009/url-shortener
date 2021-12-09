# url-shortener

A simple URL shortener application.

---

## **Features**
- The service supports 2 data stores (in-memory and PostgreSQL). The -db parameter is used to define the database (local/postgresql).
- The service accepts two requests
    - GET  /{short_url} takes a short url and returns the original.
    - POST /?url=http://vk.com saves the original url and returns the shortened.


## **Setup**

**In-memory storage:**

1. Download an image [5.32 Mb].

    ```
    docker pull fidesy/urlsapp 
    ```
2. Run an image.
    ```
    docker run -dp 80:80 --name urlsapp fidesy/urlsapp -db local
    ```
**PostreSQL**:
1. Clone repository.
    ```
     git clone github.com/fidesy/url-shortener
    ```
1. Run a docker-compose file.
    ```
    docker-compose -f docker-compose.yml up -d
    ```

## **Usage**
```bash
# POST 
curl -d "url=http://vk.com/friends" 
-X POST http://127.0.0.1

  > http://127.0.0.1/oZMiTxhQ_H

# GET
curl http://127.0.0.1/oZMiTxhQ_H
 
  > http://vk.com/friends
```