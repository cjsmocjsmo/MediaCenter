function hideStuff() {
	$('#workingLogo').hide();
	w3.slideshow(".madewith", 3000);
};

function nasaPic() {
	$.get("https://api.nasa.gov/planetary/apod?api_key=qdSdMLU2yc4wXQM9goawTdAA7sngW9KoLkDsVhWG", function (data) {
		$('#nasaPOD').attr('src', data.url);
	})
};
function nasaMarsPic() {
	$.get("https://api.nasa.gov/mars-photos/api/v1/rovers/curiosity/photos?sol=1000&camera=fhaz&api_key=qdSdMLU2yc4wXQM9goawTdAA7sngW9KoLkDsVhWG", function (data) {
		$('#nasaMarsPOD').attr('src', data.photos[0].img_src);
	})
};

$(window).on("load", ()=> {
	hideStuff();
	nasaPic();
	nasaMarsPic();
});

const omxplayer_server_playmedia = "http://192.168.0.42:8181/OmxplayerPlayMedia"
const omxplayer_server_play = "http://192.168.0.42:8181/Play"
const omxplayer_server_stop = "http://192.168.0.42:8181/Stop"
const omxplayer_server_prev = "http://192.168.0.42:8181/Previous"
const omxplayer_server_next = "http://192.168.0.42:8181/Next"
const omxplayer_server_prev_chapter = "http://192.168.0.42:8181/PreviousChapter"
const omxplayer_server_next_chapter = "http://192.168.0.42:8181/NextChapter"

const f1 = "<a href='#' ";
const f2 = "data-filepath="
const f3 = " class='tvshowLBtn list-group-item list-group-item-action'>";
const f4 = "</a>";

$(document).on('click', '#logo', ()=> {
	$('#logo').hide();
	$('#workingLogo').show();
	$.get("Update", (data)=> {
		if ( data == 0 ) {
			$('#logo').show();
			$('#workingLogo').hide();
		}
	})	
})
.on('click', '.main', ()=> {
	$('#movies').collapse("hide");
	$("#moviesGroup").collapse("hide");
	$('#mainContainer').empty();
})
.on('click', '#catagoriesBtn', ()=> {
	$('#moviesGroup').collapse("hide");
})
.on('click', '#groupsBtn', ()=> {
	$('#movies').collapse("hide");
})
.on('click', '.tvshowLBtn', ()=> {
	let enctvshow = $(this).attr("data-filepath");
	let tvshow = decodeURIComponent(enctvshow);
	$.ajax({
		"url": omxplayer_server_playmedia,
		"method": "GET",
		'cache': false,
		'dataType': "jsonp",
		"async": true,
		"crossDomain": true,
		"data": {
			"movie": tvshow,
		},
		"headers": {
			"accept": "text/html",
			"Access-Control-Allow-Origin": "*",
			"X-Content-Type-Options": "nosniff",
		},
	})
})
.on("click", ".orvilleSBtn", ()=> {
	$('#orvilleModal').modal('hide');
	const season = $(this).attr("data-season");
	$('#tvshowList').empty();
	$.get("intOrville", {
		"season": season
	},
	(data)=> {
		let epList = "";
		$.each(data, (i, v)=> {
			const endURI = encodeURIComponent(v.movfspath);
			epList = epList + f1 + f2 + endURI + f3 + v.title + f4;
		})
		$('#tvshowList').append(epList);
	})
})
.on('click', '.sttvSBtn', ()=> {
	$('#sttvModal').modal('hide');
	const season = $(this).attr("data-season");
	$('#tvshowList').empty();
	$.get("intSTTV", {
		"season": season
	},
	(data)=> {
		let epList = "";
		$.each(data, (i, v)=> {
			const endURI = encodeURIComponent(v.movfspath);
			epList = epList + f1 + f2 + endURI + f3 + v.title + f4;
		})
		$('#tvshowList').append(epList);
	})
})

.on('click', '.nextgenerationSBtn', ()=> {
	$('#nextgenerationModal').modal('hide');
	const season = $(this).attr("data-season");
	$('#tvshowList').empty();
	$.get("intTNG", {
		"season": season
	},
	(data)=> {
		let epList = "";
		$.each(data, (i, v)=> {
			const endURI = encodeURIComponent(v.movfspath);
			epList = epList + f1 + f2 + endURI + f3 + v.title + f4;
		})
		$('#tvshowList').append(epList);
	})
})

