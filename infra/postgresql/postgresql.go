package postgresql

import (
	"database/sql"
	"fmt"

	"github.com/spf13/viper"

	_ "github.com/lib/pq"
)

func init() {
	viper.BindEnv("postgres.host", "POSTGRES_HOST")
	viper.BindEnv("postgres.port", "POSTGRES_PORT")
	viper.BindEnv("postgres.user", "POSTGRES_USER")
	viper.BindEnv("postgres.password", "POSTGRES_PASSWORD")
	viper.BindEnv("postgres.dbname", "POSTGRES_DBNAME")
	viper.BindEnv("postgres.sslmode", "POSTGRES_SSLMODE")

	viper.SetDefault("postgres.host", "localhost")
	viper.SetDefault("postgres.port", "5432")
	viper.SetDefault("postgres.user", "postgres")
	viper.SetDefault("postgres.password", "password")
	viper.SetDefault("postgres.dbname", "test")
	viper.SetDefault("postgres.sslmode", "disable")
}

func OpenDB() (*sql.DB, error) {
	dbSource := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		viper.GetString("postgres.host"),
		viper.GetInt("postgres.port"),
		viper.GetString("postgres.user"),
		viper.GetString("postgres.password"),
		viper.GetString("postgres.dbname"),
		viper.GetString("postgres.sslmode"),
	)
	return sql.Open("postgres", dbSource)
}
