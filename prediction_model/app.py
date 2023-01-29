from flask import Flask, request
import pickle
import numpy as np
import sys

app = Flask(__name__)
app.config["DEBUG"] = True

# file path to the saved model
model_filepath = 'prediction_model.sav'
try:
    model = pickle.load(open(model_filepath, 'rb'))
except:
    sys.exit('Unable to load the model')


@app.route('/predict', methods=['GET'])
def predict():
    try:
        # Get parameters for temperature
        prediction_attributes = np.array([], dtype='float')
        params = ["OP_CARRIER","OP_CARRIER_FL_NUM","ORIGIN","DEST","CRS_DEP_TIME","DEP_TIME","TAXI_OUT","WHEELS_OFF","WHEELS_ON","TAXI_IN","CRS_ARR_TIME","ARR_TIME","ARR_DELAY","CANCELLED","CANCELLATION_CODE","DIVERTED","CRS_ELAPSED_TIME","ACTUAL_ELAPSED_TIME","AIR_TIME","DISTANCE","CARRIER_DELAY","WEATHER_DELAY","NAS_DELAY","SECURITY_DELAY","LATE_AIRCRAFT_DELAY"]
        #TODO: Add step for cleaning data passed in
        for param in params: 
            prediction_attributes = np.append(prediction_attributes, float(request.args.get((param))))
        features = [prediction_attributes]
        prediction = model.predict(features)
        output = round(prediction[0][0], 2)

        return {'pred_val': output}
    except Exception as e:
        print(e)
        return 'Calculation Error', 500


app.run()