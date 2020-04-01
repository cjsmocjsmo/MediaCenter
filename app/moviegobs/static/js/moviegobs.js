const HttP = "http://192.168.0.42:8181/";

const hideStuff = () => {
	$('#logo').show();
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
	if (localStorage.getItem('movieStatus') == "0") {
		$('#logo').hide();
		$('#workingLogo').hide();
	} else {
		hideStuff();
	}

	// console.log(localStorage.getItem('updateStatus'));
	
	nasaPic();
	nasaMarsPic();
});

const P1 = "<div class='row bg-dark'>";
const P2 = "<div class='pt-1 pb-1 col'>";
const P3 = "<img id='baz' class='taz d-block mx-auto' src='";
const P5 = "' ";
const P5A = "' style='width: 35%;'";

const OmxplayerServerPlaymedia = HttP + "OmxplayerPlayMedia";
const OmxplayerServerPlay = HttP + "Play";
const OmxplayerServerStop = HttP + "Stop";
const OmxplayerServerPrev = HttP + "Previous";
const OmxplayerServerNext = HttP + "Next";
const OmxplayerServerPrevChapter = HttP + "PreviousChapter";
const OmxplayerServerNextChapter = HttP + "NextChapter";

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
};

let mainFunc = (mydata) => {
	console.log("starting main function");
	console.log(mydata)
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
};

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


$(document).on('click', '#stopBtn', () => {
	myAjax(OmxplayerServerStop);
})

.on('click', '#playBtn', () => {
	let bval = $('#playBtn').text();
	console.log("this is bval");
	console.log(bval);
	if ( bval === "Pause") {
		$('#playBtn').html("Play");
	} else {
		$('#playBtn').html("Pause");
	};
	myAjax(OmxplayerServerPlay);
})

.on('click', '.taz', function() {
	console.log("taz been clicked");
	let movie = $(this).data("movie");
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
				// "movie": movie["movie"]
				"movie": movie
			},
			"headers": {
				"accept": "text/html",
				"Access-Control-Allow-Origin": "*",
				"content-type": "application/javascript;charset=utf-8"
			},
		}
	);
})

.on('click', '#logo', () => {
	$('#logo').hide();
	$('#workingLogo').show();
	$.get("Update", (data) => {
		if ( data == "0" ) {
			$('#workingLogo').hide();
			localStorage.setItem('movieStatus', JSON.stringify(data));
		}
	})
})

.on('click', '.main', () => {
	$('#movies').collapse("hide");
	$("#moviesGroup").collapse("hide");
	$('#mainContainer').empty();
})

.on('click', '#catagoriesBtn', () => {
	$('#moviesGroup').collapse("hide");
})

.on('click', '#groupsBtn', () => {
	$('#movies').collapse("hide");
})

.on('click', '#actionBtn', () => {
	$.get('intAction', (data) => {
		console.log(data);
		mainFunc(data);
	})
})

.on('click', '#cartoonsBtn', () => {
	$.get('intCartoons', (data) => {
		mainFunc(data);
	})
})

.on('click', '#comedyBtn', () => {
	$.get('intComedy', (data) => {
		mainFunc(data);
	})
})

.on('click', '#dramaBtn', () => {
	$.get('intDrama', (data) => {
		mainFunc(data);
	})
})

.on('click', '#godzillaBtn', () => {
	$.get('intGodzilla', (data) => {
		mainFunc(data);
	})
})

.on('click', '#harrypotterBtn', () => {
	$.get('intHarryPotter', (data) => {
		mainFunc(data);
	})
})

.on('click', '#indianajonesBtn', () => {
	$.get('intIndianaJones', (data) => {
		mainFunc(data);
	})
})

.on('click', '#johnwayneBtn', () => {
	$.get('intJohnWayne', (data) => {
		mainFunc(data);
	})
})

.on('click', '#jurassicparkBtn', () => {
	$.get('intJurassicPark', (data) => {
		mainFunc(data);
	})
})

.on('click', '#kingsmanBtn', () => {
	$.get('intKingsMan', (data) => {
		mainFunc(data);
	})
})

.on('click', '#meninblackBtn', () => {
	$.get('intMenInBlack', (data) => {
		mainFunc(data);
	})
})

.on('click', '#miscBtn', () => {
	$.get('intMisc', (data) => {
		console.log("this is data");
		console.log(data);
		mainFunc(data);
	})
})

.on('click', '#scifiBtn', () => {
	$.get('intSciFi', (data) => {
		console.log("this is data");
		console.log(data);
		mainFunc(data);
	})
})

.on('click', '#startrekBtn', () => {
	$.get('intStarTrek', (data) => {
		mainFunc(data);
	})
})

.on('click', '#starwarsBtn', () => {
	$.get('intStarWars', (data) => {
		mainFunc(data);
	})
})

.on('click', '#superherosBtn', () => {
	$.get('intSuperHeros', (data) => {
		mainFunc(data);
	})
})

.on('click', '#tremorsBtn', () => {
	$.get('intTremors', (data) => {
		mainFunc(data);
	})
})
	
.on('click', '#prevBtn', () => {
	myAjax(OmxplayerServerPrev);
})

.on('click', '#prevChapterBtn', () => {
	myAjax(OmxplayrServerPrevChapter);
})

.on('click', '#nextBtn', () => {
	myAjax(OmxplayerServerNext);	
})

.on('click', '#nextChapterBtn', () => {
	myAjax(OmxplayerServerNextChapter);
});