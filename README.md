# Weather Service

## Overview
- A simple weather service that provides current weather information based on geographic coordinates (latitude and longitude). 
- The service uses the National Weather Service (NWS) API to fetch weather data.
- This service has a single API endpoint and has the following signature: GET /weather/{latitude},{longitude}
- The service accepts latitude and longitude as path parameters and returns the current weather information for the specified location with a JSON Object with following attributes:

| Attribute               | Description                                              |  Type  |
|:------------------------|---------------------------------------------------------:|:------:|
| city                    | Name of the city                                         | string |
| state                   | State code                                               | string |
| temperature             | Current temperature                                      | string |
| temperatureUnit         | Unit of temperature (Fahrenheit)                         | string |
| shortForecast           | Short weather forecast description                       | string |
| weatherCharacterization | Characterization of the weather (e.g., Moderate, Severe) | string |



Example: 

``` GET /weather/41.8781,-87.6298 ``` (Coordinates for Chicago, IL)

It returns a JSON Object below:
```json
{
  city: "Chicago",
  state: "IL",
  temperature: "70.0",
  temperatureUnit: "F",
  shortForecast: "Partly Cloudy",
  weatherCharacterization: "Moderate"
}
```


## Prerequisites
Install and run docker desktop

## How to build
```docker build -t weather-service:1.0 .```

## How to run
```docker run -p 3000:3000 -t weather-service:1.0```

## How to call API
- With CURL ```curl -X GET "http://localhost:3000/weather/41.8781,-87.6298"```
- With Browser ```http://localhost:3000/weather/41.8781,-87.6298```

## What code demostrates
- Design using interface, how can we extend it using different implmentation of weather service
- Patterns like cmd, internal and provider 
- Packaging such as weather, errors provider, validation for maintanable and reuseble code.
- Error handling with custom Error objects
- Easily adapt to adding new temparature catogorization and imperial/metric consideration
- Strong containerization with multi-stage builds, effective docker layer caching, non-root containers
- As this code is built and run as a container it is easy to test on local as well as production enviroments
- Containerization enables it to be deployed either on Kubernetes offerings such as EKS, GKE or can be deployed to Serverless platforms such as AWS Lambda, GCP ClouRun enabling cost optimization with the use of services like AWS Fargate or CloudRun

## Comments on the shortcuts
- The service is designed to be simple and straightforward, focusing on fetching and returning weather data.
- It does not include advanced error handling or caching mechanisms, which could be added for production use
- No Unit Tests
- No Retry / Timeout Handling
- No graceful termination with singnal handling
- Code can be more organized.
- No OPEN-API defination and swagger endpoint






