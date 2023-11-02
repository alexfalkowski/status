# frozen_string_literal: true

When('I request to set the code {int} with HTTP') do |code|
  headers = { request_id: SecureRandom.uuid, user_agent: Status.server_config['transport']['http']['user_agent'] }

  @response = Status::V1.server_http.code(code, headers)
end

When('I request to set the invalid code {string} with HTTP') do |code|
  headers = { request_id: SecureRandom.uuid, user_agent: Status.server_config['transport']['http']['user_agent'] }

  @response = Status::V1.server_http.code(code, headers)
end

Then('I should receive a response with {int} from HTTP') do |code|
  expect(@response.code).to eq(code)
  expect(@response.body.length).to be > 0
end

Then('I should receive a internal error response from HTTP') do
  expect(@response.code).to eq(500)
  expect(@response.body.length).to be > 0
end
