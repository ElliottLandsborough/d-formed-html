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

    $(function() {
        var bgImages = [
            'bg.jpg', 'bg2.jpg', 'bg3.jpg', 'bg4.jpg', 'bg5.jpg', 'bg6.jpg', 'bg7.jpg'
        ];
        var i = Math.floor(Math.random() * bgImages.length);
        var selectedBg = bgImages[i];

        $('body').css({
            'background': "url('img/" + selectedBg + "') no-repeat",
            'background-color': '#111111',
            'background-position': 'center top',
            'background-attachment': 'fixed'
        });
    });

    (function() {
        var $contactDiv = $("#contact");
        if ($contactDiv.length) {
            // Obfuscated parts
            const a = 'hello';
            const b = 'd-formed';
            const c = 'net';
            const d = '@';
            const e = '.';
            
            // Construct email
            const email = a + d + b + e + c;
            
            // Create and insert link
            const link = document.createElement('a');
            link.href = 'mai' + 'lto:' + email;
            link.textContent = email;
            
            // Add to page after small delay
            setTimeout(function() {
                document.getElementById('contact').appendChild(link);
            }, 100);
        }
    })();
});
