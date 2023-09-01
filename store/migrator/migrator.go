package migrator

import (
	"embed"
	"errors"
	"fmt"
	"time"

	"github.com/thanos-go/log"

	"github.com/thanos-go/store/mysqlrepo"

	migrate "github.com/rubenv/sql-migrate"
)

const (
	MigrationsTable = "migrations"
)

//go:embed migrations/*.sql
var fsMigrations embed.FS

func New(repo *mysqlrepo.MysqlRepo) *Migrator {
	return &Migrator{
		repo: repo,
	}
}

type Migrator struct {
	repo *mysqlrepo.MysqlRepo
}

func (m *Migrator) Run(action string, count int) error {

	if action == "new" {
		fmt.Println(time.Now().Unix())
		return nil
	}

	if action == "fresh" {
		if err := m.dropAllTables(); err != nil {
			return fmt.Errorf("dropping all tables: %w", err)
		} else {
			log.Info("all tables dropped")
		}

		// after dropping tables we will run all the migrations
		action = "up"
		// zero count for the migrator agent means to apply/rollback all residues
		count = 0
	}

	plannedItems, err := m.getFilesSource().FindMigrations()
	if err != nil {
		return fmt.Errorf("getting list of available migrations: %w", err)
	}
	appliedItems, err := m.getAlreadyApplied()
	if err != nil {
		return fmt.Errorf("getting list of applied migrations: %w", err)
	}

	if action == "up" {
		if err := m.applyMigration(count, plannedItems, appliedItems); err != nil {
			return err
		}
	}

	if action == "down" {
		if err := m.undoMigration(count, appliedItems); err != nil {
			return err
		}
	}

	appliedItems, err = m.getAlreadyApplied()
	if err != nil {
		return fmt.Errorf("getting list of applied migrations: %w", err)
	}

	m.reportStatus(plannedItems, appliedItems)

	return nil
}

func (m *Migrator) agent() migrate.MigrationSet {
	return migrate.MigrationSet{
		TableName:     MigrationsTable,
		IgnoreUnknown: false,
	}
}

func (m *Migrator) getFilesSource() *migrate.EmbedFileSystemMigrationSource {
	return &migrate.EmbedFileSystemMigrationSource{
		FileSystem: fsMigrations,
		Root:       "migrations",
	}
}

func (m *Migrator) dropAllTables() error {

	tx, err := m.repo.DB().Begin()
	if err != nil {
		return fmt.Errorf("creating transaction: %w", err)
	}

	rows, err := tx.Query(
		"SELECT concat('DROP TABLE IF EXISTS `', table_name, '`;') AS `stmt` FROM information_schema.tables WHERE table_schema = ?;",
		m.repo.Config().Schema,
	)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	stmts := make([]string, 0)
	for rows.Next() {
		var stmt string
		if err = rows.Scan(&stmt); err != nil {
			_ = tx.Rollback()
			return err
		}
		stmts = append(stmts, stmt)
	}
	if err := rows.Close(); err != nil {
		return err
	}

	if _, err := tx.Exec("SET FOREIGN_KEY_CHECKS=0;"); err != nil {
		_ = tx.Rollback()
		return err
	}

	for _, stmt := range stmts {
		if _, err := tx.Exec(stmt); err != nil {
			_ = tx.Rollback()
			return err
		}
	}

	if _, err := tx.Exec("SET FOREIGN_KEY_CHECKS=1;"); err != nil {
		_ = tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

func (m *Migrator) getAlreadyApplied() ([]*migrate.MigrationRecord, error) {
	return m.agent().GetMigrationRecords(m.repo.DB(), mysqlrepo.Dialect)
}

func (m *Migrator) applyMigration(count int, plannedItems []*migrate.Migration, appliedItems []*migrate.MigrationRecord) error {
	if len(plannedItems) == len(appliedItems) {
		log.Info("no new migration available to apply")
		return nil
	}
	affected, err := m.agent().ExecMax(m.repo.DB(), mysqlrepo.Dialect, m.getFilesSource(), migrate.Up, count)
	if err != nil {
		return err
	}
	log.Info("%v migrations applied", affected)
	return nil
}

func (m *Migrator) undoMigration(count int, appliedItems []*migrate.MigrationRecord) error {
	if len(appliedItems) == 0 {
		return errors.New("no applied migration available to roll back")
	}
	affected, err := m.agent().ExecMax(m.repo.DB(), mysqlrepo.Dialect, m.getFilesSource(), migrate.Down, count)
	if err != nil {
		return err
	}
	log.Info("%v migrations rolled back", affected)
	return nil
}

func (m *Migrator) reportStatus(plannedItems []*migrate.Migration, appliedItems []*migrate.MigrationRecord) {
	for _, item := range plannedItems {
		var appliedAt time.Time
		for _, record := range appliedItems {
			if record.Id == item.Id {
				appliedAt = record.AppliedAt
				break
			}
		}
		if appliedAt.IsZero() {
			log.Info("%s - Not Applied", item.Id)
		} else {
			log.Info("%s - Applied at %s", item.Id, appliedAt)
		}
	}
}
