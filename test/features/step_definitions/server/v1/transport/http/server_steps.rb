# frozen_string_literal: true

When('I request to set the code {int} with HTTP') do |code|
  opts = {
    headers: {
      request_id: SecureRandom.uuid, user_agent: Status.server_config.transport.http.user_agent,
      content_type: :json, accept: :json
    },
    read_timeout: 10, open_timeout: 10
  }

  @response = Status::V1.server_http.code(code, '1ms', opts)
end

When('I request to set the invalid code {string} with HTTP') do |code|
  opts = {
    headers: {
      request_id: SecureRandom.uuid, user_agent: Status.server_config.transport.http.user_agent,
      content_type: :json, accept: :json
    },
    read_timeout: 10, open_timeout: 10
  }

  @response = Status::V1.server_http.code(code, '1ms', opts)
end

When('I request to set the code {string} and invalid {string} for HTTP') do |code, sleep|
  opts = {
    headers: {
      request_id: SecureRandom.uuid, user_agent: Status.server_config.transport.http.user_agent,
      content_type: :json, accept: :json
    },
    read_timeout: 10, open_timeout: 10
  }

  @response = Status::V1.server_http.code(code, sleep, opts)
end

Then('I should receive a bad request response from HTTP') do
  expect(@response.code).to eq(400)
  expect(@response.body.length).to be > 0
end

Then('I should receive a response with {int} from HTTP') do |code|
  expect(@response.code).to eq(code)
  expect(@response.body.length).to be > 0
end
