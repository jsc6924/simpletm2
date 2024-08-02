// Code generated by SQLBoiler 4.16.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Translate is an object representing the database table.
type Translate struct {
	GameID    string      `boil:"game_id" json:"game_id" toml:"game_id" yaml:"game_id"`
	RawWord   string      `boil:"raw_word" json:"raw_word" toml:"raw_word" yaml:"raw_word"`
	TransWord null.String `boil:"trans_word" json:"trans_word,omitempty" toml:"trans_word" yaml:"trans_word,omitempty"`

	R *translateR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L translateL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var TranslateColumns = struct {
	GameID    string
	RawWord   string
	TransWord string
}{
	GameID:    "game_id",
	RawWord:   "raw_word",
	TransWord: "trans_word",
}

var TranslateTableColumns = struct {
	GameID    string
	RawWord   string
	TransWord string
}{
	GameID:    "Translate.game_id",
	RawWord:   "Translate.raw_word",
	TransWord: "Translate.trans_word",
}

// Generated where

var TranslateWhere = struct {
	GameID    whereHelperstring
	RawWord   whereHelperstring
	TransWord whereHelpernull_String
}{
	GameID:    whereHelperstring{field: "\"Translate\".\"game_id\""},
	RawWord:   whereHelperstring{field: "\"Translate\".\"raw_word\""},
	TransWord: whereHelpernull_String{field: "\"Translate\".\"trans_word\""},
}

// TranslateRels is where relationship names are stored.
var TranslateRels = struct {
	Game string
}{
	Game: "Game",
}

// translateR is where relationships are stored.
type translateR struct {
	Game *Game `boil:"Game" json:"Game" toml:"Game" yaml:"Game"`
}

// NewStruct creates a new relationship struct
func (*translateR) NewStruct() *translateR {
	return &translateR{}
}

func (r *translateR) GetGame() *Game {
	if r == nil {
		return nil
	}
	return r.Game
}

// translateL is where Load methods for each relationship are stored.
type translateL struct{}

var (
	translateAllColumns            = []string{"game_id", "raw_word", "trans_word"}
	translateColumnsWithoutDefault = []string{"game_id", "raw_word"}
	translateColumnsWithDefault    = []string{"trans_word"}
	translatePrimaryKeyColumns     = []string{"game_id", "raw_word"}
	translateGeneratedColumns      = []string{}
)

type (
	// TranslateSlice is an alias for a slice of pointers to Translate.
	// This should almost always be used instead of []Translate.
	TranslateSlice []*Translate
	// TranslateHook is the signature for custom Translate hook methods
	TranslateHook func(context.Context, boil.ContextExecutor, *Translate) error

	translateQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	translateType                 = reflect.TypeOf(&Translate{})
	translateMapping              = queries.MakeStructMapping(translateType)
	translatePrimaryKeyMapping, _ = queries.BindMapping(translateType, translateMapping, translatePrimaryKeyColumns)
	translateInsertCacheMut       sync.RWMutex
	translateInsertCache          = make(map[string]insertCache)
	translateUpdateCacheMut       sync.RWMutex
	translateUpdateCache          = make(map[string]updateCache)
	translateUpsertCacheMut       sync.RWMutex
	translateUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var translateAfterSelectMu sync.Mutex
var translateAfterSelectHooks []TranslateHook

var translateBeforeInsertMu sync.Mutex
var translateBeforeInsertHooks []TranslateHook
var translateAfterInsertMu sync.Mutex
var translateAfterInsertHooks []TranslateHook

var translateBeforeUpdateMu sync.Mutex
var translateBeforeUpdateHooks []TranslateHook
var translateAfterUpdateMu sync.Mutex
var translateAfterUpdateHooks []TranslateHook

var translateBeforeDeleteMu sync.Mutex
var translateBeforeDeleteHooks []TranslateHook
var translateAfterDeleteMu sync.Mutex
var translateAfterDeleteHooks []TranslateHook

var translateBeforeUpsertMu sync.Mutex
var translateBeforeUpsertHooks []TranslateHook
var translateAfterUpsertMu sync.Mutex
var translateAfterUpsertHooks []TranslateHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Translate) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range translateAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Translate) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range translateBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Translate) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range translateAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Translate) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range translateBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Translate) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range translateAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Translate) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range translateBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Translate) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range translateAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Translate) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range translateBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Translate) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range translateAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddTranslateHook registers your hook function for all future operations.
func AddTranslateHook(hookPoint boil.HookPoint, translateHook TranslateHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		translateAfterSelectMu.Lock()
		translateAfterSelectHooks = append(translateAfterSelectHooks, translateHook)
		translateAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		translateBeforeInsertMu.Lock()
		translateBeforeInsertHooks = append(translateBeforeInsertHooks, translateHook)
		translateBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		translateAfterInsertMu.Lock()
		translateAfterInsertHooks = append(translateAfterInsertHooks, translateHook)
		translateAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		translateBeforeUpdateMu.Lock()
		translateBeforeUpdateHooks = append(translateBeforeUpdateHooks, translateHook)
		translateBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		translateAfterUpdateMu.Lock()
		translateAfterUpdateHooks = append(translateAfterUpdateHooks, translateHook)
		translateAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		translateBeforeDeleteMu.Lock()
		translateBeforeDeleteHooks = append(translateBeforeDeleteHooks, translateHook)
		translateBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		translateAfterDeleteMu.Lock()
		translateAfterDeleteHooks = append(translateAfterDeleteHooks, translateHook)
		translateAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		translateBeforeUpsertMu.Lock()
		translateBeforeUpsertHooks = append(translateBeforeUpsertHooks, translateHook)
		translateBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		translateAfterUpsertMu.Lock()
		translateAfterUpsertHooks = append(translateAfterUpsertHooks, translateHook)
		translateAfterUpsertMu.Unlock()
	}
}

