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

const p1 = "<div class='row bg-dark'>";
const p2 = "<div class='pt-1 pb-1 col'>";
const p3 = "<img class='taz d-block mx-auto' src='";
const p5 = "' ";
const p5a = "' style='width: 35%;'";




// const a1 = "<div class='carousel-item'>";
// const a1a = "<div class='carousel-item active'>";
// const a2 = "<img class='tazCaro d-block mx-auto w-20' src='";
// const a4 = "' ";
// const a5 = "data-movieid='";
// const a7 = "'></div>";




// const b1 = "<li data-movie='";
// const b2 = "' data-target='#demo' data-slide-to='";
// const b3a = "' class='active'></li>";
// const b3b = "' ></li>";





const omxplayer_server_playmedia = "http://192.168.0.42:8181/OmxplayerPlayMedia"
const omxplayer_server_play = "http://192.168.0.42:8181/Play"
const omxplayer_server_stop = "http://192.168.0.42:8181/Stop"
const omxplayer_server_prev = "http://192.168.0.42:8181/Previous"
const omxplayer_server_next = "http://192.168.0.42:8181/Next"
const omxplayer_server_prev_chapter = "http://192.168.0.42:8181/PreviousChapter"
const omxplayer_server_next_chapter = "http://192.168.0.42:8181/NextChapter"



function movcountOne(movcount) {
	let p4 = movcount[0].thumbpath;
	let p6 = " data-movie='";
	let p7 = movcount[0].movfspath;
	let p8 = "'></div></div>";
	only = p1 + p2 + p3 + p4 + p5a + p6 + p7 + p8;
	$('#mainContainer').append(only);
};

function movcountOdd(movdata) {
	let solo = movdata.pop();
	let p4 = solo.thumbpath;
	let p6 = " data-movie='";
	let p7 = solo.movfspath;
	let p8 = "'></div></div>";
	resultlast = p1 + p2 + p3 + p4 + p5a + p6 + p7 + p8;
	return resultlast
}

function mainFunc(mydata) {
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
		if( mcount == 1 ) {
			let p4 = V.thumbpath;
			let p6 = " data-movie='";
			let p7 = V.movfspath;
			let p8 = "'></div>";
			mres = mres + p1 + p2 + p3 + p4 + p5 + p6 + p7 + p8;
		}
		if ( mcount == 2 ) {
			let p4 = V.thumbpath;
			let p6 = " data-movie='";
			let p7 = V.movfspath;
			let p8 = "'></div></div>";
			mres = mres + p2 + p3 + p4 + p5 + p6 + p7 + p8;
			mresult = mresult + mres;
			mcount = 0;
			mres = "";
		}
	})
	mresult = mresult + mlast; 
	$('#mainContainer').append(mresult);
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

.on('click', '.taz', function () {
	var movie = $(this).data();
	console.log('this is movid');
	console.log(movie);
	$.ajax({
		"url": omxplayer_server_playmedia,
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
	})
})

