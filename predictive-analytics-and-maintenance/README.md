## Methodology of Predictive Outlier Analysis of Server Data

__This is a two step pipeline that is designed to first forecast the behavior of each variable independent of another. This independence amongst variables is a reasonable assumption and appropriate analytics decision while making the predictions for the future data i.e. forecasting target variables. And from the same forecasted variables are fed in combined manner to our unsupervised outlier detection ensemble algorithm.__  

__Below, we state the exact steps followed for the modelling and prediction process:__  

* Loading the dataset onto dataframes. Here, we have decided to work with `multi-var-four-two.csv` which describes the scenario or state features for a given single EC2 AWS VM.
* Second, we separate these variables and feed them individually for prediction into `fbprophet` addition based models. Also, we keep our forecasting range to be `4.5 days` as knowing the defect almost 1 working week earlier is more than enough time to take suitable actions for ramifications.
  * After training the model and making a forecast we do evaluate the accuracy of the model.
* After making the forecasts we use the same data for training out `pyod` outlier models from the previous training data split of `3024th` row number.
  * After training the model and making outlier predictions. We quantitatively analyze the output results by comparing outlier prediction values on actual and forecasted data in this notebook.
