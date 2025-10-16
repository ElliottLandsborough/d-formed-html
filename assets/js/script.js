$(document).ready(function() {
   jcps.fader(300, '#switcher-panel');
});

$(window).on('load', function() {
	$('div.testimonials.smallbox blockquote').quovolver();
  var d = new Date();
  var hour = d.getHours();
  var greeting;
  if (hour < 12)
  {
    greeting = 'Good morning!';
  }
  else if (hour >=12 && hour <=17)
  {
    greeting = 'Good afternoon!';
  }
  else if (hour > 17)
  {
    greeting = 'Good evening!';
  }
  $('h1#greeting').html(greeting);
});




// must be at bottom - analytics code
  var _gaq = _gaq || [];
  _gaq.push(['_setAccount', 'UA-4886492-1']);
  _gaq.push(['_trackPageview']);

  (function() {
    var ga = document.createElement('script'); ga.type = 'text/javascript'; ga.async = true;
    ga.src = ('https:' == document.location.protocol ? 'https://ssl' : 'http://www') + '.google-analytics.com/ga.js';
    var s = document.getElementsByTagName('script')[0]; s.parentNode.insertBefore(ga, s);
  })();