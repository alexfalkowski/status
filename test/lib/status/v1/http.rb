# frozen_string_literal: true

module Status
  module V1
    class HTTP < Nonnative::HTTPClient
      def code(code, headers = {})
        headers.merge!(content_type: :json, accept: :json)

        get("v1/status/#{code}", headers, 10)
      end
    end
  end
end