.on('click', '#actionBtn', function () {
	$.get('intAction', function (data) {




		// $('#r2d2').empty();
		// $('#r2d2Ind').empty();
		// let rcount = 0;
		// let result = "";
		// let result2 = "";
		// $.each(data, function (I, V) {
		// 	rcount += 1;
		// 	if (rcount == 1) {
		// 		result = result + a1a + a2 + V.carosthumbpath + a4 + a5 + V.movfspath + a7;
		// 		result2 = result2 + b1 + V.movfspath + b2 + "0" + b3a;
		// 	} else {
		// 		result = result + a1 + a2 + V.carosthumbpath + a4 + a5 + V.movfspath + a7;
		// 		c = rcount - 1;
		// 		cstring = c.toString();
		// 		result2 = result2 + b1 + V.movfspath + b2 + cstring + b3b;
		// 	}
		// })
		// $('#r2d2').append(result);
		// $('#r2d2Ind').append(result2);





		mainFunc(data);
	})
})
.on('click', '#cartoonsBtn', ()=> {
	$.get('intCartoons', (data)=> {
		console.log("this is intcartons data");
		console.log(data);




		// // $('#r2d2').empty();
		//  $('#r2d2Ind').empty();
		//  let rcount = 0;
		// // let result = "";
		// // let result2 = "";
		//  $.each(data, function (I, V) {
		//  	rcount += 1;
		//  	if (rcount == 1) {
		// // 		result = result + a1a + a2 + V.carosthumbpath + a4 + a5 + V.movfspath + a7;
		//  		result2 = result2 + b1 + V.movfspath + b2 + "0" + b3a;
	 	// } else {
		// // 		result = result + a1 + a2 + V.carosthumbpath + a4 + a5 + V.movfspath + a7;
		//  		c = rcount - 1;
		//  		cstring = c.toString();
		//  		result2 = result2 + b1 + V.movfspath + b2 + cstring + b3b;
		//  	}
		//  })
		// // $('#r2d2').append(result);
		//  $('#r2d2Ind').append(result2);



		mainFunc(data);
	})
})
.on('click', '#comedyBtn', function () {
	$.get('intComedy', function (data) {



		// $('#r2d2').empty();
		// $('#r2d2Ind').empty();
		// let rcount = 0;
		// let result = "";
		// let result2 = "";
		// $.each(data, function (I, V) {
		// 	rcount += 1;
		// 	if (rcount == 1) {
		// 		result = result + a1a + a2 + V.carosthumbpath + a4 + a5 + V.movfspath + a7;
		// 		result2 = result2 + b1 + V.movfspath + b2 + "0" + b3a;
		// 	} else {
		// 		result = result + a1 + a2 + V.carosthumbpath + a4 + a5 + V.movfspath + a7;
		// 		c = rcount - 1;
		// 		cstring = c.toString();
		// 		result2 = result2 + b1 + V.movfspath + b2 + cstring + b3b;
		// 	}
		// })
		// $('#r2d2').append(result);
		// $('#r2d2Ind').append(result2);




		mainFunc(data);
	})
})
.on('click', '#dramaBtn', function () {
	$.get('intDrama', function (data) {




		// $('#r2d2').empty();
		// $('#r2d2Ind').empty();
		// let rcount = 0;
		// let result = "";
		// let result2 = "";
		// $.each(data, function (I, V) {
		// 	rcount += 1;
		// 	if (rcount == 1) {
		// 		result = result + a1a + a2 + V.carosthumbpath + a4 + a5 + V.movfspath + a7;
		// 		result2 = result2 + b1 + V.movfspath + b2 + "0" + b3a;
		// 	} else {
		// 		result = result + a1 + a2 + V.carosthumbpath + a4 + a5 + V.movfspath + a7;
		// 		c = rcount - 1;
		// 		cstring = c.toString();
		// 		result2 = result2 + b1 + V.movfspath + b2 + cstring + b3b;
		// 	}
		// })
		// $('#r2d2').append(result);
		// $('#r2d2Ind').append(result2);



		mainFunc(data);
	})
})
.on('click', '#godzillaBtn', function () {
	$.get('intGodzilla', function (data) {



		// $('#r2d2').empty();
		// $('#r2d2Ind').empty();
		// let rcount = 0;
		// let result = "";
		// let result2 = "";
		// $.each(data, function (I, V) {
		// 	rcount += 1;
		// 	if (rcount == 1) {
		// 		result = result + a1a + a2 + V.carosthumbpath + a4 + a5 + V.movfspath + a7;
		// 		result2 = result2 + b1 + V.movfspath + b2 + "0" + b3a;
		// 	} else {
		// 		result = result + a1 + a2 + V.carosthumbpath + a4 + a5 + V.movfspath + a7;
		// 		c = rcount - 1;
		// 		cstring = c.toString();
		// 		result2 = result2 + b1 + V.movfspath + b2 + cstring + b3b;
		// 	}
		// })
		// $('#r2d2').append(result);
		// $('#r2d2Ind').append(result2);





		mainFunc(data);
	})
})
.on('click', '#harrypotterBtn', function () {
	$.get('intHarryPotter', function (data) {



		// $('#r2d2').empty();
		// $('#r2d2Ind').empty();
		// let rcount = 0;
		// let result = "";
		// let result2 = "";
		// $.each(data, function (I, V) {
		// 	rcount += 1;
		// 	if (rcount == 1) {
		// 		result = result + a1a + a2 + V.carosthumbpath + a4 + a5 + V.movfspath + a7;
		// 		result2 = result2 + b1 + V.movfspath + b2 + "0" + b3a;
		// 	} else {
		// 		result = result + a1 + a2 + V.carosthumbpath + a4 + a5 + V.movfspath + a7;
		// 		c = rcount - 1;
		// 		cstring = c.toString();
		// 		result2 = result2 + b1 + V.movfspath + b2 + cstring + b3b;
		// 	}
		// })
		// $('#r2d2').append(result);
		// $('#r2d2Ind').append(result2);



		mainFunc(data);
	})
})
.on('click', '#indianajonesBtn', function () {
	$.get('intIndianaJones', function (data) {
		console.log("this is indianajones data");
		console.log(data);



		// $('#r2d2').empty();
		// $('#r2d2Ind').empty();
		// let rcount = 0;
		// let result = "";
		// let result2 = "";
		// $.each(data, function (I, V) {
		// 	rcount += 1;
		// 	if (rcount == 1) {
		// 		result = result + a1a + a2 + V.carosthumbpath + a4 + a5 + V.movfspath + a7;
		// 		result2 = result2 + b1 + V.movfspath + b2 + "0" + b3a;
		// 	} else {
		// 		result = result + a1 + a2 + V.carosthumbpath + a4 + a5 + V.movfspath + a7;
		// 		c = rcount - 1;
		// 		cstring = c.toString();
		// 		result2 = result2 + b1 + V.movfspath + b2 + cstring + b3b;
		// 	}
		// })
		// $('#r2d2').append(result);
		// $('#r2d2Ind').append(result2);




		mainFunc(data);
	})
})
.on('click', '#johnwayneBtn', function () {
	$.get('intJohnWayne', function (data) {



		// $('#r2d2').empty();
		// $('#r2d2Ind').empty();
		// let rcount = 0;
		// let result = "";
		// let result2 = "";
		// $.each(data, function (I, V) {
		// 	rcount += 1;
		// 	if (rcount == 1) {
		// 		result = result + a1a + a2 + V.carosthumbpath + a4 + a5 + V.movfspath + a7;
		// 		result2 = result2 + b1 + V.movfspath + b2 + "0" + b3a;
		// 	} else {
		// 		result = result + a1 + a2 + V.carosthumbpath + a4 + a5 + V.movfspath + a7;
		// 		c = rcount - 1;
		// 		cstring = c.toString();
		// 		result2 = result2 + b1 + V.movfspath + b2 + cstring + b3b;
		// 	}
		// })
		// $('#r2d2').append(result);
		// $('#r2d2Ind').append(result2);



		mainFunc(data);
	})
})
.on('click', '#jurassicparkBtn', function () {
	$.get('intJurassicPark', function (data) {



		// $('#r2d2').empty();
		// $('#r2d2Ind').empty();
		// let rcount = 0;
		// let result = "";
		// let result2 = "";
		// $.each(data, function (I, V) {
		// 	rcount += 1;
		// 	if (rcount == 1) {
		// 		result = result + a1a + a2 + V.carosthumbpath + a4 + a5 + V.movfspath + a7;
		// 		result2 = result2 + b1 + V.movfspath + b2 + "0" + b3a;
		// 	} else {
		// 		result = result + a1 + a2 + V.carosthumbpath + a4 + a5 + V.movfspath + a7;
		// 		c = rcount - 1;
		// 		cstring = c.toString();
		// 		result2 = result2 + b1 + V.movfspath + b2 + cstring + b3b;
		// 	}
		// })
		// $('#r2d2').append(result);
		// $('#r2d2Ind').append(result2);


		mainFunc(data);
	})
})
.on('click', '#kingsmanBtn', function () {
	$.get('intKingsMan', function (data) {


		// $('#r2d2').empty();
		// $('#r2d2Ind').empty();
		// let rcount = 0;
		// let result = "";
		// let result2 = "";
		// $.each(data, function (I, V) {
		// 	rcount += 1;
		// 	if (rcount == 1) {
		// 		result = result + a1a + a2 + V.carosthumbpath + a4 + a5 + V.movfspath + a7;
		// 		result2 = result2 + b1 + V.movfspath + b2 + "0" + b3a;
		// 	} else {
		// 		result = result + a1 + a2 + V.carosthumbpath + a4 + a5 + V.movfspath + a7;
		// 		c = rcount - 1;
		// 		cstring = c.toString();
		// 		result2 = result2 + b1 + V.movfspath + b2 + cstring + b3b;
		// 	}
		// })
		// $('#r2d2').append(result);
		// $('#r2d2Ind').append(result2);


		mainFunc(data);
	})
})
.on('click', '#meninblackBtn', function () {

	$.get('intMenInBlack', function (data) {


		// $('#r2d2').empty();
		// $('#r2d2Ind').empty();
		// let rcount = 0;
		// let result = "";
		// let result2 = "";
		// $.each(data, function (I, V) {
		// 	rcount += 1;
		// 	if (rcount == 1) {
		// 		result = result + a1a + a2 + V.carosthumbpath + a4 + a5 + V.movfspath + a7;
		// 		result2 = result2 + b1 + V.movfspath + b2 + "0" + b3a;
		// 	} else {
		// 		result = result + a1 + a2 + V.carosthumbpath + a4 + a5 + V.movfspath + a7;
		// 		c = rcount - 1;
		// 		cstring = c.toString();
		// 		result2 = result2 + b1 + V.movfspath + b2 + cstring + b3b;
		// 	}
		// })
		// $('#r2d2').append(result);
		// $('#r2d2Ind').append(result2);



		mainFunc(data);
	})
})
.on('click', '#miscBtn', function () {

	$.get('intMisc', function (data) {


		// $('#r2d2').empty();
		// $('#r2d2Ind').empty();
		// let rcount = 0;
		// let result = "";
		// let result2 = "";

		// $.each(data, function (I, V) {
		// 	rcount += 1;
		// 	if (rcount == 1) {
		// 		result = result + a1a + a2 + V.carosthumbpath + a4 + a5 + V.movfspath + a7;
		// 		result2 = result2 + b1 + V.movfspath + b2 + "0" + b3a;
		// 	} else {
		// 		result = result + a1 + a2 + V.carosthumbpath + a4 + a5 + V.movfspath + a7;
		// 		c = rcount - 1;
		// 		cstring = c.toString();
		// 		result2 = result2 + b1 + V.movfspath + b2 + cstring + b3b;
		// 	}
		// })
		// $('#r2d2').append(result);
		// $('#r2d2Ind').append(result2);


		mainFunc(data);
	})
})
.on('click', '#scifiBtn', function () {
	$.get('intSciFi', function (data) {



		// $('#r2d2').empty();
		// $('#r2d2Ind').empty();
		// let rcount = 0;
		// let result = "";
		// let result2 = "";
		// $.each(data, function (I, V) {
		// 	rcount += 1;
		// 	if (rcount == 1) {
		// 		result = result + a1a + a2 + V.carosthumbpath + a4 + a5 + V.movfspath + a7;
		// 		result2 = result2 + b1 + V.movfspath + b2 + "0" + b3a;
		// 	} else {
		// 		result = result + a1 + a2 + V.carosthumbpath + a4 + a5 + V.movfspath + a7;
		// 		c = rcount - 1;
		// 		cstring = c.toString();
		// 		result2 = result2 + b1 + V.movfspath + b2 + cstring + b3b;
		// 	}
		// })
		// $('#r2d2').append(result);
		// $('#r2d2Ind').append(result2);



		mainFunc(data);
	})
})
.on('click', '#startrekBtn', function () {
	$.get('intStarTrek', function (data) {


		// $('#r2d2').empty();
		// $('#r2d2Ind').empty();
		// let rcount = 0;
		// let result = "";
		// let result2 = "";		
		// console.log(data);
		// $.each(data, function (I, V) {
		// 	rcount += 1;
		// 	if (rcount == 1) {
		// 		result = result + a1a + a2 + V.carosthumbpath + a4 + a5 + V.movfspath + a7;
		// 		result2 = result2 + b1 + V.movfspath + b2 + "0" + b3a;
		// 	} else {
		// 		result = result + a1 + a2 + V.carosthumbpath + a4 + a5 + V.movfspath + a7;
		// 		c = rcount - 1;
		// 		cstring = c.toString();
		// 		result2 = result2 + b1 + V.movfspath + b2 + cstring + b3b;
		// 	}
		// })
		// $('#r2d2').append(result);
		// $('#r2d2Ind').append(result2);



		mainFunc(data);
	})
})
.on('click', '#starwarsBtn', function () {
	$.get('intStarWars', function (data) {



		// $('#r2d2').empty();
		// $('#r2d2Ind').empty();
		// let rcount = 0;
		// let result = "";
		// let result2 = "";
		// console.log(data);
		// $.each(data, function (I, V) {
		// 	rcount += 1;
		// 	if (rcount == 1) {
		// 		result = result + a1a + a2 + V.carosthumbpath + a4 + a5 + V.movfspath + a7;
		// 		result2 = result2 + b1 + V.movfspath + b2 + "0" + b3a;
		// 	} else {
		// 		result = result + a1 + a2 + V.carosthumbpath + a4 + a5 + V.movfspath + a7;
		// 		c = rcount - 1;
		// 		cstring = c.toString();
		// 		result2 = result2 + b1 + V.movfspath + b2 + cstring + b3b;
		// 	}
		// })
		// $('#r2d2').append(result);
		// $('#r2d2Ind').append(result2);


		mainFunc(data);
	})
})
.on('click', '#superherosBtn', function () {
	$.get('intSuperHeros', function (data) {


		// $('#r2d2').empty();
		// $('#r2d2Ind').empty();
		// let rcount = 0;
		// let result = "";
		// let result2 = "";


		// console.log(data);
		// $.each(data, function (I, V) {
		// 	rcount += 1;
		// 	if (rcount == 1) {
		// 		result = result + a1a + a2 + V.carosthumbpath + a4 + a5 + V.movfspath + a7;
		// 		result2 = result2 + b1 + V.movfspath + b2 + "0" + b3a;
		// 	} else {
		// 		result = result + a1 + a2 + V.carosthumbpath + a4 + a5 + V.movfspath + a7;
		// 		c = rcount - 1;
		// 		cstring = c.toString();
		// 		result2 = result2 + b1 + V.movfspath + b2 + cstring + b3b;
		// 	}
		// })
		// $('#r2d2').append(result);
		// $('#r2d2Ind').append(result2);



		mainFunc(data);
	})
})


