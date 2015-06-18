{{%define "*"%}}<!-- (c) webjinn.ru -->
<!DOCTYPE html>
<html>
<head lang="en">
    <meta charset="UTF-8">
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <title>{{%.MetaTitle%}}</title>
    <meta name="author" content="{{%.MetaAuthor%}}"/>
    <meta name="copyright" content="{{%.MetaCopyright%}}"/>
    <meta name="description" content="{{%.MetaDescription%}}"/>
    <meta name="keywords" content="{{%.MetaKeywords%}}"/>

    <meta name="viewport" content="width=device-width, initial-scale=1">
    <link rel='shortcut icon' type='image/x-icon' href='//{{%.Host%}}/favicon.ico'/>
    <link rel='icon' type='image/png' href='//{{%.Host%}}/favicon.png'/>

    <link rel="stylesheet" href="/css/bootstrap.min.css">
    <link rel="stylesheet" href="/css/style.css">

    <script src="/js/bootstrap.min.js"></script>
    <script src="/js/angular.min.js"></script>
    <script src="/js/app.js"></script>
    <style>
        #center-block {width: 80%; text-align: center; margin: 0 auto; padding-top: 15%;}
    </style>
</head>
<body>
    <div id="center-block">
        <h1>{{% .content %}}</h1>
    </div>
</body>
<body>
    
</body>

</html>
{{%end%}}