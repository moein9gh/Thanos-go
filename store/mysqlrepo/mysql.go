package mysqlrepo

import (
	"database/sql"
	"fmt"
	"math"
	"time"

	"github.com/thanos-go/config"
	"github.com/thanos-go/log"

	_ "github.com/go-sql-driver/mysql"
)

const Dialect = "mysql"

type MysqlRepo struct {
	statements map[string]*sql.Stmt
	conn       *sql.DB
	config     *config.MysqlConfig
}

func New(cfg *config.MysqlConfig) *MysqlRepo {

	dsn := fmt.Sprintf(
		"%s:%s@(%s:%v)/%s?collation=utf8mb4_unicode_ci&parseTime=True",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Schema,
	)
	db, err := sql.Open(Dialect, dsn)
	if err != nil {
		log.Named("mysql").Fatal("db connection failed", err)
	}

	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetMaxIdleConns(cfg.MaxIdleConns)
	db.SetConnMaxLifetime(cfg.ConnMaxLifetime)

	if pingErr := db.Ping(); pingErr != nil {
		log.Named("mysql").Error("ping", pingErr)
	} else {
		log.Named("mysql").Debug("connection is ok")
	}

	if cfg.MaxOpenConns == 0 {
		var skip string
		var maxConnections int
		maxConErr := db.QueryRow("SHOW VARIABLES LIKE 'max_connections'").Scan(&skip, &maxConnections)
		if maxConErr != nil {
			log.Named("mysql").Error("getting the max_connections", maxConErr)
		}
		maxConnections = int(math.Floor(float64(maxConnections) * 0.9))
		if maxConnections == 0 {
			maxConnections = 100
		}
		cfg.MaxOpenConns = maxConnections
	}
	db.SetMaxOpenConns(cfg.MaxOpenConns)

	if cfg.MaxIdleConns == 0 {
		cfg.MaxIdleConns = cfg.MaxOpenConns
	}
	db.SetMaxIdleConns(cfg.MaxIdleConns)

	if cfg.ConnMaxLifetime.Seconds() == 0 {
		var skip string
		var waitTimeout int
		waitErr := db.QueryRow("SHOW VARIABLES LIKE 'wait_timeout'").Scan(&skip, &waitTimeout)
		if waitErr != nil {
			log.Named("mysql").Error("getting the wait_timeout", waitErr)
		}
		if waitTimeout == 0 {
			waitTimeout = 180
		}
		waitTimeout = int(math.Min(float64(waitTimeout), 180))
		cfg.ConnMaxLifetime = time.Duration(waitTimeout) * time.Second
	}
	db.SetConnMaxLifetime(cfg.ConnMaxLifetime)

	return &MysqlRepo{config: cfg, conn: db, statements: make(map[string]*sql.Stmt)}
}

func (mr *MysqlRepo) stmt(id string) *sql.Stmt {
	return mr.statements[id]
}

func (mr *MysqlRepo) setStmt(id string, stmt *sql.Stmt) {
	mr.statements[id] = stmt
}

func (mr *MysqlRepo) Ping() error {
	if err := mr.conn.Ping(); err != nil {
		return err
	}

	return nil
}

func (mr *MysqlRepo) DB() *sql.DB {
	return mr.conn
}

func (mr *MysqlRepo) Config() *config.MysqlConfig {
	return mr.config
}

func (mr *MysqlRepo) Close() {
	for _, stmt := range mr.statements {
		_ = stmt.Close()
	}
	_ = mr.conn.Close()
}
