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

      def code_with_location(code, location, opts = {})
        get("v1/status/#{code}?location=#{location}", opts.merge(max_redirects: 0))
      end

      def code_with_retry_after(code, retry_after, opts = {})
        get("v1/status/#{code}?retry_after=#{retry_after}", opts)
      end
    end
  end
end
