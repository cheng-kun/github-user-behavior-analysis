# -*- coding: utf-8 -*-

# Define here the models for your scraped items
#
# See documentation in:
# https://doc.scrapy.org/en/latest/topics/items.html

import scrapy


class StatlangItem(scrapy.Item):
    # define the fields for your item here like:
    # name = scrapy.Field()
    timestamp=scrapy.Field()

    repo_num=scrapy.Field()

    n1lang = scrapy.Field()
    n1num = scrapy.Field()
    n2lang = scrapy.Field()
    n2num = scrapy.Field()
    n3lang = scrapy.Field()
    n3num = scrapy.Field()
    n4lang = scrapy.Field()
    n4num = scrapy.Field()
    n5lang = scrapy.Field()
    n5num = scrapy.Field()
    n6lang = scrapy.Field()
    n6num = scrapy.Field()
    n7lang = scrapy.Field()
    n7num = scrapy.Field()
    n8lang = scrapy.Field()
    n8num = scrapy.Field()
    n9lang = scrapy.Field()
    n9num = scrapy.Field()
    n10lang = scrapy.Field()
    n10num = scrapy.Field()
    pass
