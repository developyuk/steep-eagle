"use strict";var precacheConfig=[["index.html","61d9f2a4c2d65af92c9b98d7fa033a5c"],["service-worker.js","a22fc1ca39a42fc9b71c198e46249c1f"],["static/css/app.456062065473aee1692af49c6b54d9f6.css","34281b987bc75b734efcd04d42628675"],["static/js/0.11f0fd0b2f8786f8b19e.js","b732d1353917f1eac7558e80b8a05c18"],["static/js/1.59d79a6fb15540c1aefe.js","329dbc444cab7a358eff7e29e72aa07b"],["static/js/2.0c48793c68896faeb7a6.js","83af1c114a8bbb6e38814f62c1e93116"],["static/js/5.7ccb9f8821c18798ee25.js","031d2de1ddb653af4cc28b8d0b804853"],["static/js/6.7a8d0cb7ee93a561f654.js","d5e1bf9c67df3dd967c0f7d9baee1ea6"],["static/js/7.ac841cea757c8bd9b350.js","30f63349980c7abcf444c951c581ae47"],["static/js/8.2fc521689383965415e2.js","665e386ed476d1ed5ffbce6770436c79"],["static/js/9.4b21558f9f776c4aeff7.js","77188ad34a670cc100618ff994bc10e0"],["static/js/app.ab68c6035fdaf288ad94.js","697d294f6e989b5aaaab47e7e9d67989"],["static/js/manifest.d89e02d38180ab5167f3.js","d05ed951876eeaa4d732ec7f8dfaf0b8"],["static/js/vendor.af383f5a19fe42cb81c8.js","b18a5b4ada672c78856e9f02b05da7bf"]],cacheName="sw-precache-v3-classbro-"+(self.registration?self.registration.scope:""),ignoreUrlParametersMatching=[/^utm_/],addDirectoryIndex=function(e,t){var a=new URL(e);return"/"===a.pathname.slice(-1)&&(a.pathname+=t),a.toString()},cleanResponse=function(e){return e.redirected?("body"in e?Promise.resolve(e.body):e.blob()).then(function(t){return new Response(t,{headers:e.headers,status:e.status,statusText:e.statusText})}):Promise.resolve(e)},createCacheKey=function(e,t,a,n){var r=new URL(e);return n&&r.pathname.match(n)||(r.search+=(r.search?"&":"")+encodeURIComponent(t)+"="+encodeURIComponent(a)),r.toString()},isPathWhitelisted=function(e,t){if(0===e.length)return!0;var a=new URL(t).pathname;return e.some(function(e){return a.match(e)})},stripIgnoredUrlParameters=function(e,t){var a=new URL(e);return a.hash="",a.search=a.search.slice(1).split("&").map(function(e){return e.split("=")}).filter(function(e){return t.every(function(t){return!t.test(e[0])})}).map(function(e){return e.join("=")}).join("&"),a.toString()},hashParamName="_sw-precache",urlsToCacheKeys=new Map(precacheConfig.map(function(e){var t=e[0],a=e[1],n=new URL(t,self.location),r=createCacheKey(n,hashParamName,a,!1);return[n.toString(),r]}));function setOfCachedUrls(e){return e.keys().then(function(e){return e.map(function(e){return e.url})}).then(function(e){return new Set(e)})}self.addEventListener("install",function(e){e.waitUntil(caches.open(cacheName).then(function(e){return setOfCachedUrls(e).then(function(t){return Promise.all(Array.from(urlsToCacheKeys.values()).map(function(a){if(!t.has(a)){var n=new Request(a,{credentials:"same-origin"});return fetch(n).then(function(t){if(!t.ok)throw new Error("Request for "+a+" returned a response with status "+t.status);return cleanResponse(t).then(function(t){return e.put(a,t)})})}}))})}).then(function(){return self.skipWaiting()}))}),self.addEventListener("activate",function(e){var t=new Set(urlsToCacheKeys.values());e.waitUntil(caches.open(cacheName).then(function(e){return e.keys().then(function(a){return Promise.all(a.map(function(a){if(!t.has(a.url))return e.delete(a)}))})}).then(function(){return self.clients.claim()}))}),self.addEventListener("fetch",function(e){if("GET"===e.request.method){var t,a=stripIgnoredUrlParameters(e.request.url,ignoreUrlParametersMatching),n="index.html";(t=urlsToCacheKeys.has(a))||(a=addDirectoryIndex(a,n),t=urlsToCacheKeys.has(a));0,t&&e.respondWith(caches.open(cacheName).then(function(e){return e.match(urlsToCacheKeys.get(a)).then(function(e){if(e)return e;throw Error("The cached response that was expected is missing.")})}).catch(function(t){return console.warn('Couldn\'t serve response for "%s" from cache: %O',e.request.url,t),fetch(e.request)}))}});