# frozen_string_literal: true

module Status
  module V1
    class HTTP < Nonnative::HTTPClient
      def code(code, sleep, opts = {})
        get("v1/status/#{code}?sleep=#{sleep}", opts)
      end
    end
  end
end
