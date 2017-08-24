# GOLANG Weather Project

This is a simple go application, which combines two API services into one. 
GoogleAPI (http://maps.googleapis.com/maps/api/geocode/json?address=ny) returns latitude and longitude of the given address
Darksky API (https://api.darksky.net/forecast/) returns current temperature for the given latitude and longitude

The application takes user input in a JSON format and send it to GoogleAPI, which returns latitude and longitude for that address. Then the application sends latitude and longitude to the Darksky API, which returns the current temperature for that location. Then the searched address, client IP address and current time is store the database.

# Usage
curl -d '{"address":"NY"}' -X POST http://localhost:8081/v1/weather 

curl -X GET http://localhost:8081/v1/weather-log