.on('click', '#cartoonsBtn', ()=> {
	$.get('intCartoons', (data)=> {
		console.log("this is cartoons data");
		console.log(data);




		// $('#r2d2').empty();
		// $('#r2d2Ind').empty();
		// let rcount = 0;
		// let result = "";
		// let result2 = "";
		// console.log(data);
		// $.each(data, (I, V)=> {
		// 	rcount += 1;
		// 	if (rcount == 1) {
		// 		result = result + a1a + a2 + V.carosthumbpath + a4 + a5 + V.movfspath + a7;
		// 		result2 = result2 + b1 + V.movfspath + b2 + "0" + b3a;
		// 	} else {
		// 		result = result + a1 + a2 + V.carosthumbpath + a4 + a5 + V.movfspath + a7;
		// 		c = rcount - 1;
		// 		cstring = c.toString();
		// 		result2 = result2 + b1 + V.movfspath + b2 + cstring + b3b;
		// 	}
		// })
		// $('#r2d2').append(result);
		// $('#r2d2Ind').append(result2);



		mainFunc(data);
	})
})





.on('click', '#tremorsBtn', ()=> {
	$.get('intTremors', (data)=> {



		// $('#r2d2').empty();
		// $('#r2d2Ind').empty();
		// let rcount = 0;
		// let result = "";
		// let result2 = "";


		// console.log(data);
		// $.each(data, (I, V)=> {
		// 	rcount += 1;
		// 	if (rcount == 1) {
		// 		result = result + a1a + a2 + V.carosthumbpath + a4 + a5 + V.movfspath + a7;
		// 		result2 = result2 + b1 + V.movfspath + b2 + "0" + b3a;
		// 	} else {
		// 		result = result + a1 + a2 + V.carosthumbpath + a4 + a5 + V.movfspath + a7;
		// 		c = rcount - 1;
		// 		cstring = c.toString();
		// 		result2 = result2 + b1 + V.movfspath + b2 + cstring + b3b;
		// 	}
		// })
		// $('#r2d2').append(result);
		// $('#r2d2Ind').append(result2);


		mainFunc(data);
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
	console.log(bval);
	if ( bval === "Pause") {
		$('#playBtn').html("Play");
	} else {
		$('#playBtn').html("Pause");
	}
	$.ajax({
	"url": omxplayer_server_play,
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