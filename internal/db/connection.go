package db

import (
	"github.com/ibrahaknazarov/credit_risk/internal/configs"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"os"
)

func InitConnection() (*sqlx.DB, error) {
	connectionConfigs := configs.AppSettings.PostgresParams
	dbConn, err := sqlx.Connect("postgres",
		fmt.Sprintf("port=%s host=%s user=%s password=%s dbname=%s sslmode=disable",
			connectionConfigs.Port,
			connectionConfigs.Host,
			connectionConfigs.User,
			os.Getenv("DB_PASSWORD"),
			connectionConfigs.Database))
	if err != nil {
		return nil, err
	}
	return dbConn, nil
}