window.addEventListener("load", function load(event){
    
    if (! WebAssembly.instantiateStreaming){
	
	WebAssembly.instantiateStreaming = async (resp, importObject) => {
            const source = await (await resp).arrayBuffer();
            return await WebAssembly.instantiate(source, importObject);
	};
    }
    
    const export_go = new Go();
    
    let export_mod, export_inst;
    
    var pending = 1;
    
    WebAssembly.instantiateStreaming(fetch("wasm/export_feature.wasm"), export_go.importObject).then(
	
	async result => {
	    
	    pending -= 1;
	    
	    if (pending == 0){
		enable();
	    }
	    
            export_mod = result.module;
            export_inst = result.instance;
	    await export_go.run(export_inst);
	}
    );
    
    async function enable() {
	
	console.log("OK, EXPORT");
    }
    
});
