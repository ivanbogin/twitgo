require 'mongoid'

Mongoid.load!('./config/mongoid.yml', :development)

# Cleanup database first
Before do
  Tweet.delete_all
end
