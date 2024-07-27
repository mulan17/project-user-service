# RESTful API для user-service

## Ресурси

- **`/user`**
- **`/profile`**

## Методи

### **`/user`**

- **Створення (Create)**
  - **Метод:** `POST`
  - **Опис:** Додає нового користувача.
  - **Приклад запиту:**
    ```http
    POST /user
    Content-Type: application/json

    {
      "email": "user@example.com",
      "password": "securepassword"
    }
    ```

- **Читання (Read)**
  - **Метод:** `GET`
  - **Опис:** Отримує інформацію про користувача.
  - **Приклад запиту:**
    ```http
    GET /user
    ```


### **`/profile`**

- **Читання (Read)**
  - **Метод:** `GET`
  - **Опис:** Отримує інформацію про профіль користувача.
  - **Приклад запиту:**
    ```http
    GET /profile
    ```
