Given(/^the system knows about the following tweets:$/) do |tweet|
  tweet.hashes.each do |m|
    Tweet.create(m)
  end
end

When(/^the client requests GET \/tweets$/) do
  header 'accept', 'application/json'
  get '/tweets'
end

Then(/^the response should be JSON:$/) do |json|
  JSON.parse(last_response.body).should == JSON.parse(json)
end
