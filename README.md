# URLShortner

## Build the Docker Image 
``` docker build -t shortnerapp .``` </br>
``` docker run -dp 8080:8080 shortnerapp``` 

 ## API Spec:
  ### Genrate Shortner URL:
  eg. ` curl --location --request POST 'http://localhost:8080/api/url/short' \
--header 'Content-Type: application/json' \
--data-raw '{
    "long_url":"https://www.test.com"
}' `

 ### TO redirect Actual URL:
eg.`curl --location --request GET 'http://localhost:8080/1Y0p8rG'`

 ### Get All Shortner URL:
eg.`curl --location --request GET 'http://localhost:8080/api/url/short' \
--header 'Content-Type: application/json'`
