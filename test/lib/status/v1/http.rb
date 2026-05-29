# frozen_string_literal: true

require 'securerandom'

module Status
  module V1
    class HTTP < Nonnative::HTTPClient
      def options
        {
          headers: {
            request_id: SecureRandom.uuid, user_agent: 'Status-ruby-client/1.0 HTTP/1.0',
            content_type: :json, accept: :json
          },
          read_timeout: 10, open_timeout: 10
        }
      end

      def code(code, sleep = '', opts = {})
        path = "v1/status/#{code}"
        path = "#{path}?sleep=#{sleep}" unless sleep.to_s.empty?

        get(path, opts)
      end
    end
  end
end
