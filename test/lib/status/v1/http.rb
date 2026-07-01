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

      private

      def status(code, opts, method: 'get', **params)
        path = status_path(code, params)
        args = payload_method?(method) ? [path, '', opts] : [path, opts]

        send(method, *args)
      end

      def status_path(code, params)
        query = params.filter_map { |key, value| "#{key}=#{value}" unless value.to_s.empty? }.join('&')
        path = "v1/status/#{code}"

        query.empty? ? path : "#{path}?#{query}"
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
