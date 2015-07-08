{{ define "Index" }}
<!DOCTYPE html>
<head>
	{{ template "head" }}
</head>

<body>
	This is a form
	<form method="POST" action="/show">
		<label> What is your name?</label>
		<input type="text" name="name" />
		<input type="submit" />
	</form>
</body>
</html>
{{ end }}
