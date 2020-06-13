import time
from concurrent import futures

from stats_pb2 import (
	ConditionsStatsResponse,
	ProductsStatsResponse,
	Condition
)
from stats_pb2_grpc import (
	StatsServiceServicer,
	add_StatsServiceServicer_to_server,
	grpc
)
from processor import ConditionsStatsProcessor


PORT = 6000


class PyStatsService(StatsServiceServicer):

	def GetConditionsStats(self, request, context):
		response = ConditionsStatsResponse()

		conditions_proc = ConditionsStatsProcessor()
		response.results.extend(
			conditions_proc.process(request.conditions)
		)

		return response


def serve(port):
	server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
	add_StatsServiceServicer_to_server(
		PyStatsService(),
		server
	)
	server.add_insecure_port('[::]:' + str(port))
	server.start()
	print("Listening on port {}..".format(port))
	try:
		while True:
			time.sleep(60)
	except KeyboardInterrupt:
		server.stop(0)


if __name__== "__main__":
	serve(PORT)