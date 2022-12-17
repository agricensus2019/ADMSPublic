package model

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/oleiade/reflections"
	"golang.org/x/text/language"
	"golang.org/x/text/message"

	log "github.com/sirupsen/logrus"

	"ADMSPublic/conf"
)

type Db struct {
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

func (db *Db) Init(conf conf.Config) (err error) {
	db.Conn = pg.Connect(&pg.Options{
		Addr:     conf.DbHost,
		Dialer:   nil,
		User:     conf.DbUser,
		Password: conf.DbPassword,
		Database: conf.DbDatabase,
	})
	if conf.DbLog {
		logger := dbLogger{}
		db.Conn.AddQueryHook(logger)
	}
	if conf.DbInit {
		db.createExtension()
		err = db.createSchema()
		if err != nil {
			return
		}
		db.createIndex()

	}
	return
}

func (db *Db) Close() {
	err := db.Conn.Close()
	if err != nil {
		log.Error(err)
	}
}

func (db *Db) createIndex() {
	for name, index := range index {
		log.Infof("create index : %s", name)
		_, err := db.Conn.Model((*Aggregates)(nil)).Exec(index)
		if err != nil {
			log.Error(err)
		}
	}
}

func (db *Db) createSchema() (err error) {
	for _, model := range []interface{}{
		(*Aggregates)(nil),
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

func (db *Db) createExtension() {
	_, err := db.Conn.Exec("CREATE EXTENSION if not exists ltree;")
	if err != nil {
		log.Error(err)
	}
}

type RawTableData struct {
	Data float64
	Rmo  int
}

func GetGeoRequest(division, district, upazilla, union, mouza string) (selector string, count uint, err error) {
	if division != "" {
		var divisionInt int
		divisionInt, err = strconv.Atoi(division)
		if err != nil {
			return
		}
		selector = fmt.Sprintf("%02d", divisionInt)
		count = 1
	} else {
		return
	}
	if district != "" {
		var districtInt int
		districtInt, err = strconv.Atoi(district)
		selector += "." + fmt.Sprintf("%02d", districtInt)
		count = 2
	} else {
		return
	}
	if upazilla != "" {
		var upazillaInt int
		upazillaInt, err = strconv.Atoi(upazilla)
		selector += "." + fmt.Sprintf("%02d", upazillaInt)
		count = 3
	} else {
		return
	}
	if union != "" {
		var unionInt int
		unionInt, err = strconv.Atoi(union)
		selector += "." + fmt.Sprintf("%03d", unionInt)
		count = 4
	} else {
		return
	}
	if mouza != "" {
		var mouzaInt int
		mouzaInt, err = strconv.Atoi(mouza)
		selector += "." + fmt.Sprintf("%03d", mouzaInt)
		count = 5
	} else {
		return
	}
	return
}

func (db *Db) GetAgregate(division, district, upazilla, union, mouza, tableName string) (tableData []RawTableData, err error) {
	columns := ""
	conditions := ""
	summable := false

	// sex = male
	// sex2 = female
	// sex3 = hijra

	switch tableName {
	case "1":
		columns = "hh_sno"
		summable = true
	case "2":
		columns = "sf+mf+lf"
		summable = true
	case "3":
		columns = "sf"
		summable = true
	case "4":
		columns = "mf"
		summable = true
	case "5":
		columns = "lf"
		summable = true
	case "6":
		columns = "hh_sno-c04gtrhh"
		summable = true
	case "7":
		columns = "sum(hh_a)"
		conditions = "true = true GROUP BY rmo"
		summable = true
	case "8":
		columns = "sum(hh_f)"
		conditions = "true = true GROUP BY rmo"
		summable = true
	case "9":
		columns = "SUM(c02m+c02f+c02h+c03m+c03f+c03h)::NUMERIC/SUM(hh_sno)::NUMERIC"
		summable = false
	case "10":
		columns = "SUM (c02mfarm+c02ffarm+c02hfarm+c03mfarm+c03ffarm+c03hfarm) ::decimal / SUM (sf+mf+lf) ::decimal"
		summable = false
	case "11":
		columns = "sum(c07) / SUM (hh_sno)"
		summable = false
	case "12":
		columns = "sum(c07farm) / SUM (sf + mf + lf)"
		summable = false
	case "13":
		columns = "sum(c08) / SUM (hh_sno)"
		summable = false
	case "14":
		columns = "sum(c08farm) / SUM (sf + mf + lf)"
		summable = false
	case "15":
		columns = "sum(sex)"
		conditions = "true = true GROUP BY rmo"
		summable = true
	case "16":
		columns = "sum(sex2)"
		conditions = "true = true GROUP BY rmo"
		summable = true
	case "17":
		columns = "SUM(c19gtrhh)"
		conditions = "true = true GROUP BY rmo"
		summable = true
	case "18":
		columns = "sum(c19)"
		conditions = "true = true GROUP BY rmo"
		summable = true
	case "19":
		columns = "SUM (c19smlfhh)"
		conditions = "true = true GROUP BY rmo"
		summable = true
	case "20":
		columns = "sum(c19farm)"
		conditions = "true = true GROUP BY rmo"
		summable = true
	case "21":
		columns = "c33h+c33f"
		summable = true
	case "22":
		columns = "c34h+c34f"
		summable = true
	case "23":
		columns = "sum(c35h+c35f)"
		conditions = "true = true GROUP BY rmo"
		summable = true
	case "24":
		columns = "sum(c36h+c36f)"
		conditions = "true = true GROUP BY rmo"
		summable = true
	case "25":
		columns = "sum(c28h+c28f)"
		conditions = "true = true GROUP BY rmo"
		summable = true
	case "26":
		columns = "sum(c29h+c29f)"
		conditions = "true = true GROUP BY rmo"
		summable = true
	case "27":
		columns = "sum(t101+t102+t103+t104+t105+t112+t113+t114+t121+t122+t123+t124+t125+t127+t128+t129+t130+t131+t132+t134+t135+t157+t158+t159+t160+t161+t167+t169+t175+t176+t177+t179+t182+t185+t106+t107+t108+t109+t110+t111+t115+t116+t117+t118+t119+t120+t126+t133+t136+t137+t138+t139+t140+t141+t142+t143+t144+t145+t146+t147+t148+t149+t150+t151+t152+t153+t154+t155+t156+t162+t163+t164+t165+t166+t168+t170+t171+t172+t173+t174+t178+t180+t181+t183+t184+t186+t187+t188+t189+t190+t191+t192+t193+t194+t195+t196+t197+t198+t199+t200+t201+t202+t203)/ sum(c13)"
		summable = false
	case "28":
		columns = "sum(c18-c15)"
		conditions = "true = true GROUP BY rmo"
		summable = true
	default:
		return tableData, fmt.Errorf(("don't know this table name"))
	}

	geoCodeReq, count, err := GetGeoRequest(division, district, upazilla, union, mouza)
	if err != nil {
		return
	}

	var query string
	switch summable {
	case true:
		query = fmt.Sprintf(`SELECT %s as data, rmo FROM aggregates where subpath(geocode, 0,%d) = ?;`, columns, count)
		if conditions != "" {
			query = strings.Replace(query, ";", fmt.Sprintf(" AND %s;", conditions), 1)
		}
	case false:
		query = fmt.Sprintf(`
		SELECT %s as data,
		    rmo
		FROM aggregates
		WHERE subpath(geocode, 0, %d) = ?
		GROUP BY rmo
		UNION
		SELECT %s as data,
		    -1 as rmo
		from aggregates;
		`, columns, count, columns,
		)
	}
	_, err = db.Conn.Query(&tableData, query,
		geoCodeReq)
	if err != nil {
		log.Error(err)
		return
	}
	return
}

type OccupationHouseHoldHead struct {
	Occ   uint
	Occ2  uint
	Occ3  uint
	Occ4  uint
	Occ5  uint
	Total uint
}

func (db *Db) GetOccupationOfHouseHold(division, district, upazilla, union, mouza string) (data OccupationHouseHoldHead, err error) {
	geoCodeReq, count, err := GetGeoRequest(division, district, upazilla, union, mouza)
	if err != nil {
		return
	}
	query := fmt.Sprintf(`
	select sum(occ) as occ,
    sum(occ2) as occ2,
    sum(occ3) as occ3,
    sum(occ4) as occ4,
    sum(occ5) as occ5,
	(sum(occ) + sum(occ2) + sum(occ3) + sum(occ4) + sum(occ5)) as total
	from aggregates
	where subpath(geocode, 0, %d) = ?;`,
		count)
	_, err = db.Conn.QueryOne(&data, query,
		geoCodeReq)
	if err != nil {
		log.Error(err)
		return
	}
	return
}

type EducationOfTheHouseholdHead struct {
	NoEducation               uint
	Class1                    uint
	Class2                    uint
	Class3                    uint
	Class4                    uint
	Class5                    uint
	Class6                    uint
	Class7                    uint
	Class8                    uint
	Class9                    uint
	Ssc                       uint
	Hsc                       uint
	BachelorEquivalent        uint
	MastersEquivalentOrHigher uint
	Total                     uint
}

func (db *Db) GetEducationOfTheHouseholdHead(division, district, upazilla, union, mouza string) (data EducationOfTheHouseholdHead, err error) {
	geoCodeReq, count, err := GetGeoRequest(division, district, upazilla, union, mouza)
	if err != nil {
		return
	}
	query := fmt.Sprintf(`
	select 
	sum(edu) as no_education,
    sum(edu1) as Class1,
    sum(edu2) as Class2,
    sum(edu3) as Class3,
    sum(edu4) as Class4,
    sum(edu5) as Class5,
    sum(edu6) as Class6,
    sum(edu7) as Class7,
    sum(edu8) as Class8,
    sum(edu9) as Class9, 
    sum(edu10) as Ssc,
    sum(edu12) as Hsc,
    sum(edu15) as Bachelor_equivalent,
    sum(edu18) as Masters_Equivalent_Or_Higher,
    (
        sum(edu) + 
		sum(edu1) + 
		sum(edu2) + 
		sum(edu3) + 
		sum(edu4) + 
		sum(edu5) + 
		sum(edu6) + 
		sum(edu7) + 
		sum(edu8) + 
		sum(edu9) + 
		sum(edu10) + 
		sum(edu12) + 
		sum(edu15) + 
		sum(edu18)
    ) as Total
from aggregates where subpath(geocode, 0,%d) = ?;`,
		count)
	_, err = db.Conn.QueryOne(&data, query,
		geoCodeReq)
	if err != nil {
		log.Error(err)
		return
	}
	return
}

type GenderOfTheHouseholdHead struct {
	Male   uint
	Female uint
	Hijra  uint
	Total  uint
}

func (db *Db) GetGenderOfTheHouseholdHead(division, district, upazilla, union, mouza string) (data GenderOfTheHouseholdHead, err error) {
	geoCodeReq, count, err := GetGeoRequest(division, district, upazilla, union, mouza)
	if err != nil {
		return
	}
	query := fmt.Sprintf(`
	select sum(sex) as male,
    sum(sex2) as female,
    sum(sex3) as hijra,
    (sum(sex) + sum(sex2) + sum(sex3)) as total
	from aggregates
	where subpath(geocode, 0, %d) = ?;`,
		count)
	_, err = db.Conn.QueryOne(&data, query,
		geoCodeReq)
	if err != nil {
		log.Error(err)
		return
	}
	return
}

type FisheryHolding struct {
	NumberOfFisheryHousehold uint
	Percentage               float64
}

func (db *Db) GetFisheryHolding(division, district, upazilla, union, mouza string) (data FisheryHolding, err error) {
	geoCodeReq, count, err := GetGeoRequest(division, district, upazilla, union, mouza)
	if err != nil {
		return
	}
	query := fmt.Sprintf(`
	SELECT sum(hh_f) as Number_Of_Fishery_Household,
    	(sum(hh_f)::NUMERIC / sum(hh_sno)::NUMERIC)::NUMERIC * 100 as Percentage
	FROM aggregates
	WHERE   subpath(geocode, 0, %d) = ?;`,
		count)
	_, err = db.Conn.QueryOne(&data, query,
		geoCodeReq)
	if err != nil {
		log.Error(err)
		return
	}
	return
}

type AgriculuralLaborHolding struct {
	NumberOfAgriLaborHouseHold uint
	Percentage                 float64
}

func (db *Db) GetAgriculuralLaborHolding(division, district, upazilla, union, mouza string) (data AgriculuralLaborHolding, err error) {
	geoCodeReq, count, err := GetGeoRequest(division, district, upazilla, union, mouza)
	if err != nil {
		return
	}
	query := fmt.Sprintf(`
	SELECT sum(hh_a) as Number_Of_Agri_Labor_House_Hold,
    	(sum(hh_f)::NUMERIC / sum(hh_sno)::NUMERIC)::NUMERIC * 100 as Percentage
	FROM aggregates
	WHERE  subpath(geocode, 0, %d) = ?;`,
		count)
	_, err = db.Conn.QueryOne(&data, query,
		geoCodeReq)
	if err != nil {
		log.Error(err)
		return
	}
	return
}

type HouseholdHeadInformation struct {
	NoEducation                   uint
	Class_I_V                     uint
	Class_VI_IX                   uint
	SccPassed                     uint
	HscPassed                     uint
	DegreePassed                  uint
	MasterPassed                  uint
	TotalEducation                uint
	Agriculture                   uint
	Industry                      uint
	Service                       uint
	Business                      uint
	Other                         uint
	TotalOccupation               uint
	FisheryHolding                uint
	FisheryHoldingPercentage      float64
	AgriculturalHolding           uint
	AgriculturalHoldingPercentage float64
	HouseholdMemberMale           uint
	HouseholdMemberFemale         uint
	HouseholdMemberHijra          uint
	HouseholdMemberTotal          uint
	HouseholdWorkerMale           uint
	HouseholdWorkerFemale         uint
	HouseholdWorkerHijra          uint
	HouseholdWorkerTotal          uint
	HouseholdWorker_10_14Male     uint
	HouseholdWorker_10_14Female   uint
	HouseholdWorker_10_14Hijra    uint
	HouseholdWorker_10_14Total    uint
	HouseholdWorker_15PlusMale    uint
	HouseholdWorker_15PlusFemale  uint
	HouseholdWorker_15PlusHijra   uint
	HouseholdWorker_15PlusTotal   uint
}

func (db *Db) GetHouseholdHeadInformation(division, district, upazilla, union, mouza string) (data HouseholdHeadInformation, err error) {
	geoCodeReq, count, err := GetGeoRequest(division, district, upazilla, union, mouza)
	if err != nil {
		return
	}
	query := fmt.Sprintf(`
	SELECT sum(edu) as no_education,
    	(sum(edu1) + sum(edu2) + sum(edu3) + sum(edu4) + sum(edu5)) as class_I_V,
    	(sum(edu6) + sum(edu7) + sum(edu8) + sum(edu9)) as class_VI_IX,
    	sum(edu10) as Scc_Passed,
    	sum(edu12) as Hsc_Passed,
    	sum(edu15) as Degree_Passed,
    	sum(edu18) as Master_Passed,
    	(
    	    sum(edu) + sum(edu1) + sum(edu2) + sum(edu3) + sum(edu4) + sum(edu5) + sum(edu6) + sum(edu7) + sum(edu8) + sum(edu9) + sum(edu10) + sum(edu12) + sum(edu15) + sum(edu18)
    	) as Total_Education,
    	sum(occ) as Agriculture,
    	sum(occ2) as Industry,
    	sum(occ3) as Service,
    	sum(occ4) as Business,
    	sum(occ5) as Other,
    	(sum(occ) + sum(occ2) + sum(occ3) + sum(occ4) + sum(occ5)) as Total_Occupation,
    	sum(c01m) as Household_Member_Male,
    	sum(c01f) as Household_Member_Female,
    	sum(c01h) as Household_Member_Hijra,
    	(sum(c01m) + sum(c01f) + sum(c01h)) as Household_Member_Total,
    	(sum(c02m) + sum(c03m)) as Household_Worker_Male,
    	(sum(c02f) + sum(c03f)) as Household_Worker_Female,
    	(sum(c02h) + sum(c03h)) as Household_Worker_Hijra,
    	(sum(c02m) + sum(c03m) + sum(c02f) + sum(c03f) + sum(c02h) + sum(c03h)) as Household_Worker_Total,
    	sum(c02m) as Household_Worker_10_14_Male,
    	sum(c02f) as Household_Worker_10_14_Female,
    	sum(c02h) as Household_Worker_10_14_Hijra,
    	(sum(c02m) + sum(c02f) + sum(c02h)) as Household_Worker_10_14_Total,
    	sum(c03m) as Household_Worker_15_Plus_Male,
    	sum(c03f) as Household_Worker_15_Plus_Female,
    	sum(c03h) as Household_Worker_15_Plus_Hijra,
    	(sum(c03m) + sum(c03f) + sum(c03h)) as Household_Worker_15_Plus_Total
	FROM aggregates
	WHERE subpath(geocode, 0, %d) = ?;`,
		count)
	_, err = db.Conn.QueryOne(&data, query,
		geoCodeReq)
	if err != nil {
		log.Error(err)
		return
	}
	return
}

type OccupationOfHouseholdHead struct {
	Agriculture uint
	Industry    uint
	Service     uint
	Business    uint
	Others      uint
	Total       uint
}

func (db *Db) GetOccupationOfHouseholdHead(division, district, upazilla, union, mouza string) (data OccupationOfHouseholdHead, err error) {
	geoCodeReq, count, err := GetGeoRequest(division, district, upazilla, union, mouza)
	if err != nil {
		return
	}
	query := fmt.Sprintf(`
	SELECT sum(occ) as Agriculture,
    	sum(occ2) as Industry,
    	sum(occ3) as Service,
    	sum(occ4) as Business,
    	sum(occ5) as Others,
    	(
    	    sum(occ) + sum(occ2) + sum(occ3) + sum(occ4) + sum(occ5)
    	) as Total
	FROM aggregates
	WHERE subpath(geocode, 0, %d) = ?;`,
		count)
	_, err = db.Conn.QueryOne(&data, query,
		geoCodeReq)
	if err != nil {
		log.Error(err)
		return
	}
	return
}

type TotalNumberOfHouseholdMembers struct {
	Male   uint
	Female uint
	Hijra  uint
	Total  uint
}

func (db *Db) GetTotalNumberOfHouseholdMembers(division, district, upazilla, union, mouza string) (data TotalNumberOfHouseholdMembers, err error) {
	geoCodeReq, count, err := GetGeoRequest(division, district, upazilla, union, mouza)
	if err != nil {
		return
	}
	query := fmt.Sprintf(`
	SELECT sum(c01m) as Male,
    	sum(c01f) as Female,
    	sum(c01h) as Hijra,
    	(sum(c01m) + sum(c01f) + sum(c01h)) as Total
	FROM aggregates
	WHERE subpath(geocode, 0, %d) = ?;`,
		count)
	_, err = db.Conn.QueryOne(&data, query,
		geoCodeReq)
	if err != nil {
		log.Error(err)
		return
	}
	return
}

func (db *Db) GetTotalNumberOfHouseholdWorkers(division, district, upazilla, union, mouza string) (data TotalNumberOfHouseholdMembers, err error) {
	geoCodeReq, count, err := GetGeoRequest(division, district, upazilla, union, mouza)
	if err != nil {
		return
	}
	query := fmt.Sprintf(`
	SELECT sum(c02m + c03m) as Male,
    	sum(c02f + c03f) as Female,
    	sum(c02h + c03h) as Hijra,
    	(sum(c02m + c03m) + sum(c02f + c03f) + sum(c02h + c03h)) as Total
	FROM aggregates
	WHERE subpath(geocode, 0, %d) = ?;`,
		count)
	_, err = db.Conn.QueryOne(&data, query,
		geoCodeReq)
	if err != nil {
		log.Error(err)
		return
	}
	return
}

func (db *Db) GetTotalNumberOfHouseholdWorkers1014(division, district, upazilla, union, mouza string) (data TotalNumberOfHouseholdMembers, err error) {
	geoCodeReq, count, err := GetGeoRequest(division, district, upazilla, union, mouza)
	if err != nil {
		return
	}
	query := fmt.Sprintf(`
	SELECT sum(c02m) as Male,
    	sum(c02f) as Female,
    	sum(c02h) as Hijra,
    	(sum(c02m) + sum(c02f) + sum(c02h)) as Total
	FROM aggregates
	WHERE subpath(geocode, 0, %d) = ?;`,
		count)
	_, err = db.Conn.QueryOne(&data, query,
		geoCodeReq)
	if err != nil {
		log.Error(err)
		return
	}
	return
}

func (db *Db) GetTotalNumberOfHouseholdWorkers15plus(division, district, upazilla, union, mouza string) (data TotalNumberOfHouseholdMembers, err error) {
	geoCodeReq, count, err := GetGeoRequest(division, district, upazilla, union, mouza)
	if err != nil {
		return
	}
	query := fmt.Sprintf(`
	SELECT sum(c03m) as Male,
    	sum(c03f) as Female,
    	sum(c03h) as Hijra,
    	(sum(c03m) + sum(c03f) + sum(c03h)) as Total
	FROM aggregates
	WHERE subpath(geocode, 0, %d) = ?;`,
		count)
	_, err = db.Conn.QueryOne(&data, query,
		geoCodeReq)
	if err != nil {
		log.Error(err)
		return
	}
	return
}

func (db *Db) GetGeoCode(geoCodeNumber string) (geoCode GeoCodes, err error) {
	geoCode = GeoCodes{
		GeocodeID: geoCodeNumber,
	}
	err = db.Conn.Model(&geoCode).WherePK().Select()
	return
}

type HouseholdLandInformation struct {
	Name                           string
	NumberOfReportingHoldingColumn string
	NumberOfReportingHoldings      uint
	NumberOfFarmHoldingsColumn     string
	NumberOfFarmHoldings           uint
	TotalAreaOfOwnLandColumn       string
	TotalAreaOfOwnLand             float64
	TotalFarmHoldingAreaColumn     string
	TotalFarmHoldingArea           float64
	AverageAreaPerFarmHolding      float64
}

func (db *Db) GetHouseholdLandInformation(division, district, upazilla, union, mouza string) (data []HouseholdLandInformation, err error) {
	geoCodeReq, count, err := GetGeoRequest(division, district, upazilla, union, mouza)
	if err != nil {
		return
	}
	data = []HouseholdLandInformation{
		{
			Name:                           "Own land",
			NumberOfReportingHoldingColumn: "c04gtrhh",
			NumberOfFarmHoldingsColumn:     "c04smlfhh",
			TotalAreaOfOwnLandColumn:       "c04",
			TotalFarmHoldingAreaColumn:     "c04smlf",
		},
		{
			Name:                           "Given land",
			NumberOfReportingHoldingColumn: "c05gtrhh",
			NumberOfFarmHoldingsColumn:     "c05smlfhh",
			TotalAreaOfOwnLandColumn:       "c05",
			TotalFarmHoldingAreaColumn:     "c05smlf",
		},
		{
			Name:                           "Taken land",
			NumberOfReportingHoldingColumn: "c06gtrhh",
			NumberOfFarmHoldingsColumn:     "c06smlfhh",
			TotalAreaOfOwnLandColumn:       "c06",
			TotalFarmHoldingAreaColumn:     "c06smlf",
		},
		{
			Name:                           "Operated land",
			NumberOfReportingHoldingColumn: "c07gtrhh",
			NumberOfFarmHoldingsColumn:     "c07smlfhh",
			TotalAreaOfOwnLandColumn:       "c07",
			TotalFarmHoldingAreaColumn:     "c07smlf",
		},
		{
			Name:                           "Homestead land",
			NumberOfReportingHoldingColumn: "c08gtrhh",
			NumberOfFarmHoldingsColumn:     "c08smlfhh",
			TotalAreaOfOwnLandColumn:       "c08",
			TotalFarmHoldingAreaColumn:     "c08smlf",
		},
		{
			Name:                           "Permanent Fellow land",
			NumberOfReportingHoldingColumn: "c11gtrhh",
			NumberOfFarmHoldingsColumn:     "c11smlfhh",
			TotalAreaOfOwnLandColumn:       "c11",
			TotalFarmHoldingAreaColumn:     "c11smlf",
		},
		{
			Name:                           "Uncultivated land",
			NumberOfReportingHoldingColumn: "c12gtrhh",
			NumberOfFarmHoldingsColumn:     "c12smlfhh",
			TotalAreaOfOwnLandColumn:       "c12",
			TotalFarmHoldingAreaColumn:     "c12smlf",
		},
		{
			Name:                           "Land under temporary crops",
			NumberOfReportingHoldingColumn: "c13gtrhh",
			NumberOfFarmHoldingsColumn:     "c13smlfhh",
			TotalAreaOfOwnLandColumn:       "c13",
			TotalFarmHoldingAreaColumn:     "c13smlf",
		},
		{
			Name:                           "Land under permanent crops",
			NumberOfReportingHoldingColumn: "c14gtrhh",
			NumberOfFarmHoldingsColumn:     "c14smlfhh",
			TotalAreaOfOwnLandColumn:       "c14",
			TotalFarmHoldingAreaColumn:     "c14smlf",
		},
		{
			Name:                           "Land under nursery",
			NumberOfReportingHoldingColumn: "c16gtrhh",
			NumberOfFarmHoldingsColumn:     "c16smlfhh",
			TotalAreaOfOwnLandColumn:       "c16",
			TotalFarmHoldingAreaColumn:     "c16smlf",
		},
		{
			Name:                           "Land under current fallow",
			NumberOfReportingHoldingColumn: "c17gtrhh",
			NumberOfFarmHoldingsColumn:     "c17smlfhh",
			TotalAreaOfOwnLandColumn:       "c17",
			TotalFarmHoldingAreaColumn:     "c17smlf",
		},
		{
			Name:                           "Total cultivated land",
			NumberOfReportingHoldingColumn: "c18gtrhh",
			NumberOfFarmHoldingsColumn:     "c18smlfhh",
			TotalAreaOfOwnLandColumn:       "c18",
			TotalFarmHoldingAreaColumn:     "c18smlf",
		},
		{
			Name:                           "Land under irrigation",
			NumberOfReportingHoldingColumn: "c19gtrhh",
			NumberOfFarmHoldingsColumn:     "c19smlfhh",
			TotalAreaOfOwnLandColumn:       "c19",
			TotalFarmHoldingAreaColumn:     "c19smlf",
		},
		{
			Name:                           "Land under salt cultivation",
			NumberOfReportingHoldingColumn: "c20gtrhh",
			NumberOfFarmHoldingsColumn:     "c20smlfhh",
			TotalAreaOfOwnLandColumn:       "c20",
			TotalFarmHoldingAreaColumn:     "c20smlf",
		},
	}

	for i, c := range data {
		query := fmt.Sprintf(`
		SELECT (
			SELECT sum(%s)
			FROM aggregates
			WHERE subpath(geocode, 0, %d) = ?
		) AS number_of_reporting_holdings,
		(
			SELECT sum(%s)
			FROM aggregates
			WHERE subpath(geocode, 0, %d) = ?
		) AS number_of_farm_holdings,
		(
			SELECT sum(%s)
			FROM aggregates
			WHERE subpath(geocode, 0, %d) = ?
		) AS total_area_of_own_land,
		(
			SELECT sum(%s)
			FROM aggregates
			WHERE subpath(geocode, 0, %d) = ?
		) AS total_farm_holding_area;`,
			c.NumberOfReportingHoldingColumn, count,
			c.NumberOfFarmHoldingsColumn, count,
			c.TotalAreaOfOwnLandColumn, count,
			c.TotalFarmHoldingAreaColumn, count)
		_, err = db.Conn.QueryOne(&c, query,
			geoCodeReq, geoCodeReq, geoCodeReq, geoCodeReq)
		if err != nil {
			log.Error(err)
			return
		}
		data[i] = c
	}
	return
}

type HouseholdFisheryLandInformation struct {
	Name                            string
	NumberOfReportingHoldingsColumn string
	NumberOfReportingHoldings       uint
	NumberOfFarmHoldingsColumn      string
	NumberOfFarmHoldings            uint
	TotalAreaOfOwnLandColumn        string
	TotalAreaOfOwnLand              float64
	TotalFarmHoldingAreaColumn      string
	TotalFarmHoldingArea            float64
	AverageAreaPerFarmHolding       float64
}

func (db *Db) GetHouseholdFisheryLandInformation(division, district, upazilla, union, mouza string) (data []HouseholdFisheryLandInformation, err error) {
	geoCodeReq, count, err := GetGeoRequest(division, district, upazilla, union, mouza)
	if err != nil {
		return
	}
	data = []HouseholdFisheryLandInformation{
		{
			Name:                            "Land under ponds/digi",
			NumberOfReportingHoldingsColumn: "c21gtrhh",
			NumberOfFarmHoldingsColumn:      "c21smlfhh",
			TotalAreaOfOwnLandColumn:        "c21",
			TotalFarmHoldingAreaColumn:      "c21smlf",
		},
		{
			Name:                            "Fishery Land other than ponds",
			NumberOfReportingHoldingsColumn: "c22gtrhh+c23gtrhh+c24gtrhh",
			NumberOfFarmHoldingsColumn:      "c22smlfhh+c23smlfhh+c24smlfhh",
			TotalAreaOfOwnLandColumn:        "c22+c23+c24",
			TotalFarmHoldingAreaColumn:      "c22smlf+c23smlf+c24smlf",
		},

		{
			Name:                            "Fishery Land under salt cultivation",
			NumberOfReportingHoldingsColumn: "c25gtrhh",
			NumberOfFarmHoldingsColumn:      "c25smlfhh",
			TotalAreaOfOwnLandColumn:        "c25",
			TotalFarmHoldingAreaColumn:      "c25smlf",
		},

		{
			Name:                            "Fishery Land cultivated under pan/cage",
			NumberOfReportingHoldingsColumn: "c26gtrhh",
			NumberOfFarmHoldingsColumn:      "c26smlfhh",
			TotalAreaOfOwnLandColumn:        "c26",
			TotalFarmHoldingAreaColumn:      "c26smlf",
		},

		{
			Name:                            "Fishery Land under fish cultivation by Creek",
			NumberOfReportingHoldingsColumn: "c27gtrhh",
			NumberOfFarmHoldingsColumn:      "c27smlfhh",
			TotalAreaOfOwnLandColumn:        "c27",
			TotalFarmHoldingAreaColumn:      "c27smlf",
		},
	}

	for i, c := range data {
		query := fmt.Sprintf(`
		SELECT (
			SELECT sum(%s)
			FROM aggregates
			WHERE subpath(geocode, 0, %d) = ?
		) AS number_of_reporting_holdings,
		(
			SELECT sum(%s)
			FROM aggregates
			WHERE subpath(geocode, 0, %d) = ?
		) AS number_of_farm_holdings,
		(
			SELECT sum(%s)
			FROM aggregates
			WHERE subpath(geocode, 0, %d) = ?
		) AS total_area_of_own_land,
		(
			SELECT sum(%s)
			FROM aggregates
			WHERE subpath(geocode, 0, %d) = ?
		) AS total_farm_holding_area`,
			c.NumberOfReportingHoldingsColumn, count,
			c.NumberOfFarmHoldingsColumn, count,
			c.TotalAreaOfOwnLandColumn, count,
			c.TotalFarmHoldingAreaColumn, count)
		_, err = db.Conn.QueryOne(&c, query,
			geoCodeReq, geoCodeReq, geoCodeReq, geoCodeReq)
		if err != nil {
			log.Error(err)
			return
		}
		data[i] = c
	}
	return
}

type HouseholdPoultryInformation struct {
	Name                                     string
	NumberOfHousehold                        string
	Column                                   string
	NumberOfHouseholdPoultryColumn           string
	NumberOfHouseholdAttachFarmPoultryColumn string
	NumberOfReportingHoldings                uint
	TotalNumberOfPoultry                     uint
	AverageTypeOfPoultryPerHolding           float64
}

func (db *Db) GetHouseholdPoultryInformation(division, district, upazilla, union, mouza string) (data []HouseholdPoultryInformation, err error) {
	geoCodeReq, count, err := GetGeoRequest(division, district, upazilla, union, mouza)
	if err != nil {
		return
	}
	data = []HouseholdPoultryInformation{
		{
			Name:                                     "Cock/Hen",
			Column:                                   "c28gtrhh",
			NumberOfHouseholdPoultryColumn:           "c28h + c28f",
			NumberOfHouseholdAttachFarmPoultryColumn: "sum(c28h + c28f)/sum(hh_sno)",
		},
		{
			Name:                                     "Duck",
			Column:                                   "c29gtrhh",
			NumberOfHouseholdPoultryColumn:           "c29h + c29f",
			NumberOfHouseholdAttachFarmPoultryColumn: "sum(c29h + c29f) / sum(hh_sno)",
		},
		{
			Name:                                     "Pigeon",
			Column:                                   "c30gtrhh",
			NumberOfHouseholdPoultryColumn:           "c30h + c30f",
			NumberOfHouseholdAttachFarmPoultryColumn: "sum(c30h + c30f) / sum(hh_sno)",
		},
		{
			Name:                                     "Quail",
			Column:                                   "c31gtrhh",
			NumberOfHouseholdPoultryColumn:           "c31h + c31f",
			NumberOfHouseholdAttachFarmPoultryColumn: "sum(c31h + c31f) / sum(hh_sno)",
		},
		{
			Name:                                     "Turkey",
			Column:                                   "c32gtrhh",
			NumberOfHouseholdPoultryColumn:           "c32h + c32f",
			NumberOfHouseholdAttachFarmPoultryColumn: "sum(c32h + c32f) / sum(hh_sno)",
		},
	}

	for i, c := range data {
		query := fmt.Sprintf(`
		SELECT (
			SELECT sum(%s)
			FROM aggregates
			WHERE subpath(geocode, 0, %d) = ?
		) AS number_of_reporting_holdings,
		(
			SELECT sum(%s)
			FROM aggregates
			WHERE subpath(geocode, 0, %d) = ?
		) AS total_number_of_poultry,
		(
			SELECT %s::FLOAT
			FROM aggregates
			WHERE subpath(geocode, 0, %d) = ?
		) AS average_type_of_poultry_per_holding;`,
			c.Column, count,
			c.NumberOfHouseholdPoultryColumn, count,
			c.NumberOfHouseholdAttachFarmPoultryColumn, count)
		_, err = db.Conn.QueryOne(&c, query,
			geoCodeReq, geoCodeReq, geoCodeReq, geoCodeReq)
		if err != nil {
			log.Error(err)
			return
		}
		data[i] = c
	}
	return
}

type HouseholdCattleInformation struct {
	Name                                    string
	Column                                  string
	NumberOfHouseholdCattleColumn           string
	NumberOfHouseholdAttachFarmCattleColumn string
	NumberOfReportingHoldings               uint
	TotalNumberOfCattle                     uint
	AverageTypeOfCattlePerHolding           float64
}

func (db *Db) GetHouseholdCattlenformation(division, district, upazilla, union, mouza string) (data []HouseholdCattleInformation, err error) {
	geoCodeReq, count, err := GetGeoRequest(division, district, upazilla, union, mouza)
	if err != nil {
		return
	}
	data = []HouseholdCattleInformation{
		{
			Name:                                    "Cow",
			Column:                                  "c33gtrhh",
			NumberOfHouseholdCattleColumn:           "c33h + c33f",
			NumberOfHouseholdAttachFarmCattleColumn: "sum(c33h + c33f) / sum(hh_sno)",
		},
		{
			Name:                                    "Buffalo",
			Column:                                  "c34gtrhh",
			NumberOfHouseholdCattleColumn:           "c34h + c34f",
			NumberOfHouseholdAttachFarmCattleColumn: "sum(c34h + c34f) / sum(hh_sno)",
		},
		{
			Name:                                    "Goat",
			Column:                                  "c35gtrhh",
			NumberOfHouseholdCattleColumn:           "c35h + c35f",
			NumberOfHouseholdAttachFarmCattleColumn: "sum(c35h + c35f) / sum(hh_sno)",
		},
		{
			Name:                                    "Sheep",
			Column:                                  "c36gtrhh",
			NumberOfHouseholdCattleColumn:           "c36h + c36f",
			NumberOfHouseholdAttachFarmCattleColumn: "sum(c36h + c36f) / sum(hh_sno)",
		},
		{
			Name:                                    "Pig",
			Column:                                  "c37gtrhh",
			NumberOfHouseholdCattleColumn:           "c37h + c37f",
			NumberOfHouseholdAttachFarmCattleColumn: "sum(c37h + c37f) / sum(hh_sno)",
		},
		{
			Name:                                    "Horse",
			Column:                                  "c38gtrhh",
			NumberOfHouseholdCattleColumn:           "c38h + c38f",
			NumberOfHouseholdAttachFarmCattleColumn: "sum(c38h + c38f) / sum(hh_sno)",
		},
	}

	for i, c := range data {
		query := fmt.Sprintf(`
		SELECT (
			SELECT sum(%s)
			FROM aggregates
			WHERE subpath(geocode, 0, %d) = ?
		) AS number_of_reporting_holdings,
		(
			SELECT sum(%s)
			FROM aggregates
			WHERE subpath(geocode, 0, %d) = ?
		) AS total_number_of_cattle,
		(
			SELECT %s::FLOAT
			FROM aggregates
			WHERE subpath(geocode, 0, %d) = ?
		) AS average_type_of_cattle_per_holding;`,
			c.Column, count,
			c.NumberOfHouseholdCattleColumn, count,
			c.NumberOfHouseholdAttachFarmCattleColumn, count)
		_, err = db.Conn.QueryOne(&c, query,
			geoCodeReq, geoCodeReq, geoCodeReq, geoCodeReq)
		if err != nil {
			log.Error(err)
			return
		}
		data[i] = c
	}
	return
}

type HouseholdAgricultureEquipement struct {
	Name                              string
	NumberOfReportingHoldingsColumn   string
	NumberOfReportingHoldings         uint
	TotalNumberColumn                 string
	TotalNumber                       uint
	NumberOfNonMechanicalDeviceColumn string
	NumberOfNonMechanicalDevice       uint
	NumberOfDieselDeviceColumn        string
	NumberOfDieselDevice              uint
	NumberOfElectricalDeviceColumn    string
	NumberOfElectricalDevice          uint
}

func (db *Db) GetHouseholdAgricultureEquipement(division, district, upazilla, union, mouza string) (data []HouseholdAgricultureEquipement, err error) {
	geoCodeReq, count, err := GetGeoRequest(division, district, upazilla, union, mouza)
	if err != nil {
		return
	}
	data = []HouseholdAgricultureEquipement{
		{
			Name:                            "Tractor",
			NumberOfReportingHoldingsColumn: "c39gtrhh",
			TotalNumberColumn:               "c39",
			NumberOfDieselDeviceColumn:      "c39",
		},
		{
			Name:                            "Power tiller",
			NumberOfReportingHoldingsColumn: "c40gtrhh",
			TotalNumberColumn:               "c40",
			NumberOfDieselDeviceColumn:      "c40",
		},
		{
			Name:                            "Power pump",
			NumberOfReportingHoldingsColumn: "c41gtrhh",
			TotalNumberColumn:               "(c41a + c41b)",
			NumberOfDieselDeviceColumn:      "c41a",
			NumberOfElectricalDeviceColumn:  "c41b",
		},
		{
			Name:                            "Deep/Shallow tube well",
			NumberOfReportingHoldingsColumn: "c42gtrhh",
			TotalNumberColumn:               "(c42a + c42b)",
			NumberOfDieselDeviceColumn:      "c42a",
			NumberOfElectricalDeviceColumn:  "c42b",
		},
		{
			Name:                              "Crop planting machine",
			NumberOfReportingHoldingsColumn:   "c43gtrhh",
			TotalNumberColumn:                 "(c43a + c43b)",
			NumberOfNonMechanicalDeviceColumn: "c43a",
			NumberOfDieselDeviceColumn:        "c43b",
		},
		{
			Name:                              "Crop cutting machine",
			NumberOfReportingHoldingsColumn:   "c44gtrhh",
			TotalNumberColumn:                 "(c44a + c44b)",
			NumberOfNonMechanicalDeviceColumn: "c44a",
			NumberOfDieselDeviceColumn:        "c44b",
		},
		{
			Name:                              "Crop threshing machine",
			NumberOfReportingHoldingsColumn:   "c45gtrhh",
			TotalNumberColumn:                 "(c45a + c45b + c45c)",
			NumberOfNonMechanicalDeviceColumn: "c45a",
			NumberOfDieselDeviceColumn:        "c45b",
			NumberOfElectricalDeviceColumn:    "c45c",
		},
		{
			Name:                              "Fertilizer Appling machine",
			NumberOfReportingHoldingsColumn:   "c46gtrhh",
			TotalNumberColumn:                 "(c46a + c46b)",
			NumberOfNonMechanicalDeviceColumn: "c46a",
			NumberOfDieselDeviceColumn:        "c46b",
		},
		{
			Name:                              "Fish catching boat/trailer",
			NumberOfReportingHoldingsColumn:   "c47gtrhh",
			TotalNumberColumn:                 "(c47a + c47b)",
			NumberOfNonMechanicalDeviceColumn: "c47a",
			NumberOfDieselDeviceColumn:        "c47b",
		},
		{
			Name:                              "Fish catching net (business)",
			NumberOfReportingHoldingsColumn:   "c48gtrhh",
			TotalNumberColumn:                 "c48",
			NumberOfNonMechanicalDeviceColumn: "c48",
		},
		{
			Name:                              "Plough",
			NumberOfReportingHoldingsColumn:   "c49gtrhh",
			TotalNumberColumn:                 "c49",
			NumberOfNonMechanicalDeviceColumn: "c49",
		},
	}

	for i, c := range data {
		query := fmt.Sprintf(`
		SELECT (
			SELECT sum(%s)
			FROM aggregates
			WHERE subpath(geocode, 0, %d) = ?
		) AS number_of_reporting_holdings,
		(
			SELECT sum(%s)
			FROM aggregates
			WHERE subpath(geocode, 0, %d) = ?
		) AS total_number`,
			c.NumberOfReportingHoldingsColumn, count,
			c.TotalNumberColumn, count)

		if c.NumberOfNonMechanicalDeviceColumn != "" {
			query += fmt.Sprintf(`
			,(
				SELECT sum(%s)
				FROM aggregates
				WHERE subpath(geocode, 0, %d) = ?
			) AS number_of_non_mechanical_device
			`, c.NumberOfNonMechanicalDeviceColumn, count)
		}
		if c.NumberOfDieselDeviceColumn != "" {
			query += fmt.Sprintf(`
			,(
				SELECT sum(%s)
				FROM aggregates
				WHERE subpath(geocode, 0, %d) = ?
			) AS number_of_diesel_device
			`, c.NumberOfDieselDeviceColumn, count)
		}
		if c.NumberOfElectricalDeviceColumn != "" {
			query += fmt.Sprintf(`
			,(
				SELECT sum(%s)
				FROM aggregates
				WHERE subpath(geocode, 0, %d) = ?
			) AS number_of_electrical_device
			`, c.NumberOfElectricalDeviceColumn, count)
		}
		_, err = db.Conn.QueryOne(&c, query,
			geoCodeReq, geoCodeReq, geoCodeReq, geoCodeReq, geoCodeReq)
		if err != nil {
			log.Error(err)
			return
		}
		data[i] = c
	}
	return
}

type TemporaryCrops struct {
	NumberOfFarmHoldings uint
	CropArea             float64
	T101                 float64
	T102                 float64
	T103                 float64
	T104                 float64
	T105                 float64
	T106                 float64
	T107                 float64
	T108                 float64
	T109                 float64
	T110                 float64
	T111                 float64
	T112                 float64
	T113                 float64
	T114                 float64
	T115                 float64
	T116                 float64
	T117                 float64
	T118                 float64
	T119                 float64
	T120                 float64
	T121                 float64
	T122                 float64
	T123                 float64
	T124                 float64
	T125                 float64
	T126                 float64
	T127                 float64
	T128                 float64
	T129                 float64
	T130                 float64
	T131                 float64
	T132                 float64
	T133                 float64
	T134                 float64
	T135                 float64
	T136                 float64
	T137                 float64
	T138                 float64
	T139                 float64
	T140                 float64
	T141                 float64
	T142                 float64
	T143                 float64
	T144                 float64
	T145                 float64
	T146                 float64
	T147                 float64
	T148                 float64
	T149                 float64
	T150                 float64
	T151                 float64
	T152                 float64
	T153                 float64
	T154                 float64
	T155                 float64
	T156                 float64
	T157                 float64
	T158                 float64
	T159                 float64
	T160                 float64
	T161                 float64
	T162                 float64
	T163                 float64
	T164                 float64
	T165                 float64
	T166                 float64
	T167                 float64
	T168                 float64
	T169                 float64
	T170                 float64
	T171                 float64
	T172                 float64
	T173                 float64
	T174                 float64
	T175                 float64
	T176                 float64
	T177                 float64
	T178                 float64
	T179                 float64
	T180                 float64
	T181                 float64
	T182                 float64
	T183                 float64
	T184                 float64
	T185                 float64
	T186                 float64
	T187                 float64
	T188                 float64
	T189                 float64
	T190                 float64
	T191                 float64
	T192                 float64
	T193                 float64
	T194                 float64
	T195                 float64
	T196                 float64
	T197                 float64
	T198                 float64
	T199                 float64
	T200                 float64
	T201                 float64
	T202                 float64
	T203                 float64
}

func (c TemporaryCrops) PercentageOfCropArea(cropArea string) string {
	p := message.NewPrinter(language.English)
	r, err := reflections.GetField(c, cropArea)
	if err != nil {
		log.Errorf("unable to get %s field from Crops struct : %s", cropArea, err)
		return "err"
	}
	return p.Sprintf("%.2f%%", (r.(float64)/
		(c.T101+
			c.T102+
			c.T103+
			c.T104+
			c.T105+
			c.T106+
			c.T107+
			c.T108+
			c.T109+
			c.T110+
			c.T111+
			c.T112+
			c.T113+
			c.T114+
			c.T115+
			c.T116+
			c.T117+
			c.T118+
			c.T119+
			c.T120+
			c.T121+
			c.T122+
			c.T123+
			c.T124+
			c.T125+
			c.T126+
			c.T127+
			c.T128+
			c.T129+
			c.T130+
			c.T131+
			c.T132+
			c.T133+
			c.T134+
			c.T135+
			c.T136+
			c.T137+
			c.T138+
			c.T139+
			c.T140+
			c.T141+
			c.T142+
			c.T143+
			c.T144+
			c.T145+
			c.T146+
			c.T147+
			c.T148+
			c.T149+
			c.T150+
			c.T151+
			c.T152+
			c.T153+
			c.T154+
			c.T155+
			c.T156+
			c.T157+
			c.T158+
			c.T159+
			c.T160+
			c.T161+
			c.T162+
			c.T163+
			c.T164+
			c.T165+
			c.T166+
			c.T167+
			c.T168+
			c.T169+
			c.T170+
			c.T171+
			c.T172+
			c.T173+
			c.T174+
			c.T175+
			c.T176+
			c.T177+
			c.T178+
			c.T179+
			c.T180+
			c.T181+
			c.T182+
			c.T183+
			c.T184+
			c.T185+
			c.T186+
			c.T187+
			c.T188+
			c.T189+
			c.T190+
			c.T191+
			c.T192+
			c.T193+
			c.T194+
			c.T195+
			c.T196+
			c.T197+
			c.T198+
			c.T199+
			c.T200+
			c.T201+
			c.T202+
			c.T203))*100)
}

func (db *Db) GetTemporaryCrops(division, district, upazilla, union, mouza string) (data TemporaryCrops, err error) {
	geoCodeReq, count, err := GetGeoRequest(division, district, upazilla, union, mouza)
	if err != nil {
		return
	}

	query := fmt.Sprintf(`
	SELECT 
    	sum(sf + mf + lf) as number_of_farm_holdings,
    	sum(c13) as crop_area,
		sum(t101) as t101,
		sum(t102) as t102,
		sum(t103) as t103,
		sum(t104) as t104,
		sum(t105) as t105,
		sum(t106) as t106,
		sum(t107) as t107,
		sum(t108) as t108,
		sum(t109) as t109,
		sum(t110) as t110,
		sum(t111) as t111,
		sum(t112) as t112,
		sum(t113) as t113,
		sum(t114) as t114,
		sum(t115) as t115,
		sum(t116) as t116,
		sum(t117) as t117,
		sum(t118) as t118,
		sum(t119) as t119,
		sum(t120) as t120,
		sum(t121) as t121,
		sum(t122) as t122,
		sum(t123) as t123,
		sum(t124) as t124,
		sum(t125) as t125,
		sum(t126) as t126,
		sum(t127) as t127,
		sum(t128) as t128,
		sum(t129) as t129,
		sum(t130) as t130,
		sum(t131) as t131,
		sum(t132) as t132,
		sum(t133) as t133,
		sum(t134) as t134,
		sum(t135) as t135,
		sum(t136) as t136,
		sum(t137) as t137,
		sum(t138) as t138,
		sum(t139) as t139,
		sum(t140) as t140,
		sum(t141) as t141,
		sum(t142) as t142,
		sum(t143) as t143,
		sum(t144) as t144,
		sum(t145) as t145,
		sum(t146) as t146,
		sum(t147) as t147,
		sum(t148) as t148,
		sum(t149) as t149,
		sum(t150) as t150,
		sum(t151) as t151,
		sum(t152) as t152,
		sum(t153) as t153,
		sum(t154) as t154,
		sum(t155) as t155,
		sum(t156) as t156,
		sum(t157) as t157,
		sum(t158) as t158,
		sum(t159) as t159,
		sum(t160) as t160,
		sum(t161) as t161,
		sum(t162) as t162,
		sum(t163) as t163,
		sum(t164) as t164,
		sum(t165) as t165,
		sum(t166) as t166,
		sum(t167) as t167,
		sum(t168) as t168,
		sum(t169) as t169,
		sum(t170) as t170,
		sum(t171) as t171,
		sum(t172) as t172,
		sum(t173) as t173,
		sum(t174) as t174,
		sum(t175) as t175,
		sum(t176) as t176,
		sum(t177) as t177,
		sum(t178) as t178,
		sum(t179) as t179,
		sum(t180) as t180,
		sum(t181) as t181,
		sum(t182) as t182,
		sum(t183) as t183,
		sum(t184) as t184,
		sum(t185) as t185,
		sum(t186) as t186,
		sum(t187) as t187,
		sum(t188) as t188,
		sum(t189) as t189,
		sum(t190) as t190,
		sum(t191) as t191,
		sum(t192) as t192,
		sum(t193) as t193,
		sum(t194) as t194,
		sum(t195) as t195,
		sum(t196) as t196,
		sum(t197) as t197,
		sum(t198) as t198,
		sum(t199) as t199,
		sum(t200) as t200,
		sum(t201) as t201,
		sum(t202) as t202,
		sum(t203) as t203
	FROM aggregates
	WHERE subpath(geocode, 0, %d) = ?;
	`, count)

	_, err = db.Conn.QueryOne(&data, query,
		geoCodeReq)
	if err != nil {
		log.Error(err)
		return
	}
	return

}

type PermanantCrops struct {
	NumberOfFarmHoldings uint
	CropArea             float64
	P501a                float64
	P502a                float64
	P503a                float64
	P504a                float64
	P505a                float64
	P506a                float64
	P507a                float64
	P508a                float64
	P509a                float64
	P510a                float64
	P511a                float64
	P512a                float64
	P513a                float64
	P514a                float64
	P515a                float64
	P516a                float64
	P517a                float64
	P518a                float64
	P519a                float64
	P520a                float64
	P521a                float64
	P522a                float64
	P523a                float64
	P524a                float64
	P525a                float64
	P526a                float64
	P527a                float64
	P528a                float64
	P529a                float64
	P530a                float64
	P531a                float64
	P532a                float64
	P533a                float64
	P534a                float64
	P535a                float64
	P536a                float64
	P537a                float64
	P538a                float64
	P539a                float64
	P540a                float64
	P541a                float64
	P542a                float64
	P543a                float64
	P544a                float64
	P545a                float64
	P546a                float64
	P547a                float64
	P548a                float64
	P549a                float64
	P550a                float64
	P551a                float64
	P552a                float64
	P553a                float64
	P554a                float64
	P555a                float64
	P556a                float64
	P557a                float64
	P559a                float64
	P560a                float64
	P561a                float64
	P562a                float64
	P563a                float64
	P564a                float64
	P565a                float64
	P566a                float64
	P567a                float64
	P568a                float64
	P569a                float64
	P570a                float64
	P571a                float64
	P572a                float64
	P573a                float64
	P574a                float64
	P575a                float64
	P577a                float64
	P579a                float64
	P580a                float64
	P581a                float64
	P582a                float64
	P583a                float64
	P584a                float64
	P585a                float64
}

func (c PermanantCrops) PercentageOfPermantCropArea(cropArea string) string {
	p := message.NewPrinter(language.English)
	r, err := reflections.GetField(c, cropArea)
	if err != nil {
		log.Errorf("unable to get %s field from Crops struct : %s", cropArea, err)
		return "err"
	}
	return p.Sprintf("%.2f%%", (r.(float64)/
		(c.P501a+
			c.P502a+
			c.P503a+
			c.P504a+
			c.P505a+
			c.P506a+
			c.P507a+
			c.P508a+
			c.P509a+
			c.P510a+
			c.P511a+
			c.P512a+
			c.P513a+
			c.P514a+
			c.P515a+
			c.P516a+
			c.P517a+
			c.P518a+
			c.P519a+
			c.P520a+
			c.P521a+
			c.P522a+
			c.P523a+
			c.P524a+
			c.P525a+
			c.P526a+
			c.P527a+
			c.P528a+
			c.P529a+
			c.P530a+
			c.P531a+
			c.P532a+
			c.P533a+
			c.P534a+
			c.P535a+
			c.P536a+
			c.P537a+
			c.P538a+
			c.P539a+
			c.P540a+
			c.P541a+
			c.P542a+
			c.P543a+
			c.P544a+
			c.P545a+
			c.P546a+
			c.P547a+
			c.P548a+
			c.P549a+
			c.P550a+
			c.P551a+
			c.P552a+
			c.P553a+
			c.P554a+
			c.P555a+
			c.P556a+
			c.P557a+
			c.P559a+
			c.P560a+
			c.P561a+
			c.P562a+
			c.P563a+
			c.P564a+
			c.P565a+
			c.P566a+
			c.P567a+
			c.P568a+
			c.P569a+
			c.P570a+
			c.P571a+
			c.P572a+
			c.P573a+
			c.P574a+
			c.P575a+
			c.P577a+
			c.P579a+
			c.P580a+
			c.P581a+
			c.P582a+
			c.P584a+
			c.P585a))*100)
}

func (db *Db) GetPermanantCrops(division, district, upazilla, union, mouza string) (data PermanantCrops, err error) {
	geoCodeReq, count, err := GetGeoRequest(division, district, upazilla, union, mouza)
	if err != nil {
		return
	}

	query := fmt.Sprintf(`
	SELECT 
    	sum(sf + mf + lf) as number_of_farm_holdings,
    	sum(c14) as crop_area,
		sum(p501a) as p501a,
		sum(p502a) as p502a,
		sum(p503a) as p503a,
		sum(p504a) as p504a,
		sum(p505a) as p505a,
		sum(p506a) as p506a,
		sum(p507a) as p507a,
		sum(p508a) as p508a,
		sum(p509a) as p509a,
		sum(p510a) as p510a,
		sum(p511a) as p511a,
		sum(p512a) as p512a,
		sum(p513a) as p513a,
		sum(p514a) as p514a,
		sum(p515a) as p515a,
		sum(p516a) as p516a,
		sum(p517a) as p517a,
		sum(p518a) as p518a,
		sum(p519a) as p519a,
		sum(p520a) as p520a,
		sum(p521a) as p521a,
		sum(p522a) as p522a,
		sum(p523a) as p523a,
		sum(p524a) as p524a,
		sum(p525a) as p525a,
		sum(p526a) as p526a,
		sum(p527a) as p527a,
		sum(p528a) as p528a,
		sum(p529a) as p529a,
		sum(p530a) as p530a,
		sum(p531a) as p531a,
		sum(p532a) as p532a,
		sum(p533a) as p533a,
		sum(p534a) as p534a,
		sum(p535a) as p535a,
		sum(p536a) as p536a,
		sum(p537a) as p537a,
		sum(p538a) as p538a,
		sum(p539a) as p539a,
		sum(p540a) as p540a,
		sum(p541a) as p541a,
		sum(p542a) as p542a,
		sum(p543a) as p543a,
		sum(p544a) as p544a,
		sum(p545a) as p545a,
		sum(p546a) as p546a,
		sum(p547a) as p547a,
		sum(p548a) as p548a,
		sum(p549a) as p549a,
		sum(p550a) as p550a,
		sum(p551a) as p551a,
		sum(p552a) as p552a,
		sum(p553a) as p553a,
		sum(p554a) as p554a,
		sum(p555a) as p555a,
		sum(p556a) as p556a,
		sum(p557a) as p557a,
		sum(p559a) as p559a,
		sum(p560a) as p560a,
		sum(p561a) as p561a,
		sum(p562a) as p562a,
		sum(p563a) as p563a,
		sum(p564a) as p564a,
		sum(p565a) as p565a,
		sum(p566a) as p566a,
		sum(p567a) as p567a,
		sum(p568a) as p568a,
		sum(p569a) as p569a,
		sum(p570a) as p570a,
		sum(p571a) as p571a,
		sum(p572a) as p572a,
		sum(p573a) as p573a,
		sum(p574a) as p574a,
		sum(p575a) as p575a,
		sum(p577a) as p577a,
		sum(p579a) as p579a,
		sum(p580a) as p580a,
		sum(p581a) as p581a,
		sum(p582a) as p582a,
		sum(p584a) as p584a,
		sum(p585a) as p585a
	FROM aggregates
	WHERE subpath(geocode, 0, %d) = ?;
	`, count)

	_, err = db.Conn.QueryOne(&data, query,
		geoCodeReq)
	if err != nil {
		log.Error(err)
		return
	}
	return

}
