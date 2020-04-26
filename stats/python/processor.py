import os
import math

import numpy as np
import pandas as pd
import matplotlib.pyplot as plt
from pandas.plotting import register_matplotlib_converters
from stats_pb2 import ConditionsStatsResults


RESOURCES_PATH = "images"


class ConditionsStatsProcessor():

	def process(self, conditions):
		df = self.conditions_to_df(conditions)
		stats_resutls = []
		for metric_id in df["metric_id"].unique():
			series = self.prepare_series(df, metric_id)

			stats = ConditionsStatsResults()
			self.calculate_stats(series, stats)
			stats.image_url = self.make_plot(series, metric_id)
			stats.metric_id = metric_id

			stats_resutls.append(stats)

		return stats_resutls

	def conditions_to_df(self, conditions):
		register_matplotlib_converters()

		data = {
			"metric_id": [],
			"value": [],
			"unix_timestamp": []
		}
		for cond in conditions:
			data["metric_id"].append(cond.metric_id)
			data["value"].append(cond.value)
			data["unix_timestamp"].append(cond.unix_timestamp)
		
		df = pd.DataFrame(data)
		df["date"] = pd.to_datetime(
			df["unix_timestamp"],
			unit="s",
		)
		df.set_index('date', inplace=True)

		return df

	def prepare_series(self, df, metric_id, time_interval="D", smoothing_f=None, **kwargs):
		series = df \
			.loc[df["metric_id"] == metric_id, "value"] \
			.resample(time_interval) \
			.mean()
		
		series = series.fillna(series.mean())
		if smoothing_f is not None:
			series = smoothing_f(series, **kwargs)
		return series
	
	def calculate_stats(self, series, stats):
		stats.mean = series.mean()
		stats.variance = series.var() if series.shape[0] > 1 else 0
		stats.max = series.max()
		stats.min = series.min()


	def make_plot(self, series, metric_id, save_path=RESOURCES_PATH):
		plt.figure(figsize=(20,10))
		plt.plot(series)

		save_path = os.path.join(save_path, f"metric_{metric_id}.png")
		plt.savefig(save_path)

		return save_path


class TimeSeriesProcessor():

	@staticmethod
	def mean_smoothing(series, kernel=3, anomaly_ix=None):
		seriess = series.copy()
		p = (kernel - 1) / 2
		
		if anomaly_ix is None:
			anomaly_ix = np.array(range(int(p) + 1, int(series.shape[0] - p)))
			
		for i in anomaly_ix:
			seriess[i] = np.mean(series[int(i-p): int(i+p)])
			
		return seriess

	@staticmethod
	def mean_chrono_smoothing(series, t, anomaly_ix=None):
		seriess = series.copy()
		
		if anomaly_ix is None:
			anomaly_ix = np.array(range(int(t/2 + 1), int(series.shape[0] - t/2)))
			
		for i in anomaly_ix:
			a = range(int(i - t/2 + 1), int(i + t/2))
			numerator = sum([series[j] for j in a])
			seriess[i] = (series[int(i - t/2)]/2 + numerator + series[int(i + t/2)]/2) / t
			
		return seriess
