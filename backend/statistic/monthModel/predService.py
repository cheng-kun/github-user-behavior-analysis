import numpy as np
import datetime
import pandas as pd
import keras
from keras.models import Sequential, load_model
from keras.layers import Dense, LSTM
from sklearn.preprocessing import MinMaxScaler
from createDataset import parseJson, create_dataset
from pymongo import MongoClient
import os

path = os.path.dirname(os.path.realpath(__file__))

look_back = 12


def train_model(lang, total=False, output=1, look_back=12):
    language = lang
    trainSeq, testSeq, dataset = parseJson(language, total)
    scaler = MinMaxScaler(feature_range=(0, 1))

    dataSequence = scaler.fit_transform(dataset)
    inputX, inputY = create_dataset(dataSequence, look_back=look_back)

    inputX = np.reshape(inputX, (inputX.shape[0], 1, inputX.shape[1]))
    # inverse=scaler.inverse_transform(dataSequence)

    model = Sequential()
    model.add(LSTM(10, input_shape=(1, look_back)))
    model.add(Dense(output))
    model.compile(loss='mean_squared_error', optimizer='adam')
    model.fit(inputX, inputY, epochs=1000, batch_size=100, verbose=10)
    model.save(path + '/models/%smodel_avg.h5' % language)
    del model


def predMonth(months, lang, lookback=look_back):
    # check constraint
    lang_list = ['C#', 'C++', 'CSS', 'HTML', 'Java', 'JavaScript', 'PHP', 'Python', 'Ruby', 'TypeScript', 'Perl', 'C']
    if lang not in lang_list:
        return lang_list
    # load data
    connection = MongoClient('0.0.0.0', 27017)

    db = connection.GHUserAnalyse

    set = db.MonAvgLang

    results = set.find({'language': lang})

    dataset = []
    for result in results:
        dataset.append([result['timestamp'], result['month_avg']])

    dataset = np.array(dataset)

    connection.close()

    time_value = dataset

    # clear old model and load new model
    keras.backend.clear_session()
    model = load_model(path + '/models/%smodel_avg.h5' % lang)
    # output
    pred = []

    while (months != 0):
        # get last 7 days' data
        datas = np.array(time_value[time_value.shape[0] - lookback:time_value.shape[0], 1])
        datas = np.reshape(datas, [lookback, 1])
        # scale data
        scaler = MinMaxScaler(feature_range=(0, 1))
        datas = scaler.fit_transform(datas)
        dataSequence = np.reshape(datas, (1, 1, lookback))
        # predict
        pre = model.predict(dataSequence)
        # inverse data
        pre = scaler.inverse_transform(pre)

        update = np.reshape(['pre', pre[0][0]], [1, 2])
        time_value = np.append(time_value, update, axis=0)

        pred = np.append(pred, int(pre))
        months -= 1

    del model
    del time_value

    return pred


def evaluate():
    pass


def predTop10():
    lang_list = ['C#', 'C++', 'CSS', 'HTML', 'Java', 'JavaScript', 'PHP', 'Python', 'Ruby', 'TypeScript']
    top10 = {}

    connection = MongoClient('0.0.0.0', 27017)

    db = connection.GHUserAnalyse

    set = db.Top10Pred

    timestamp=datetime.date.today()
    timestamp=datetime.datetime(timestamp.year,timestamp.month,1)
    if (set.count({'timestamp': timestamp}) != 0):
        results=set.find({'timestamp':timestamp})
        for result in results:
            top10[result['language']]=result['amount']
    else:
        for lang in lang_list:
            pred = predMonth(1, lang=lang, lookback=12)
            # top10.append([lang,pred[0]])
            top10[lang] = pred[0]
            language={
                'timestamp':timestamp,
                "language":lang,
                'amount':pred[0]

            }
            set.insert(language)
    connection.close()
    # sorted(top10,key=lambda top10:top10[1])
    # top10_sort = sorted(top10.items(), key=lambda top10: top10[1],reverse=True)
    #
    # top10={}
    # for tuple in top10_sort:
    #     top10[tuple[0]]=tuple[1]

    # top10 = np.array(top10)
    return top10


def getAverage(timestamp):
    timestamp = datetime.datetime.strptime(str(timestamp), "%Y%m%d")
    connection = MongoClient('0.0.0.0', 27017)

    db = connection.GHUserAnalyse

    set = db.MonAvgLang

    # results = set.find({'timestamp':timestamp})
    results = set.find({'timestamp': timestamp})

    # str=str(timestamp.year)+'-'+str(timestamp.month)+'-'+str(timestamp.day)+"T00:00:00Z"
    #
    # dict={}
    # dict['amount']=results['month_avg']
    # dict['time_stamp':]
    languages = {}
    for result in results:
        language = result['language']
        num = result['month_avg']
        languages[language] = num

    connection.close()

    return languages

def draw(model,language):
    trainSeq, testSeq, dataset = parseJson(language, True)
    pass


# a=predTill(datetime.datetime(2018,11,23),'JavaScript')
# print(a)
# predTop10()
# train_model('Perl')
# train_model(lang="TypeScript")
#
# a=predDays(12,'Python',lookback=12)
# a = predTop10()
# a=getAverage("Python",datetime.datetime(2018,1,1))
# for item in a:
#     print(item.value)
# a = getAverage(20180201)
