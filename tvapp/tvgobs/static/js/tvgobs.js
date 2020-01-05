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

function epiFunc(epiData) {
	let epList = "";
	$.each(epiData, (i, v)=> {
		const endURI = encodeURIComponent(v.movfspath);
		epList = epList + f1 + f2 + endURI + f3 + v.title + f4;
	})
	$('#tvshowList').append(epList);
}

function myAjax(omxCMD) {
	$.ajax({
		"url": omxCMD,
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
}

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
	let tvshow = decodeURIComponent($(this).attr("data-filepath"));
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
	$('#tvshowList').empty();
	$.get("intOrville", {
		"season": $(this).attr("data-season")
	},
	(data)=> {
		epiFunc(data);
	})
})
.on('click', '.sttvSBtn', ()=> {
	$('#sttvModal').modal('hide');
	$('#tvshowList').empty();
	$.get("intSTTV", {
		"season": $(this).attr("data-season")
	},
	(data)=> {
		epiFunc(data);
	})
})
.on('click', '.nextgenerationSBtn', ()=> {
	$('#nextgenerationModal').modal('hide');
	$('#tvshowList').empty();
	$.get("intTNG", {
		"season": $(this).attr("data-season")
	},
	(data)=> {
		epiFunc(data);
	})
})
.on('click', '.enterpriseSBtn', ()=> {
	$('#enterpriseModal').modal('hide');
	$('#tvshowList').empty();
	$.get("intEnterprise", {
		"season": $(this).attr("data-season")
	},
	(data)=> {
		epiFunc(data);
	})
})
.on('click', '.discoverySBtn', ()=> {
	$('#discoveryModal').modal('hide');
	$('#tvshowList').empty();
	$.get("intDiscovery", {
		"season": $(this).attr("data-season")
	},
	(data)=> {
		epiFunc(data);
	})
})
.on('click', '.voyagerSBtn', ()=> {
	$('#voyagerModal').modal('hide');
	$('#tvshowList').empty();
	$.get("intVoyager", {
		"season": $(this).attr("data-season")
	},
	(data)=> {
		epiFunc(data);
	})
})
.on('click', '.lastshipSBtn', ()=> {
	$('#lastshipModal').modal('hide');
	$('#tvshowList').empty();
	$.get("intLastShip", {
		"season": $(this).attr("data-season")
	},
	(data)=> {
		epiFunc(data);
	})
})
.on('click', '#prevBtn', ()=> {
	myAjax(omxplayer_server_prev);
})
.on('click', '#prevChapterBtn', ()=> {
	myAjax(omxplayer_server_prev_chapter);
})
.on('click', '#nextBtn', ()=> {
	myAjax(omxplayer_server_next);
})
.on('click', '#nextChapterBtn', ()=> {
	myAjax(omxplayer_server_next_chapter);
})
.on('click', '#playBtn', ()=> {
	let bval = $('#playBtn').text();
	if ( bval === "Pause") {
		$('#playBtn').html("Play");
	} else {
		$('#playBtn').html("Pause");
	}
	myAjax(omxplayer_server_play);
})
.on('click', '#stopBtn', ()=> {
	myAjax(omxplayer_server_stop);
});