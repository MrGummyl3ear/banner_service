# Avito banner

## Content
1. [Команды](#Команды)
2. [Handlers](#Handlers) \
    2.1 [Authorization](#Authorization) \
    2.2.[Banners API](#Banners-API)
3. [Проблемы, с которыми столкнулся и их решения](#Проблемы-с-которыми-столкнулся-и-их-решения)
4. [Load Testing](#Load-Testing)

# Команды 
 - Создание docker image
```bash
    make image_up
```
 - Запуск docker container
```bash
    make service_up
```
 - Запуск сервера без поднятия контейнера
```bash
    make run
```
 - Запуск линтера(golangci-lint)
```bash
    make linter
```

# Handlers

## Authorization
1. POST: /auth/sign-up \
    Регистрация пользователя
2. POST: /auth/sign-in \
    Аутентификация. При успешной авторизации возвращает JWT Token.
   
## Banners API
1. GET: /banners \
   Получение всех баннеров с фильтрацией по фиче и тегу и параметрам limit, offset.Требуется админский токен 
2. POST: /banners \
   Создание нового банера..Требуется админский токен 
3. DELETE: /banners/{id} \
   Удаление баннера по id..Требуется админский токен 
4. PATCH: /banners/{id} \
   Обновление баннера по id..Требуется админский токен 
5. GET: /user_banner \
   Получение баннера по фиче и тегу.Достаточно пользовательского токена 
6. DELETE: /delete \
    Удаление баннера по фиче и тегу. .Требуется админский токен

# Проблемы, с которыми столкнулся и их решения
- В техническом задании не было указано, как добавлять пользователей в базу данных и какой уровень доступа им присваивать. Исходя из этого, каждый пользователь, зарегистрированный через обработчик /auth/sign-up, получает обычный токен. Однако, если необходимо получить административный токен, необходимо войти через /auth/sign-in, используя логин и пароль, указанные в конфигурации env. Также в jwt.Claims записывается массив ролей пользователя. Это позволяет проводить дальнейшее масштабирование. К примеру у пользователя могут быть роли ["Admin","System_admin"],в котором у каждой роли свои права и ограничения  
- Вопрос был также в том, что заносить в кэш. Ведь получать все записи из бд непрактично,но  при этом нет критериев,чтобы определять популярные запросы. Мой подход был следующим: при запросе пользователя,если записи нет в кэше,то запись вносится в кэш.При этом через определенное время кэш очищается
- В api складывается впечатление,что любой параметр баннера можно поменять, в том числе id фичи и тэга. При этом считать ли после таких изменений баннер новой версией или считать его уже другим баннером? Мне кажется, в patch с учетом доп. задания по различным версиям логично было бы запретить изменять фичи или тэги, иначе надо было бы удалять и заново ставить баннер

# Load Testing
  Для нагрузочного тестирования использовал технологии k6. В папке loadtest можно посмотреть результаты тестов(есть как выводы в консоли,так и графические представления)
