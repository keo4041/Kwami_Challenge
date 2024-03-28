import requests

def test_http_redirect_to_https():
    # Replace the IP address with your EC2 instance's public IP address
    server_ip = "52.91.35.195"

    # Send an HTTP GET request
    response = requests.get(f'http://{server_ip}', allow_redirects=False)

    # Check if the response status code is 301 (Moved Permanently)
    assert response.status_code == 301, "Expected HTTP status code 301, got {}".format(response.status_code)

    # Check if the 'Location' header in the response points to the HTTPS URL
    expected_location = f'https://{server_ip}/'
    assert response.headers['Location'] == expected_location, "Expected Location header to be '{}', got '{}'".format(expected_location, response.headers['Location'])

    print("HTTP to HTTPS redirect test passed!")

if __name__ == '__main__':
    test_http_redirect_to_https()
