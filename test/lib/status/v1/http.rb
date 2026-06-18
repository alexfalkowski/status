# frozen_string_literal: true

module Status
  module V1
    class HTTP < Nonnative::HTTPClient
      def options
        Status.http_options(
          headers: {
            user_agent: 'Status-ruby-client/1.0 HTTP/1.0',
            content_type: :json, accept: :json
          }
        )
      end

      def code(code, sleep = '', opts = {})
        path = "v1/status/#{code}"
        path = "#{path}?sleep=#{sleep}" unless sleep.to_s.empty?

        get(path, opts)
      end
    end
  end
end
