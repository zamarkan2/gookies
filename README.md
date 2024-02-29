# Gookies - HTTP Cookie Parser
This module assumes that cookies are retrieved from the HTTP request header as a string, with the format:

Cookie:[fruit=banana; color=red; spice=curry; city=memphis]

This module has one function that receives a string parameter and returns a pointer to a map.