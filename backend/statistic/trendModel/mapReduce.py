from pymongo import MongoClient
from bson.code import Code
import datetime
import calendar

connection = MongoClient('0.0.0.0', 27017)
# connection = MongoClient('52.14.21.138', 27017)

db = connection.GHUserAnalyse

set = db.Top10Lang

set2 = db.MonAvgLang


def add_months(sourcedate, months):
    month = sourcedate.month - 1 + months
    year = sourcedate.year + month // 12
    month = month % 12 + 1
    day = min(sourcedate.day, calendar.monthrange(year, month)[1])
    return datetime.datetime(year, month, day)


map = Code(

    "function() {"
    "   emit(this.language, this.number);"
    "}"

)

reduce = Code(

    "function(key,values){"
    "   var sum = Array.sum(values);"
    "   var days = values.length;"
    "   var avg = Math.round(sum/days);"
    "   return avg;"
    "}"
)
date = datetime.datetime(2009, 1, 1)

while (date < datetime.datetime.now()):
    date_start = date
    date_end = add_months(date_start, 1)
    results = set.map_reduce(map, reduce, "results",
                             query={'$and': [{"timestamp": {'$gte': date_start}}, {"timestamp": {'$lt': date_end}}]})

    for doc in results.find():
        language = doc['_id']
        avgnum = doc['value']

        file={
            "language":language,
            "timestamp":date_start,
            "month_avg":avgnum
        }

        result=set2.insert(file)

    date = add_months(date, 1)
