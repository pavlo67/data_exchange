package records_sqlite

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/pavlo67/common/common/strlib"
	"github.com/pavlo67/data_exchange/components/ns"

	"github.com/pavlo67/common/common"
	"github.com/pavlo67/common/common/auth"
	"github.com/pavlo67/common/common/db"
	"github.com/pavlo67/common/common/errors"
	"github.com/pavlo67/common/common/selectors"
	"github.com/pavlo67/common/common/sqllib"
	"github.com/pavlo67/data_exchange/components/tags"
	"github.com/pavlo67/data_exchange/types/records"
)

var fieldsToInsert = []string{"title", "summary", "type_key", "data", "embedded", "tags", "urn", "owner_id", "viewer_id", "history"}
var fieldsToInsertStr = strings.Join(fieldsToInsert, ", ")

var fieldsToUpdate = append(fieldsToInsert, "updated_at")
var fieldsToUpdateStr = strings.Join(fieldsToUpdate, " = ?, ") + " = ?"

var fieldsToRead = append(fieldsToUpdate, "created_at")
var fieldsToReadStr = strings.Join(fieldsToRead, ", ")

var fieldsToList = append(fieldsToRead, "id")
var fieldsToListStr = strings.Join(fieldsToList, ", ")

var _ records.Operator = &recordsSQLite{}

type recordsSQLite struct {
	db    *sql.DB
	table string

	sqlInsert, sqlUpdate, sqlRead, sqlRemove, sqlClean string
	stmInsert, stmUpdate, stmRead, stmRemove, stmClean *sql.Stmt
}

const onNew = "on recordsSQLite.New(): "

func New(db *sql.DB, table string) (records.Operator, db.Cleaner, error) {
	if table == "" {
		table = records.CollectionDefault
	}

	recordsOp := recordsSQLite{
		db:    db,
		table: table,

		sqlInsert: "INSERT INTO " + table + " (" + fieldsToInsertStr + ") VALUES (" + strings.Repeat(",? ", len(fieldsToInsert))[1:] + ")",
		sqlUpdate: "UPDATE " + table + " SET " + fieldsToUpdateStr + " WHERE id = ?",
		sqlRemove: "DELETE FROM " + table + " where id = ?",
		sqlRead:   "SELECT " + fieldsToReadStr + " FROM " + table + " WHERE id = ?",

		sqlClean: "DELETE FROM " + table,
	}

	sqlStmts := []sqllib.SqlStmt{
		{&recordsOp.stmInsert, recordsOp.sqlInsert},
		{&recordsOp.stmUpdate, recordsOp.sqlUpdate},
		{&recordsOp.stmRead, recordsOp.sqlRead},
		{&recordsOp.stmRemove, recordsOp.sqlRemove},
		{&recordsOp.stmClean, recordsOp.sqlClean},
	}

	for _, sqlStmt := range sqlStmts {
		if err := sqllib.Prepare(db, sqlStmt.Sql, sqlStmt.Stmt); err != nil {
			return nil, nil, errors.Wrap(err, onNew)
		}
	}

	return &recordsOp, &recordsOp, nil
}

const onSave = "on recordsSQLite.Save(): "

//Save(Item, *auth.Identity) (ID, error)
//Read(ID, *auth.Identity) (*Item, error)
//Remove(ID, *auth.Identity) error
//
//
//
//Tags(*auth.Identity) (tags.StatMap, error)
//AddParent(ts []tags.Item, id ID) ([]tags.Item, error)
//

