window.addEventListener("load", function load(event){

    var raw = document.getElementById("raw");
    var exported = document.getElementById("exported");
    var feedback = document.getElementById("feedback");
    
    var do_export = function(){

	feedback.innerText = ""
	exported.innerText = "";

	console.log("WTF '" + raw.innerText + "'");

	try {
	    var f = JSON.parse(raw.innerText);
	} catch(err) {
	    feedback.innerText = "Failed to parse feature: " + err;
	    return;
	}

	var str_f = JSON.stringify(f);
	
	export_feature(str_f).then(rsp => {

	    exported.innerText = rsp;
	}).catch(err => {

	    feedback.innerText = "Failed to export feature: " + err;
	});
    };
    
    var init = function(){

	var btn = document.getElementById("submit");

	if (! btn){
	    console.log("Unable to load submit button");
	    return;
	}

	btn.onclick = function(){
	    do_export();
	    return false;
	};

	btn.innerText = "Export";	
	btn.removeAttribute("disabled");
    };
    
    whosonfirst.export.feature.init().then(rsp => {
	init();	
    });
    
});
