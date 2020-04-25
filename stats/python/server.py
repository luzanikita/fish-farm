import time
from concurrent import futures

from stats_pb2 import *
from stats_pb2_grpc import *


class PyStatsService(StatsServiceServicer):

	def GetStats(self, request, context):
		# Text we got from the client
		text = request.text

		response = StatsResponse()
		response.text = text + "_!"

		return response

def serve(port):
	server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
	add_StatsServiceServicer_to_server(
		PyStatsService(), server)
	server.add_insecure_port('[::]:' + str(port))
	server.start()
	print("Listening on port {}..".format(port))
	try:
		while True:
			time.sleep(10000)
	except KeyboardInterrupt:
		server.stop(0)


if __name__== "__main__":
	serve(6000)