func (recordsOp *recordsSQLite) Save(item records.Item, identity *auth.Identity) (records.ID, error) {

	if identity == nil {

		identity = &auth.Identity{}
		// TODO!!!
		// return "", errors.CommonError(common.NoRightsKey)
	}

	// TODO!!! rbac check

	if item.ID == "" {
		// TODO!!!
		item.OwnerNSS = ns.NSS(identity.ID)
	}

	var err error

	embeddedBytes := []byte{} // to satisfy NOT NULL constraint
	if len(item.Content.Embedded) > 0 {
		if embeddedBytes, err = json.Marshal(item.Content.Embedded); err != nil {
			return "", errors.Wrapf(err, onSave+"can't marshal .Embedded(%#v)", item.Content.Embedded)
		}
	}

	tagsBytes := []byte{} // to to satisfy NOT NULL constraint
	if len(item.Tags) > 0 {
		if tagsBytes, err = json.Marshal(item.Tags); err != nil {
			return "", errors.Wrapf(err, onSave+"can't marshal .Tags(%#v)", item.Tags)
		}
	}

	// TODO!!! append to .History

	historyBytes := []byte{} // to satisfy NOT NULL constraint
	if len(item.History) > 0 {
		historyBytes, err = json.Marshal(item.History)
		if err != nil {
			return "", errors.Wrapf(err, onSave+"can't marshal .History(%#v)", item.History)
		}
	}

	// "title", "summary", "type_key", "data", "embedded", "tags",
	// "urn", "owner_id", "viewer_id", "history"
	values := []interface{}{
		item.Content.Title, item.Content.Summary, item.Content.TypeKey, item.Content.Data, embeddedBytes, tagsBytes,
		item.URN, item.OwnerNSS, item.ViewerNSS, historyBytes}

	if item.ID == "" {
		res, err := recordsOp.stmInsert.Exec(values...)
		if err != nil {
			return "", errors.Wrapf(err, onSave+sqllib.CantExec, recordsOp.sqlInsert, strlib.Stringify(values))
		}

		idSQLite, err := res.LastInsertId()
		if err != nil {
			return "", errors.Wrapf(err, onSave+sqllib.CantGetLastInsertId, recordsOp.sqlInsert, strlib.Stringify(values))
		}
		return records.ID(strconv.FormatInt(idSQLite, 10)), nil

	}

	values = append(values, time.Now().Format(time.RFC3339), item.ID)
	if _, err := recordsOp.stmUpdate.Exec(values...); err != nil {
		return "", errors.Wrapf(err, onSave+sqllib.CantExec, recordsOp.sqlUpdate, strlib.Stringify(values))
	}

	return item.ID, nil
}

const onRead = "on recordsSQLite.Read(): "

func (recordsOp *recordsSQLite) Read(id records.ID, options *auth.Identity) (*records.Item, error) {
	idNum, err := strconv.ParseUint(string(id), 10, 64)
	if err != nil {
		return nil, fmt.Errorf(onRead+"wrong id (%s)", id)
	}

	item := records.Item{ID: id}

	var embeddedBytes, tagsBytes, historyBytes []byte

	// "title", "summary", "type_key", "data", "embedded", "tags",
	// "urn", "owner_id", "viewer_id", "history", "updated_at", "created_at"

	if err = recordsOp.stmRead.QueryRow(idNum).Scan(
		&item.Content.Title, &item.Content.Summary, &item.Content.TypeKey, &item.Content.Data, &embeddedBytes, &tagsBytes,
		&item.URN, &item.OwnerNSS, &item.ViewerNSS, &historyBytes, &item.UpdatedAt, &item.CreatedAt); err == sql.ErrNoRows {
		return nil, common.ErrNotFound
	} else if err != nil {
		return nil, errors.Wrapf(err, onRead+sqllib.CantScanQueryRow, recordsOp.sqlRead, idNum)
	}

	if len(embeddedBytes) > 0 {
		if err = json.Unmarshal(embeddedBytes, &item.Content.Embedded); err != nil {
			return &item, errors.Wrapf(err, onRead+"can't unmarshal .Embedded (%s)", embeddedBytes)
		}
	}

	if len(tagsBytes) > 0 {
		if err = json.Unmarshal(tagsBytes, &item.Tags); err != nil {
			return &item, errors.Wrapf(err, onRead+"can't unmarshal .Tags (%s)", tagsBytes)
		}
	}

	if len(historyBytes) > 0 {
		if err = json.Unmarshal(historyBytes, &item.History); err != nil {
			return &item, errors.Wrapf(err, onRead+"can't unmarshal .History (%s)", historyBytes)
		}
	}

	return &item, nil
}

const onRemove = "on recordsSQLite.Remove()"

func (recordsOp *recordsSQLite) Remove(id records.ID, options *auth.Identity) error {

	// TODO!!! rbac check

	idNum, err := strconv.ParseUint(string(id), 10, 64)
	if err != nil {
		return fmt.Errorf(onRemove+"wrong id (%s)", id)
	}

	if _, err = recordsOp.stmRemove.Exec(idNum); err != nil {
		return errors.Wrapf(err, onRemove+sqllib.CantExec, recordsOp.sqlRemove, idNum)
	}

	return nil
}

const onList = "on recordsSQLite.List()"

