window.log = function f(){ log.history = log.history || []; log.history.push(arguments); if(this.console) { var args = arguments, newarr; args.callee = args.callee.caller; newarr = [].slice.call(args); if (typeof console.log === 'object') log.apply.call(console.log, console, newarr); else console.log.apply(console, newarr);}};
(function(a){function b(){}for(var c="assert,count,debug,dir,dirxml,error,exception,group,groupCollapsed,groupEnd,info,log,markTimeline,profile,profileEnd,time,timeEnd,trace,warn".split(","),d;!!(d=c.pop());){a[d]=a[d]||b;}})
(function(){try{console.log();return window.console;}catch(a){return (window.console={});}}());

/* jQuery Content Panel Switcher JS */
var jcps = {};
jcps.fader = function(speed, target, panel) {
	jcps.show(target, panel);
    if (panel == null) {panel = ''};
	$('.switcher' + panel).click(function() {
		var _contentId = '#' + $(this).attr('id') + '-content';
		var _content = $(_contentId).html();
		if (speed == 0) {
			$(target).html(_content);
		}
		else {	
			$(target).fadeToggle(speed, function(){$(this).html(_content);}).fadeToggle(speed);
		}
	});
};
jcps.slider = function(speed, target, panel) { 
	jcps.show(target, panel);
    if (panel == null) {panel = ''};
	$('.switcher' + panel).click(function() {
		var _contentId = '#' + $(this).attr('id') + '-content';
		var _content = $(_contentId).html();
		if (speed == 0) {
			$(target).html(_content);
		}
		else {	
			$(target).slideToggle(speed, function(){$(this).html(_content);}).slideToggle(speed);
		}
	});
};
jcps.show = function (target, panel) {
$('.show').each(function() {
	if (panel == null) {
		$(target).append($(this).html() + '<br/>');
	}
	else {
		var trimPanel = panel.replace('.', '');
		if ($(this).hasClass(trimPanel) == true){$(target).append($(this).html() + '<br/>');}
	}
});
}
