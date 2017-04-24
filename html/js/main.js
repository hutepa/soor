$(document).ready(function(){
	var d = new Date();
	var n = d.getHours();
	if (n > 18 || n < 6)
	  // If time is after 7PM or before 6AM, apply night theme to ‘body’
	  document.body.className = "night";
	
	else
	  // Else use ‘day’ theme
	  document.body.className = "day";

    $('#english').load("terms.txt").html();
    $('#arabic').load("arabic.html").html();
});
