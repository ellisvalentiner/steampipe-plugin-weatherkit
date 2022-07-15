connection "weatherkit" {
    plugin    = "ellisvalentiner/weatherkit"

    # WeatherKit requires authorization using a signed developer token
    # See the Apple Developer documentation
    # https://developer.apple.com/documentation/weatherkitrestapi/request_authentication_for_weatherkit_rest_api

    # The 10-character key identifier from your developer account.
    # key_id = "STJY7HX969"

    # The service identifier.
    # service_id = "com.ellisvalentiner.weatherkit-client"

    # The Apple Developer Program (ADP) team identifier.
    # team_id = "JS4JVS2JBT"

    # Path to your private key for signing the JWT.
    # private_key_path = "~/.auth/AuthKey_STJY7HX969.p8"
}
