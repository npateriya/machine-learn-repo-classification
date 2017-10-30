# cisco-err-pages
Customized error pages for Cisco Web Security Appliance (Cisco WSA).

Creates minified html pages for Cisco WSA without external dependiences. Css files are minified, unused styles are removed and result injected in HTML. All external images are convert to base64 and injected into HTML page. The final HTML is minidied.

### Sources
`src/`

### Result
`build/index.html`

## Dependencies
- Node.js

## Tested on
- Ubuntu 17.04
- Node.js 8.2.1
- NPM 5.3.0

## Initialization
`make init`

## Build project
`make build`

## Example

### Source *./src/index.html*
```html
<!DOCTYPE html>
<html>

<head>
  <title>Служба по ИТ, АО "КОНАР"</title>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <!-- shortcut::src/img/favicon.ico -->

  <link rel="stylesheet" href="./css/bootstrap.min.css">
  <link rel="stylesheet" href="./css/style.css">
</head>

<body>
  <div class="container">
    <p class="text-center">
      <img src="img/logo.gif" class="img-fluid" alt="АО КОНАР" />
    </p>
    <hr/>
    <div class="row">
      <div class="col-md-2">
        <p class="text-center">
          <img src="img/man.jpg" class="img-fluid" alt="Warning" />
        </p>
      </div>
      <div class="col-md-10">
        <p>К сожалению, Вам не разрешено подключение к локальной сети ПГ КОНАР.</p>
        <p>Это могло произойти по следующим причинам:</p>
        <ul>
          <li>в текущем местоположении вашему компьютеру или учетной записи запрещено подключаться к сети;</li>
          <li>подключенный к сети компьютер не является корпоративным;</li>
          <li>в работе сети в настоящий момент имеются проблемы.</li>
        </ul>
        <p>Вы можете обратиться в техническую поддержку по телефону <em>+7 (351) 216-81-11</em></p>
      </div>
    </div>
    <hr/>
    <p class="text-center">
      <a href="mailto:support@konar.ru">support@konar.ru</a><br/>
      <em>+7 (351) 216-81-11</em>
    </p>
  </div>
</body>

</html>
```

