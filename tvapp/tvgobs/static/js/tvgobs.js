let nasaPic = () => {
	$.get("https://api.nasa.gov/planetary/apod?api_key=qdSdMLU2yc4wXQM9goawTdAA7sngW9KoLkDsVhWG", (data) => {
		$('#nasaPOD').attr('src', data.url);
	})
};
let nasaMarsPic = () => {
	$.get("https://api.nasa.gov/mars-photos/api/v1/rovers/curiosity/photos?sol=1000&camera=fhaz&api_key=qdSdMLU2yc4wXQM9goawTdAA7sngW9KoLkDsVhWG", (data) => {
		$('#nasaMarsPOD').attr('src', data.photos[0].img_src);
	})
};

const checkUpdate = () => {
	$.get("Update", (data) => {
		if ( data === 0 ) {
			return true;
		} else {
			return false;
		}
	})
};

let hideStuff = () => {
	if ( checkUpdate ) {
		$('#logo').hide();
		$('#workingLogo').hide();
	} else {
		$('#logo').show();
	}	
};

$(window).on("load", () => {
	checkUpdate();
	hideStuff();
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

$(document).ready(() => {
	$('#logo').click(() => {
		$('#logo').hide();
		$('#workingLogo').show();
	})

	$('.main').click(() => {
		$('#movies').collapse("hide");
		$("#moviesGroup").collapse("hide");
		$('#mainContainer').empty();
	})

	$('#catagoriesBtn').click(() => {
		$('#moviesGroup').collapse("hide");
	})

	$('#groupsBtn').click(() => {
		$('#movies').collapse("hide");
	})

	$('.tvshowLBtn').click(() => {
		let enctvshow = $(this).attr("data-filepath");
		let tvshow = decodeURIComponent(enctvshow);
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

	$(".orvilleSBtn").click(() => {
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
				epList = epList + F1 + F2 + endURI + F3 + v.title + F4;
			})
			$('#tvshowList').append(epList);
		})
	})

	$('.sttvSBtn').click(() => {
		$('#sttvModal').modal('hide');
		const season = $(this).attr("data-season");
		$('#tvshowList').empty();
		$.get("intSTTV", {
			"season": season
		},
		(data)=> {
			let epList = "";
			$.each(data, (i, v) => {
				const endURI = encodeURIComponent(v.movfspath);
				epList = epList + F1 + F2 + endURI + F3 + v.title + F4;
			})
			$('#tvshowList').append(epList);
		})
	})

	$('.nextgenerationSBtn').click(() => {
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
				epList = epList + F1 + F2 + endURI + F3 + v.title + F4;
			})
			$('#tvshowList').append(epList);
		})
	})

	$('.enterpriseSBtn').click(() => {
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
				epList = epList + F1 + F2 + endURI + F3 + v.title + F4;
			})
			$('#tvshowList').append(epList);
		})
	})

	$('.discoverySBtn').click(() => {
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
				epList = epList + F1 + F2 + endURI + F3 + v.title + F4;
			})
			$('#tvshowList').append(epList);
		})
	})

	$('.voyagerSBtn').click(() => {
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
				epList = epList + F1 + F2 + endURI + F3 + v.title + F4;
			})
			$('#tvshowList').append(epList);
		})
	})

	$('.lastshipSBtn').click(() => {
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
				epList = epList + F1 + F2 + endURI + F3 + v.title + F4;
			})
			$('#tvshowList').append(epList);
		})
	})

	$('.prevBtn').click(() => {
		$.ajax(
			{
				"url": OmxPrev,
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
			}
		)
	})

	$('#prevChapterBtn').click(() => {
		$.ajax(
			{
				"url": OmxPrev_chapter,
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
			}
		)
	})

	$('#nextBtn').click(() => {
		$.ajax(
			{
				"url": OmxNext,
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
			}
		)
	})

	$('#nextChapterBtn').click(() => {
		$.ajax(
			{
				"url": OmxNext_chapter,
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
			}
		)
	})

	$('#playBtn').click(() => {
		let bval = $('#playBtn').text();
		if ( bval === "Pause") {
			$('#playBtn').html("Play");
		} else {
			$('#playBtn').html("Pause");
		}
		$.ajax(
			{
				"url": OmxPlay,
				"method": "GET",
				'cache': true,
				"async": true,
				"crossDomain": true,
				"headers": {
					"accept": "text/html",
					"Access-Control-Allow-Origin": "*",
					"Content-Type": "text/html",
				},
			}
		)
	})

	$('#stopBtn').click(() => {
		$.ajax(
			{
				"url": OmxStop,
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
			}
		)
	})
});