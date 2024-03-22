# frozen_string_literal: true

require 'securerandom'
require 'yaml'
require 'base64'

require 'grpc/health/v1/health_services_pb'

require 'status/v1/http'

module Status
  class << self
    def observability
      @observability ||= Nonnative::Observability.new('http://localhost:11000')
    end

    def server_config
      @server_config ||= Nonnative.configurations('.config/server.yml')
    end

    def health_grpc
      @health_grpc ||= Grpc::Health::V1::Health::Stub.new('localhost:12000', :this_channel_is_insecure, channel_args: Status.user_agent)
    end

    def user_agent
      @user_agent ||= Nonnative::Header.grpc_user_agent(server_config.transport.grpc.user_agent)
    end
  end

  module V1
    class << self
      def server_http
        @server_http ||= Status::V1::HTTP.new('http://localhost:11000')
      end
    end
  end
end
