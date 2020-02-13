const hideStuff = () => {
	$('#workingLogo').hide();
	w3.slideshow(".madewith", 3000);
};

const nasaPic = () => {
	$.get("https://api.nasa.gov/planetary/apod?api_key=qdSdMLU2yc4wXQM9goawTdAA7sngW9KoLkDsVhWG", function (data) {
		$('#nasaPOD').attr('src', data.url);
	})
};

const nasaMarsPic = () => {
	$.get("https://api.nasa.gov/mars-photos/api/v1/rovers/curiosity/photos?sol=1000&camera=fhaz&api_key=qdSdMLU2yc4wXQM9goawTdAA7sngW9KoLkDsVhWG", function (data) {
		$('#nasaMarsPOD').attr('src', data.photos[0].img_src);
	})
};

$(window).on("load", () => {
	hideStuff();
	nasaPic();
	nasaMarsPic();
});

const P1 = "<div class='row bg-dark'>";
const P2 = "<div class='pt-1 pb-1 col'>";
const P3 = "<img class='taz d-block mx-auto' src='";
const P5 = "' ";
const P5A = "' style='width: 35%;'";

const HttP = "http://192.168.0.42:8181/"
const OmxplayerServerPlaymedia = HttP + "OmxplayerPlayMedia"
const OmxplayerServerPlay = HttP + "Play"
const OmxplayerServerStop = HttP + "Stop"
const OmxplayerServerPrev = HttP + "Previous"
const OmxplayerServerNext = HttP + "Next"
const OmxplayerServerPrevChapter = HttP + "PreviousChapter"
const OmxplayerServerNextChapter = HttP + "NextChapter"


let movcountOne = (movcount) => {
	let P4 = movcount[0].thumbpath;
	let P6 = " data-movie='";
	let P7 = movcount[0].movfspath;
	let P8 = "'></div></div>";
	only = P1 + P2 + P3 + P4 + P5A + P6 + P7 + P8;
	$('#mainContainer').append(only);
};

let movcountOdd = (movdata) => {
	let solo = movdata.pop();
	let P4 = solo.thumbpath;
	let P6 = " data-movie='";
	let P7 = solo.movfspath;
	let P8 = "'></div></div>";
	return P1 + P2 + P3 + P4 + P5A + P6 + P7 + P8
}

let mainFunc = (mydata) => {
	let mlast = "";
	let mresult = "";
	let mres = '';
	let mcount = 0;
	if ( mydata.length === 1 ) {
		movcountOne(mydata);
	}
	if ( mydata.length % 2 != 0 ) {
		mlast = movcountOdd(mydata);
	}
	$.each(mydata, function (I, V) {
		mcount += 1;
		if ( mcount == 1 ) {
			let P4 = V.thumbpath;
			let P6 = " data-movie='";
			let P7 = V.movfspath;
			let P8 = "'></div>";
			mres = mres + P1 + P2 + P3 + P4 + P5 + P6 + P7 + P8;
		}
		if ( mcount == 2 ) {
			let P4 = V.thumbpath;
			let P6 = " data-movie='";
			let P7 = V.movfspath;
			let P8 = "'></div></div>";
			mres = mres + P2 + P3 + P4 + P5 + P6 + P7 + P8;
			mresult = mresult + mres;
			mcount = 0;
			mres = "";
		}
	})
	mresult = mresult + mlast; 
	$('#mainContainer').append(mresult);
}


$(document).ready(() => {

	$('#logo').click(() => {
		$('#logo').hide();
		$('#workingLogo').show();
		$.get("Update", (data)=> {
			if ( data == 0 ) {
				$('#logo').show();
				$('#workingLogo').hide();
			}
		})
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

	$('.taz').click(() => {
		let movie = $(this).data();
		console.log('this is movid');
		console.log(movie);
		$.ajax(
			{
				"url": OmxplayerServerPlaymedia,
				"method": "GET",
				'cache': false,
				'dataType': "jsonp",
				"async": true,
				"crossDomain": true,
				"data": {
					"movie": movie["movie"]
				},
				"headers": {
					"accept": "text/html",
					"Access-Control-Allow-Origin": "*",
					"X-Content-Type-Options": "nosniff",
				},
			}
		)
	})

	$('#actionBtn').click(() => {
		$.get('intAction', (data) => {
			mainFunc(data);
		})
	})

	$('#cartoonsBtn').click(() => {
		$.get('intCartoons', (data) => {
			mainFunc(data);
		})
	})

	$('#comedyBtn').click(()=> {
		$.get('intComedy', (data) => {
			mainFunc(data);
		})
	})

	$('#dramaBtn').click(() => {
		$.get('intDrama', (data) => {
			mainFunc(data);
		})
	})

	$('#godzillaBtn').click(() => {
		$.get('intGodzilla', (data) => {
			mainFunc(data);
		})
	})

	$('#harrypotterBtn').click(() => {
		$.get('intHarryPotter', (data) => {
			mainFunc(data);
		})
	})

	$('#indianajonesBtn').click(() => {
		$.get('intIndianaJones', (data) => {
			mainFunc(data);
		})
	})

	$('#johnwayneBtn').click(() => {
		$.get('intJohnWayne', (data) => {
			mainFunc(data);
		})
	})

	$('#jurassicparkBtn').click(() => {
		$.get('intJurassicPark', (data) => {
			mainFunc(data);
		})
	})

	$('#kingsmanBtn').click(() => {
		$.get('intKingsMan', (data) => {
			mainFunc(data);
		})
	})

	$('#meninblackBtn').click(() => {
		$.get('intMenInBlack', (data) => {
			mainFunc(data);
		})
	})

	$('#miscBtn').click(() => {
		$.get('intMisc', (data) => {
			mainFunc(data);
		})
	})

	$('#scifiBtn').click(() => {
		$.get('intSciFi', (data) => {
			mainFunc(data);
		})
	})

	$('#startrekBtn').click(() => {
		$.get('intStarTrek', (data) => {
			mainFunc(data);
		})
	})

	$('#starwarsBtn').click(() => {
		$.get('intStarWars', (data) => {
			mainFunc(data);
		})
	})

	$('#superherosBtn').click(() => {
		$.get('intSuperHeros', (data) => {
			mainFunc(data);
		})
	})

	$('#tremorsBtn').click(() => {
		console.log("tremors has been clicked");
		$.get('intTremors', (data) => {
			mainFunc(data);
		})
	})
	
	$('#prevBtn').click(() => {
		$.ajax(
			{
				"url": OmxplayerServerPrev,
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
				"url": OmxplayrServerPrevChapter,
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
				"url": OmxplayerServerNext,
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
				"url": OmxplayerServerNextChapter,
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
		console.log("this is bval");
		console.log(bval);
		if ( bval === "Pause") {
			$('#playBtn').html("Play");
		} else {
			$('#playBtn').html("Pause");
		}
		$.ajax(
			{
				"url": OmxplayerServerPlay,
				"method": "GET",
				'cache': true,
				//'dataType': "jsonp",
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
				"url": OmxplayerServerStop,
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