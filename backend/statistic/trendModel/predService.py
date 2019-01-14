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


def train_model(lang,total=False,output=1,look_back=7):
    language = lang
    trainSeq, testSeq, dataset = parseJson(language,total)
    scaler = MinMaxScaler(feature_range=(0, 1))
    dataSequence = scaler.fit_transform(dataset)
    inputX, inputY = create_dataset(dataSequence, look_back=look_back)

    inputX = np.reshape(inputX, (inputX.shape[0], 1, inputX.shape[1]))
    # inverse=scaler.inverse_transform(dataSequence)

    model = Sequential()
    model.add(LSTM(10, input_shape=(1, look_back)))
    model.add(Dense(output))
    model.compile(loss='mean_squared_error', optimizer='adam')
    model.fit(inputX, inputY, epochs=int(float(dataset.shape[0]) / 10), batch_size=100, verbose=10)
    model.save(path + '/models/%smodel.h5' % language)
    del model


def pred(timestamp, lang):
    # time_value[timestamp,repo_number]
    time_value = pd.read_csv('datas/%sdata.csv' % lang, encoding='gb18030')
    # bottom of time_value
    time_last = time_value.loc[time_value.shape[0] - 1]['timestamp']
    # input of model
    datas = np.array(time_value.loc[time_value.shape[0] - 7:time_value.shape[0] - 1]['repo_number'])
    datas = np.reshape(datas, [7, 1])
    # normalize input

    scaler = MinMaxScaler(feature_range=(0, 1))
    datas = scaler.fit_transform(datas)
    dataSequence = np.reshape(datas, (1, 1, 7))
    # load model
    model = load_model(path + '/models/%smodel.h5' % lang)
    # predict
    pre = scaler.inverse_transform(model.predict(dataSequence))
    # write in csv
    time_last = datetime.datetime.strptime(time_last, "%Y-%m-%d %H:%M:%S")
    time_new = time_last + datetime.timedelta(days=1)
    time_value.append([time_new, pre])
    time_value.to_csv(path + 'datas/%sdata.csv' % lang, encoding='gb18030')

    # 清理内存
    del time_value
    del model

    if ((timestamp - time_last).days == 1):
        tempSeq = None

        return 1
    else:
        return predTill(timestamp + datetime.timedelta(days=-1), lang)


def predTill(timestamp, lang):
    # check constraint
    lang_list=['C#','C++','CSS','HTML','Java','JavaScript','PHP','Python','Ruby','TypeScript','Perl','C']
    if lang not in lang_list:
        return "[ERROR]: No model found for %s"%lang
    # load data
    time_value = pd.read_csv(path + '/datas/%sdata.csv' % lang, encoding='gb18030')
    time_value = pd.DataFrame(time_value, columns=['timestamp', 'repo_number'])
    time_value = np.array(time_value)

    # the date of latest data in csv file
    time_last = time_value[time_value.shape[0] - 1][0]
    time_last = datetime.datetime.strptime(time_last, "%Y-%m-%d %H:%M:%S")

    if (timestamp <= time_last):
        return "[ERROR]: Please select a date ahead of today"

    # clear old model and load new model
    keras.backend.clear_session()
    model = load_model(path + '/models/%smodel.h5' % lang)
    # output
    pred = []

    while ((timestamp - time_last).days != 0):
        # get last 7 days' data
        datas = np.array(time_value[time_value.shape[0] - 7:time_value.shape[0], 1])
        datas = np.reshape(datas, [7, 1])
        scaler = MinMaxScaler(feature_range=(0, 1))
        datas = scaler.fit_transform(datas)
        dataSequence = np.reshape(datas, (1, 1, 7))
        pre=model.predict(dataSequence)
        pre = scaler.inverse_transform(pre)
        print(pre[0][0])

        time_last = time_last + datetime.timedelta(days=1)
        print(time_value.shape)
        update = np.reshape(['pre', pre[0][0]], [1, 2])
        time_value = np.append(time_value, update, axis=0)
        print(time_value.shape)
        # output
        # pred.append(pre)
        pred = np.append(pred, int(pre))

    del model
    del time_value

    return pred

def predDays(days, lang):
    # check constraint
    lang_list=['C#','C++','CSS','HTML','Java','JavaScript','PHP','Python','Ruby','TypeScript','Perl','C']
    if lang not in lang_list:
        return lang_list
    # load data
    # time_value = pd.read_csv(path + '/datas/%sdata.csv' % lang, encoding='gb18030')
    # time_value = pd.DataFrame(time_value, columns=['timestamp', 'repo_number'])
    #
    # time_value = np.array(time_value)

    connection = MongoClient('0.0.0.0', 27017)

    db = connection.GHUserAnalyse

    set = db.Top10Lang

    results = set.find({'language': lang})

    dataset = []
    for result in results:
        dataset.append([result['timestamp'], result['number']])

    dataset = np.array(dataset)

    connection.close()

    time_value = dataset



    # clear old model and load new model
    keras.backend.clear_session()
    model = load_model(path + '/models/%smodel.h5' % lang)
    # output
    pred = []

    while (days!= 0):
        # get last 7 days' data
        datas = np.array(time_value[time_value.shape[0] - 7:time_value.shape[0], 1])
        datas = np.reshape(datas, [7, 1])
        # scale data
        scaler = MinMaxScaler(feature_range=(0, 1))
        datas = scaler.fit_transform(datas)
        dataSequence = np.reshape(datas, (1, 1, 7))
        # predict
        pre=model.predict(dataSequence)
        # inverse data
        pre = scaler.inverse_transform(pre)

        update = np.reshape(['pre', pre[0][0]], [1, 2])
        time_value = np.append(time_value, update, axis=0)

        pred = np.append(pred, int(pre))
        days-=1

    del model
    del time_value

    return pred

def predTotal(days):
# load data

    connection = MongoClient('0.0.0.0', 27017)

    db = connection.GHUserAnalyse

    set = db.TotalRepoAmount

    results = set.find({})


    dataset = []
    for result in results:
        dataset.append([result['timestamp'], result['amount']])

    dataset = np.array(dataset)
    print(dataset)
    connection.close()


    time_value = dataset
    keras.backend.clear_session()
    model = load_model(path + '/models/Totalmodel.h5')
    pred = []

    while (days!= 0):
        # get last 7 days' data
        datas = np.array(time_value[time_value.shape[0] - 7:time_value.shape[0], 1])
        datas = np.reshape(datas, [7, 1])
        # scale data
        scaler = MinMaxScaler(feature_range=(0, 1))
        datas = scaler.fit_transform(datas)
        dataSequence = np.reshape(datas, (1, 1, 7))
        # predict
        pre=model.predict(dataSequence)
        # inverse data
        pre = scaler.inverse_transform(pre)

        update = np.reshape(['pre', pre[0][0]], [1, 2])
        time_value = np.append(time_value, update, axis=0)

        pred = np.append(pred, int(pre))
        days-=1

    del model
    del time_value

    return pred


# a=predTill(datetime.datetime(2018,11,23),'JavaScript')
# print(a)
# train_model('JavaScript')
# train_model('Perl')
# train_model(lang="Total",total=True)
# train_model("Total",total=True)
# a=predTotal(12)
# print(a)