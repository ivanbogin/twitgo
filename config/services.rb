require 'service_manager'

ServiceManager.define_service "start_mongodb" do |s|
  s.start_cmd = "mongod --pidfilepath mongod.pid --dbpath ./data/db"
  s.cwd = Dir.pwd
  s.pid_file = 'mongod.pid'
end

ServiceManager.define_service "start_server" do |s|
  s.host = "127.0.0.1"
  s.port = 8080
  s.start_cmd = "go run server.go -env=test -host=127.0.0.1 -port=8080"
  s.cwd = Dir.pwd
end
