# frozen_string_literal: true

require 'securerandom'
require 'status/v1/http'

module Status
  class << self
    ##
    # Returns bounded per-call options for HTTP feature-harness requests.
    #
    # Each call includes a generated `request_id` header. Caller-provided
    # headers are merged afterward, so scenarios can override that value or add
    # transport-specific headers such as content type and user agent.
    #
    # @param headers [Hash] HTTP headers merged after the generated request id
    # @param read_timeout [Integer] read timeout in seconds
    # @param open_timeout [Integer] connection open timeout in seconds
    # @return [Hash] options compatible with `Nonnative::HTTPClient` calls
    def http_options(headers: {}, read_timeout: 10, open_timeout: 10)
      {
        headers: { request_id: SecureRandom.uuid }.merge(headers),
        read_timeout:,
        open_timeout:
      }
    end
  end

  module V1
    class << self
      def http
        @http ||= Status::V1::HTTP.new('http://localhost:11000')
      end
    end
  end
end
