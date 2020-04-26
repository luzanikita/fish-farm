# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
import grpc

import stats_pb2 as stats__pb2


class StatsServiceStub(object):
    """Missing associated documentation comment in .proto file"""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.GetConditionsStats = channel.unary_unary(
                '/stats.StatsService/GetConditionsStats',
                request_serializer=stats__pb2.ConditionsStatsRequest.SerializeToString,
                response_deserializer=stats__pb2.ConditionsStatsResponse.FromString,
                )
        self.GetProductsStats = channel.unary_unary(
                '/stats.StatsService/GetProductsStats',
                request_serializer=stats__pb2.ProductsStatsRequest.SerializeToString,
                response_deserializer=stats__pb2.ProductsStatsResponse.FromString,
                )


class StatsServiceServicer(object):
    """Missing associated documentation comment in .proto file"""

    def GetConditionsStats(self, request, context):
        """Missing associated documentation comment in .proto file"""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetProductsStats(self, request, context):
        """Missing associated documentation comment in .proto file"""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_StatsServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'GetConditionsStats': grpc.unary_unary_rpc_method_handler(
                    servicer.GetConditionsStats,
                    request_deserializer=stats__pb2.ConditionsStatsRequest.FromString,
                    response_serializer=stats__pb2.ConditionsStatsResponse.SerializeToString,
            ),
            'GetProductsStats': grpc.unary_unary_rpc_method_handler(
                    servicer.GetProductsStats,
                    request_deserializer=stats__pb2.ProductsStatsRequest.FromString,
                    response_serializer=stats__pb2.ProductsStatsResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'stats.StatsService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class StatsService(object):
    """Missing associated documentation comment in .proto file"""

    @staticmethod
    def GetConditionsStats(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/stats.StatsService/GetConditionsStats',
            stats__pb2.ConditionsStatsRequest.SerializeToString,
            stats__pb2.ConditionsStatsResponse.FromString,
            options, channel_credentials,
            call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetProductsStats(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/stats.StatsService/GetProductsStats',
            stats__pb2.ProductsStatsRequest.SerializeToString,
            stats__pb2.ProductsStatsResponse.FromString,
            options, channel_credentials,
            call_credentials, compression, wait_for_ready, timeout, metadata)
