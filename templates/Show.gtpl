{{ define "Show" }}
<!DOCTYPE html>
<head>
	{{ template "head" }}
</head>

<body>
	You typed: {{ .FName }}
	
<br />
	<a href="/">Back</a>
</body>
</html>
{{ end }}