.on('click', '.enterpriseSBtn', ()=> {
	$('#enterpriseModal').modal('hide');
	const season = $(this).attr("data-season");
	$('#tvshowList').empty();
	$.get("intEnterprise", {
		"season": season
	},
	(data)=> {
		let epList = "";
		$.each(data, (i, v)=> {
			const endURI = encodeURIComponent(v.movfspath);
			epList = epList + f1 + f2 + endURI + f3 + v.title + f4;
		})
		$('#tvshowList').append(epList);
	})
})

.on('click', '.discoverySBtn', ()=> {
	$('#discoveryModal').modal('hide');
	const season = $(this).attr("data-season");
	$('#tvshowList').empty();
	$.get("intDiscovery", {
		"season": season
	},
	(data)=> {
		let epList = "";
		$.each(data, (i, v)=> {
			const endURI = encodeURIComponent(v.movfspath);
			epList = epList + f1 + f2 + endURI + f3 + v.title + f4;
		})
		$('#tvshowList').append(epList);
	})
})

.on('click', '.voyagerSBtn', ()=> {
	$('#voyagerModal').modal('hide');
	const season = $(this).attr("data-season");
	$('#tvshowList').empty();
	$.get("intVoyager", {
		"season": season
	},
	(data)=> {
		let epList = "";
		$.each(data, (i, v)=> {
			const endURI = encodeURIComponent(v.movfspath);
			epList = epList + f1 + f2 + endURI + f3 + v.title + f4;
		})
		$('#tvshowList').append(epList);
	})
})

.on('click', '.lastshipSBtn', ()=> {
	$('#lastshipModal').modal('hide');
	const season = $(this).attr("data-season");
	$('#tvshowList').empty();
	$.get("intLastShip", {
		"season": season
	},
	(data)=> {
		let epList = "";
		$.each(data, (i, v)=> {
			const endURI = encodeURIComponent(v.movfspath);
			epList = epList + f1 + f2 + endURI + f3 + v.title + f4;
		})
		$('#tvshowList').append(epList);
	})
})

.on('click', '#prevBtn', ()=> {
	$.ajax({
	"url": omxplayer_server_prev,
	"method": "GET",
	'cache': false,
	'dataType': "jsonp",
	"async": true,
	"crossDomain": true,
	"headers": {
		"accept": "text/html",
		"Access-Control-Allow-Origin": "*",
		"X-Content-Type-Options": "nosniff",
	},
	})
})
.on('click', '#prevChapterBtn', ()=> {
	$.ajax({
	"url": omxplayer_server_prev_chapter,
	"method": "GET",
	'cache': false,
	'dataType': "jsonp",
	"async": true,
	"crossDomain": true,
	"headers": {
		"accept": "text/html",
		"Access-Control-Allow-Origin": "*",
		"X-Content-Type-Options": "nosniff",
	},
	})
})
.on('click', '#nextBtn', ()=> {
	$.ajax({
	"url": omxplayer_server_next,
	"method": "GET",
	'cache': false,
	'dataType': "jsonp",
	"async": true,
	"crossDomain": true,
	"headers": {
		"accept": "text/html",
		"Access-Control-Allow-Origin": "*",
		"X-Content-Type-Options": "nosniff",
	},
	})
})
.on('click', '#nextChapterBtn', ()=> {
	$.ajax({
	"url": omxplayer_server_next_chapter,
	"method": "GET",
	'cache': false,
	'dataType': "jsonp",
	"async": true,
	"crossDomain": true,
	"headers": {
		"accept": "text/html",
		"Access-Control-Allow-Origin": "*",
		"X-Content-Type-Options": "nosniff",
	},
	})
})
.on('click', '#playBtn', ()=> {
	let bval = $('#playBtn').text();
	if ( bval === "Pause") {
		$('#playBtn').html("Play");
	} else {
		$('#playBtn').html("Pause");
	}
	$.ajax({
	"url": omxplayer_server_play,
	"method": "GET",
	'cache': true,
	"async": true,
	"crossDomain": true,
	"headers": {
		"accept": "text/html",
		"Access-Control-Allow-Origin": "*",
		"Content-Type": "text/html",
	},

	})
})
.on('click', '#stopBtn', ()=> {
	$.ajax({
	"url": omxplayer_server_stop,
	"method": "GET",
	'cache': false,
	'dataType': "jsonp",
	"async": true,
	"crossDomain": true,
	"headers": {
		"accept": "text/html",
		"Access-Control-Allow-Origin": "*",
		"X-Content-Type-Options": "nosniff",
	},
	})
});