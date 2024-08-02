// Code generated by SQLBoiler 4.16.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testTranslates(t *testing.T) {
	t.Parallel()

	query := Translates()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testTranslatesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Translate{}
	if err = randomize.Struct(seed, o, translateDBTypes, true, translateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Translate struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Translates().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTranslatesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Translate{}
	if err = randomize.Struct(seed, o, translateDBTypes, true, translateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Translate struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Translates().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Translates().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTranslatesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Translate{}
	if err = randomize.Struct(seed, o, translateDBTypes, true, translateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Translate struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := TranslateSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Translates().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTranslatesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Translate{}
	if err = randomize.Struct(seed, o, translateDBTypes, true, translateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Translate struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := TranslateExists(ctx, tx, o.GameID, o.RawWord)
	if err != nil {
		t.Errorf("Unable to check if Translate exists: %s", err)
	}
	if !e {
		t.Errorf("Expected TranslateExists to return true, but got false.")
	}
}

func testTranslatesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Translate{}
	if err = randomize.Struct(seed, o, translateDBTypes, true, translateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Translate struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	translateFound, err := FindTranslate(ctx, tx, o.GameID, o.RawWord)
	if err != nil {
		t.Error(err)
	}

	if translateFound == nil {
		t.Error("want a record, got nil")
	}
}

func testTranslatesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Translate{}
	if err = randomize.Struct(seed, o, translateDBTypes, true, translateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Translate struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Translates().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testTranslatesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Translate{}
	if err = randomize.Struct(seed, o, translateDBTypes, true, translateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Translate struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Translates().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testTranslatesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	translateOne := &Translate{}
	translateTwo := &Translate{}
	if err = randomize.Struct(seed, translateOne, translateDBTypes, false, translateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Translate struct: %s", err)
	}
	if err = randomize.Struct(seed, translateTwo, translateDBTypes, false, translateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Translate struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = translateOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = translateTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Translates().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testTranslatesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	translateOne := &Translate{}
	translateTwo := &Translate{}
	if err = randomize.Struct(seed, translateOne, translateDBTypes, false, translateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Translate struct: %s", err)
	}
	if err = randomize.Struct(seed, translateTwo, translateDBTypes, false, translateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Translate struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = translateOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = translateTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Translates().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func translateBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Translate) error {
	*o = Translate{}
	return nil
}

func translateAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Translate) error {
	*o = Translate{}
	return nil
}

func translateAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Translate) error {
	*o = Translate{}
	return nil
}

func translateBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Translate) error {
	*o = Translate{}
	return nil
}

func translateAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Translate) error {
	*o = Translate{}
	return nil
}

func translateBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Translate) error {
	*o = Translate{}
	return nil
}

func translateAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Translate) error {
	*o = Translate{}
	return nil
}

func translateBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Translate) error {
	*o = Translate{}
	return nil
}

func translateAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Translate) error {
	*o = Translate{}
	return nil
}

func testTranslatesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Translate{}
	o := &Translate{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, translateDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Translate object: %s", err)
	}

	AddTranslateHook(boil.BeforeInsertHook, translateBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	translateBeforeInsertHooks = []TranslateHook{}

	AddTranslateHook(boil.AfterInsertHook, translateAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	translateAfterInsertHooks = []TranslateHook{}

	AddTranslateHook(boil.AfterSelectHook, translateAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	translateAfterSelectHooks = []TranslateHook{}

	AddTranslateHook(boil.BeforeUpdateHook, translateBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	translateBeforeUpdateHooks = []TranslateHook{}

	AddTranslateHook(boil.AfterUpdateHook, translateAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	translateAfterUpdateHooks = []TranslateHook{}

	AddTranslateHook(boil.BeforeDeleteHook, translateBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	translateBeforeDeleteHooks = []TranslateHook{}

	AddTranslateHook(boil.AfterDeleteHook, translateAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	translateAfterDeleteHooks = []TranslateHook{}

	AddTranslateHook(boil.BeforeUpsertHook, translateBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	translateBeforeUpsertHooks = []TranslateHook{}

	AddTranslateHook(boil.AfterUpsertHook, translateAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	translateAfterUpsertHooks = []TranslateHook{}
}

func testTranslatesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Translate{}
	if err = randomize.Struct(seed, o, translateDBTypes, true, translateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Translate struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Translates().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testTranslatesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Translate{}
	if err = randomize.Struct(seed, o, translateDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Translate struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(translateColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Translates().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testTranslateToOneGameUsingGame(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Translate
	var foreign Game

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, translateDBTypes, false, translateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Translate struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, gameDBTypes, false, gameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Game struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.GameID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Game().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	ranAfterSelectHook := false
	AddGameHook(boil.AfterSelectHook, func(ctx context.Context, e boil.ContextExecutor, o *Game) error {
		ranAfterSelectHook = true
		return nil
	})

	slice := TranslateSlice{&local}
	if err = local.L.LoadGame(ctx, tx, false, (*[]*Translate)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Game == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Game = nil
	if err = local.L.LoadGame(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Game == nil {
		t.Error("struct should have been eager loaded")
	}

	if !ranAfterSelectHook {
		t.Error("failed to run AfterSelect hook for relationship")
	}
}

func testTranslateToOneSetOpGameUsingGame(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Translate
	var b, c Game

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, translateDBTypes, false, strmangle.SetComplement(translatePrimaryKeyColumns, translateColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, gameDBTypes, false, strmangle.SetComplement(gamePrimaryKeyColumns, gameColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, gameDBTypes, false, strmangle.SetComplement(gamePrimaryKeyColumns, gameColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Game{&b, &c} {
		err = a.SetGame(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Game != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.GameTranslates[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.GameID != x.ID {
			t.Error("foreign key was wrong value", a.GameID)
		}

		if exists, err := TranslateExists(ctx, tx, a.GameID, a.RawWord); err != nil {
			t.Fatal(err)
		} else if !exists {
			t.Error("want 'a' to exist")
		}

	}
}

func testTranslatesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Translate{}
	if err = randomize.Struct(seed, o, translateDBTypes, true, translateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Translate struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testTranslatesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Translate{}
	if err = randomize.Struct(seed, o, translateDBTypes, true, translateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Translate struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := TranslateSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testTranslatesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Translate{}
	if err = randomize.Struct(seed, o, translateDBTypes, true, translateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Translate struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Translates().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	translateDBTypes = map[string]string{`GameID`: `TEXT`, `RawWord`: `TEXT`, `TransWord`: `TEXT`}
	_                = bytes.MinRead
)

func testTranslatesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(translatePrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(translateAllColumns) == len(translatePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Translate{}
	if err = randomize.Struct(seed, o, translateDBTypes, true, translateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Translate struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Translates().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, translateDBTypes, true, translatePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Translate struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testTranslatesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(translateAllColumns) == len(translatePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Translate{}
	if err = randomize.Struct(seed, o, translateDBTypes, true, translateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Translate struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Translates().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, translateDBTypes, true, translatePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Translate struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(translateAllColumns, translatePrimaryKeyColumns) {
		fields = translateAllColumns
	} else {
		fields = strmangle.SetComplement(
			translateAllColumns,
			translatePrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := TranslateSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testTranslatesUpsert(t *testing.T) {
	t.Parallel()
	if len(translateAllColumns) == len(translatePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Translate{}
	if err = randomize.Struct(seed, &o, translateDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Translate struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Translate: %s", err)
	}

	count, err := Translates().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, translateDBTypes, false, translatePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Translate struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Translate: %s", err)
	}

	count, err = Translates().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
