# frozen_string_literal: true

When('I request to set the code {int}') do |code|
  @response = Status::V1.http.code(code, '1ms', Status::V1.http.options)
end

When('I request to set the code {int} without sleep') do |code|
  @response = Status::V1.http.code(code, '', Status::V1.http.options)
end

When('I request to set the code {int} and sleep {string}') do |code, sleep|
  start = Process.clock_gettime(Process::CLOCK_MONOTONIC)
  @response = Status::V1.http.code(code, sleep, Status::V1.http.options)
  @elapsed = Process.clock_gettime(Process::CLOCK_MONOTONIC) - start
end

When('I request to set the code {int} and location {string}') do |code, location|
  @response = Status::V1.http.code_with_location(code, location, Status::V1.http.options)
end

When('I request to set the code {int} and retry after {string}') do |code, retry_after|
  @response = Status::V1.http.code_with_retry_after(code, retry_after, Status::V1.http.options)
end

When('I request to set the invalid code {string}') do |code|
  @response = Status::V1.http.code(code, '1ms', Status::V1.http.options)
end

When('I request to set the code {string} and invalid {string}') do |code, sleep|
  @response = Status::V1.http.code(code, sleep, Status::V1.http.options)
end

Then('I should receive a bad request response') do
  expect(@response.code).to eq(400)
  expect(@response.body.length).to be > 0
end

Then('I should receive a response with {int} and {string}') do |code, body|
  expect(@response.code).to eq(code)
  expect(@response.body.strip).to eq(body)
end

Then('I should receive a location {string}') do |location|
  expect(@response.headers[:location]).to eq(location)
end

Then('I should receive a retry after {string}') do |retry_after|
  expect(@response.headers[:retry_after]).to eq(retry_after)
end

Then('I should receive the response in at least {int} ms') do |time|
  expect(@elapsed * 1000).to be >= time
end

Then('I should receive the response in less than {int} ms') do |time|
  expect(@elapsed * 1000).to be < time
end
