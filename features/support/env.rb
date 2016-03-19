# Cleanup database first
Before do
  client = Mongo::Client.new(['127.0.0.1:27017'], :database => 'tweeter_test')
  client[:tweets].drop
end
