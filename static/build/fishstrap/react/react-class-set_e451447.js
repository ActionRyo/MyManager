define("fishstrap/react/react-class-set.js",function(r,e,s){"use strict";function t(){for(var r="",e=0;e<arguments.length;e++){var s=arguments[e];if(s){var a=typeof s;if("string"===a||"number"===a)r+=" "+s;else if(Array.isArray(s))r+=" "+t.apply(null,s);else if("object"===a)for(var n in s)s.hasOwnProperty(n)&&s[n]&&(r+=" "+n)}}return r.substr(1)}s.exports=t});