package gookies

func CookieParser(cookies_raw string) *map[string]string {

	// map to hold 'name' and 'value' of all cookies
	cookies := make(map[string]string)

	// variable to know where we are in the loop
	var status string = "name"

	// variable to build up the name of cookie letter by letter
	var tmp_name_string string = ""

	// variable to hold the completed name of cookie
	var name string

	// variable to build up the value of cookie letter by letter
	var tmp_value_string string = ""

	// variable to hold the completed value of cookie
	var value string

	// loop through HTTP cookie string
	for i := 0; i < len(cookies_raw); i++ {

		if string(cookies_raw[i]) == ";" && status == "value" {
			value = tmp_value_string
			tmp_value_string = ""
			status = "next"
			cookies[name] = value
			name = ""
			value = ""
		} else if string(cookies_raw[i]) == " " && status == "next" {
			status = "name"
		} else if string(cookies_raw[i]) != "=" && status == "name" {
			tmp_name_string += string(cookies_raw[i])
		} else if string(cookies_raw[i]) == "=" && status == "name" {
			name = tmp_name_string
			tmp_name_string = ""
			status = "value"
		} else if string(cookies_raw[i]) != "=" && status == "value" {
			tmp_value_string += string(cookies_raw[i])
		}
	}

	value = tmp_value_string
	cookies[name] = value

	return &cookies
}
