connection "weatherkit" {
    plugin    = "ellisvalentiner/weatherkit"

    # WeatherKit requires authorization using a signed developer token
    # You must either provide the information to generate a signed JSON web token (JWT) or supply a pre-generated JWT.
    # See the Apple Developer documentation
    # https://developer.apple.com/documentation/weatherkitrestapi/request_authentication_for_weatherkit_rest_api

    # Option 1: Generate an JWT
    # The 10-character key identifier from your developer account.
    # key_id = "STJY7HX969"

    # The service identifier.
    # service_id = "com.ellisvalentiner.weatherkit-client"

    # The Apple Developer Program (ADP) team identifier.
    # team_id = "JS4JVS2JBT"

    # Path to your private key for signing the JWT.
    # private_key_path = "~/.auth/AuthKey_STJY7HX969.p8"

    # Option 2: Use a pre-generated JWT
    # token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
}