### No minified result
```html
<!DOCTYPE html>
<html>

  <head>
    <title>Служба по ИТ, АО "КОНАР"</title>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width,initial-scale=1">
    <style type="text/css">
      hr,
      p,
      ul {
        margin-bottom: 1rem
      }

      .col-md-10,
      .col-md-2,
      .container {
        position: relative;
        padding-right: 15px;
        padding-left: 15px
      }

      html {
        font-family: sans-serif;
        line-height: 1.15;
        -ms-text-size-adjust: 100%;
        -webkit-text-size-adjust: 100%;
        -webkit-box-sizing: border-box;
        box-sizing: border-box;
        -ms-overflow-style: scrollbar;
        -webkit-tap-highlight-color: transparent
      }

      hr {
        -webkit-box-sizing: content-box;
        box-sizing: content-box;
        height: 0;
        overflow: visible;
        margin-top: 1rem;
        border: 0;
        border-top: 1px solid rgba(0, 0, 0, .1)
      }

      a:active,
      a:hover {
        outline-width: 0
      }

      img {
        border-style: none;
        vertical-align: middle
      }

      ::-webkit-file-upload-button {
        -webkit-appearance: button;
        font: inherit
      }

      @media print {
        *,
        ::after,
        ::before,
        div::first-letter,
        div::first-line,
        li::first-letter,
        li::first-line,
        p::first-letter,
        p::first-line {
          text-shadow: none!important;
          -webkit-box-shadow: none!important;
          box-shadow: none!important
        }
        a,
        a:visited {
          text-decoration: underline
        }
        img {
          page-break-inside: avoid
        }
        p {
          orphans: 3;
          widows: 3
        }
      }

      *,
      ::after,
      ::before {
        -webkit-box-sizing: inherit;
        box-sizing: inherit
      }

      @-ms-viewport {
        width: device-width
      }

      body {
        margin: 0;
        font-family: -apple-system, system-ui, BlinkMacSystemFont, "Segoe UI", Roboto, "Helvetica Neue", Arial, sans-serif;
        font-size: 1rem;
        font-weight: 400;
        line-height: 1.5;
        color: #292b2c;
        background-color: #fff;
        padding-top: 150px
      }

      p,
      ul {
        margin-top: 0
      }

      a {
        background-color: transparent;
        -webkit-text-decoration-skip: objects;
        color: #0275d8;
        text-decoration: none;
        -ms-touch-action: manipulation;
        touch-action: manipulation
      }

      a:focus,
      a:hover {
        color: #014c8c;
        text-decoration: underline
      }

      .img-fluid {
        max-width: 100%;
        height: auto
      }

      .container {
        margin-left: auto;
        margin-right: auto
      }

      @media (min-width:576px) {
        .container {
          padding-right: 15px;
          padding-left: 15px;
          width: 540px;
          max-width: 100%
        }
        .row {
          margin-right: -15px;
          margin-left: -15px
        }
      }

      @media (min-width:768px) {
        .container {
          padding-right: 15px;
          padding-left: 15px;
          width: 720px;
          max-width: 100%
        }
        .row {
          margin-right: -15px;
          margin-left: -15px
        }
      }

      .row {
        display: -webkit-box;
        display: -webkit-flex;
        display: -ms-flexbox;
        display: flex;
        -webkit-flex-wrap: wrap;
        -ms-flex-wrap: wrap;
        flex-wrap: wrap;
        margin-right: -15px;
        margin-left: -15px
      }

      @media (min-width:992px) {
        .container {
          padding-right: 15px;
          padding-left: 15px;
          width: 960px;
          max-width: 100%
        }
        .row {
          margin-right: -15px;
          margin-left: -15px
        }
        .col-md-10,
        .col-md-2 {
          padding-right: 15px;
          padding-left: 15px
        }
      }

      @media (min-width:1200px) {
        .container {
          padding-right: 15px;
          padding-left: 15px;
          width: 1140px;
          max-width: 100%
        }
        .row {
          margin-right: -15px;
          margin-left: -15px
        }
        .col-md-10,
        .col-md-2 {
          padding-right: 15px;
          padding-left: 15px
        }
      }

      .col-md-10,
      .col-md-2 {
        width: 100%;
        min-height: 1px
      }

      @media (min-width:576px) {
        .col-md-10,
        .col-md-2 {
          padding-right: 15px;
          padding-left: 15px
        }
      }

      @media (min-width:768px) {
        .col-md-10,
        .col-md-2 {
          padding-right: 15px;
          padding-left: 15px;
          -webkit-box-flex: 0;
          -webkit-flex: 0 0 16.666667%;
          -ms-flex: 0 0 16.666667%;
          flex: 0 0 16.666667%;
          max-width: 16.666667%
        }
        .col-md-10 {
          -webkit-flex: 0 0 83.333333%;
          -ms-flex: 0 0 83.333333%;
          flex: 0 0 83.333333%;
          max-width: 83.333333%
        }
      }

      .form-control::placeholder {
        color: #636c72;
        opacity: 1
      }

      .text-center {
        text-align: center!important
      }

    </style>
  </head>

  <body>
    <div class="container">
      <p class="text-center"><img src="data:image/gif;base64,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"
        class="img-fluid" alt="АО КОНАР"></p>
      <hr>
      <div class="row">
        <div class="col-md-2">
          <p class="text-center"><img src="data:image/jpeg;base64,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"
            class="img-fluid" alt="Warning"></p>
        </div>
        <div class="col-md-10">
          <p>К сожалению, Вам не разрешено подключение к локальной сети ПГ КОНАР.</p>
          <p>Это могло произойти по следующим причинам:</p>
          <ul>
            <li>в текущем местоположении вашему компьютеру или учетной записи запрещено подключаться к сети;</li>
            <li>подключенный к сети компьютер не является корпоративным;</li>
            <li>в работе сети в настоящий момент имеются проблемы.</li>
          </ul>
          <p>Вы можете обратиться в техническую поддержку по телефону <em>+7 (351) 216-81-11</em></p>
        </div>
      </div>
      <hr>
      <p class="text-center"><a href="mailto:support@konar.ru">support@konar.ru</a>
        <br><em>+7 (351) 216-81-11</em></p>
    </div>
  </body>

</html>
```

