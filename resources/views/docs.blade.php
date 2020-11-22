<!doctype html>
<!--

  Instructions:

  - Save this file.
  - Replace "USER" with your GitHub username.
  - Replace "REPO" with your GitHub repo name.
  - Replace "Your Project" with your project name.
  - Upload this file (or commit to GitHub Pages).

  Customize as you see fit!

-->
<html>
<head>
  <meta charset='utf-8'>
  <meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1">
  <meta name="viewport" content="width=device-width">

  <title>Meta Grabber</title>

  <!-- Flatdoc -->
  <script src="/js/jquery.min.js"></script>
  <script src='/js/flatdoc/v0.9.0/legacy.js'></script>
  <script src='/js/flatdoc/v0.9.0/flatdoc.js'></script>

  <!-- Flatdoc theme -->
  <link href='/theme-kitech/style.css' rel='stylesheet'>
  <script src='/theme-kitech/script.js'></script>
  
  <!-- Meta -->
  <meta content="Meta Grabber Api" property="og:title">
  <meta content="Meta Grabber gets meta data from a website for creating sharable cards." name="description">

  <!-- Initializer -->
  <script>
    Flatdoc.run({
      fetcher: Flatdoc.file('/docs/Readme.md')
    });
  </script>
</head>
<body role='flatdoc'>

  <div class='header'>
    <div class='left'>
      <h1>Meta Grabber Api</h1>
    </div>
    <div class='right'>
      <!-- GitHub buttons: see http://ghbtns.com -->

    </div>
  </div>

  <div class='content-root'>
    <div class='menubar'>
      <div class='menu section' role='flatdoc-menu'></div>
    </div>
    <div role='flatdoc-content' class='content'></div>
  </div>

</body>
</html>
