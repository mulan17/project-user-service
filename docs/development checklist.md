# Чекліст по розробці сервісу

1. **Описуємо бізнес моделі та їх поведінку**
   - Якими сутностями оперує наш продукт?
   - Що можна робити з цими сутностями?

2. **Описуємо RESTful API на основі бізнес моделей та поведінки**
   - Які у нас є ресурси?
     - `/user`
     - `/user{id}`
     - `/admin/block/{id}`
     - `/admin/limit/{id}`
   - Які операції ці ресурси підтримують?
     - `POST /users` - Adds a new user.
     - `GET /users` - Retrieves information about all users.- - `PATCH /users/{id}` - Updates information about a specific user.
     - `GET /users/{id}` - Retrieves information about a specific user by ID. Authentication is required.
     - `POST /admin/block/{id}` - Blocks a specific user by ID. Authentication is required.
     - `POST /admin/limit/{id}` - Limits a specific user by ID. Authentication is required.
     - `POST /login` - Authenticates a user and provides a session or token.

3. **Описуємо архітектуру сервісу з точки зору Go**
   - Які у нас є шари та пакети?
     - Наприклад, `rest`, `database`, `core`
   - Які ми використовуємо сторонні бібліотеки / фреймворки?

4. **Створюємо репозиторій, домовляємося про git flow та ініціалізуємо Go проєкт** - DONE

5. **Розділити початкові задачі (дуже прості) на розробку між розробниками**
   - Може не бути бази даних
   - Може не бути ідеального тестування
   - Може не бути коректної авторизації
   - Може не бути взаємодії між сервісами (mocks)

6. **Ітеративно додаємо все більше нових функцій**
   - Підтримка бази даних
   - Більше тестів
   - Інші функції