### Minified result *./build/index.html*
```html
<!DOCTYPE html><html><head><title>Служба по ИТ, АО "КОНАР"</title><meta charset="UTF-8"><meta name="viewport" content="width=device-width,initial-scale=1"><style type="text/css">hr,p,ul{margin-bottom:1rem}.col-md-10,.col-md-2,.container{position:relative;padding-right:15px;padding-left:15px}html{font-family:sans-serif;line-height:1.15;-ms-text-size-adjust:100%;-webkit-text-size-adjust:100%;-webkit-box-sizing:border-box;box-sizing:border-box;-ms-overflow-style:scrollbar;-webkit-tap-highlight-color:transparent}hr{-webkit-box-sizing:content-box;box-sizing:content-box;height:0;overflow:visible;margin-top:1rem;border:0;border-top:1px solid rgba(0,0,0,.1)}a:active,a:hover{outline-width:0}img{border-style:none;vertical-align:middle}::-webkit-file-upload-button{-webkit-appearance:button;font:inherit}@media print{*,::after,::before,div::first-letter,div::first-line,li::first-letter,li::first-line,p::first-letter,p::first-line{text-shadow:none!important;-webkit-box-shadow:none!important;box-shadow:none!important}a,a:visited{text-decoration:underline}img{page-break-inside:avoid}p{orphans:3;widows:3}}*,::after,::before{-webkit-box-sizing:inherit;box-sizing:inherit}@-ms-viewport{width:device-width}body{margin:0;font-family:-apple-system,system-ui,BlinkMacSystemFont,"Segoe UI",Roboto,"Helvetica Neue",Arial,sans-serif;font-size:1rem;font-weight:400;line-height:1.5;color:#292b2c;background-color:#fff;padding-top:150px}p,ul{margin-top:0}a{background-color:transparent;-webkit-text-decoration-skip:objects;color:#0275d8;text-decoration:none;-ms-touch-action:manipulation;touch-action:manipulation}a:focus,a:hover{color:#014c8c;text-decoration:underline}.img-fluid{max-width:100%;height:auto}.container{margin-left:auto;margin-right:auto}@media (min-width:576px){.container{padding-right:15px;padding-left:15px;width:540px;max-width:100%}.row{margin-right:-15px;margin-left:-15px}}@media (min-width:768px){.container{padding-right:15px;padding-left:15px;width:720px;max-width:100%}.row{margin-right:-15px;margin-left:-15px}}.row{display:-webkit-box;display:-webkit-flex;display:-ms-flexbox;display:flex;-webkit-flex-wrap:wrap;-ms-flex-wrap:wrap;flex-wrap:wrap;margin-right:-15px;margin-left:-15px}@media (min-width:992px){.container{padding-right:15px;padding-left:15px;width:960px;max-width:100%}.row{margin-right:-15px;margin-left:-15px}.col-md-10,.col-md-2{padding-right:15px;padding-left:15px}}@media (min-width:1200px){.container{padding-right:15px;padding-left:15px;width:1140px;max-width:100%}.row{margin-right:-15px;margin-left:-15px}.col-md-10,.col-md-2{padding-right:15px;padding-left:15px}}.col-md-10,.col-md-2{width:100%;min-height:1px}@media (min-width:576px){.col-md-10,.col-md-2{padding-right:15px;padding-left:15px}}@media (min-width:768px){.col-md-10,.col-md-2{padding-right:15px;padding-left:15px;-webkit-box-flex:0;-webkit-flex:0 0 16.666667%;-ms-flex:0 0 16.666667%;flex:0 0 16.666667%;max-width:16.666667%}.col-md-10{-webkit-flex:0 0 83.333333%;-ms-flex:0 0 83.333333%;flex:0 0 83.333333%;max-width:83.333333%}}.form-control::placeholder{color:#636c72;opacity:1}.text-center{text-align:center!important}</style></head><body><div class="container"><p class="text-center"><img src="data:image/gif;base64,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" class="img-fluid" alt="АО КОНАР"></p><hr><div class="row"><div class="col-md-2"><p class="text-center"><img src="data:image/jpeg;base64,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" class="img-fluid" alt="Warning"></p></div><div class="col-md-10"><p>К сожалению, Вам не разрешено подключение к локальной сети ПГ КОНАР.</p><p>Это могло произойти по следующим причинам:</p><ul><li>в текущем местоположении вашему компьютеру или учетной записи запрещено подключаться к сети;</li><li>подключенный к сети компьютер не является корпоративным;</li><li>в работе сети в настоящий момент имеются проблемы.</li></ul><p>Вы можете обратиться в техническую поддержку по телефону <em>+7 (351) 216-81-11</em></p></div></div><hr><p class="text-center"><a href="mailto:support@konar.ru">support@konar.ru</a><br><em>+7 (351) 216-81-11</em></p></div></body></html>
```
