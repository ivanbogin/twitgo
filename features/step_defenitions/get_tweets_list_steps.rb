require 'httparty'
require 'mongo'

Given(/^the system knows about the following tweets:$/) do |tweets|
  @mongo = Mongo::Client.new(['127.0.0.1:27017'], :database => 'tweeter_test')
  tweets.hashes.each do |m|
    @mongo[:tweets].insert_one({
      body: m[:body],
      created_at: DateTime.parse(m[:created_at])
    })
  end
end

When(/^the client requests GET \/tweets$/) do
  @last_response = HTTParty.get('http://127.0.0.1:8080/tweets/')
end

Then(/^the response should be JSON:$/) do |json|
  expect(JSON.parse(@last_response.body)).to eq JSON.parse(json)
end
