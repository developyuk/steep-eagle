"use strict";var precacheConfig=[["index.html","948d47ce2da26289a611d1d712a9915d"],["service-worker.js","4c171fec88a45e2480a6b26c62e66c80"],["static/css/app.456062065473aee1692af49c6b54d9f6.css","34281b987bc75b734efcd04d42628675"],["static/js/0.30eb9857a34be0ae410f.js","a2ef094fce112ee039197a2a2e5857e3"],["static/js/1.17464cf1d86affa19044.js","ce6bd670ad928542c46514c742fb83e8"],["static/js/2.fd070babb4bb509ea40e.js","fc1132cf3f3f5f75978fe59bfae8b719"],["static/js/5.92011f85fab918d175ff.js","24163bb6a449153f88705a7e1e2feeb1"],["static/js/6.51385943c5508081434c.js","7b85afb29f2b0daf0cfc451d7787da91"],["static/js/7.9a8e2d7127480bb94962.js","bc09845f0db39e48ac22fcb8acebf67e"],["static/js/8.a07e0e9bbb6ad88db635.js","e1cda7af65af3d758dbbe472b69fb296"],["static/js/9.23c6d3c37d9cea9f523c.js","429f3adc39718e4c1dce44d9437bb650"],["static/js/app.b801478ba89b3188fa70.js","4ae829981fe1ac74d75b8bd018131277"],["static/js/manifest.6b21a124fad164702d53.js","ca49052924e91773f45aa307c6454476"],["static/js/vendor.af383f5a19fe42cb81c8.js","b18a5b4ada672c78856e9f02b05da7bf"]],cacheName="sw-precache-v3-classbro-"+(self.registration?self.registration.scope:""),ignoreUrlParametersMatching=[/^utm_/],addDirectoryIndex=function(e,t){var a=new URL(e);return"/"===a.pathname.slice(-1)&&(a.pathname+=t),a.toString()},cleanResponse=function(e){return e.redirected?("body"in e?Promise.resolve(e.body):e.blob()).then(function(t){return new Response(t,{headers:e.headers,status:e.status,statusText:e.statusText})}):Promise.resolve(e)},createCacheKey=function(e,t,a,n){var r=new URL(e);return n&&r.pathname.match(n)||(r.search+=(r.search?"&":"")+encodeURIComponent(t)+"="+encodeURIComponent(a)),r.toString()},isPathWhitelisted=function(e,t){if(0===e.length)return!0;var a=new URL(t).pathname;return e.some(function(e){return a.match(e)})},stripIgnoredUrlParameters=function(e,t){var a=new URL(e);return a.hash="",a.search=a.search.slice(1).split("&").map(function(e){return e.split("=")}).filter(function(e){return t.every(function(t){return!t.test(e[0])})}).map(function(e){return e.join("=")}).join("&"),a.toString()},hashParamName="_sw-precache",urlsToCacheKeys=new Map(precacheConfig.map(function(e){var t=e[0],a=e[1],n=new URL(t,self.location),r=createCacheKey(n,hashParamName,a,!1);return[n.toString(),r]}));function setOfCachedUrls(e){return e.keys().then(function(e){return e.map(function(e){return e.url})}).then(function(e){return new Set(e)})}self.addEventListener("install",function(e){e.waitUntil(caches.open(cacheName).then(function(e){return setOfCachedUrls(e).then(function(t){return Promise.all(Array.from(urlsToCacheKeys.values()).map(function(a){if(!t.has(a)){var n=new Request(a,{credentials:"same-origin"});return fetch(n).then(function(t){if(!t.ok)throw new Error("Request for "+a+" returned a response with status "+t.status);return cleanResponse(t).then(function(t){return e.put(a,t)})})}}))})}).then(function(){return self.skipWaiting()}))}),self.addEventListener("activate",function(e){var t=new Set(urlsToCacheKeys.values());e.waitUntil(caches.open(cacheName).then(function(e){return e.keys().then(function(a){return Promise.all(a.map(function(a){if(!t.has(a.url))return e.delete(a)}))})}).then(function(){return self.clients.claim()}))}),self.addEventListener("fetch",function(e){if("GET"===e.request.method){var t,a=stripIgnoredUrlParameters(e.request.url,ignoreUrlParametersMatching),n="index.html";(t=urlsToCacheKeys.has(a))||(a=addDirectoryIndex(a,n),t=urlsToCacheKeys.has(a));0,t&&e.respondWith(caches.open(cacheName).then(function(e){return e.match(urlsToCacheKeys.get(a)).then(function(e){if(e)return e;throw Error("The cached response that was expected is missing.")})}).catch(function(t){return console.warn('Couldn\'t serve response for "%s" from cache: %O',e.request.url,t),fetch(e.request)}))}});