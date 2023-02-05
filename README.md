# Weather API
Work with external weather forecast service "Open Weather API".

The goals are:
- define weather forecast area with latitude and longitude;
- return the weather conditions in the selected area (snow, rain, etc.);
- approximate whether it’s hot, cold, or moderate outside;
- define whether there are any weather alerts in the selected area;
- show the weather conditions related to the alert, if any.

## Work on a repository:
-	To clone repository use:

	``` git clone https://github.com/vgolskiy/WeatherAPI.git```

-	To build an application, please, run the following command inside the repository root directory:

	``` docker-compose up -d weather_api ```

-	To verify results you can use Postman requests collection (WeatherAPI.postman_collection.json) or a command below:
     
    ``` curl '127.0.0.1:2000?lon=-170.37&lat=63.46' ```

## Weather API contract

**Forecast object**
```
{
  	weather_conditions: []string
	feels_like: string
	have_alerts: bool
	weather_alert (optional): []weather_alert_object
}
```

Possible values list for `weather_conditions`:
* Thunderstorm
* Drizzle
* Rain
* Snow
* Clear
* Clouds
* Mist
* Smoke
* Haze
* Dust
* Fog
* Sand
* Dust
* Ash
* Squall
* Tornado

Possible values list for `feels_like`:
* hot - over 30C/303K/86F,
* moderate - over -5C/268K/23F up to 30C/303K/86F,
* cold - up to -5C/268K/23F.

**Weather Alert object**
```
{
  	name: string
	weather_conditions: []string
}
```

**GET /**
____
Returns weather forecast for the area selected by latitude and longitude.
- **URL Parameters**
  *Required:*
  - `lat=[float(2)]`
  - `lon=[float(2)]`
- **Headers**
  - Content-Type: application/json
- **Success Response:**
  - **Code:** 200  
    **Content:**
```
{
  weather_conditions: [
           {<string>},
           {<string>},
           ...
         ],
  feels_like: {<string>},
  have_alerts: {<boolean>},
  weather_alert: [
           {<weather_alert_object>},
           {<weather_alert_object>},
           ...
         ]
}
```
- **Error Response:**
  - **Code:** *Depends on type of error*  
    **Content:** `{ message : "external server error" }` 
