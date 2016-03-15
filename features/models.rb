class Tweet
  include Mongoid::Document
  field :body, type: String
  field :created_at, type: DateTime
end
