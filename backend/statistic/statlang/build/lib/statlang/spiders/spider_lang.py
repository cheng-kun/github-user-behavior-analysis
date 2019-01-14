import scrapy
import datetime
from scrapy.selector import Selector
import re
from statlang.items import StatlangItem
import time

def timeadd(timein):
    if isinstance(timein,str):
        timein = datetime.datetime.strptime(timein, "%Y-%m-%dT%H:%M:%S")
    timeout=timein+datetime.timedelta(days=1)
    return timeout


def convert_time(timein):

    if isinstance(timein,str):
        timein=datetime.datetime.strptime(timein,"%Y-%m-%dT%H:%M:%S")
    date_start = timein
    date_end = date_start + datetime.timedelta(days=1)

    dt_start = date_start.strftime('%Y-%m-%dT%H:%M:%S')
    dt_end = date_end.strftime('%Y-%m-%dT%H:%M:%S')

    return dt_start, dt_end


date_begin = datetime.datetime(2009, 1, 1, 0, 0, 0)
date_start, date_end = convert_time(date_begin)
count=0

class StatsticLang(scrapy.Spider):
    name = "Stat_lang"
    allowed_domains = ['github.com']
    headers = {
        'Authorization': 'token 5aa6cadda5bdc1145e6bc978b62daeb61738584a',
        "Accept": "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8",
        "Accept-Encoding": "gzip,deflate",
        "Accept-Language": "en-US,en;q=0.8",
        "Connection": "keep-alive",
        "Content-Type": " application/x-www-form-urlencoded",
        "User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/38.0.2125.111 Safari/537.36",
    }
    start_urls = [
        "https://github.com/search?q=created%3A" + date_start + ".." + date_end + "&type=Repositories",
    ]

    def parse(self, response):
        print(response.request.headers["User-Agent"])
        global date_start, date_end,count
        statlang = StatlangItem()
        selector = Selector(response)
        if (response.status == 200):
            repos_per_day = selector.xpath(
                './/div[@class="d-flex flex-column flex-md-row flex-justify-between border-bottom pb-3 position-relative"]/h3/text()'
            ).extract()[0]

            languages = selector.xpath(
                './/ul[@class="filter-list small"]/li/a/text()'
            ).extract()

            lang_num = selector.xpath(
                './/ul[@class="filter-list small"]/li/a/span/text()'
            ).extract()

            # repos number pushed to Github at this date
            statlang['repo_num'] = int(re.sub(r'\s+', '', re.sub(' repository results', '', repos_per_day)))

            statlang['timestamp'] = convert_time(date_start)[0]

            for i in range(1, 11):
                statlang["n%dlang" % (i)] = re.sub(r'\s+', '', languages[2 * i - 1])
                statlang["n%dnum" % (i)] = int(re.sub(r'\s', '', lang_num[i - 1]))

            # n1lang = re.sub(r'\s+', '', languages[1])  # 2i-1
            # n1num = int(re.sub(r'\s', '', lang_num[0]))  # i-1
            # n2lang = re.sub(r'\s+', '', languages[3])  # 2i-1
            # n2num = int(re.sub(r'\s', '', lang_num[1]))  # i-1
            # n3lang = re.sub(r'\s+', '', languages[5])  # 2i-1
            # n3num = int(re.sub(r'\s', '', lang_num[2]))  # i-1
            # n4lang = re.sub(r'\s+', '', languages[7])  # 2i-1
            # n4num = int(re.sub(r'\s', '', lang_num[3]))  # i-1
            # n5lang = re.sub(r'\s+', '', languages[9])  # 2i-1
            # n5num = int(re.sub(r'\s', '', lang_num[4]))  # i-1
            # n6lang = re.sub(r'\s+', '', languages[11])  # 2i-1
            # n6num = int(re.sub(r'\s', '', lang_num[5]))  # i-1
            # n7lang = re.sub(r'\s+', '', languages[13])  # 2i-1
            # n7num = int(re.sub(r'\s', '', lang_num[6]))  # i-1
            # n8lang = re.sub(r'\s+', '', languages[15])  # 2i-1
            # n8num = int(re.sub(r'\s', '', lang_num[7]))  # i-1
            # n9lang = re.sub(r'\s+', '', languages[17])  # 2i-1
            # n9num = int(re.sub(r'\s', '', lang_num[8]))  # i-1
            # n10lang = re.sub(r'\s+', '', languages[19])  # 2i-1
            # n10num = int(re.sub(r'\s', '', lang_num[9]))  # i-1
            yield statlang
        else:
            print(date_start)

        date_start,date_end=convert_time(date_end)

        next_url="https://github.com/search?q=created%3A" + date_start + ".." + date_end + "&type=Repositories"
        yield response.follow(next_url,self.parse)



