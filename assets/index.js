fetch('/api')
    .then(response => response.json())
    .then(data => document.write(data));