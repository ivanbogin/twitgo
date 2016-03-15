require 'service_manager'

ServiceManager.define_service 'mongo' do |s|
  s.host = '127.0.0.1'
  s.port = 27017
  s.start_cmd = 'mongod --dbpath ./data/db --bind_ip 127.0.0.1 --port 27017'
  s.cwd = Dir.pwd
end

ServiceManager.define_service 'twitgo' do |s|
  s.host = '127.0.0.1'
  s.port = 8080
  s.start_cmd = 'go run server.go -env=test -host=127.0.0.1 -port=8080'
  s.cwd = Dir.pwd
end
