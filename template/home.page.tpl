
{{template "base" .}}

{{define "content"}}

<body>
    <div class="container">
        
        <form action="" method="get">
            <label for="website">Write the ticker of a Stock</label>
            <input type="text" name="website" id="website">
            <input type="submit" value="Submit">
        </form>

    </div>

   
</body>


{{end}}