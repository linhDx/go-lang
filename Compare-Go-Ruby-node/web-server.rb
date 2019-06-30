require "webrick"

class MyServlet < WEBrick::HTTPServlet::AbstractServlet
    def do_POST (request, response)
            response.status = 200
            response.body = "Hello world!"
    end
end

server = WEBrick::HTTPServer.new(:Port => 8000)

server.mount "/", MyServlet

server.start