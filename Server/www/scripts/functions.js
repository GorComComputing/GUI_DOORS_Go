function  HttpRequest(s, body) {
  	//console.log('Log from JS: ['+s+']');
	var response = "";
	var xhttp = new XMLHttpRequest();
    xhttp.onreadystatechange = function() {
    	if (this.readyState == 4 && this.status == 200) {  
        	console.log(xhttp.responseText);
            response = xhttp.responseText;
        } else {
        	response = "Response FAIL"
        }
      };
      //xhttp.withCredentials = true;
      xhttp.open("POST", s, false);
      //xhttp.setRequestHeader('Content-Type', 'application/x-www-form-urlencode');
	  xhttp.send(body);
     
      return {response: response};
}




		


		

