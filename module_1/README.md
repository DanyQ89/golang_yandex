# Calculation service

- ### клонируем репозиторий
- ### Устанавливаем Postman (https://www.postman.com/downloads/)
- ### пишем в терминал команду: 
    `go run ./cmd/calc_service/...`
- ### Принимаем все условия сервера
- ### пишем запросы в postman
    - New - http    
    - Пишем строку подключения:<br>
      `http://localhost:8080/api/v1/calculate`
    - Выбираем метод POST
    - Переходим в параметр Body, выбираем raw и пишем любой из 3 запросов:
        - code 200 + result<br>
        `{
            "expression": "2+2*2"
        }`
        - code 422 + error<br>
        `{
            "expression": "2+2*2a"
        }`
        - code 500 + error<br>
        `{
            "expression": 22
        }`

![image](https://github.com/user-attachments/assets/d3c14530-ee70-4ceb-9ae4-ff08fb07d524)
