# frozen_string_literal: true

When('the system requests the {string}') do |name|
  @response = Nonnative.observability.send(name, Status.http_options)
end

Then('the system should respond with a healthy status') do
  expect(@response.code).to eq(200)
  expect(@response.body.strip).to eq('SERVING')
end

Then('the system should respond with metrics') do
  expect(@response.code).to eq(200)
  expect(@response.body).to include('go_info')
end