func (recordsOp *recordsSQLite) List(selector *selectors.Term, identity *auth.Identity) ([]records.Item, error) {
	condition, values, err := Conditions(selector, identity)
	if err != nil {
		return nil, fmt.Errorf(onList+": wrong selector: %#v", selector)
	}

	query := sqllib.SQLList(recordsOp.table, fieldsToListStr, condition, identity)
	stm, err := recordsOp.db.Prepare(query)
	if err != nil {
		return nil, errors.Wrapf(err, onList+": can't db.Prepare(%s)", query)
	}

	//l.Infof("%s / %#v\n%s", condition, values, query)

	rows, err := stm.Query(values...)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, errors.Wrapf(err, onList+": "+sqllib.CantQuery, query, values)
	}
	defer rows.Close()

	var items []records.Item

	for rows.Next() {
		var idNum int64
		var item records.Item
		var embeddedBytes, tagsBytes, historyBytes []byte

		// "title", "summary", "type_key", "data", "embedded", "tags",
		// "urn", "owner_id", "viewer_id", "history", "updated_at", "created_at",
		// "id"

		if err := rows.Scan(
			&item.Content.Title, &item.Content.Summary, &item.Content.TypeKey, &item.Content.Data, &embeddedBytes, &tagsBytes,
			&item.URN, &item.OwnerNSS, &item.ViewerNSS, &historyBytes, &item.UpdatedAt, &item.CreatedAt,
			&idNum); err != nil {
			return items, errors.Wrapf(err, onList+": "+sqllib.CantScanQueryRow, query, values)
		}

		if len(embeddedBytes) > 0 {
			if err = json.Unmarshal(embeddedBytes, &item.Content.Embedded); err != nil {
				return items, errors.Wrapf(err, onList+": can't unmarshal .Embedded (%s)", embeddedBytes)
			}
		}

		if len(tagsBytes) > 0 {
			if err = json.Unmarshal(tagsBytes, &item.Tags); err != nil {
				return items, errors.Wrapf(err, onList+": can't unmarshal .Tags (%s)", tagsBytes)
			}
		}

		if len(historyBytes) > 0 {
			if err = json.Unmarshal(historyBytes, &item.History); err != nil {
				return items, errors.Wrapf(err, onList+": can't unmarshal .History (%s)", historyBytes)
			}
		}

		item.ID = records.ID(strconv.FormatInt(idNum, 10))
		items = append(items, item)
	}

	if err = rows.Err(); err != nil {
		return items, errors.Wrapf(err, onList+": "+sqllib.RowsError, query, values)
	}

	return items, nil
}

const onTags = "on recordsSQLite.Tags()"

func (recordsOp *recordsSQLite) Tags(selector *selectors.Term, identity *auth.Identity) (tags.StatMap, error) {
	condition, values, err := Conditions(selector, identity)
	if err != nil {
		return nil, fmt.Errorf(onList+": wrong selector: %#v", selector)
	}

	query := sqllib.SQLList(recordsOp.table, "tags", condition, identity)
	stm, err := recordsOp.db.Prepare(query)
	if err != nil {
		return nil, errors.Wrapf(err, onTags+": can't db.Prepare(%s)", query)
	}

	rows, err := stm.Query(values...)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, errors.Wrapf(err, onTags+": "+sqllib.CantQuery, query, values)
	}
	defer rows.Close()

	tagsStat := tags.StatMap{}

	for rows.Next() {
		var tagsBytes []byte

		if err := rows.Scan(&tagsBytes); err != nil {
			return tagsStat, errors.Wrapf(err, onTags+": "+sqllib.CantScanQueryRow, query, values)
		}

		if len(tagsBytes) > 0 {
			var ts []tags.Item
			if err = json.Unmarshal(tagsBytes, &ts); err != nil {
				// TODO!!! collect errors
				l.Errorf(onTags+": can't unmarshal ts (%s): %s", tagsBytes, err)
				continue
			}

			for _, tag := range ts {
				// l.Info(tag, tagsStat[tag])

				// ts := tagsStat[tag]
				tagsStat[tag] = tagsStat[tag] + 1
			}
		}

	}

	if err = rows.Err(); err != nil {
		return tagsStat, errors.Wrapf(err, onTags+": "+sqllib.RowsError, query, values)
	}

	return tagsStat, nil
}

func (recordsOp *recordsSQLite) Close() error {
	return errors.Wrap(recordsOp.db.Close(), "on recordsSQLite.Close()")
}

const onStat = "on recordsSQLite.StatMap(): "

func (recordsOp *recordsSQLite) Stat(*selectors.Term, *auth.Identity) (db.StatMap, error) {
	return nil, common.ErrNotImplemented
	//	condition, values, err := selectors_sql.Use(term)
	//	if err != nil {
	//		termStr, _ := json.Marshal(term)
	//		return 0, errors.Wrapf(err, onCount+": can't selectors_sql.Use(%s)", termStr)
	//	}
	//
	//	query := sqllib.SQLCount(recordsOp.table, condition, options)
	//	stm, err := recordsOp.db.Prepare(query)
	//	if err != nil {
	//		return 0, errors.Wrapf(err, onCount+": can't db.Prepare(%s)", query)
	//	}
	//
	//	var num uint64
	//
	//	err = stm.QueryRow(values...).Scan(&num)
	//	if err != nil {
	//		return 0, errors.Wrapf(err, onCount+sqllib.CantScanQueryRow, query, values)
	//	}
	//
	//	return nil, common.ErrNotImplemented
}
