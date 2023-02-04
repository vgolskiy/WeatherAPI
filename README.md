# Weather API
Work with external weather forecast service "Open Weather API".

Goals are: 
- define weather forecast area with latitude and longitude;
- return the weather conditions in selected area (snow, rain, etc);
- approximate whether it’s hot, cold, or moderate outside;
- define whether there are any weather alerts in selected area;
- show the weather conditions related to the alert if any.

Working with repository:
-	To clone repository use:

	``` git clone https://github.com/vgolskiy/WeatherAPI.git```

-	To build an application, please, run the following command inside the repository root directory:

	``` docker-compose up -d weather_api ```

-	To verify results you can use Postman requests collection (WeatherAPI.postman_collection.json)
