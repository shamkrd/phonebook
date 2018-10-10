package models

//Configuration is структура для конфигурации приложения
type Configuration struct {
	DbUser     string
	DBPassword string
	DBName     string
	Port       string
}

//GetDbConnectionString is собираем строку подключения для работы с бд
func (c *Configuration) GetDbConnectionString() string {

	connectionString := c.DbUser + ":" + c.DBPassword + "@/" + c.DBName
	return connectionString
}
