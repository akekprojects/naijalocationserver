# Naija Location Server
![code coverage badge](https://github.com/muhammadolammi/naijalocationserver/actions/workflows/ci.yml/badge.svg)

This is the Repository for http://naijalocationserver.com/, an api endpoint that provides the necesary infomation about Nigeria locations, ranging from States, Capitals, Local Government Areas and more to come.

## Motivation Behind NaijaLocationSever
The development of NaijaLocationSever is guided by four key principles:

* Correctness: This project is personally managed by me, and information will be consistently updated to reflect the latest data.

* Simplicity: The process is straightforward; all that's required is to send an HTTP GET request to one of our endpoints. (Details on how to do this will be explained in the next section.)

* Accessibility: The server's primary goal is to make information easily accessible for use. We handle the heavy lifting in terms of data sourcing, and all you need to do is use HttpClient.Get() on your desired endpoint.

* Performance: Being a Go application, I chose this language for its exceptional performance and speed. If you grasp concepts like memory management and concurrency, you'll appreciate the efficiency of this server.

## 🚀 Quick Start

Navigate to http://naijalocationserver.com/api/states to access a list of all Nigerian states and their capitals in "application/json" format.


## Examples
### Note: All my examples will be in Golang, You can use any http Client From any Language
* Get All States in Nigeria

```
func getFromEndpoint() interface{} {
	httpClient := http.Client{
		Timeout: time.Second * 10, // Maximum of 10 secs

	}
	resp, err := httpClient.Get(" http://naijalocationserver.com/api/states")
	if err != nil {
		log.Printf("error getting resp from endpoint: %v", err)
		return "error getting resp from endpoint"
	}

	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)

	var bodyData interface{}
	//this is a bad practice, you should create a struct for body data to be marshalled into that correlate with the json data from endpoint. To manage type better.
	err = decoder.Decode(&bodyData)
	if err != nil {
		log.Printf("error creating decoder from resp.Body(): %v", err)
		return "error creating decoder from resp.Body()"
	}

	return bodyData

}

```
### The code here will return the following json
```
```