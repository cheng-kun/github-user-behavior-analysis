1. Get Daily Rank By Date

HTTP GET
HOST:52.14.21.138:8080
ENDPOINT:/language/date/{:date}

RESPONSE: JSON

type Ranking struct {
	RepoNum int64 `json:"repo_num"`
	TimeStamp string`json:"timestamp"`
	N1lang string	`json:"n1lang"`
	N1num int64		`json:"n1num"`
	N2lang string	`json:"n2lang"`
	N2num int64		`json:"n2num"`
	N3lang string	`json:"n3lang"`
	N3num int64		`json:"n3num"`
	N4lang string	`json:"n4lang"`
	N4num int64		`json:"n4num"`
	N5lang string	`json:"n5lang"`
	N5num int64		`json:"n5num"`
	N6lang string	`json:"n6lang"`
	N6num int64		`json:"n6num"`
	N7lang string	`json:"n7lang"`
	N7num int64		`json:"n7num"`
	N8lang string	`json:"n8lang"`
	N8num int64		`json:"n8num"`
	N9lang string	`json:"n9lang"`
	N9num int64		`json:"n9num"`
	N10lang string	`json:"n10lang"`
	N10num int64	`json:"n10num"`
}


EXAMPLE

GET: 52.14.21.138:8080/language/date/20140101

RESPONSE:
{
    "repo_num": 3914,
    "timestamp": "2014-01-01T00:00:00Z",
    "n1lang": "JavaScript",
    "n1num": 458,
    "n2lang": "Java",
    "n2num": 349,
    "n3lang": "Python",
    "n3num": 277,
    "n4lang": "Ruby",
    "n4num": 275,
    "n5lang": "CSS",
    "n5num": 207,
    "n6lang": "PHP",
    "n6num": 161,
    "n7lang": "C++",
    "n7num": 145,
    "n8lang": "C",
    "n8num": 138,
    "n9lang": "C#",
    "n9num": 92,
    "n10lang": "Shell",
    "n10num": 91
}

=============================
2. Get Rank Info By Language Name

HTTP GET
HOST:52.14.21.138:8080
ENDPOINT:/language/nameandday/{:name}/{:datetime}

RESPONSE: JSON

type LangaugeRank struct {
	Amount int64 `json:"amount"`
	TimeStamp string `json:"time_stamp"`
	Rank int64 `json:"rank"`
} 


EXAMPLE

GET: 52.14.21.138:8080/language/nameandday/go/20091118

RESPONSE:
[
    {
        "amount": 4,
        "time_stamp": "2009-11-18T00:00:00Z",
        "rank": 10
    }
]

=============================
3. Get Top User

HTTP GET
HOST:52.14.21.138:8080
ENDPOINT: /topuser/{:amount}

RESPONSE: JSON

type UserFollower struct {
	User string `json:"user"`
	Follower int64 `json:"follower"`
	Rank int64 `json:"rank"`
} 

EXAMPLE

GET: 52.14.21.138:8080/topuser/3

RESPONSE:
[
    {
        "user": "4148",
        "follower": 164129,
        "rank": 1
    },
    {
        "user": "cusspvz",
        "follower": 114438,
        "rank": 2
    },
    {
        "user": "alex-cory",
        "follower": 85929,
        "rank": 3
    }
]

=============================
4.
Get Top Country

HTTP GET
HOST:52.14.21.138:8080
ENDPOINT: /topcountry/{:amount}

RESPONSE: JSON

type UserCountry struct {
	Country string `json:"country"`
	UserAmount int64 `json:"user_amount"`
	RepoAmount int64 `json:"repo_amount"`
	PushOct	int64 `json:"push_oct"`
} 

EXAMPLE

GET: 52.14.21.138:8080/topcountry/3

RESPONSE:
[
    {
        "country": "us",
        "user_amount": 183319,
        "repo_amount": 827123,
        "push_oct": 1045587
    },
    {
        "country": "gb",
        "user_amount": 33907,
        "repo_amount": 155739,
        "push_oct": 244199
    },
    {
        "country": "cn",
        "user_amount": 31793,
        "repo_amount": 127811,
        "push_oct": 149626
    }
]