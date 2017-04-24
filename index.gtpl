<html>
    <head>
        <meta charset="utf-8">
        <meta http-equiv="x-ua-compatible" content="ie=edge">
        <title>Welcome to Sama</title>
        <meta name="description" content="">
        <meta name="viewport" content="initial-scale=1">
        <link rel="stylesheet" href="css/normalize.css">
        <link rel="stylesheet" href="css/mystyle.css">
        <script src="js/jquery-1.12.0.min.js"></script>
        <script src="js/jquery-ui.js"></script>
        <script src="js/plugins.js"></script>
        <script src="js/main.js"></script>
        <script src="js/css-pop.js"></script>

    </head>
    <body >

                  <div>
                <button id="btn1" onclick="window.location.href='./signin.wifi'">Connect</button>
        </div><!--end of container-->

        <div >
           <div id="cbox"></div>

            <div id="blanket" style="display:none;"></div>
            <div id="popUpDiv" style="display:none;">
                <a href="#" onclick="popup('popUpDiv')" >Close</a>
                <div id="english">

                </div>
                <div id="arabic" >


                </div>
            </div>

           <div id="mytext">
              <p>I accept <span><a href="#" onclick="popup('popUpDiv')">terms and conditions</a></span></p>
           </div>
        </div>






    </body>
</html>