// One returns a single translate record from the query.
func (q translateQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Translate, error) {
	o := &Translate{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for Translate")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Translate records from the query.
func (q translateQuery) All(ctx context.Context, exec boil.ContextExecutor) (TranslateSlice, error) {
	var o []*Translate

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Translate slice")
	}

	if len(translateAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Translate records in the query.
func (q translateQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count Translate rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q translateQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if Translate exists")
	}

	return count > 0, nil
}

// Game pointed to by the foreign key.
func (o *Translate) Game(mods ...qm.QueryMod) gameQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.GameID),
	}

	queryMods = append(queryMods, mods...)

	return Games(queryMods...)
}

// LoadGame allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (translateL) LoadGame(ctx context.Context, e boil.ContextExecutor, singular bool, maybeTranslate interface{}, mods queries.Applicator) error {
	var slice []*Translate
	var object *Translate

	if singular {
		var ok bool
		object, ok = maybeTranslate.(*Translate)
		if !ok {
			object = new(Translate)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeTranslate)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeTranslate))
			}
		}
	} else {
		s, ok := maybeTranslate.(*[]*Translate)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeTranslate)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeTranslate))
			}
		}
	}

	args := make(map[interface{}]struct{})
	if singular {
		if object.R == nil {
			object.R = &translateR{}
		}
		args[object.GameID] = struct{}{}

	} else {
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &translateR{}
			}

			args[obj.GameID] = struct{}{}

		}
	}

	if len(args) == 0 {
		return nil
	}

	argsSlice := make([]interface{}, len(args))
	i := 0
	for arg := range args {
		argsSlice[i] = arg
		i++
	}

	query := NewQuery(
		qm.From(`Game`),
		qm.WhereIn(`Game.id in ?`, argsSlice...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Game")
	}

	var resultSlice []*Game
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Game")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for Game")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for Game")
	}

	if len(gameAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Game = foreign
		if foreign.R == nil {
			foreign.R = &gameR{}
		}
		foreign.R.GameTranslates = append(foreign.R.GameTranslates, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.GameID == foreign.ID {
				local.R.Game = foreign
				if foreign.R == nil {
					foreign.R = &gameR{}
				}
				foreign.R.GameTranslates = append(foreign.R.GameTranslates, local)
				break
			}
		}
	}

	return nil
}

