package db

import (
	"database/sql"
	"fmt"
	"github.com/github-user-behavior-analysis/logs"
	"github.com/github-user-behavior-analysis/models"
	"strings"
	"time"

	"github.com/github-user-behavior-analysis/conf"
	"github.com/google/go-github/github"
	"github.com/lib/pq"
)

// Database connection to the github postgres database
type Database struct {
	*sql.DB
}

var ConnDB *Database

func init()  {
	cfg, err := conf.LoadConfigFile("/home/nebula-ai-chengkun/gopath/src/github.com/github-user-behavior-analysis/conf/config.toml")
	if err != nil {
		logs.PrintLogger().Error(err)
		return
	}

	ConnDB, err = Connect(*cfg)
	if err != nil {
		logs.PrintLogger().Error(err)
		return
	}
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

// InsertUserStatus updates the statuscode/fetchtime associated with a user in the case that it
// can't be fetched
func (conn *Database) InsertUserStatus(id int64, login string, statuscode int, fetchtime time.Time) error {
	sql := `INSERT INTO users (id, login, fetched, statuscode) VALUES ($1, $2, $3, $4)
		ON CONFLICT(id) DO UPDATE SET fetched=$3, statuscode=$4`

	// TODO: don't use prepare?
	stmt, err := conn.Prepare(sql)
	if err != nil {
		fmt.Printf("Failed to prepare: %s\n", err.Error())
		return err
	}

	_, err = stmt.Exec(id, login, fetchtime, statuscode)
	if err != nil {
		fmt.Printf("Failed to exec: %s\n", err.Error())
		return err
	}
	return err
}

// InsertUser inserts a github.User object into the database
func (conn *Database) InsertUser(statuscode *int, fetchtime *time.Time, user *github.User, upsert bool) error {
	sql := `INSERT INTO users (id, login, name, company, location, bio, email, type, followers, following,
		   created, modified, fetched, statuscode, blog)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15) `
	if upsert {
		sql += `ON CONFLICT(id) DO UPDATE SET login=$2, name=$3, company=$4, location=$5, bio=$6, email=$7, type=$8,
				 followers=$9, following=$10, created=$11, modified=$12, fetched=$13, statuscode=$14, blog=$15`
	} else {
		sql += `ON CONFLICT(id) DO NOTHING`
	}

	var modified *time.Time
	if user.UpdatedAt != nil {
		modified = &user.UpdatedAt.Time
	}

	var created *time.Time
	if user.CreatedAt != nil {
		created = &user.CreatedAt.Time
	}

	_, err := conn.Exec(sql, user.GetID(),
		user.GetLogin(),
		user.Name,
		user.Company,
		user.Location,
		user.Bio,
		user.Email,
		user.Type,
		user.Followers,
		user.Following,
		created,
		modified,
		fetchtime,
		statuscode,
		user.Blog)
	return err
}

// InsertRepoStatus updates the statuscode/fetchtime associated with a repo in the case that it
// can't be fetched
func (conn *Database) InsertRepoStatus(repoid int64, reponame string, statuscode int, fetchtime time.Time, upsert bool) error {
	sql := `INSERT INTO repos (id, name, fetched, statuscode) VALUES ($1, $2, $3, $4)`

	if upsert {
		sql += ` ON CONFLICT(id) DO UPDATE SET name=$2, fetched=$3, statuscode=$4`
	} else {
		sql += ` ON CONFLICT(id) DO NOTHING`
	}

	stmt, err := conn.Prepare(sql)
	if err != nil {
		fmt.Printf("Failed to prepare: %s\n", err.Error())
		return err
	}

	_, err = stmt.Exec(repoid, reponame, fetchtime, statuscode)
	return err
}

// InsertRepo inserts a github.Repository object into the database
func (conn *Database) InsertRepo(statuscode *int, fetchtime *time.Time, repo *github.Repository, upsert bool) error {
	sql := `INSERT INTO repos (id, name, language, description, size, stars, forks, topics, parentid,
							   ownerid, created, modified, fetched, statuscode, license, homepage)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16)`
	if upsert {
		sql += ` ON CONFLICT(id) DO UPDATE SET name=$2, language=$3, description=$4, size=$5, stars=$6,
			    forks=$7, topics=$8, parentid=$9, ownerid=$10, created=$11, modified=$12, fetched=$13,
			    statuscode=$14, license=$15, homepage=$16`
	} else {
		sql += ` ON CONFLICT(id) DO NOTHING`
	}

	owner := repo.GetOwner()
	var ownerid *int64
	if owner != nil {
		ownerid = owner.ID
		err := conn.InsertUser(nil, nil, owner, false)
		if err != nil {
			return err
		}
	}

	var parentid *int64
	parent := repo.GetParent()
	if parent != nil {
		// If we have a parent, insert the parent if its missing
		err := conn.InsertRepo(nil, nil, parent, false)
		if err != nil {
			return err
		}
		parentid = parent.ID
	}

	var modified *time.Time
	if repo.PushedAt != nil {
		modified = &repo.PushedAt.Time
	}

	var created *time.Time
	if repo.CreatedAt != nil {
		created = &repo.CreatedAt.Time
	}

	_, err := conn.Exec(sql, repo.GetID(),
		repo.GetFullName(),
		repo.Language,
		repo.Description,
		repo.GetSize(),
		repo.GetStargazersCount(),
		repo.GetForksCount(),
		pq.Array(repo.Topics),
		parentid,
		ownerid,
		created,
		modified,
		fetchtime,
		statuscode,
		repo.GetLicense().GetKey(),
		repo.GetHomepage())

	if err != nil {
		fmt.Printf("Failed to insert github repo: %s", err.Error())
		return err
	}

	return nil
}

// HasRepo returns whether the repo has been fetched
func (conn *Database) HasRepo(repoid int64) (bool, error) {
	// TODO: this doesn't seem all that good
	rows, err := conn.Query("SELECT fetched from repos where id=$1 and fetched is not null", repoid)
	if err != nil {
		return false, err
	}

	defer rows.Close()
	for rows.Next() {
		var fetched time.Time
		if err := rows.Scan(&fetched); err != nil {
			return false, err
		}
		return true, nil
	}

	return false, nil
}

// HasUser returns whether the user has been fetched
func (conn *Database) HasUser(userid int64) (bool, error) {
	// TODO: this doesn't seem all that good
	rows, err := conn.Query("SELECT fetched from users where id=$1 and fetched is not null", userid)
	if err != nil {
		return false, err
	}

	defer rows.Close()
	for rows.Next() {
		var fetched time.Time
		if err := rows.Scan(&fetched); err != nil {
			return false, err
		}
		return true, nil
	}

	return false, nil
}

// InsertOrganizationMembers inserts or updates the list of public organization members in the database
func (conn *Database) InsertOrganizationMembers(orgid int64, orgname string, members []*github.User, statuscode int, fetchtime time.Time, upsert bool) error {
	// Insert a stub user if not already fetched for each user in the organization
	// Note purposefully setting statuscode/fetched time to nil here to mark as not fetched
	// and setting upsert flag to false to prevent overwriting good data
	var memberids []int64
	for _, user := range members {
		memberids = append(memberids, *user.ID)
		conn.InsertUser(nil, nil, user, false)
	}

	fmt.Printf("Writing: %s - %d members\n", orgname, len(memberids))

	sql := `INSERT INTO organization_members (organization, members, fetched, statuscode) VALUES ($1, $2, $3, $4)`
	if upsert {
		sql += ` ON CONFLICT(organization) DO UPDATE SET members=$2, fetched=$3, statuscode=$4`
	} else {
		sql += ` ON CONFLICT(id) DO NOTHING`
	}

	_, err := conn.Exec(sql, orgid, pq.Array(memberids), fetchtime, statuscode)
	return err
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

func (conn *Database) GetRankInfoByLanguage(lan string) ([]*models.LangaugeRank, error) {
	sqlQuery := `select n1num as amount, time_stamp, 1 as rank from top_ten where lower(n1lang) = $1
union
select n2num as amount, time_stamp, 2 as rank from top_ten where lower(n2lang) = $1
union
select n3num as amount, time_stamp, 3 as rank from top_ten where lower(n3lang) = $1
union
select n4num as amount, time_stamp, 4 as rank from top_ten where lower(n4lang) = $1
union
select n5num as amount, time_stamp, 5 as rank from top_ten where lower(n5lang) = $1
union
select n6num as amount, time_stamp, 6 as rank from top_ten where lower(n6lang) = $1
union
select n7num as amount, time_stamp, 7 as rank from top_ten where lower(n7lang) = $1
union
select n8num as amount, time_stamp, 8 as rank from top_ten where lower(n8lang) = $1
union
select n9num as amount, time_stamp, 9 as rank from top_ten where lower(n9lang) = $1
union
select n10num as amount, time_stamp, 10 as rank from top_ten where lower(n10lang) = $1
order by time_stamp;`

	lanLowerCase := strings.ToLower(strings.TrimSpace(lan))

	rows, err := conn.Query(sqlQuery, lanLowerCase)
	if err != nil {
		logs.PrintLogger().Error(err)
		return nil, err
	}
	defer rows.Close()

	languageRanks := make([]*models.LangaugeRank,0)

	var amountN, rankN sql.NullInt64
	var timeStampN sql.NullString

	for rows.Next() {

		languageRank := &models.LangaugeRank{}

		err = rows.Scan(&amountN, &timeStampN, &rankN )

		languageRank.Amount = amountN.Int64
		languageRank.TimeStamp = timeStampN.String
		languageRank.Rank = rankN.Int64

		languageRanks = append(languageRanks, languageRank)
	}

	return languageRanks, err
}
