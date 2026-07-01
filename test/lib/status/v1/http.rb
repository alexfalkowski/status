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
        status(code, opts, sleep: sleep)
      end

      def code_with_method(code, method, opts = {})
        status(code, opts, method: method.downcase)
      end

      def code_with_location(code, location, opts = {})
        status(code, opts.merge(max_redirects: 0), location: location)
      end

      def code_with_retry_after(code, retry_after, opts = {})
        status(code, opts, retry_after: retry_after)
      end

      def code_with_header(code, header, value, opts = {})
        status(code, opts, header: "#{header}:#{value}")
      end

      def code_with_headers(code, headers, opts = {})
        status(code, opts, header: headers.map { |header, value| "#{header}:#{value}" })
      end

      def code_with_raw_header(code, header, opts = {})
        status(code, opts, header: header)
      end

      private

      def status(code, opts, method: 'get', **params)
        path = status_path(code, params)
        args = payload_method?(method) ? [path, '', opts] : [path, opts]

        send(method, *args)
      end

      def status_path(code, params)
        query = params.flat_map { |key, value| query_values(key, value) }.join('&')
        path = "v1/status/#{code}"

        query.empty? ? path : "#{path}?#{query}"
      end

      def query_values(key, value)
        Array(value).filter_map { |item| "#{key}=#{item}" unless item.to_s.empty? }
      end

      def payload_method?(method)
        %w[post put patch].include?(method)
      end

      def patch(path, payload, opts = {})
        with_exception { resource(path, opts).patch(payload) }
      end
    end
  end
end
