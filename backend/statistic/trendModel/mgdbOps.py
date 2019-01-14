from pymongo import MongoClient
import jsonlines as jl
import datetime
import os
path=os.path.dirname(os.path.realpath(__file__))

connection = MongoClient('0.0.0.0',27017)

db = connection.GHUserAnalyse

set = db.Top10Lang

def write():
    x = 0
    filename = path + '/datas/iterm.jl'
    with open(filename, 'r+', encoding='utf8') as f:
        for item in jl.Reader(f):

            # total={
            #         'timestamp':datetime.datetime.strptime(item['timestamp'], "%Y-%m-%dT%H:%M:%S"),
            #         'language':'Total',
            #         'number':item['repo_num']
            #     }
            # result=set.insert(total)

            for i in range(0, 10):

                language={
                    'timestamp':datetime.datetime.strptime(item['timestamp'], "%Y-%m-%dT%H:%M:%S"),
                    'language':item['n%dlang' % (i + 1)],
                    'number':item['n%dnum' % (i + 1)]
                }
                result=set.insert(language)
                x=x+1
                if x%100==0:
                    print(x)

def find():
    for i in set.find({'language':'Python'}):
        print(i)

write()
