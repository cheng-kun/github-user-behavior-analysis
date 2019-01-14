from flask import Flask, request, jsonify

import datetime
import numpy as np
import sys
import os

sys.path.append(os.path.abspath('../'))
sys.path.append(os.path.abspath('../trendModel/'))
sys.path.append(os.path.abspath('../monthModel/'))
from trendModel.predService import predDays, predTotal
from monthModel.predService import predMonth, predTop10, getAverage
from flask_cors import *

app = Flask(__name__)
CORS(app, supports_credentials=True)


@app.route('/')
def home():
    return app.send_static_file('index.html')


@app.route('/api/pred/days/', methods=['GET'])
def getPredTill():
    args_days = request.args.get('days')
    args_days = int(args_days)
    args_lang = request.args.get('language')
    # print(args_lang)
    # date=datetime.datetime.strptime(args_date,'%Y%m%d')
    # check constraint
    lang_list = ['C#', 'C++', 'CSS', 'HTML', 'Java', 'JavaScript', 'PHP', 'Python', 'Ruby', 'TypeScript', 'Perl', 'C']
    if args_lang not in lang_list:
        return jsonify(
            "[ERROR]: No model found for %s, please try another languages" % args_lang
        )
    if args_days <= 0:
        return jsonify(
            "[ERROR]: Illegal number, days must be bigger than 0."
        )

    pred = predDays(args_days, args_lang)
    pred = pred.tolist()
    return jsonify(pred)


@app.route('/api/pred/months/', methods=['GET'])
def getPredMonth():
    args_month = request.args.get('months')
    args_month = int(args_month)
    args_lang = request.args.get('language')
    # print(args_lang)
    # date=datetime.datetime.strptime(args_date,'%Y%m%d')
    # check constraint
    lang_list = ['C#', 'C++', 'CSS', 'HTML', 'Java', 'JavaScript', 'PHP', 'Python', 'Ruby', 'TypeScript', 'Perl', 'C']
    if args_lang not in lang_list:
        return jsonify(
            "[ERROR]: No model found for %s, please try another languages" % args_lang
        )
    if args_month <= 0:
        return jsonify(
            "[ERROR]: Illegal number, days must be bigger than 0."
        )

    pred = predMonth(args_month, args_lang, 12)
    pred = pred.tolist()
    return jsonify(pred)


@app.route('/api/top10/', methods=['GET'])
def pred_Top10():
    top10 = predTop10()
    return jsonify(top10)


@app.route('/api/monthavg/', methods=['GET'])
def getMonthAvg():
    args_month = request.args.get('timestamp')
    args_month = int(args_month)
    get = getAverage(args_month)
    return jsonify(get)


@app.route('/api/pred/days/total/', methods=['GET'])
def pred_Total():
    args = request.args.get('days')
    total = predTotal(int(args))
    total=total.tolist()

    return jsonify(total)


if __name__ == '__main__':
    # get localhost:2222/api/date=20180101&language=Python
    app.run(host="172.31.31.247", port=2222)
    # app.run(host="0.0.0.0", port=2222)
    # "+"和"#"转成ASCII码值 '%2B'和'%23'