// SetGame of the translate to the related item.
// Sets o.R.Game to related.
// Adds o to related.R.GameTranslates.
func (o *Translate) SetGame(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Game) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"Translate\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, []string{"game_id"}),
		strmangle.WhereClause("\"", "\"", 0, translatePrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.GameID, o.RawWord}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.GameID = related.ID
	if o.R == nil {
		o.R = &translateR{
			Game: related,
		}
	} else {
		o.R.Game = related
	}

	if related.R == nil {
		related.R = &gameR{
			GameTranslates: TranslateSlice{o},
		}
	} else {
		related.R.GameTranslates = append(related.R.GameTranslates, o)
	}

	return nil
}

// Translates retrieves all the records using an executor.
func Translates(mods ...qm.QueryMod) translateQuery {
	mods = append(mods, qm.From("\"Translate\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"Translate\".*"})
	}

	return translateQuery{q}
}

// FindTranslate retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindTranslate(ctx context.Context, exec boil.ContextExecutor, gameID string, rawWord string, selectCols ...string) (*Translate, error) {
	translateObj := &Translate{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"Translate\" where \"game_id\"=? AND \"raw_word\"=?", sel,
	)

	q := queries.Raw(query, gameID, rawWord)

	err := q.Bind(ctx, exec, translateObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from Translate")
	}

	if err = translateObj.doAfterSelectHooks(ctx, exec); err != nil {
		return translateObj, err
	}

	return translateObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Translate) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no Translate provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(translateColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	translateInsertCacheMut.RLock()
	cache, cached := translateInsertCache[key]
	translateInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			translateAllColumns,
			translateColumnsWithDefault,
			translateColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(translateType, translateMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(translateType, translateMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"Translate\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"Translate\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into Translate")
	}

	if !cached {
		translateInsertCacheMut.Lock()
		translateInsertCache[key] = cache
		translateInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Translate.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Translate) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	translateUpdateCacheMut.RLock()
	cache, cached := translateUpdateCache[key]
	translateUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			translateAllColumns,
			translatePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update Translate, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"Translate\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 0, wl),
			strmangle.WhereClause("\"", "\"", 0, translatePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(translateType, translateMapping, append(wl, translatePrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update Translate row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for Translate")
	}

	if !cached {
		translateUpdateCacheMut.Lock()
		translateUpdateCache[key] = cache
		translateUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q translateQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for Translate")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for Translate")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o TranslateSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), translatePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"Translate\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, translatePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in translate slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all translate")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Translate) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no Translate provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(translateColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	translateUpsertCacheMut.RLock()
	cache, cached := translateUpsertCache[key]
	translateUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			translateAllColumns,
			translateColumnsWithDefault,
			translateColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			translateAllColumns,
			translatePrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert Translate, could not build update column list")
		}

		ret := strmangle.SetComplement(translateAllColumns, strmangle.SetIntersect(insert, update))

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(translatePrimaryKeyColumns))
			copy(conflict, translatePrimaryKeyColumns)
		}
		cache.query = buildUpsertQuerySQLite(dialect, "\"Translate\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(translateType, translateMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(translateType, translateMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert Translate")
	}

	if !cached {
		translateUpsertCacheMut.Lock()
		translateUpsertCache[key] = cache
		translateUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Translate record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Translate) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Translate provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), translatePrimaryKeyMapping)
	sql := "DELETE FROM \"Translate\" WHERE \"game_id\"=? AND \"raw_word\"=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from Translate")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for Translate")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q translateQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no translateQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from Translate")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for Translate")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o TranslateSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(translateBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), translatePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"Translate\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, translatePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from translate slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for Translate")
	}

	if len(translateAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Translate) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindTranslate(ctx, exec, o.GameID, o.RawWord)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *TranslateSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := TranslateSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), translatePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"Translate\".* FROM \"Translate\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, translatePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in TranslateSlice")
	}

	*o = slice

	return nil
}

// TranslateExists checks if the Translate row exists.
func TranslateExists(ctx context.Context, exec boil.ContextExecutor, gameID string, rawWord string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"Translate\" where \"game_id\"=? AND \"raw_word\"=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, gameID, rawWord)
	}
	row := exec.QueryRowContext(ctx, sql, gameID, rawWord)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if Translate exists")
	}

	return exists, nil
}

// Exists checks if the Translate row exists.
func (o *Translate) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return TranslateExists(ctx, exec, o.GameID, o.RawWord)
}