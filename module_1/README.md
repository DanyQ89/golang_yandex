# Calculation service

### клонируем проект себе в репозиторий
2) для проверки запросов я использовал Postman (https://www.postman.com/downloads/)
3) пишем в терминал команду: go run ./cmd/calc_service/...
4)Принимаем все условия сервера
5) пишем запросы в postman 
•Пишем строку подключения:<br>
http://localhost:8080/api/v1/calculate
•Выбираем метод POST
•Переходим в параметр Body, выбираем raw и пишем по очереди:
1. code 200 + result
{
    "expression": "2+2*2"
}

2. code 422 + error
{
    "expression": "2+2*2a"
}

3. code 500 + error
{
    "expression": 22
}

![image](https://github.com/user-attachments/assets/d3c14530-ee70-4ceb-9ae4-ff08fb07d524)
