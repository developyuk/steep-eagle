"use strict";var precacheConfig=[["index.html","002374434d6650f749ff1a5d22989db8"],["service-worker.js","ecb5f0d5f3a529d48ff283b9fb6e6b7f"],["static/css/app.995fcdf04e7133e6115feade93696d72.css","2f9fe4364bdff91c88fe908fb173d7a5"],["static/js/0.a4420a3f4d50851fc290.js","5b04fe3e0b61d1da93a735ba385baeb6"],["static/js/1.ee1afec965ebf5e95c7f.js","777e8a6215d297ba264f3a2845fd10bd"],["static/js/10.9cda9c4633f366dccd8f.js","100bb517aa7bd4bff71820bbfce18ed4"],["static/js/2.6eeeaea9870211be5fee.js","1e01ec0d40001191069e3a8b435a0131"],["static/js/5.2bcd2f0c29218744cd62.js","2f2482bb8dc11bebb78164c56c4dee28"],["static/js/6.fa0a0f91b4990c98d2bd.js","9333b908729e2aee822d7846c8de5ece"],["static/js/7.3c67d8abd73787404b29.js","42ea926d3d4d1e9d24aefc288aa54ab7"],["static/js/8.8ac63046643ee42478a9.js","231d2c82d2bc18118b169d8204d9cab5"],["static/js/9.62502f6e0350e3ef68a0.js","5a11d86bf24f03222b1d3ea0045c6496"],["static/js/app.3d204e1537852daf7399.js","2c8a3af634df783fbcb939d9d15f70d9"],["static/js/manifest.cbcce8451aecc766e242.js","638d0434070ee0a48ff51278cbf7e121"],["static/js/vendor.eae59248b017ad1ae731.js","9063cf3f48adf7ae966a51a6d249d845"]],cacheName="sw-precache-v3-classbro-"+(self.registration?self.registration.scope:""),ignoreUrlParametersMatching=[/^utm_/],addDirectoryIndex=function(e,t){var a=new URL(e);return"/"===a.pathname.slice(-1)&&(a.pathname+=t),a.toString()},cleanResponse=function(e){return e.redirected?("body"in e?Promise.resolve(e.body):e.blob()).then(function(t){return new Response(t,{headers:e.headers,status:e.status,statusText:e.statusText})}):Promise.resolve(e)},createCacheKey=function(e,t,a,n){var r=new URL(e);return n&&r.pathname.match(n)||(r.search+=(r.search?"&":"")+encodeURIComponent(t)+"="+encodeURIComponent(a)),r.toString()},isPathWhitelisted=function(e,t){if(0===e.length)return!0;var a=new URL(t).pathname;return e.some(function(e){return a.match(e)})},stripIgnoredUrlParameters=function(e,t){var a=new URL(e);return a.hash="",a.search=a.search.slice(1).split("&").map(function(e){return e.split("=")}).filter(function(e){return t.every(function(t){return!t.test(e[0])})}).map(function(e){return e.join("=")}).join("&"),a.toString()},hashParamName="_sw-precache",urlsToCacheKeys=new Map(precacheConfig.map(function(e){var t=e[0],a=e[1],n=new URL(t,self.location),r=createCacheKey(n,hashParamName,a,!1);return[n.toString(),r]}));function setOfCachedUrls(e){return e.keys().then(function(e){return e.map(function(e){return e.url})}).then(function(e){return new Set(e)})}self.addEventListener("install",function(e){e.waitUntil(caches.open(cacheName).then(function(e){return setOfCachedUrls(e).then(function(t){return Promise.all(Array.from(urlsToCacheKeys.values()).map(function(a){if(!t.has(a)){var n=new Request(a,{credentials:"same-origin"});return fetch(n).then(function(t){if(!t.ok)throw new Error("Request for "+a+" returned a response with status "+t.status);return cleanResponse(t).then(function(t){return e.put(a,t)})})}}))})}).then(function(){return self.skipWaiting()}))}),self.addEventListener("activate",function(e){var t=new Set(urlsToCacheKeys.values());e.waitUntil(caches.open(cacheName).then(function(e){return e.keys().then(function(a){return Promise.all(a.map(function(a){if(!t.has(a.url))return e.delete(a)}))})}).then(function(){return self.clients.claim()}))}),self.addEventListener("fetch",function(e){if("GET"===e.request.method){var t,a=stripIgnoredUrlParameters(e.request.url,ignoreUrlParametersMatching),n="index.html";(t=urlsToCacheKeys.has(a))||(a=addDirectoryIndex(a,n),t=urlsToCacheKeys.has(a));0,t&&e.respondWith(caches.open(cacheName).then(function(e){return e.match(urlsToCacheKeys.get(a)).then(function(e){if(e)return e;throw Error("The cached response that was expected is missing.")})}).catch(function(t){return console.warn('Couldn\'t serve response for "%s" from cache: %O',e.request.url,t),fetch(e.request)}))}});