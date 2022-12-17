package goPg

import (
	"ADMSPublic/conf"
	"context"
	"fmt"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	log "github.com/sirupsen/logrus"
	//"../AgriInject/conf"
)

type GoPg struct {
	Conn *pg.DB
}

type dbLogger struct{}

func (d dbLogger) BeforeQuery(c context.Context, _ *pg.QueryEvent) (context.Context, error) {
	return c, nil
}

func (d dbLogger) AfterQuery(_ context.Context, q *pg.QueryEvent) error {
	query, err := q.FormattedQuery()
	if err != nil {
		log.Error(err)
	}
	log.Info(string(query))
	return nil
}

func New(conf conf.Config) (db GoPg, err error) {
	db = GoPg{}
	db.Conn = pg.Connect(&pg.Options{
		Addr:     fmt.Sprintf("%s:%d", conf.DbHost, "conf.DbPort"),
		Dialer:   nil,
		User:     conf.DbUser,
		Password: conf.DbPassword,
		Database: conf.DbDatabase,
	})
	// logger := dbLogger{}
	// db.Conn.AddQueryHook(logger)
	err = db.createSchema()
	if err != nil {
		return
	}
	return
}

func (db *GoPg) createSchema() (err error) {
	for _, model := range []interface{}{
		(*TallySheet)(nil),
		(*Questionnaire)(nil),
		(*OverwrittenQuestionnaire)(nil),
	} {
		err := db.Conn.Model(model).CreateTable(&orm.CreateTableOptions{
			IfNotExists:   true,
			FKConstraints: true,
		})
		if err != nil {
			return err
		}
	}
	return
}

func (db *GoPg) Insert(collection Root) (tx *pg.Tx, err error) {
	tx, err = db.Conn.Begin()
	if err != nil {
		return tx, fmt.Errorf("error when trying to initiate transaction : %w", err)
	}

	tallySheet := collection.TallySheet
	_, err = tx.Model(tallySheet).Insert()
	if err != nil {
		return tx, fmt.Errorf("error when insert tally sheet %s : %w",
			tallySheet.TallySheetNo,
			err)
	}
	questionnaires := collection.Questionnaire
	for _, questionnaire := range questionnaires {
		_, err := tx.Model(questionnaire).Insert()
		if err != nil {
			return tx, fmt.Errorf("error when trying to insert questionnaire %s : %w",
				questionnaire.FormName,
				err)
		}
	}

	// TODO update tally sheet

	_, err = tx.QueryOne(tallySheet, `
SELECT *
FROM (select count(q.questionnaire_num) as updated_agri_professionals
      from questionnaires q,
           tally_sheets t
      where q.tally_sheet_no = t.tally_sheet_no
        and t.tally_sheet_no = ?id
	    and q.agri_labor_code = 1) as updated_agri_professionals,
     (select count(q.questionnaire_num) as updated_fishing_professionals
      from questionnaires q,
           tally_sheets t
      where q.tally_sheet_no = t.tally_sheet_no
        and t.tally_sheet_no = ?id
        and q.is_fishing_related = 1) as updated_fishing_professionals,
     (select count(q.questionnaire_num) as updated_house5_more
      from questionnaires q,
           tally_sheets t
      where q.tally_sheet_no = t.tally_sheet_no
        and t.tally_sheet_no = ?id
        and q.cultivated_land > 0.05) as updated_house5_more,
     (select count(q.questionnaire_num) as updated_house_fisheries
      from questionnaires q,
           tally_sheets t
      where q.tally_sheet_no = t.tally_sheet_no
        and t.tally_sheet_no = ?id
        and q.pond_land > 0) as updated_house_fisheries,
     (select count(q.questionnaire_num) as updated_house_no_land
      from questionnaires q,
           tally_sheets t
      where q.tally_sheet_no = t.tally_sheet_no
        and t.tally_sheet_no = ?id
        and (q.total_land = 0 or q.total_land is null)) as updated_house_no_land,
     (select count(q.questionnaire_num) as updated_house_received_land
      from questionnaires q,
           tally_sheets t
      where q.tally_sheet_no = t.tally_sheet_no
        and t.tally_sheet_no = ?id
        and q.land_taken > 0) as updated_house_received_land,
     (select count(q.house_serial) as updated_total_house
      from questionnaires q,
           tally_sheets t
      where q.tally_sheet_no = t.tally_sheet_no
        and q.house_serial is not null
        and t.tally_sheet_no = ?id) as updated_total_house,
     (select (sum(greatest(q.buffalo_at_farm,0)) + sum(greatest(q.buffalo_at_home,0))) as updated_buffalo_count
      from questionnaires q,
           tally_sheets t
      where q.tally_sheet_no = t.tally_sheet_no
        and t.tally_sheet_no = ?id) as updated_buffalo_count,
     (select (sum(greatest(q.cock_at_farm,0)) + sum(greatest(q.cock_at_home,0))) as updated_cock_count
      from questionnaires q,
           tally_sheets t
      where q.tally_sheet_no = t.tally_sheet_no
        and t.tally_sheet_no = ?id) as updated_cock_count,
     (select (sum(greatest(q.cow_at_farm,0)) + sum(greatest(q.cow_at_home,0))) as updated_cow_count
      from questionnaires q,
           tally_sheets t
      where q.tally_sheet_no = t.tally_sheet_no
        and t.tally_sheet_no = ?id) as updated_cow_count,
     (select (sum(greatest(q.duck_at_farm,0)) + sum(greatest(q.duck_at_home,0))) as updated_duck_count
      from questionnaires q,
           tally_sheets t
      where q.tally_sheet_no = t.tally_sheet_no
        and t.tally_sheet_no = ?id) as updated_duck_count,
     (select (sum(greatest(q.goat_at_farm,0)) + sum(greatest(q.goat_at_home,0))) as updated_goat_count
      from questionnaires q,
           tally_sheets t
      where q.tally_sheet_no = t.tally_sheet_no
        and t.tally_sheet_no = ?id) as updated_goat_count,
     (select (sum(greatest(q.sheep_at_farm,0)) + sum(greatest(q.sheep_at_home,0))) as updated_sheep_count
      from questionnaires q,
           tally_sheets t
      where q.tally_sheet_no = t.tally_sheet_no
        and t.tally_sheet_no = ?id) as updated_sheep_count,
     (select (sum(greatest(q.turkey_at_farm,0)) + sum(greatest(q.turkey_at_home,0))) as updated_turkey_count
      from questionnaires q,
           tally_sheets t
      where q.tally_sheet_no = t.tally_sheet_no
        and t.tally_sheet_no = ?id) as updated_turkey_count;`,
		struct {
			Id string
		}{
			Id: tallySheet.TallySheetNo,
		})
	if err != nil {
		return tx, fmt.Errorf("error when trying to update tally sheet %s data: %w",
			tallySheet.TallySheetNo, err)
	}
	_, err = tx.Model(tallySheet).WherePK().Update()
	if err != nil {
		return tx, fmt.Errorf("error when trying to update tally sheet %s data: %w",
			tallySheet.TallySheetNo, err)
	}
	return
}
