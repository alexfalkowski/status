# frozen_string_literal: true

When('I request with HTTP which performs in {int} ms') do |time|
  opts = {
    headers: {
      request_id: SecureRandom.uuid, user_agent: 'Status-ruby-client/1.0 HTTP/1.0',
      content_type: :json, accept: :json
    },
    read_timeout: 10, open_timeout: 10
  }

  expect { Status::V1.http.code(200, '1ms', opts) }.to perform_under(time).ms
end
