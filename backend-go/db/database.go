package db

import (
	"database/sql"
	"fmt"
	"github.com/github-user-behavior-analysis/conf"
	"github.com/github-user-behavior-analysis/logs"
	"github.com/github-user-behavior-analysis/models"
	"strings"
	_ "github.com/lib/pq"
)

// Database connection to the github postgres database
type Database struct {
	*sql.DB
}

var ConnDB *Database

func init()  {

	logs.PrintLogger().Info("initilizing database connection ... ")

	cfg, err := conf.LoadConfigFile("./config.toml")
	if err != nil {
		logs.PrintLogger().Error(err)
		return
	}

	ConnDB, err = Connect(*cfg)
	if err != nil {
		logs.PrintLogger().Error(err)
		return
	}

	logs.PrintLogger().Info("Setup database connection sucessfully! ")
}

// Connect to the database
func Connect(cfg conf.Config) (*Database, error) {
	connStr := fmt.Sprintf("user='%s' password='%s' dbname='%s' host='%s' sslmode=disable",
		cfg.Database.Username, cfg.Database.Password, cfg.Database.DBName, cfg.Database.Host)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	return &Database{DB: db}, nil
}

func (conn *Database) SaveTopTenRanking(ranking *models.Ranking) error {
	sql := `INSERT INTO top_ten (repo_num, time_stamp, n1lang, n1num, n2lang, n2num, n3lang, n3num, n4lang, n4num, n5lang, n5num, n6lang, n6num, n7lang, n7num, n8lang, n8num, n9lang, n9num, n10lang, n10num)
			VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18,$19,$20,$21,$22)`



	_, err := conn.Exec(sql, ranking.RepoNum,
		ranking.TimeStamp,
		ranking.N1lang,
		ranking.N1num,
		ranking.N2lang,
		ranking.N2num,
		ranking.N3lang,
		ranking.N3num,
		ranking.N4lang,
		ranking.N4num,
		ranking.N5lang,
		ranking.N5num,
		ranking.N6lang,
		ranking.N6num,
		ranking.N7lang,
		ranking.N7num,
		ranking.N8lang,
		ranking.N8num,
		ranking.N9lang,
		ranking.N9num,
		ranking.N10lang,
		ranking.N10num)

	if err == nil {
		logs.PrintLogger().Infof("Successfully insert repo_num %d", ranking.RepoNum)
	}


	return err
}

func (conn *Database) GetRankInfoByLanguage(lan, dateTime string) (*models.LangaugeRank, error) {
	sqlQuery := `select n1num as amount, time_stamp, 1 as rank from top_ten where lower(n1lang) = $1 and time_stamp = $2
union
select n2num as amount, time_stamp, 2 as rank from top_ten where lower(n2lang) = $1 and time_stamp = $2
union
select n3num as amount, time_stamp, 3 as rank from top_ten where lower(n3lang) = $1 and time_stamp = $2
union
select n4num as amount, time_stamp, 4 as rank from top_ten where lower(n4lang) = $1 and time_stamp = $2
union
select n5num as amount, time_stamp, 5 as rank from top_ten where lower(n5lang) = $1 and time_stamp = $2
union
select n6num as amount, time_stamp, 6 as rank from top_ten where lower(n6lang) = $1 and time_stamp = $2
union
select n7num as amount, time_stamp, 7 as rank from top_ten where lower(n7lang) = $1 and time_stamp = $2
union
select n8num as amount, time_stamp, 8 as rank from top_ten where lower(n8lang) = $1 and time_stamp = $2
union
select n9num as amount, time_stamp, 9 as rank from top_ten where lower(n9lang) = $1 and time_stamp = $2
union
select n10num as amount, time_stamp, 10 as rank from top_ten where lower(n10lang) = $1 and time_stamp = $2
`

	lanLowerCase := strings.ToLower(strings.TrimSpace(lan))

	rows, err := conn.Query(sqlQuery, lanLowerCase, dateTime)
	if err != nil {
		logs.PrintLogger().Error(err)
		return nil, err
	}
	defer rows.Close()

	//languageRanks := make([]*models.LangaugeRank,0)

	languageRank := &models.LangaugeRank{}

	var amountN, rankN sql.NullInt64
	var timeStampN sql.NullString

	for rows.Next() {

		err = rows.Scan(&amountN, &timeStampN, &rankN )

		languageRank.Amount = amountN.Int64
		languageRank.TimeStamp = timeStampN.String
		languageRank.Rank = rankN.Int64

		//languageRanks = append(languageRanks, languageRank)
	}

	//return languageRanks, err
	return languageRank, err
}

func (conn *Database) GetDailyRankByDate(date string) (*models.Ranking, error)  {
	sqlQuery := `select * from top_ten where time_stamp = $1`

	rows, err := conn.Query(sqlQuery, date)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var timeStampN, n1langN, n2langN, n3langN, n4langN, n5langN, n6langN, n7langN, n8langN, n9langN, n10langN sql.NullString
	var repoNumN, n1numN, n2numN, n3numN, n4numN, n5numN, n6numN, n7numN, n8numN, n9numN, n10numN sql.NullInt64

	dailyRank := &models.Ranking{}

	if rows.Next() {
		err = rows.Scan(&repoNumN, &timeStampN, &n1langN, &n1numN, &n2langN, &n2numN, &n3langN, &n3numN, &n4langN, &n4numN ,&n5langN, &n5numN ,&n6langN, &n6numN ,&n7langN, &n7numN, &n8langN, &n8numN, &n9langN, &n9numN ,&n10langN, &n10numN )
		if err != nil {
			return nil, err
		}

		dailyRank.RepoNum = repoNumN.Int64
		dailyRank.TimeStamp = timeStampN.String
		dailyRank.N1lang = n1langN.String
		dailyRank.N1num = n1numN.Int64
		dailyRank.N2lang = n2langN.String
		dailyRank.N2num = n2numN.Int64
		dailyRank.N3lang = n3langN.String
		dailyRank.N3num = n3numN.Int64
		dailyRank.N4lang = n4langN.String
		dailyRank.N4num = n4numN.Int64
		dailyRank.N5lang = n5langN.String
		dailyRank.N5num = n5numN.Int64
		dailyRank.N6lang = n6langN.String
		dailyRank.N6num = n6numN.Int64
		dailyRank.N7lang = n7langN.String
		dailyRank.N7num = n7numN.Int64
		dailyRank.N8lang = n8langN.String
		dailyRank.N8num = n8numN.Int64
		dailyRank.N9lang = n9langN.String
		dailyRank.N9num = n9numN.Int64
		dailyRank.N10lang = n10langN.String
		dailyRank.N10num = n10numN.Int64

	}

	return dailyRank, err
}
