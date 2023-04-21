package usecase

import(
	"github.com/VladMak/golang_test/internal/domain"
	//"database/sql"
	_ "github.com/lib/pq"
	//"fmt"
)

type Db struct {
	Db domain.PostgresDB
}

// Инициализация базы данных
func (db *Db) CreateDb(config domain.ConfigDb) {
	db.Db = domain.PostgresDB{}
	db.Db.CreateDb(config.Db.Username, config.Db.Password, config.Db.Host, config.Db.Dbname, config.Db.Port)
}

// Соединение с базой данных
func (db *Db) Connect() {
	db.Db.DB, _ = db.Db.Connect()
}

// Запись в базу данных информации об изменении
func (db *Db) Insert(path string) {
	db.Db.Exec(path)
}
