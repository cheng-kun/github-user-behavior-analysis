package db

import (
	"database/sql"
	"fmt"
	"github.com/github-user-behavior-analysis/backend/conf"
	"github.com/github-user-behavior-analysis/backend/logs"
	"github.com/github-user-behavior-analysis/backend/models"
	_ "github.com/lib/pq"
	"strings"
	"time"
)

// Database connection to the github postgres ConnDatabase
type Database struct {
	*sql.DB
}

var ConnDB *Database

func init()  {

	logs.PrintLogger().Info("initilizing ConnDatabase connection ... ")

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

	logs.PrintLogger().Info("Setup ConnDatabase connection sucessfully! ")
}

// Connect to the ConnDatabase
func Connect(cfg conf.Config) (*Database, error) {
	connStr := fmt.Sprintf("user='%s' password='%s' dbname='%s' host='%s' sslmode=disable",
		cfg.Database.Username, cfg.Database.Password, cfg.Database.DBName, cfg.Database.Host)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	return &Database{DB: db}, nil
}

func (conn *Database) SaveTopUsers(userfollower *models.UserFollower) error {
	sql := `INSERT INTO top_users( login_user, followers, rank) VALUES($1,$2,$3)`

	_, err := conn.Exec(sql, userfollower.User, userfollower.Follower, userfollower.Rank)
	if err == nil {
		logs.PrintLogger().Infof("Successfully insert user_login name %s", userfollower.User)
	}

	return err
}

func (conn *Database) GetTopUsers(amount string) ([]*models.UserFollower, error) {
	sql1 := `SELECT rank, login_user, followers FROM top_users Order by rank limit $1`

	rows, err := conn.Query(sql1, amount)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	topUsers := make([]*models.UserFollower, 0)

	var loginN sql.NullString
	var followersN, rankN sql.NullInt64

	for rows.Next() {

		err = rows.Scan(&rankN, &loginN, &followersN)

		topUser := &models.UserFollower{loginN.String, followersN.Int64, rankN.Int64}

		topUsers = append(topUsers, topUser)
	}

	return topUsers, err


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

func (conn *Database) SaveProjectLanguage(timeStamp string, language string, amount string) error {
	sqlQuery := `INSERT INTO project_language (time_stamp, language, amount) VALUES ($1, $2, $3)`

	_, err := conn.Exec(sqlQuery, timeStamp, language, amount)
	if err != nil {
		logs.PrintLogger().Error(err)
		return err
	}

	return err
}

func (conn *Database) GetAllDailyRanking() ([]*models.Ranking, error)  {
	sqlQuery := `select * from top_ten `
	rows, err := conn.Query(sqlQuery)
	if err != nil {
		return nil, err
	}

	rankings := make([]*models.Ranking, 0)

	for rows.Next() {
		ranking := &models.Ranking{}
		var repoN, n1numN, n2numN, n3numN, n4numN, n5numN, n6numN, n7numN, n8numN, n9numN, n10numN sql.NullInt64
		var n1langN, n2langN, n3langN, n4langN, n5langN, n6langN, n7langN, n8langN, n9langN, n10langN sql.NullString
		var timeStamp time.Time

		rows.Scan(&repoN,&timeStamp,&n1langN,&n1numN,&n2langN,&n2numN,&n3langN,&n3numN,&n4langN,&n4numN,&n5langN,&n5numN,&n6langN,&n6numN,&n7langN,&n7numN,&n8langN,&n8numN,&n9langN, &n9numN,&n10langN,&n1numN)

			ranking.RepoNum =repoN.Int64
			ranking.TimeStamp = timeStamp.String()
			ranking.N1lang = n1langN.String
			ranking.N1num = n1numN.Int64
			ranking.N2lang = n2langN.String
			ranking.N2num = n2numN.Int64
			ranking.N3lang = n3langN.String
			ranking.N3num = n3numN.Int64
			ranking.N4lang = n4langN.String
			ranking.N4num = n4numN.Int64
			ranking.N5lang = n5langN.String
			ranking.N5num = n5numN.Int64
			ranking.N6lang = n6langN.String
			ranking.N6num = n6numN.Int64
			ranking.N7lang = n7langN.String
			ranking.N7num = n7numN.Int64
			ranking.N8lang = n8langN.String
			ranking.N8num = n8numN.Int64
			ranking.N9lang = n9langN.String
			ranking.N9num = n9numN.Int64
			ranking.N10lang = n10langN.String
			ranking.N10num = n10numN.Int64

		rankings = append(rankings, ranking)
	}

	return rankings, err

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

func (conn *Database) SaveCountryRepos(country, repos string) error  {

	selectQ := `select * from user_country where country = $1`
	rows, err := conn.Query(selectQ, country)
	if err != nil {
		return err
	}
	defer rows.Close()

	if rows.Next() {
		updateQ := `update user_country SET repo_amount = $2 WHERE country = $1`

		_, err = conn.Exec(updateQ, country, repos)
		if err != nil {
			logs.PrintLogger().Info(err)
			return err
		}

	} else {
		insertQ := `INSERT INTO user_country (country, repo_amount) VALUES($1, $2)`

		_, err = conn.Exec(insertQ, country, repos)
		if err != nil {
			logs.PrintLogger().Info(err)
			return err
		}
	}

	return err
}

func (conn *Database) SaveCountryUsers(country, user string) error  {

	selectQ := `select * from user_country where country = $1`
	rows, err := conn.Query(selectQ, country)
	if err != nil {
		return err
	}
	defer rows.Close()

	if rows.Next() {
		updateQ := `update user_country SET user_amount = $2 WHERE country = $1`

		_, err = conn.Exec(updateQ, country, user)
		if err != nil {
			logs.PrintLogger().Info(err)
			return err
		}

	} else {
		insertQ := `INSERT INTO user_country (country, user_amount) VALUES($1, $2)`

		_, err = conn.Exec(insertQ, country, user)
		if err != nil {
			logs.PrintLogger().Info(err)
			return err
		}
	}

	return err
}

func (conn *Database) SaveCountryPushs(country, push string) error  {

	selectQ := `select * from user_country where country = $1`
	rows, err := conn.Query(selectQ, country)
	if err != nil {
		return err
	}
	defer rows.Close()

	if rows.Next() {
		updateQ := `update user_country SET push_amount = $2 WHERE country = $1`

		_, err = conn.Exec(updateQ, country, push)
		if err != nil {
			logs.PrintLogger().Info(err)
			return err
		}

	} else {
		insertQ := `INSERT INTO user_country (country, push_amount) VALUES($1, $2)`

		_, err = conn.Exec(insertQ, country, push)
		if err != nil {
			logs.PrintLogger().Info(err)
			return err
		}
	}

	return err
}

func (conn *Database) GetCountryUser(amount string) ([]*models.UserCountry, error) {
	sql1 := `SELECT country, user_amount, repo_amount, push_amount FROM user_country Order by user_amount desc limit $1`

	rows, err := conn.Query(sql1, amount)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	topUsers := make([]*models.UserCountry, 0)

	var countryN sql.NullString
	var userN, repoN, pushN sql.NullInt64

	for rows.Next() {

		err = rows.Scan(&countryN, &userN, &repoN, &pushN)

		topUser := &models.UserCountry{countryN.String, userN.Int64, repoN.Int64, pushN.Int64}

		logs.PrintLogger().Info(topUser)

		topUsers = append(topUsers, topUser)
	}

	return topUsers, err
}

