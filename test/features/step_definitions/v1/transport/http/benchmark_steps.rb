# frozen_string_literal: true

When('I request which performs in {int} ms') do |time|
  opts = Status.http_options(
    headers: {
      user_agent: 'Status-ruby-client/1.0 HTTP/1.0',
      content_type: :json, accept: :json
    }
  )

  expect { Status::V1.http.code(200, '1ms', opts) }.to perform_under(time).ms
end
