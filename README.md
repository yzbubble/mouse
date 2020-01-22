# mouse

可以指定任意目录起个服务，将 markdown 渲染成 html。

## 快速开始

```bash
# 不指定参数，默认为命令程序所在目录，http 地址为 ":8080"
mouse

# 指定目录 
mouse -r /path/to/root
# 指定目录的另一种简写方式
mouse /path/to/root

# 指定 http 地址，默认为 ":8080"
mouse -a :8181

# 查看版本号
mouse -v

# 查看帮助文档
mouse -h
```

## 自定义模板

```bash
# 指定自定义模板路径
mouse -m /path/to/template

# 'default' 具有特殊含义，意为使用内嵌默认模板
mouse -m default
```

默认内嵌模板为：

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width,initial-scale=1.0, minimum-scale=1.0, maximum-scale=1.0, user-scalable=no"/>
    <title>{{.FileName}}</title>
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/github-markdown-css/3.0.1/github-markdown.min.css" integrity="sha256-HbgiGHMLxHZ3kkAiixyvnaaZFNjNWLYKD/QG6PWaQPc=" crossorigin="anonymous" />
    <link rel="stylesheet" href="//cdn.jsdelivr.net/gh/highlightjs/cdn-release@9.17.1/build/styles/default.min.css">
    <style>
        .markdown-body {
            box-sizing: border-box;
            min-width: 200px;
            max-width: 980px;
            margin: 0 auto;
            padding: 45px;
        }
        @media (max-width: 767px) {
            .markdown-body {
                padding: 15px;
            }
        }
    </style>
</head>
<body>
    <article class="markdown-body">{{.Content}}</article>
    <script src="//cdn.jsdelivr.net/gh/highlightjs/cdn-release@9.17.1/build/highlight.min.js"></script>
    <script>hljs.initHighlightingOnLoad();</script>
</body>
</html>
```

