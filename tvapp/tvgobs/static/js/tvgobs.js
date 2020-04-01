let nasaPic = function() {
	$.get("https://api.nasa.gov/planetary/apod?api_key=qdSdMLU2yc4wXQM9goawTdAA7sngW9KoLkDsVhWG", (data) => {
		$('#nasaPOD').attr('src', data.url);
	})
};
let nasaMarsPic = function() {
	$.get("https://api.nasa.gov/mars-photos/api/v1/rovers/curiosity/photos?sol=1000&camera=fhaz&api_key=qdSdMLU2yc4wXQM9goawTdAA7sngW9KoLkDsVhWG", (data) => {
		$('#nasaMarsPOD').attr('src', data.photos[0].img_src);
	})
};
// let UDS = localStorage.getItem('updateStatus');

let hideStuff = function() {
	$('#logo').show();
	$('#workingLogo').hide();
	// localStorage.setItem('updateStatus', JSON.stringify(1));
};

$(window).on("load", function() {
	if (localStorage.getItem("TVStatus") == "0") {
		$('#logo').hide();
		$('#workingLogo').hide();
	} else {
		hideStuff();
	};
	nasaPic();
	nasaMarsPic();
});

const HTTP = "http://192.168.0.42:8181/"
const OmxPlaymedia = HTTP + "OmxplayerPlayMedia"
const OmxPlay = HTTP + "Play"
const OmxStop = HTTP + "Stop"
const OmxPrev = HTTP + "Previous"
const OmxNext = HTTP + "Next"
const OmxPrevChapter = HTTP + "PreviousChapter"
const OmxNextChapter = HTTP + "NextChapter"

const F1 = "<a href='#' ";
const F2 = "data-filepath="
const F3 = " class='tvshowLBtn list-group-item list-group-item-action'>";
const F4 = "</a>";

let myAjax = (OMX) => {
	console.log("this is omx");
	console.log(OMX);
	$.ajax(
		{
			"url": OMX,
			"method": "GET",
			'cache': false,
			'dataType': "jsonp",
			"async": true,
			"crossDomain": true,
			"headers": {
				"accept": "text/html",
				"Access-Control-Allow-Origin": "*",
				"content-type": "application/javascript;charset=utf-8",
			},
		}
	)
};

$(document).on('click', '#logo', function() {
	$('#logo').hide();
	$('#workingLogo').show();
	$.get("Update", (data) => {
		if (data == "0") {
			$('#workingLogo').hide();
			localStorage.setItem('TVStatus', JSON.stringify(data));
		}
		
	})
})

.on('click', '.main', function() {
	$('#movies').collapse("hide");
	$("#moviesGroup").collapse("hide");
	$('#mainContainer').empty();
})

.on('click', '#catagoriesBtn', function() {
	$('#moviesGroup').collapse("hide");
})

.on('click', '#groupsBtn', function() {
	$('#movies').collapse("hide");
})

.on('click', '.tvshowLBtn', function() {
	let enctvshow = $(this).attr("data-filepath");
	let tvshow = decodeURIComponent(enctvshow);
	console.log("this is tvshow");
	console.log(tvshow);
	$.ajax(
		{
			"url": OmxPlaymedia,
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
		}
	)
})

.on('click', ".orvilleSBtn", function() {
	$('#orvilleModal').modal('hide');
	const season = $(this).attr("data-season");
	$('#tvshowList').empty();
	$.get("intOrville", {
		"season": season
	},
	(data) => {
		let epList = "";
		$.each(data, (i, v)=> {
			const endURI = encodeURIComponent(v.movfspath);
			epList = epList + F1 + F2 + endURI + F3 + v.title + F4;
		})
		$('#tvshowList').append(epList);
	})
})

.on('click', '.sttvSBtn', function() {
	$('#sttvModal').modal('hide');
	const season = $(this).attr("data-season");
	console.log("this is season");
	console.log(season);
	$('#tvshowList').empty();
	$.get("intSTTV", {
		"season": season
	},
	(data) => {
		let epList = "";
		$.each(data, (i, v) => {
			const endURI = encodeURIComponent(v.movfspath);
			epList = epList + F1 + F2 + endURI + F3 + v.title + F4;
		})
		$('#tvshowList').append(epList);
	})
})

.on('click', '.nextgenerationSBtn', function() {
	$('#nextgenerationModal').modal('hide');
	const season = $(this).attr("data-season");
	$('#tvshowList').empty();
	$.get("intTNG", {
		"season": season
	},
	(data) => {
		let epList = "";
		$.each(data, (i, v) => {
			const endURI = encodeURIComponent(v.movfspath);
			epList = epList + F1 + F2 + endURI + F3 + v.title + F4;
		})
		$('#tvshowList').append(epList);
	})
})

.on('click', '.enterpriseSBtn', function() {
	$('#enterpriseModal').modal('hide');
	const season = $(this).attr("data-season");
	$('#tvshowList').empty();
	$.get("intEnterprise", {
		"season": season
	},
	(data) => {
		let epList = "";
		$.each(data, (i, v)=> {
			const endURI = encodeURIComponent(v.movfspath);
			epList = epList + F1 + F2 + endURI + F3 + v.title + F4;
		})
		$('#tvshowList').append(epList);
	})
})

.on('click', '.discoverySBtn', function() {
	$('#discoveryModal').modal('hide');
	const season = $(this).attr("data-season");
	$('#tvshowList').empty();
	$.get("intDiscovery", {
		"season": season
	},
	(data) => {
		let epList = "";
		$.each(data, (i, v)=> {
			const endURI = encodeURIComponent(v.movfspath);
			epList = epList + F1 + F2 + endURI + F3 + v.title + F4;
		})
		$('#tvshowList').append(epList);
	})
})

.on('click', '.voyagerSBtn', function() {
	$('#voyagerModal').modal('hide');
	const season = $(this).attr("data-season");
	$('#tvshowList').empty();
	$.get("intVoyager", {
		"season": season
	},
	(data) => {
		let epList = "";
		$.each(data, (i, v)=> {
			const endURI = encodeURIComponent(v.movfspath);
			epList = epList + F1 + F2 + endURI + F3 + v.title + F4;
		})
		$('#tvshowList').append(epList);
	})
})

.on('click', '.lastshipSBtn', function() {
	$('#lastshipModal').modal('hide');
	const season = $(this).attr("data-season");
	$('#tvshowList').empty();
	$.get("intLastShip", {
		"season": season
	},
	(data) => {
		let epList = "";
		$.each(data, (i, v)=> {
			const endURI = encodeURIComponent(v.movfspath);
			epList = epList + F1 + F2 + endURI + F3 + v.title + F4;
		})
		$('#tvshowList').append(epList);
	})
})

.on('click', '.prevBtn', function() {
	myAjax(OmxPrev);
})

.on('click', '#prevChapterBtn', function() {
	myAjax(OmxPrev_chapter);
})

.on('click', '#nextBtn', function() {
	myAjax(OmxNext);
})

.on('click', '#nextChapterBtn', function() {
	myAjax(OmxNext_chapter);
})

.on('click', '#playBtn', function() {
	let bval = $('#playBtn').text();
	if ( bval === "Pause") {
		$('#playBtn').html("Play");
	} else {
		$('#playBtn').html("Pause");
	}
	myAjax(OmxPlay);
	})

.on('click', '#stopBtn', function() {
	myAjax(OmxStop);
});