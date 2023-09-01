package mysqlrepo

import (
	"context"
	"fmt"
	"strings"

	"github.com/p3ym4n/re"
)

func (mr *MysqlRepo) Unique(c context.Context, entity, field string, value interface{}) (bool, re.Error) {
	const op = re.Op("store.Unique")

	count := 0
	query := fmt.Sprintf("select count(*) from %s where %s = ?", entity, field)
	if err := mr.conn.QueryRowContext(c, query, value).Scan(&count); err != nil {
		return false, re.New(op, fmt.Errorf("query error : %w", err), re.KindUnexpected)
	}

	return count == 0, nil
}

func (mr *MysqlRepo) UniqueExceptCertainID(c context.Context, entity, field string, value interface{}, ignoredID uint) (bool, re.Error) {
	const op = re.Op("store.UniqueExceptCertainID")

	count := 0
	query := fmt.Sprintf("select count(*) from %s where %s = ? and id != %v", entity, field, ignoredID)
	if err := mr.conn.QueryRowContext(c, query, value).Scan(&count); err != nil {
		return false, re.New(op, fmt.Errorf("query error : %w", err), re.KindUnexpected)
	}

	return count == 0, nil
}

func (mr *MysqlRepo) Exists(c context.Context, entity, field string, value interface{}) (bool, re.Error) {
	const op = re.Op("store.Exists")

	shouldBe := 1
	values := make([]interface{}, 0)
	switch typed := value.(type) {
	case []string:
		shouldBe = len(typed)
		for _, v := range typed {
			values = append(values, v)
		}
	case []int:
		shouldBe = len(typed)
		for _, v := range typed {
			values = append(values, v)
		}
	case []int64:
		shouldBe = len(typed)
		for _, v := range typed {
			values = append(values, v)
		}
	case []uint:
		shouldBe = len(typed)
		for _, v := range typed {
			values = append(values, v)
		}
	case []uint64:
		shouldBe = len(typed)
		for _, v := range typed {
			values = append(values, v)
		}
	default:
		values = append(values, value)
	}

	query := fmt.Sprintf("select count(*) from %s where", entity)
	if shouldBe == 0 {
		return false, nil
	} else if shouldBe == 1 {
		query = fmt.Sprintf("%s %s = ?", query, field)
	} else {
		query = fmt.Sprintf("%s %s in (?%s)", query, field, strings.Repeat(",?", shouldBe-1))
	}

	count := 0
	err := mr.conn.QueryRowContext(c, query, values...).Scan(&count)
	if err != nil {
		return false, re.New(op, fmt.Errorf("query error : %w", err), re.KindUnexpected)
	}

	return count == shouldBe, nil
}
