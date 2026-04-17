# frozen_string_literal: true

require 'securerandom'

require 'status/v1/http'

module Status
  module V1
    class << self
      def http
        @http ||= Status::V1::HTTP.new('http://localhost:11000')
      end
    end
  end
end
