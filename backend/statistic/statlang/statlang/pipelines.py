# -*- coding: utf-8 -*-

# Define your item pipelines here
#
# Don't forget to add your pipeline to the ITEM_PIPELINES setting
# See: https://doc.scrapy.org/en/latest/topics/item-pipeline.html
import json
import os
import datetime
from pymongo import MongoClient
filepath=os.path.abspath('../trendModel/datas/iterm.jl')

class StatlangPipeline(object):

    def open_spider(self,spider):
        self.connection = MongoClient('0.0.0.0', 27017)
        db = self.connection.GHUserAnalyse
        self.set = db.Top10Lang
        self.set2 = db.TotalRepoAmount

        self.file = open(filepath, 'a+')
    def close_spider(self,spider):
        self.connection.close()

        self.file.close()
    def process_item(self, item, spider):
        # line=json.dumps(dict(item))+"\n"
        # # time_new=line[3]
        # # time_new=datetime.datetime.strptime(time_new,['%Y-%m-%dT%H:%M:%S'])
        # # self.file[-1]['timestamp']
        # self.file.write(line)
        timestamp=datetime.datetime.strptime(item['timestamp'], "%Y-%m-%dT%H:%M:%S")
        if (self.set.count({'timestamp':timestamp})!=0):
            print('[LOG]: data of %s all ready exists'%item['timestamp'])
        else:
            total = {
                'timestamp': timestamp,

                'amount': item['repo_num']
            }
            self.set2.insert(total)
            for i in range(0, 10):
                language = {
                    'timestamp': timestamp,
                    'language': item['n%dlang' % (i + 1)],
                    'number': item['n%dnum' % (i + 1)]
                }
                result = self.set.insert(language)

        line = json.dumps(dict(item)) + "\n"
        self.file.write(line)
        return item
