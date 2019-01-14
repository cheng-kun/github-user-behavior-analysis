import h5py
import numpy as np
from keras.models import Sequential, load_model
from keras.layers import Dense, LSTM
from sklearn.preprocessing import MinMaxScaler
from sklearn.metrics import mean_squared_error
from createDataset import parseJson, create_dataset
import pandas as pd
import datetime
import matplotlib.pyplot as plt

look_back = 7
language = 'Python'
trainSeq, testSeq, dataset = parseJson(language)


def train_model(trainseq,language=language):
    scaler = MinMaxScaler(feature_range=(0, 1))
    dataSequence = scaler.fit_transform(trainseq)
    inputX, inputY = create_dataset(dataSequence, look_back=look_back)

    inputX = np.reshape(inputX, (inputX.shape[0], 1, inputX.shape[1]))
    # inverse=scaler.inverse_transform(dataSequence)

    model = Sequential()
    model.add(LSTM(10, input_shape=(1, look_back)))
    model.add(Dense(1))
    model.compile(loss='mean_squared_error', optimizer='adam')
    model.fit(inputX, inputY, epochs=int(float(trainseq.shape[0]) / 10), batch_size=100, verbose=10)
    model.save('models/%smodel.h5' % language)


def predict(testseq,language=language):
    scaler = MinMaxScaler(feature_range=(0, 1))
    model = load_model('models/%smodel.h5' % language)
    #[-1,1]
    dataSequence = scaler.fit_transform(testseq)
    testX, testY = create_dataset(dataSequence, look_back=look_back)
    print(testX.shape)
    testX = np.reshape(testX, (testX.shape[0], 1, testX.shape[1]))
    print(testX.shape)
    Pre = scaler.inverse_transform(model.predict(testX))
    Real = scaler.inverse_transform(np.reshape(testY, [len(testY), 1]))
    return Pre, Real


def predTill(timestamp,lang):
    #timestamp:datetime类
    dataX=[]

    time_value=pd.read_csv('datas/%sdata.csv'%lang,encoding='gb18030')
    time_last=time_value.loc[time_value.shape[0]]['timestamp']
    datas=time_value.loc[time_value.shape[0]-7:time_value.shape[0]]['repo_number']
    datas=np.reshape(datas,(1,1,7))
    scaler = MinMaxScaler(feature_range=(0, 1))
    dataSequence=scaler.fit_transform(datas)
    model = load_model('models/%smodel.h5' % language)
    pre=scaler.inverse_transform(model.predict(dataSequence))

    predict()
    time_last=datetime.datetime.strptime(time_last,"%Y-%m-%d %H:%M:%S")
    timestamp=datetime.datetime(timestamp)
    if ((timestamp-time_last).days==1):
        tempSeq=None

        return 1
    else:
        return predTill(timestamp+datetime.timedelta(days=-1),lang)



def plot(dataset, trainPredict):
    trainPredictPlot = np.empty_like(dataset)
    trainPredictPlot[:, :] = np.nan
    trainPredictPlot[look_back:len(trainPredict) + look_back, :] = trainPredict
    plt.subplot(dataset, label='Actual')
    plt.subplot()
    plt.show()
    return trainPredictPlot


def plotall():
    p, r = predict(testSeq)
    p2, r2 = predict(trainSeq)
    trainPredictPlot = np.empty_like(dataset)
    trainPredictPlot[:, :] = np.nan
    trainPredictPlot[look_back:len(p2) + look_back, :] = p2

    testPredictPlot = np.empty_like(dataset)
    testPredictPlot[:, :] = np.nan
    testPredictPlot[len(p2) + (look_back * 2) + 1:len(dataset) - 1, :] = p
    plt.plot(dataset, label='Actual')
    plt.plot(trainPredictPlot, label='trainset')
    plt.plot(testPredictPlot, label='test')
    plt.show()
plotall()