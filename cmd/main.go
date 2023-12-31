package main

import (
    "github.com/gin-gonic/gin"
    "github.com/alina-boichenko/go-gin-api-medium/pkg/books"
    "github.com/alina-boichenko/go-gin-api-medium/pkg/common/db"
    "github.com/spf13/viper"
)

func main() {
    viper.SetConfigFile("./pkg/common/envs/.env")
    viper.ReadInConfig()

    port := viper.Get("PORT").(string)
    dbUrl := viper.Get("DB_URL").(string)

    r := gin.Default()
    h := db.Init(dbUrl)

    books.RegisterRoutes(r, h)
    // здесь регистрируются дополнительные маршруты

    r.Run(port)
}
