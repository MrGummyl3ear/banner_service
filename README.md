# 🏷️ Avito Banner Service

A scalable, robust backend service for managing banners, featuring authentication, admin/user roles, and high performance.

---

## 🌟 Overview

**banner_service** is a backend solution for handling dynamic banners with features like fine-grained authorization, efficient caching, and load testing. The project uses Docker, Makefile automation, and modern Go practices to ensure reliability and scalability.

---

## 🚀 Features

- **JWT-based Authentication**: Secure sign-up and sign-in endpoints.
- **RESTful Banner API**: CRUD operations for banners, with admin and user token support.
- **Makefile Automation**: Easy commands for building, running, linting, and containerization.
- **Load Testing**: k6 scripts included for performance validation.
- **Problem-Solving Notes**: Real-world engineering decisions and challenges explained.

---

## 📦 Content

1. [Commands](#Commands)
2. [Handlers](#Handlers)  
    2.1 [Authorization](#Authorization)  
    2.2 [Banners API](#Banners-API)
3. [Problems Encountered and Their Solutions](#Problems-Encountered-and-Their-Solutions)
4. [Load Testing](#Load-Testing)

---

## 🛠️ Commands

- Creating a docker image
    ```bash
    make image_up
    ```
- Starting a docker container
    ```bash
    make service_up
    ```
- Running the server without a container
    ```bash
    make run
    ```
- Running the linter (golangci-lint)
    ```bash
    make linter
    ```

---

## 🔌 Handlers

### Authorization
1. POST: /auth/sign-up  
    Register a new user
2. POST: /auth/sign-in  
    Authentication. Returns JWT Token upon successful authorization.

### Banners API
1. GET: /banners  
   Retrieve all banners with filtering by feature and tag, and limit/offset parameters. Requires admin token.
2. POST: /banners  
   Create a new banner. Requires admin token.
3. DELETE: /banners/{id}  
   Delete banner by ID. Requires admin token.
4. PATCH: /banners/{id}  
   Update banner by ID. Requires admin token.
5. GET: /user_banner  
   Get banner by feature and tag. User token is sufficient.
6. DELETE: /delete  
   Delete banner by feature and tag. Requires admin token.

---

## 🧑‍💻 Problems Encountered and Their Solutions

- **User Management**:  
  The technical specification did not specify how to add users to the database or what level of access to assign.  
  **Solution**: Implemented user registration (`/auth/sign-up`) and authentication (`/auth/sign-in`) endpoints, with role-based access using JWT tokens.
- **Caching Strategy**:  
  There was a question of what to cache. Retrieving all records from the database is impractical, but there is no critical need for real-time updates.  
  **Solution**: Cached only the most frequently accessed banners by feature and tag, optimizing for speed without unnecessary memory usage.
- **API Flexibility and Data Integrity**:  
  The API gives the impression that any banner parameter can be changed, including the feature ID and tag. In practice, this can create inconsistencies.  
  **Solution**: Restricted updates to only editable fields and validated inputs to preserve referential integrity.

---

## 📈 Load Testing

For load testing, k6 technology was used. You can find test results in the loadtest folder.

---

_Questions or feedback? Open an issue or reach out via [GitHub](https://github.com/MrGummyl3ear/banner_service)._
