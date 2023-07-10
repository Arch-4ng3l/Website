
function send(path) {

	var xhr = new XMLHttpRequest();
	
	xhr.onreadystatechange = function() {
		if (xhr.readyState == XMLHttpRequest.DONE) {
			console.log(xhr.responseText);
		}
	}




	var username = document.getElementsByName('user_name')[0].value;
	var password = document.getElementsByName('password')[0].value;

	var requestdata = {
		user_name: username,
		password: password
	};

	xhr.open('POST', path);
	xhr.setRequestHeader('Content-Type', 'application/json');

	xhr.send(JSON.stringify(requestdata));
	console.log(requestdata);				
				
	console.log(xhr.getAllResponseHeaders());

}

