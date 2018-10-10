# phonebook

Для корректной работы приложения необходимо сделать следующее:
1. Установить mySqlServer
2. Востановить базу данных из дампа phonebook "20181009 2241.sql"
3. Установить следующие зависимости:
 github.com/tkanos/gonfig,
 github.com/go-sql-driver/mysql
4. В файле config.development.json изменить параметры доступа к базе

Для запуска приложения необходимо выполнить команду: go run main.go
