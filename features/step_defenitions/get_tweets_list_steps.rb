require 'httparty'

Given(/^the system knows about the following tweets:$/) do |tweets|
  tweets.hashes.each do |m|
    Tweet.create(m)
  end
end

When(/^the client requests GET \/tweets$/) do
  @last_response = HTTParty.get('http://127.0.0.1:8080/tweets/')
end

Then(/^the response should be JSON:$/) do |json|
  expect(JSON.parse(@last_response.body)).to eq JSON.parse(json)
end
