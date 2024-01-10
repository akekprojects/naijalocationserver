# Naija Location Server
![code coverage badge](https://github.com/muhammadolammi/naijalocationserver/actions/workflows/ci.yml/badge.svg)

This is the Repository for http://naijalocationserver.com/, an api endpoint that provides the necesary infomation about Nigeria locations, ranging from States, Capitals, Local Government Areas and more to come.

# Motivation Behind NaijaLocationSever
The development of NaijaLocationSever is guided by four key principles:

* Correctness: This project is personally managed by me, and information will be consistently updated to reflect the latest data.

* Simplicity: The process is straightforward; all that's required is to send an HTTP GET request to one of our endpoints. (Details on how to do this will be explained in the next section.)

* Accessibility: The server's primary goal is to make information easily accessible for use. We handle the heavy lifting in terms of data sourcing, and all you need to do is use HttpClient.Get() on your desired endpoint.

* Performance: Being a Go application, I chose this language for its exceptional performance and speed. If you grasp concepts like memory management and concurrency, you'll appreciate the efficiency of this server.

## 🚀 Quick Start

Navigate to http://naijalocationserver.com/api/states to access a list of all Nigerian states and their capitals in "application/json" format.


